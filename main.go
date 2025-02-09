package main

import (
    "fmt"
    "regexp"
    "github.com/go-rod/rod"
)

func main() {
    browser := rod.New().MustConnect()
    page := browser.MustPage("https://portfolio-p4qj-git-main-trajicks-projects.vercel.app/")
    page.MustWaitLoad() // Wait for JS to load content

    content := page.MustElement("body").MustText()
    emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
    matches := emailRegex.FindAllString(content, -1)

    for _, file
     := range matches {
        fmt.Println("Found Email:", email)
    }
}
