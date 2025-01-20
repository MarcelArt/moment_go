package main

import (
	"log"
	"os"

	"github.com/MarcelArt/moment_go/pkg/goment"
)

func main() {
	date, err := goment.New("31/08/2024", "DD/MM/YYYY")
	if err != nil {
		log.Panicln(err.Error())
		os.Exit(1)
	}

	log.Println(date)
	log.Println(date.Formats("YYYY-MM-DD"))
}
