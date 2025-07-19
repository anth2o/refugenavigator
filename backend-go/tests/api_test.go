package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/anth2o/refugenavigator/internal/server"
	"github.com/gin-gonic/gin"
)

func TestGetGPX(t *testing.T) {
	// Read the reference GPX file
	refGPX, err := os.ReadFile("../data/exported_enriched.gpx")
	if err != nil {
		t.Fatalf("Failed to read reference GPX file: %v", err)
	}

	// Create a test server
	gin.SetMode(gin.TestMode)
	engine := server.SetupRoutes()

	// Create a test request with the bounding box from getBoundingBoxStringTest
	req, err := http.NewRequest("GET", "/api/gpx", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	bboxTest := getBoundingBoxTest()
	// Add the bounding box query parameters
	req.URL.RawQuery = "SouthWest.Latitude=" + strconv.FormatFloat(bboxTest.SouthWest.Latitude(), 'f', -1, 64) + "&SouthWest.Longitude=" + strconv.FormatFloat(bboxTest.SouthWest.Longitude(), 'f', -1, 64) + "&NorthEast.Latitude=" + strconv.FormatFloat(bboxTest.NorthEast.Latitude(), 'f', -1, 64) + "&NorthEast.Longitude=" + strconv.FormatFloat(bboxTest.NorthEast.Longitude(), 'f', -1, 64)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Perform the request
	engine.ServeHTTP(rec, req)

	// Check the response status
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	// Check the content type
	if rec.Header().Get("Content-Type") != "application/gpx+xml" {
		t.Errorf("Expected Content-Type application/gpx+xml, got %s", rec.Header().Get("Content-Type"))
	}

	// Compare the response with the reference GPX
	responseGPX := rec.Body.Bytes()
	refGPXStr := string(refGPX)
	responseGPXStr := string(responseGPX)
	if refGPXStr != responseGPXStr {
		diff, err := diffLines(refGPXStr, responseGPXStr)
		if err != nil {
			panic(err)
		}
		t.Errorf("GPX content mismatch:\n%s", diff)
	}
}

func TestGitTag(t *testing.T) {
	// Create a test server
	gin.SetMode(gin.TestMode)
	guard := monkey.Patch(server.GetGitTag, func() string {
		return "mocked-tag"
	})
	defer guard.Unpatch()
	engine := server.SetupRoutes()

	// Create a test request
	req, err := http.NewRequest("GET", "/api/git-tag", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Perform the request
	engine.ServeHTTP(rec, req)

	// Check the response status
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	// Check the response body
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	if response["tag"] != "mocked-tag" {
		t.Errorf("Expected tag mocked-tag, got %s", response["tag"])
	}
}
