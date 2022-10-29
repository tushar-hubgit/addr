package main

import (
	"addr/functions"
	"addr/logic"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	HOST = "0.0.0.0"
	// HOST = "127.0.0.1"
)

func main() {
	PORT := os.Getenv("PORT")
	log.Printf("Server started at http://%s:%s", HOST, PORT)

	mux := http.NewServeMux()

	appLogic := logic.AppLogic{
		AppFunctions: functions.AppFunctions{},
	}

	mux.HandleFunc("/", appLogic.GetIP)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), mux))
}
