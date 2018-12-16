package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

//colors
var (
	green         = color.New(color.FgHiGreen)
	boldcyan      = color.New(color.FgCyan, color.Bold)
	boldred       = color.New(color.FgRed, color.Bold)
	boldblue      = color.New(color.FgHiBlue, color.Bold)
	boldyellow    = color.New(color.FgYellow, color.Bold)
	boldwhite     = color.New(color.FgHiWhite, color.Bold)
	boldgreen     = color.New(color.FgHiGreen, color.Bold)
	italicmagenta = color.New(color.FgHiMagenta)
	italicblue    = color.New(color.FgBlue, color.Italic)
)

func main() {
	doc, err := goquery.NewDocument("https://spotifycharts.com/regional/jp/daily/latest")
	if err != nil {
		fmt.Println(err)
	}
	//search
	doc.Find(".chart-table tbody tr").Each(func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}
		position := s.Find(".chart-table-position").Text()
		song := s.Find(".chart-table-track strong").Text()
		artist := s.Find(".chart-table-track span").Text()
		stream := s.Find("td.chart-table-streams").Text()

		if i == 0 {
			boldyellow.Printf("%s "+"%s"+" %s"+" %s\n", position, song, artist, stream)
		} else if i == 1 {
			boldgreen.Printf("%s "+"%s"+" %s"+" %s\n", position, song, artist, stream)
		} else if i == 2 {
			boldcyan.Printf("%s "+"%s"+" %s"+" %s\n", position, song, artist, stream)
		} else {
			fmt.Printf("%s "+"%s"+" %s"+" %s\n", position, song, artist, stream)
		}
	})
}
