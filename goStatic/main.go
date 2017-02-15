// This small program is just a small web server created in static mode
// in order to provide the smallest docker image possible

package main

import (
	"flag"
	"log"
	"strconv"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// Def of flags
	portPtr = flag.Int("p", 8043, "The listening port")
	pathPtr = flag.String("static", "/srv/http", "The path for the static files")
	HTTPPtr = flag.Bool("forceHTTP", false, "Forcing HTTP and not HTTPS")
)

func main() {

	flag.Parse()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Routes
	e.Static("/", *pathPtr)

	log.Println("Starting goStatic")

	port := ":" + strconv.FormatInt(int64(*portPtr), 10)

	// Start server with unsecure HTTP
	if *HTTPPtr {
		log.Println("Starting serving", *pathPtr, "on", *portPtr)

		// Start server
		e.Logger.Fatal(e.Start(port))
		// or with awesome TLS
	} else {
		// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
		// Cache certificates
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

		log.Println("Starting serving", *pathPtr, "on", *portPtr)
		e.Logger.Fatal(e.StartAutoTLS(port))
	}

}
