package main

import (
	"bufio"
	"fmt"
	"github.com/TheForgotten69/goinsta/v2"
	"github.com/dhowden/raspicam"
	"log"
	"os"
	"time"
)

func main() {
	dt := time.Now()

	var filename string
	filename = fmt.Sprintf("%s.jpg", dt.Format("2006-01-02T15:04:05-0700"))

	if err := takePicture(filename); err != nil {
		log.Fatalln(err)
	}

	readableDt := dt.Format("2006-01-02 15:04:05")

	var message string
	message = fmt.Sprintf(`%s #im_bot #golang #봇 #바질키우기 #바질
	`, readableDt)

	uploadToInsta(filename, message)

	// TODO: notify to slack
}

func uploadToInsta(imagePath, message string) error {
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Oh, %s not exists.", imagePath)
		return err
	}

	// TODO: 인스타 인자 정보를 가변 인자로 받을 수 있도록 수정이 필요함
	var insta *goinsta.Instagram
	insta, err := goinsta.Import("./.goinsta");
	if err != nil {
		log.Fatalln(err)
		return err
	}

	file, err := os.Open(imagePath)
	defer file.Close()

	buffer := bufio.NewReader(file)
	item, err := insta.UploadPhoto(buffer, message, 0, 0)
	if err != nil {
		log.Fatal("Error on Upload photo", err)
		return err
	}
	log.Println("Upload Complete. ID : ", item.ID)
	return nil
}

func takePicture(filename string) error {
	if _, err := os.Stat(filename); os.IsExist(err) {
		log.Fatalf("Oh, %s already exists.", filename)
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	s := raspicam.NewStill()
	errCh := make(chan error)
	go func() {
		for x := range errCh {
			log.Fatal(errCh, x)
		}
	}()
	log.Println("Capturing Image")
	raspicam.Capture(s, f, errCh)

	if err := <-errCh; err != nil {
		log.Fatalln("Oh, Sorry. Capture Failure. :( ", err)
		return err
	}

	return nil
}
