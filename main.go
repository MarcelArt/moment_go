package main

import (
	"log"

	"github.com/MarcelArt/moment_go/pkg/moment"
)

func main() {
	date, err := moment.New("31/08/2024", "DD/MM/YYYY")
	if err != nil {
		panic(err.Error())
	}

	log.Println(date)
	log.Println(date.Format("YYYY-MM-DD"))
}
