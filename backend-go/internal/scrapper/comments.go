package scrapper

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gage-technologies/mistral-go"
)

func ScrapeComments(html string) ([]Comment, error) {
	var comments []Comment
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	doc.Find(".bloc_commentaire").Each(func(i int, commentDiv *goquery.Selection) {
		comment := Comment{}

		anchor := commentDiv.Find("a").First()
		if id, exists := anchor.Attr("id"); exists {
			comment.ID = strings.TrimPrefix(id, "C")
		}

		metaInfo := commentDiv.Find(".commentaire_metainfo").Text()
		if parts := strings.Split(metaInfo, "par"); len(parts) > 1 {
			comment.Date = strings.TrimSpace(parts[0])
			comment.Author = strings.TrimSpace(parts[1])
		}

		content := commentDiv.Find("blockquote p").Text()
		comment.Content = strings.TrimSpace(content)

		photoDiv := commentDiv.Find(".photos")
		if photoDiv.Length() > 0 {
			img := photoDiv.Find("img").First()
			if url, exists := img.Attr("src"); exists {
				comment.PhotoURL = url
			}
			if photoDate := photoDiv.Find(".texte_sur_image").Text(); photoDate != "" {
				comment.PhotoDate = photoDate
			}
		}

		comments = append(comments, comment)
	})

	return comments, nil
}

var prompt = "Je vais te donner une liste de commentaires ecrits par des utilisateurs sur un point d'eau. Je voudrais que tu me donnes un resume des commentaires en francais en me donnant le debit moyen attendu en fonction de la saison. Voici la liste des commentaires : "

func (feature Feature) GetCommentsSummaryFromLlm() string {
	client := mistral.NewMistralClientDefault(os.Getenv("MISTRAL_API_KEY"))
	var chatMessages []mistral.ChatMessage
	chatMessages = append(chatMessages, mistral.ChatMessage{Content: prompt, Role: mistral.RoleUser})
	for _, comment := range feature.CommentData.Comments {
		chatMessages = append(chatMessages, mistral.ChatMessage{Content: comment.Date + ": " + comment.Content, Role: mistral.RoleUser})
	}
	chatRes, err := client.Chat("mistral-small", chatMessages, nil)
	if err != nil {
		fmt.Printf("Error getting chat completion: %v", err)
		return ""
	}
	return chatRes.Choices[0].Message.Content
}
