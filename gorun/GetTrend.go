package main

import (
    "fmt"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New()
    trends, err := scraper.GetTrends()
    if err != nil {
        panic(err)
    }
    fmt.Println(trends)
}