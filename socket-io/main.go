package main

import (
	"github.com/joho/godotenv"
	"github.com/letsgo-framework/examples/socket-io/database"
	letslog "github.com/letsgo-framework/examples/socket-io/log"
	"github.com/letsgo-framework/examples/socket-io/routes"
	"os"
)

func main() {

	// Load env
	err := godotenv.Load()

	// Setup log writing
	letslog.InitLogFuncs()

	if err != nil {
		letslog.Error("Error loading .env file")
	} else {
		letslog.Info("env loaded")
	}

	database.Connect()

	srv := routes.PaveRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if os.Getenv("SERVE_TLS") == "true" {
		srv.RunTLS(port,os.Getenv("CERTIFICATE_LOCATION"),"KEY_FILE_LOCATION")
	} else {
		srv.Run(port)
	}

}
