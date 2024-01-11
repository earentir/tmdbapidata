// Package tmdbapidata : This package is responsible for getting data from the TMDB API.
package tmdbapidata

import (
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
)

// SearchTMDB : Search the TMDB API for a given string
func SearchTMDB(APIKey, searchstring string) (string, string, interface{}) {
	tmdbClient, err := tmdb.Init(APIKey)

	if err != nil {
		fmt.Println(err)
		return "Error", "", err
	}

	options := make(map[string]string)
	options["language"] = "en-US"

	// Multi Search
	search, err := tmdbClient.GetSearchMulti(searchstring, options)

	if err != nil {
		fmt.Println(err)
		return "Error", "", err
	}

	// Iterate
	for _, v := range search.Results {
		mediaInfo := tmdbMediaInfo{
			ID:               v.ID,
			Adult:            v.Adult,
			Name:             v.Name,
			Overview:         v.Overview,
			PosterPath:       v.PosterPath,
			ProfilePath:      v.ProfilePath,
			FirstAirDate:     v.FirstAirDate,
			ReleaseDate:      v.ReleaseDate,
			OriginCountry:    v.OriginCountry,
			OriginalLanguage: v.OriginalLanguage,
			MediaType:        v.MediaType,
		}

		if v.MediaType == "movie" {
			mediaInfo.Title = v.Title
			return v.MediaType, v.Title, mediaInfo
		} else if v.MediaType == "tv" {
			mediaInfo.Title = v.Name
			return v.MediaType, v.Name, mediaInfo
		} else if v.MediaType == "person" {
			mediaInfo.Name = v.Name
			return v.MediaType, v.Name, mediaInfo
		}
	}

	return "No results found", "", nil
}
