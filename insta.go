package main

import (
	"bufio"
	"github.com/TheForgotten69/goinsta/v2"
	"log"
	"os"
)


func main () {
	//insta := goinsta.New("id", "password")
	//if err := insta.Login(); err != nil {
	//	log.Fatalln(err)
	//}
	//insta.Export("./.goinsta")

	insta, err := goinsta.Import("./.goinsta")
	if err != nil {
		log.Fatalln(err)
	}


	var filename string = "./test.jpg"

	file, err := os.Open(filename)
	buff := bufio.NewReader(file)
	_, err = insta.UploadPhoto(buff, "반갑다. 휴--먼 #i_am_bot", 0,0)

	if err != nil {
		log.Fatalln("Error at Upload Photo", err)
	}

	//if err != nil {
	//	return
	//}
}
