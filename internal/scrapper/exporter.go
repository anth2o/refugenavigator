package scrapper

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/twpayne/go-gpx"
)

func (fc FeatureCollection) ToGpx() *gpx.GPX {
	exportedGpx := gpx.GPX{
		Version: "1.1",
		Creator: "Export gpx standard de refuges.info",
		Metadata: &gpx.MetadataType{
			Name: "Points de refuges.info. The data included in this document is from www.refuges.info. The data is made available under CC By-Sa 2.0",
			Desc: "grp_rep:refuges.info " + time.Now().Format("2006-01-02 15:04:05"), // important so that features are in the same group in iPhigenie
			Author: &gpx.PersonType{
				Name: "Contributeurs refuges.info",
			},
			Copyright: &gpx.CopyrightType{
				Author:  "Contributeurs refuges.info",
				Year:    time.Now().Year(),
				License: "http://creativecommons.org/licenses/by-sa/2.0/deed.fr",
			},
			Link: []*gpx.LinkType{{
				HREF: "https://www.refuges.info",
			}},
		},
	}
	for _, feature := range fc.Features {
		exportedGpx.Wpt = append(exportedGpx.Wpt, feature.ToGpx())
	}
	return &exportedGpx
}
func formatStringForGpx(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "'", " ")
	return s
}

func (f Feature) ToGpx() *gpx.WptType {
	return &gpx.WptType{
		Lat:  f.Geometry.Coordinates.latitude(),
		Lon:  f.Geometry.Coordinates.longitude(),
		Name: f.Properties.Name,
		Desc: formatStringForGpx(f.Properties.Description.Value),
	}
}

func ExportFeatureCollection(featureCollection *FeatureCollection, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(xml.Header); err != nil {
		fmt.Printf("err == %v", err)
		return
	}
	if err := featureCollection.ToGpx().WriteIndent(f, "", "  "); err != nil {
		fmt.Printf("err == %v", err)
		return
	}
	fmt.Printf("GPX was successfully exported to %s\n", outputFile)
}
