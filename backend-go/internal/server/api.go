package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anth2o/refugenavigator/internal/scrapper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:5173"},
		AllowMethods: []string{"GET"},
	}))
	engine.GET("/api/gpx", getGPX)
	return engine
}

func getQuery(c *gin.Context, key string) string {
	value, ok := c.GetQuery(key)
	if !ok {
		c.Error(errors.New("Missing query parameter " + key))
		return ""
	}
	return value
}

func getGPX(c *gin.Context) {
	swLat, _ := strconv.ParseFloat(getQuery(c, "SouthWest.Latitude"), 64)
	swLon, _ := strconv.ParseFloat(getQuery(c, "SouthWest.Longitude"), 64)
	neLat, _ := strconv.ParseFloat(getQuery(c, "NorthEast.Latitude"), 64)
	neLon, _ := strconv.ParseFloat(getQuery(c, "NorthEast.Longitude"), 64)

	if c.Errors != nil {
		var errors []string
		for _, err := range c.Errors {
			errors = append(errors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	swPoint := scrapper.Point{swLon, swLat}
	nePoint := scrapper.Point{neLon, neLat}
	bbox := scrapper.BoundingBox{
		SouthWest: swPoint,
		NorthEast: nePoint,
	}
	fmt.Printf("bbox: %s\n", bbox)

	featureCollection := scrapper.GetFeatureCollection(bbox, nil)
	scrapper.EnrichFeatureCollection(featureCollection, nil)

	gpxBytes, err := scrapper.ExportFeatureCollection(featureCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/gpx+xml")
	c.Header("Content-Disposition", "attachment; filename=route.gpx")
	c.Data(200, "application/gpx+xml", gpxBytes)
}
