package main

import (
	"os"
	"strings"

	"github.com/hellominchan/Indeed-Crawler-Website/crawler"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleCrawler(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(crawler.CleanString(c.FormValue("term")))
	crawler.Crawler(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/crawler", handleCrawler)
	e.Logger.Fatal(e.Start(":1323"))
}
