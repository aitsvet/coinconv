package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aitsvet/coinconv/internal/app"
)

func main() {
	a := app.Configure(
		os.Getenv("CMC_PRO_API_URL"),
		os.Getenv("CMC_PRO_API_KEY"),
	)
	result, err := a.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
