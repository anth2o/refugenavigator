package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func main() {
	runtime.SetBlockProfileRate(1)
	start := time.Now()
	bbox := scrapper.BoundingBox{
		NorthEast: scrapper.Point{5.52315, 44.9159},
		SouthWest: scrapper.Point{5.49826, 44.8983},
	}

	featureCollection := scrapper.GetFeatureCollection(bbox)
	scrapper.EnrichFeatureCollection(featureCollection)
	gpxBytes, err := scrapper.ExportFeatureCollection(featureCollection)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("gpxBytes: %d\n", len(gpxBytes))
	fmt.Printf("time: %s\n", time.Since(start))
	blockFile, err := os.Create("block.prof")
	if err != nil {
		log.Fatal("could not create block profile: ", err)
	}
	defer blockFile.Close()

	p := pprof.Lookup("block")
	if p == nil {
		log.Fatal("block profile is not enabled")
	}
	if err := p.WriteTo(blockFile, 0); err != nil {
		log.Fatal("could not write block profile: ", err)
	}
}
