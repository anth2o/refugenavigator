package server

import (
	"errors"
	"net/url"
)

func getGpxUrl(originalURL string) (string, error) {
	if originalURL == "" {
		return originalURL, errors.New("Missing URL")
	}
	parsedURL, err := url.Parse(originalURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return originalURL, err
	}

	// see https://manuels.iphigen.ie/fr/article/imports-gpx-iphigenie-ios-12d3iji/#3-importer-directement-un-fichier-qui-se-trouve-sur-le-net
	switch parsedURL.Scheme {
	case "http":
		parsedURL.Scheme = "gpx"
	case "https":
		parsedURL.Scheme = "gpxs"
	}

	return parsedURL.String(), nil
}
