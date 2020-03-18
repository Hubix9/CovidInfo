package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//DownloadWebsite function used to make GET request to selected url
func DownloadWebsite(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("got an error while processing request: %d %s", res.StatusCode, res.Status)
	}
	return res, err
}

//HandleScrape Function used to handle website data scraping and printing
func HandleScrape(res *http.Response, CountryToFind string) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var CountryRow *goquery.Selection
	QueryString := ":contains('" + CountryToFind + "')"
	doc.Find("td").Each(func(i int, selection *goquery.Selection) {
		selection.Filter(QueryString).Each(func(i int, selection *goquery.Selection) {
			CountryRow = selection.Parent()
		})
	})
	if CountryRow == nil {
		fmt.Println("Sorry couldn't find the country you are looking for")
		return
	}
	CountryParams := CountryRow.Children()
	Country := strings.TrimSpace(CountryParams.Eq(0).Text())
	fmt.Println("Country: " + Country)
	TotalCases := strings.TrimSpace(CountryParams.Eq(1).Text())
	fmt.Println("Total Cases: " + TotalCases)
	NewCases := strings.TrimSpace(CountryParams.Eq(2).Text())
	fmt.Println("New Cases: " + NewCases)
	TotalDeaths := strings.TrimSpace(CountryParams.Eq(3).Text())
	fmt.Println("Total Deaths: " + TotalDeaths)
	NewDeaths := strings.TrimSpace(CountryParams.Eq(4).Text())
	fmt.Println("New Deaths: " + NewDeaths)
	TotalRecovered := strings.TrimSpace(CountryParams.Eq(5).Text())
	fmt.Println("Total Recovered: " + TotalRecovered)
	ActiveCases := strings.TrimSpace(CountryParams.Eq(6).Text())
	fmt.Println("Active Cases: " + ActiveCases)
	SeriousCritical := strings.TrimSpace(CountryParams.Eq(7).Text())
	fmt.Println("Serious Critical: " + SeriousCritical)
	TotalCasesPerMil := strings.TrimSpace(CountryParams.Eq(8).Text())
	fmt.Println("Total Cases per 1 milion of population: " + TotalCasesPerMil)

}

//main Main function of script, running all the logic
func main() {
	res, _ := DownloadWebsite("https://www.worldometers.info/coronavirus/")
	inputArgs := os.Args[1:]
	HandleScrape(res, inputArgs[0])
}
