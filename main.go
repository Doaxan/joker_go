package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	// Number of attempts to get a unique joke
	jokeGetRetryCount = 5

	urlJokesRandom         = "https://api.chucknorris.io/jokes/random"
	urlJokesCategories     = "https://api.chucknorris.io/jokes/categories"
	urlJokesRandomCategory = "https://api.chucknorris.io/jokes/random?category="

	helpMessage = `Usage of ./joker:
  random: Retrieve a random chuck joke
  dump: Get 5 random unique jokes for each of the existing categories and saves them to text files - one for each of the categories
  dump -n int: Get n random unique jokes for each of the existing categories and saves them to text files - one for each of the categories`
)

type JokesRandom struct {
	Joke string `json:"value"`
}

type JokesCategories []string

func main() {
	// flag.Parse() does not allow non-flag arguments, so use os.Args with switch
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(randomJoke())
		return
	}

	switch args[0] {
	case "random":
		fmt.Println(randomJoke())
	case "dump":
		switch argsLen := len(args); {
		case argsLen == 1:
			dumpJokesToFiles(5)
		case argsLen == 3:
			n, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("-n value must be an int value")
				return
			}
			if n > 0 {
				dumpJokesToFiles(n)
			} else {
				fmt.Println("-n value must be greater than 0")
			}
		default:
			fmt.Println(helpMessage)
		}
	case "-h":
		fallthrough
	case "--help":
		fallthrough
	default:
		fmt.Println(helpMessage)
	}

}

// Get jokesCount random unique jokes for each of the existing categories
// and saves them to text files - one for each of the categories
func dumpJokesToFiles(jokesCount int) {
	var categories JokesCategories
	decodeJsonToTarget(urlJokesCategories, &categories)

	err := writeMapToFile(jokesCategoriesToMap(jokesCount, categories))
	if err != nil {
		log.Fatal("writeMapToFile error occurred")
	}
}

// Return random joke with provided category
func randomJokeCategory(category string) string {
	var jokeCategory JokesRandom
	decodeJsonToTarget(urlJokesRandomCategory+category, &jokeCategory)

	return jokeCategory.Joke
}

// Return random joke string
func randomJoke() string {
	var chuck JokesRandom
	decodeJsonToTarget(urlJokesRandom, &chuck)

	return chuck.Joke
}
