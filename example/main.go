package main

import (
	"fmt"
	"github.com/kogonia/maxmind"
	"log"
	"os"
)

func main() {
	if err := maxmind.Init(0); err != nil {
		log.Fatal(err)
	}
	org, err := maxmind.GetByIP("8.8.8.8")
	if err != nil {
		log.Printf("ERROR %v", err)
		os.Exit(1)
	}
	fmt.Println(org.Json())
}
