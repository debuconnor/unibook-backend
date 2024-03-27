package unibookBackend

import (
	"strconv"
	"sync"

	"github.com/debuconnor/dbcore"
	sc "github.com/debuconnor/sc-scraper"
)

func runScScraper(userid, platform string) {
	db := initDb()
	db.ConnectMysql()
	defer db.DisconnectMysql()

	dml := dbcore.NewDml()
	dml.SelectAll()
	dml.From(SCHEMA_NAME_SYSTEM)
	dml.Where("", "config_key", "=", "sc_scrape_page_count")
	dml.Where("OR", "config_key", "=", "sc_scrape_url")
	dml.Where("OR", "config_key", "=", "sc_scrape_selector")
	queryResult := dml.Execute(db.GetDb())

	session := initSc(db, userid, platform)
	scrapeUrl := ""
	scrapeSelector := ""
	pages := 0

	for _, row := range queryResult {
		key := row["config_key"]
		value := row["config_value"]

		if key == "sc_scrape_url" {
			scrapeUrl = value
		} else if key == "sc_scrape_selector" {
			scrapeSelector = value
		} else if key == "sc_scrape_page_count" {
			pages, _ = strconv.Atoi(value)
		}
	}

	var wait sync.WaitGroup
	wait.Add(pages)

	for i := 1; i <= pages; i++ {
		go func() {
			defer wait.Done()
			page := strconv.Itoa(i)
			sd := sc.NewScrapeData(scrapeUrl+page, scrapeSelector)
			sd.ScrapeTarget(session)
			sd.SaveReservation(db.GetDb())
		}()
	}
	wait.Wait()
}
