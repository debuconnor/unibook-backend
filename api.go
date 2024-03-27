package unibookBackend

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Run() {
	// Start DB Connection Check
	log.Println("Checking DB Connection...")
	db := initDb()
	db.ConnectMysql()

	if errCode := checkTables(db); errCode != ERROR_NOT_FOUND {
		Error(errCode)
		return
	}

	db.DisconnectMysql()
	log.Println("DB Connection Check Complete")
	// End DB Connection Check

	r := router.New()
	log.Println("Starting API Server...")
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Welcome to Unibook Backend API Server")
		log.Println("Request: /")
	})
	r.GET("/scscraper", runScScraperHandler)
	r.GET("/nvrscraper", runNvrScraperHandler)

	log.Println("API Server Started. url: http://localhost:9090")
	if err := fasthttp.ListenAndServe(":9090", r.Handler); err != nil {
		SaveLog("error.log", err)
		return
	}
}

func runScScraperHandler(ctx *fasthttp.RequestCtx) {
	log.Println("Request: /scscraper")
	userid, platform := getUserInfo(ctx)

	if checkUserInfo(userid, platform) {
		fmt.Fprintf(ctx, "Please wait... Running SC Scraper for %s, Platform Id: %s.\n\n", userid, platform)
		runScScraper(userid, platform)
		fmt.Fprintf(ctx, "Scraping completed.")
	}
}

func runNvrScraperHandler(ctx *fasthttp.RequestCtx) {
	log.Println("Request: /nvrscraper")
	userid, platform := getUserInfo(ctx)

	fmt.Fprintf(ctx, "Running NVR Scraper for %s, %s", userid, platform)
}

func getUserInfo(ctx *fasthttp.RequestCtx) (userid, platform string) {
	queryArgs := ctx.QueryArgs()
	userid = string(queryArgs.Peek("userid"))
	platform = string(queryArgs.Peek("platform"))
	Log("userid: " + userid + ", platform: " + platform)

	return
}

func checkUserInfo(userid, platform string) bool {
	if userid == "" || platform == "" {
		return false
	}

	return true
}
