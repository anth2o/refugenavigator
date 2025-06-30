package scrapper

import (
	"encoding/xml"
	"fmt"
	"os"
)

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
