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
		Waypoints:        make([]gpx.GPXPoint, 0),
	}

	for _, feature := range fc.Features {
		wp := feature.ToGpx()
		exportedGpx.Waypoints = append(exportedGpx.Waypoints, *wp)
	}
	return exportedGpx
}
func formatStringForGpx(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

func (f Feature) ToGpx() *gpx.GPXPoint {
	point := gpx.Point{
		Latitude:  f.Geometry.Coordinates.latitude(),
		Longitude: f.Geometry.Coordinates.longitude(),
	}
	return &gpx.GPXPoint{
		Point:       point,
		Name:        f.Properties.Name,
		Description: gpx.CDATA(formatStringForGpx(f.Properties.Description.Value)),
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
