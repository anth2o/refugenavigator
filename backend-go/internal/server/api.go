package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anth2o/refugenavigator/internal/scrapper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getMode() string {
	return os.Getenv("GIN_MODE")
}
func getHost() string {
	mode := getMode()
	if mode != "release" {
		// for hot reload with air
		return "127.0.0.1"
	}
	return ""
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

const staticPath = "../static/"
const staticUrl = "/static"

func setupRoutes() *gin.Engine {
	fmt.Println("Setting up routes")
	defer func() { fmt.Println("Routes set up") }()
	engine := gin.Default()
	mode := getMode()
	if mode != "release" {
		// for local dev with yarn run dev, could be optimized by removing it from prod docker
		engine.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://127.0.0.1:5173"},
			AllowMethods: []string{"GET"},
		}))
	}
	engine.GET("/api/points", getPoints)
	engine.GET("/api/gpx/download", sendGpxFile)
	engine.GET("/api/gpx/static", sendGpxStaticUrl)
	engine.GET("/api/git-tag", getGitTag)
	engine.Static("/site", "../frontend/dist")
	engine.Static(staticUrl, staticPath)
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/site")
	})
	return engine
}

func Run() {
	engine := setupRoutes()
	if err := engine.Run(getHost() + ":" + getPort()); err != nil {
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

func getFeatureCollection(c *gin.Context) (scrapper.BoundingBox, *scrapper.FeatureCollection) {
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
		return bbox, nil
	}
	return bbox, scrapper.GetFeatureCollection(bbox, nil)

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
	_, featureCollection := getFeatureCollection(c)
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

func getGpxBytes(c *gin.Context) (scrapper.BoundingBox, []byte) {
	bbox, featureCollection := getFeatureCollection(c)
	if featureCollection == nil {
		returnGinErrors(c)
		return bbox, nil
	}
	scrapper.EnrichFeatureCollection(featureCollection, nil)
	gpxBytes, err := scrapper.ExportFeatureCollection(featureCollection)
	if err != nil {
		c.Error(err)
		returnGinErrors(c)
		return bbox, nil
	}
	return bbox, gpxBytes
}

func sendGpxFile(c *gin.Context) {
	_, gpxBytes := getGpxBytes(c)
	c.Header("Content-Type", "application/gpx+xml")
	c.Header("Content-Disposition", "attachment; filename=route.gpx")
	c.Data(200, "application/gpx+xml", gpxBytes)

}

func sendGpxStaticUrl(c *gin.Context) {
	// gpxsURL, err := getGpxUrl(c.Request.URL.String())
	// if err != nil {
	// 	c.Error(err)
	// 	returnGinErrors(c)
	// 	return
	// }
	bbox, gpxBytes := getGpxBytes(c)
	if gpxBytes == nil {
		returnGinErrors(c)
		return
	}
	fileName := filepath.Join("gpx", bbox.String()+".gpx")
	filePath := filepath.Join(staticPath, fileName)
	if err := WriteFile(gpxBytes, filePath); err != nil {
		c.Error(err)
		returnGinErrors(c)
		return
	}
	// https://manuels.iphigen.ie/fr/article/imports-gpx-iphigenie-ios-12d3iji/#3-importer-directement-un-fichier-qui-se-trouve-sur-le-net
	fileUrl := "https://" + filepath.Join(c.Request.Host, staticUrl, fileName)
	fmt.Println(fileUrl)
	c.JSON(http.StatusOK, gin.H{"url": fileUrl})
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
