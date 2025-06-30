package tests

import (
	"fmt"
	"strings"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func getBoundingBoxTest() scrapper.BoundingBox {
	return scrapper.BoundingBox{
		NorthEast: scrapper.Point{5.52315, 44.9159},
		SouthWest: scrapper.Point{5.49826, 44.8983},
	}
}

func getBoundingBoxStringTest() string {
	return "bbox=5.49826,44.89830,5.52315,44.91590"
}

func getFeatureCollectionTest() *scrapper.FeatureCollection {
	var features []scrapper.Feature = []scrapper.Feature{}
	features = append(features, scrapper.Feature{
		Id:         28,
		Properties: scrapper.Properties{Name: "Refuge de la Jasse du Play", Coord: scrapper.Coord{Altitude: 1629}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5021, 44.91067}},
	})
	features = append(features, scrapper.Feature{
		Id:         1198,
		Properties: scrapper.Properties{Name: "Fontaine du Play", Coord: scrapper.Coord{Altitude: 1670}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.51051, 44.90526}},
	})
	features = append(features, scrapper.Feature{
		Id:         1199,
		Properties: scrapper.Properties{Name: "Deuxième fontaine du Play", Coord: scrapper.Coord{Altitude: 1670}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5093, 44.9035}},
	})
	features = append(features, scrapper.Feature{
		Id:         1987,
		Properties: scrapper.Properties{Name: "Rocher de Séguret", Coord: scrapper.Coord{Altitude: 2051}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.52081, 44.90792}},
	})
	features = append(features, scrapper.Feature{
		Id:         1986,
		Properties: scrapper.Properties{Name: "Pas de Bèrrièves", Coord: scrapper.Coord{Altitude: 1887}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5173, 44.90996}},
	})
	return &scrapper.FeatureCollection{Features: features}
}

func getFeatureCollectionEnrichedTest() *scrapper.FeatureCollection {
	featureCollection := getFeatureCollectionTest()
	featureCollection.Features[0].Properties.Description = scrapper.Description{Value: "Places prévues pour dormir: 10\nBon état général, \r\n1 table, 2 bancs, pelle et balai, fil à linge.\r\nCouchage sur plancher à l'étage.\r\nManque un livre d'or.\r\nPlus de scie\r\n[->1198] à proximité\r\n\r\nIl y a du reseau GSM disponible, il faut se rendre sur la petite colline à 50m au nord ouest de la cabane pour capter de la 4G avec un assez bon débit. Sur place, aucune indication pour la source n'est donnée, retenez bien sa position avant de vous y rendre. Ou alors regardez sur internet depuis la colline !\nLe GR passe devant, a disons 4h ou 5h de Pré Peyret au sud, sans doute une journée de Corrençon au nord\r\nle refuge est bien en vue au milieu d'un endroit assez dégagé\nParc Regional du Vercors (OT Villard ou La chapelle)\n"}
	featureCollection.Features[1].Properties.Description = scrapper.Description{Value: "Que chaque randonneur qui l'a vue donne sa date de passage et l'état de la source, cela permet à ceux qui veulent y aller de prévoir leur chance d'y trouver de l'eau. \r\n\r\n Captage sécurisé avec grillages afin d'éviter les souillures par les animaux (mais la source reste accessible aux randonneurs) - avril 2014\nA partir de la cabane de la Jasse du Play, continuer sur le GR 91 vers le Sud-Est le GR passe maintenant (08/2018) devant la source.\r\n\r\nLa fontaine se trouve à gauche au bas d’un pan rocheux.\n\n"}
	featureCollection.Features[2].Properties.Description = scrapper.Description{Value: "Pas vraiment une fontaine mais plutôt une simple source non aménagée d'où coule assez peu d'eau\n200m au sud de la première fontaine du Play. Pas de chemin.\r\nde la première fontaine continuer sous la petite falaise rocheuse\r\nRedescendre légèrement pour traverser un petit pierrier \r\nEntrer dans la forêt, la fontaine est là au pied d'épicéas\r\nJe l'ai trouvé à l'oreille car on ne la voit pas de loin\n\n"}
	featureCollection.Features[3].Properties.Description = scrapper.Description{Value: "Comme pour le sommet de Pierre Blanche, par les crètes on peut atteindre le Pas de la Ville.\r\nRéservés toutefois à des randonneurs ayant l'habitude de progresser en terrains de crètes et hors sentiers...!\nDu hameau les Petits Deux.\r\nSuivre l'itinéraire qui vous mène au Pas de Bèrrièves.\r\nEnsuite vers la gauche(sud est), on gagne le sommet du Rocher de Séguret.\r\nLa descente s'effectue par le mème itinéraire...!\n\n"}
	featureCollection.Features[4].Properties.Description = scrapper.Description{Value: "Du Pas de Bèrrièves, on a la possibilité de descendre sur le refuge de la Jasse du Play, bien visible.\r\nOu par les crètes de rejoindre le Pas de la Ville vers le(sud).\r\nUn des nombreux Passage pour accéder sur le plateau du Vercors...!\nDeux kilomètres avant Gresse en Vercors, prendre la route à droite qui vous emmène au hameau les Petits Deux,et à droite le Col des Deux(1222)m. Parking le long de la route à gauche.\r\nRevenir sur ces pas, entre les maisons démarre notre sentier.\r\nVers l'(ouest), il longe de grandes prairies, et entre franchement en forèt...!\n\n"}
	return featureCollection
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func diffLines(expected, actual string) (string, error) {
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")
	diff := ""
	for i := 0; i < len(expectedLines) || i < len(actualLines); i++ {
		expectedLine := ""
		if i < len(expectedLines) {
			expectedLine = expectedLines[i]
		}
		actualLine := ""
		if i < len(actualLines) {
			actualLine = actualLines[i]
		}
		if expectedLine != actualLine {
			diff += fmt.Sprintf("< %s\n", expectedLine)
			diff += fmt.Sprintf("> %s\n", actualLine)
		}
	}
	return diff, nil
}
