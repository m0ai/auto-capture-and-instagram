package main

import (
	"github.com/dhowden/raspicam"
	"log"
	"os"
)


func main() {
	f, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := raspicam.NewStill()
	errCh := make(chan error )
	go func() {
		for x := range errCh {
			log.Fatal(errCh, x)
		}
	}()

	log.Println("Capturing Image...")
	raspicam.Capture(s, f, errCh)
}
