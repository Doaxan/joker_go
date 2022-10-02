package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Check if a slice contains a string
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Return GET http.Request with provided url
func createRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("createRequest error occurred")
	}

	return req
}

// Call getJsonToTarget and decode JSON-encoded value into some interface
func decodeJsonToTarget(url string, target interface{}) {
	if err := getJsonToTarget(url, target); err != nil {
		log.Fatal("Can't decode json")
	}
}

// Sends an HTTP request, return JSON-encoded value
func getJsonToTarget(url string, target interface{}) error {
	client := &http.Client{}
	res, err := client.Do(createRequest(url))
	if err != nil {
		log.Fatal("Possible no internet connection")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("res.Body.Close() error occurred")
		}
	}(res.Body)

	return json.NewDecoder(res.Body).Decode(target)
}

// Get randomJokeCategory and assigns it to map
func jokesCategoriesToMap(jokesCount int, categories []string) map[string][]string {
	categoriesJokes := map[string][]string{}
	for _, category := range categories {
		var jokes []string
		for retryCount := 0; len(jokes) < jokesCount; retryCount++ {
			joke := randomJokeCategory(category)
			if contains(jokes, joke) {
				if retryCount == jokeGetRetryCount {
					break
				}
				retryCount++
				continue
			}
			retryCount = 0
			jokes = append(jokes, joke)
		}
		categoriesJokes[category] = jokes
	}
	return categoriesJokes
}

// Write dataMap value to a file in the following format: <key>.txt
func writeMapToFile(dataMap map[string][]string) error {
	for key, slice := range dataMap {
		err := os.MkdirAll("jokes", os.ModePerm)
		if err != nil {
			return err
		}

		file, err := os.Create("jokes/" + key + ".txt")
		if err != nil {
			return err
		}

		for index, value := range slice {
			// Add a newline except the last one
			if index != len(slice)-1 {
				value = value + "\n"
			}

			_, err := file.WriteString(value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
