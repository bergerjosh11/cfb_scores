package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    // Create a new collector
    c := colly.NewCollector()

    // URL of ESPN's NCAAF scores and schedules
    url := "https://www.espn.com/college-football/scoreboard"

    // Find and print NCAAF scores
    c.OnHTML(".scoreboard", func(e *colly.HTMLElement) {
        e.ForEach(".scoreCell", func(_ int, el *colly.HTMLElement) {
            team1 := el.ChildText(".teamName span")
            team2 := el.ChildText(".teamName span:last-child")
            score := el.ChildText(".total")

            fmt.Printf("%s vs %s: %s\n", team1, team2, score)
        })
    })

    // Find and print NCAAF schedules
    c.OnHTML(".gamePods", func(e *colly.HTMLElement) {
        e.ForEach(".gamePod", func(_ int, el *colly.HTMLElement) {
            date := el.ChildText(".gamePod-header .date")
            teams := el.ChildText(".competitors")

            fmt.Printf("%s - %s\n", date, teams)
        })
    })

    // Start scraping
    c.Visit(url)
}
