package scrapper

import (
	"fmt"
	"strings"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

func (fc FeatureCollection) ToGpx() *gpx.GPX {
	exportedGpx := &gpx.GPX{
		Version:          "1.1",
		Creator:          "Export gpx standard de refuges.info",
		XMLNs:            "http://www.topografix.com/GPX/1/1",
		XmlNsXsi:         "http://www.w3.org/2001/XMLSchema-instance",
		XmlSchemaLoc:     "http://www.topografix.com/GPX/1/1 https://www.topografix.com/GPX/1/1/gpx.xsd",
		Name:             "Points de refuges.info. The data included in this document is from www.refuges.info. The data is made available under CC By-Sa 2.0",
		Description:      "grp_rep:refuges.info " + time.Now().Format("2006-01-02 15:04:05"),
		AuthorName:       "Contributeurs refuges.info",
		Copyright:        "Contributeurs refuges.info",
		CopyrightYear:    fmt.Sprintf("%d", time.Now().Year()),
		CopyrightLicense: "http://creativecommons.org/licenses/by-sa/2.0/deed.fr",
		Link:             "https://www.refuges.info",
		LinkText:         "refuges.info",
		LinkType:         "website",
		Waypoints:        make([]gpx.GPXPoint, len(fc.Features)),
	}

	syncChannel := make(chan int, len(fc.Features))
	for i, feature := range fc.Features {
		go func(f Feature, index int) {
			wp := f.ToGpx()
			exportedGpx.Waypoints[index] = *wp
			syncChannel <- 1
		}(feature, i)
	}
	for i := 0; i < len(fc.Features); i++ {
		<-syncChannel
	}
	return exportedGpx
}
func formatStringForGpx(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

func (f Feature) ToGpx() *gpx.GPXPoint {
	point := gpx.Point{
		Latitude:  f.Geometry.Coordinates.Latitude(),
		Longitude: f.Geometry.Coordinates.Longitude(),
	}
	description := ""
	if f.Properties.Description.Valeur != "" {
		description += fmt.Sprintf("Description: \n\n%s", f.Properties.Description.Valeur)
	}
	if f.Properties.Remarque.Valeur != "" {
		description += "\n\n*****\n\n"
		description += fmt.Sprintf("Remarque: \n\n%s", f.Properties.Remarque.Valeur)
	}
	if f.Properties.Acces.Valeur != "" {
		description += "\n\n*****\n\n"
		description += fmt.Sprintf("Accès: \n\n%s", f.Properties.Acces.Valeur)
	}
	if f.Comments != nil {
		comments := SummarizeComments(f.Comments)
		if comments != "" {
			description += "\n\n*****\n\n"
			description += "Voici un résumé des commentaires:\n\n"
			description += comments
		}
	}
	return &gpx.GPXPoint{
		Point:       point,
		Name:        f.Properties.Name,
		Description: gpx.CDATA(formatStringForGpx(description)),
		Comment:     f.Properties.Link,
	}
}

func ExportFeatureCollection(featureCollection *FeatureCollection) ([]byte, error) {
	gpxData := featureCollection.ToGpx()
	gpxBytes, err := gpxData.ToXml(gpx.ToXmlParams{
		Version: "1.1",
		Indent:  true,
	})
	if err != nil {
		return nil, err
	}
	return gpxBytes, nil
}
