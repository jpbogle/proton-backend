package main

import (
	"log"
	"net/http"
	"os"
	"proton/controllers"
	h "proton/handlers"
    "proton/utils"
	"sync"
    // "github.com/fvbock/endless"
)

var IS_DEBUG = false

func main() {
	defer controllers.Close()
	log.SetFlags(0)

	parseArgs()

	var wg sync.WaitGroup
	wg.Add(1)

    go http.ListenAndServe(":8080", h.RootHandler)
	// go endless.ListenAndServe(":8080", h.RootHandler)
	log.Print("\n> Now listening on localhost:8080\n\n")
	wg.Wait()
}

func parseArgs() {
	for _, v := range os.Args {
		if v == "--debug" {
			utils.SetDebug(true)
		}
		if v == "--drop" && IS_DEBUG {
			controllers.DropTables()
			log.Println("Successfully dropped tables")
			os.Exit(0)
		}
	}
}


func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
    target := "https://" + req.Host + req.URL.Path
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    log.Printf("redirect to: %s", target)
    http.Redirect(w, req, target, http.StatusMovedPermanently)
}
