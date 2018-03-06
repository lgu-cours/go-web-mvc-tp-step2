package main

import (
	"log"
	"net/http"
	
	"webapp2/utils"
)

const webDir  = "www"
const webPort = ":80"

func serveStaticFile(w http.ResponseWriter, r *http.Request) {		
	s := webDir + r.URL.Path
	log.Print("Static file requested. URL path '" + r.URL.Path + "' --> server file '" + s + "'" )
	http.ServeFile(w, r, s)
}

func main() {
	log.Print("Starting... ")

	utils.PrintEnv()

	log.Print("Setting static file handler... ")
	
	// Set handler to serve static files
	http.HandleFunc("/", serveStaticFile)			

	log.Print("Setting application handlers... ")
	initControllers()

	// Launch the http server
	log.Print("Launching http server (port=" + webPort + ") ... ")
	log.Fatal(http.ListenAndServe(webPort, nil))

}