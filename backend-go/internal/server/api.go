package server

import (
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
	engine.GET("/api/health", getHealth)
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

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
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

func getGitTag(c *gin.Context) {
	gitHeadBytes, err := os.ReadFile("../.git-tag")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gitHead := strings.TrimSpace(string(gitHeadBytes))
	c.JSON(http.StatusOK, gin.H{"tag": gitHead})
}
