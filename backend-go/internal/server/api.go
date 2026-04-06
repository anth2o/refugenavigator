package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/anth2o/refugenavigator/internal/scrapper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func setupRoutes() *gin.Engine {
	fmt.Println("Setting up routes")
	defer func() { fmt.Println("Routes set up") }()
	engine := gin.Default()
	if mode := os.Getenv("GIN_MODE"); mode != "release" {
		// for local dev with yarn run dev, could be optimized by removing it from prod docker
		engine.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://127.0.0.1:5173"},
			AllowMethods: []string{"GET"},
		}))
	}
	engine.GET("/api/points", getPoints)
	engine.GET("/api/gpx", getGPX)
	engine.GET("/api/git-tag", getGitTag)
	engine.Static("/site", "../frontend/dist")
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/site")
	})
	return engine
}

func Run() {
	engine := setupRoutes()
	if err := engine.Run(":" + getPort()); err != nil {
		log.Fatal(err)
	}
}

func getQuery(c *gin.Context, key string) string {
	value, ok := c.GetQuery(key)
	if !ok {
		c.Error(errors.New("Missing query parameter " + key))
		return ""
	}
	return value
}

func getFeatureCollection(c *gin.Context) *scrapper.FeatureCollection {
	swLat, _ := strconv.ParseFloat(getQuery(c, "SouthWest.Latitude"), 64)
	swLon, _ := strconv.ParseFloat(getQuery(c, "SouthWest.Longitude"), 64)
	neLat, _ := strconv.ParseFloat(getQuery(c, "NorthEast.Latitude"), 64)
	neLon, _ := strconv.ParseFloat(getQuery(c, "NorthEast.Longitude"), 64)

	swPoint := scrapper.Point{swLon, swLat}
	nePoint := scrapper.Point{neLon, neLat}
	bbox := scrapper.BoundingBox{
		SouthWest: swPoint,
		NorthEast: nePoint,
	}

	if bbox.Area() >= 1 { // this value was chosen so that the response time from refuges.info API is almost instantaneous
		c.Error(errors.New("Area is too large: try selecting a smaller one."))
		return nil
	}
	return scrapper.GetFeatureCollection(bbox, nil)

}

func returnGinErrors(c *gin.Context) {
	if c.Errors != nil {
		var errors []string
		for _, err := range c.Errors {
			errors = append(errors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
	}
}

func getPoints(c *gin.Context) {
	featureCollection := getFeatureCollection(c)
	if featureCollection != nil {
		bytes, err := json.Marshal(featureCollection)
		if err == nil {
			c.Header("Content-Type", "application/json")
			c.Data(200, "application/json", bytes)
			return
		} else {
			c.Error(err)
		}
	}
	returnGinErrors(c)
}

func getGPX(c *gin.Context) {
	featureCollection := getFeatureCollection(c)
	if featureCollection != nil {
		scrapper.EnrichFeatureCollection(featureCollection, nil)
		gpxBytes, err := scrapper.ExportFeatureCollection(featureCollection)
		if err == nil {

			c.Header("Content-Type", "application/gpx+xml")
			c.Header("Content-Disposition", "attachment; filename=route.gpx")
			c.Data(200, "application/gpx+xml", gpxBytes)
		} else {
			c.Error(err)
		}

	}
	returnGinErrors(c)
}

func getGitTag(c *gin.Context) {
	gitHeadBytes, err := os.ReadFile("../.git-tag")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gitHead := strings.TrimSpace(string(gitHeadBytes))
	c.JSON(http.StatusOK, gin.H{"tag": gitHead})
}
