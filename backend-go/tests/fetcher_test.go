package tests

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/anth2o/refugenavigator/internal/scrapper"
	"github.com/google/go-cmp/cmp"
)

func TestQueryUrl(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
		wantErr        bool
	}{
		{
			name:           "Successful request",
			url:            "https://example.com",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello World",
			wantErr:        false,
		},
		{
			name:           "404 Not Found",
			url:            "https://example.com/not-found",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not Found",
			wantErr:        false,
		},
		{
			name:    "Network error",
			url:     "https://invalid-url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/not-found" {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte(tt.expectedBody))
				} else {
					w.WriteHeader(tt.expectedStatus)
					w.Write([]byte(tt.expectedBody))
				}
			}))
			defer ts.Close()

			// Override the default http.Client with our test server
			originalClient := http.DefaultClient
			http.DefaultClient = ts.Client()
			defer func() {
				http.DefaultClient = originalClient
			}()

			// Test the QueryUrl method
			querier := scrapper.DefaultQuerier{}
			// Remove the base URL from the test URL since we're using a test server
			testUrl := tt.url
			testUrl = strings.TrimPrefix(testUrl, "https://example.com")
			result := querier.QueryUrl(ts.URL + testUrl)

			if tt.wantErr {
				if result != nil {
					t.Errorf("QueryUrl() should have returned nil for error case, got %v", result)
				}
			} else {
				if string(result) != tt.expectedBody {
					t.Errorf("QueryUrl() = %v, want %v", string(result), tt.expectedBody)
				}
			}
		})
	}
}

type UnifiedQuerier struct {
	t *testing.T
}

func (q UnifiedQuerier) QueryUrl(url string) []byte {
	// Check if this is a bbox query
	if strings.Contains(url, "/bbox") {
		expectedUrl := scrapper.GetBaseUrl() + "/bbox?" + getBoundingBoxStringTest()
		if url != expectedUrl {
			q.t.Errorf("QueryUrl() = %s, want %s", url, expectedUrl)
		}
		data, err := os.ReadFile("../data/bbox.json")
		checkError(err)
		return data
	}

	// Check if this is a point query
	if strings.Contains(url, "/point") {
		id, err := extractId(url)
		if err != nil {
			q.t.Errorf("QueryUrl() = %s, want to be able to extract an id", url)
		}
		expectedUrl := scrapper.GetBaseUrl() + "/point?id=" + strconv.Itoa(id) + "&detail=complet&format=geojson&format_texte=markdown"
		if url != expectedUrl {
			q.t.Errorf("QueryUrl() = %s, want %s", url, expectedUrl)
		}
		data, err := os.ReadFile("../data/" + strconv.Itoa(id) + ".json")
		checkError(err)
		return data
	}

	q.t.Errorf("QueryUrl() = %s, want either bbox or point query", url)
	return nil
}

func extractId(rawurl string) (int, error) {
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		return 0, fmt.Errorf("url %s is not valid", rawurl)
	}
	idStr, ok := parsedUrl.Query()["id"]
	if !ok {
		return 0, fmt.Errorf("url %s does not contain an id", rawurl)
	}
	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		return 0, fmt.Errorf("cannot convert %s to int", idStr[0])
	}
	return id, nil
}

func TestExtractId(t *testing.T) {
	tests := []struct {
		name    string
		rawurl  string
		want    int
		wantErr bool
	}{
		{
			name:   "valid url",
			rawurl: "https://example.com/point?id=1234&detail=complet&format=geojson",
			want:   1234,
		},
		{
			name:    "url without id",
			rawurl:  "https://example.com/point?detail=complet&format=geojson",
			want:    0,
			wantErr: true,
		},
		{
			name:    "url with id not convertible to int",
			rawurl:  "https://example.com/point?id=abc&detail=complet&format=geojson",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractId(tt.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := UnifiedQuerier{t: t}
	featureCollection := scrapper.GetFeatureCollection(bbox, querier)
	expectedFeatureCollection := getFeatureCollectionTest()
	diff := cmp.Diff(featureCollection, expectedFeatureCollection)
	if diff != "" {
		t.Errorf("GetFeatureCollection() did not produce the expected FeatureCollection:\n%s", diff)
	}
}

func TestGetFeature(t *testing.T) {
	querier := UnifiedQuerier{t: t}
	feature := *scrapper.GetFeature(1198, querier)
	expectedFeature := getFeatureCollectionEnrichedTest().Features[1]
	diff := cmp.Diff(feature, expectedFeature)
	if diff != "" {
		t.Errorf("GetFeature() did not produce the expected Feature:\n%s", diff)
	}
}

func TestEnrichFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := UnifiedQuerier{t: t}
	featureCollection := scrapper.GetFeatureCollection(bbox, querier)
	scrapper.EnrichFeatureCollection(featureCollection, querier)
	expectedFeatureCollection := getFeatureCollectionEnrichedTest()
	diff := cmp.Diff(featureCollection, expectedFeatureCollection)
	if diff != "" {
		t.Errorf("EnrichFeatureCollection() did not produce the expected FeatureCollection:\n%s", diff)
	}
}
