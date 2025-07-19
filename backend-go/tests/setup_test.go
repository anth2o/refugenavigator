package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func TestMain(m *testing.M) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/api/bbox"):
			bbox := r.URL.Query().Get("bbox")
			expectedBbox := getBoundingBoxStringTest()
			if bbox != expectedBbox {
				fmt.Printf("Got %q, expected %q", bbox, expectedBbox)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			data, err := os.ReadFile("../data/bbox.json")
			if err != nil {
				fmt.Printf("Failed to read bbox.json: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case strings.Contains(r.URL.Path, "/api/point"):
			// fetching point from API
			id := r.URL.Query().Get("id")
			if id == "" {
				fmt.Printf("Failed to extract ID from URL: %v", r.URL.Query())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if id == "-1" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			data, err := os.ReadFile(fmt.Sprintf("../data/%s.json", id))
			if err != nil {
				fmt.Printf("Failed to read feature data: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		case strings.Contains(r.URL.Path, "/point"):
			// fetching point from HTML
			id, err := strconv.Atoi(r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1 : len(r.URL.Path)])
			if err != nil {
				fmt.Printf("Failed to extract ID from URL: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if id == -1 {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			data, err := os.ReadFile(fmt.Sprintf("../data/%d.html", id))
			if err != nil {
				fmt.Printf("Failed to read feature data: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html")
			w.Write(data)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	originalBaseUrl := scrapper.GetBaseUrl()
	scrapper.SetBaseUrl(testServer.URL + "/")
	defer scrapper.SetBaseUrl(originalBaseUrl)

	guardTime := monkey.Patch(time.Now, func() time.Time {
		return time.Date(2025, 6, 29, 12, 13, 24, 0, time.UTC)
	})
	defer guardTime.Unpatch()

	guardLlm := monkey.Patch(scrapper.SummarizeComments, func(comments []scrapper.Comment) string {
		return "This is a mocked summary:\n\n" + fmt.Sprintf("%v", comments)
	})
	defer guardLlm.Unpatch()
	exitCode := m.Run()

	os.Exit(exitCode)
}
