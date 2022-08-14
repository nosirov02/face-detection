package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("unable to init web cam: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("face detection in Go")
	defer window.Close()

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from web cam")
			continue
		}

		window.IMShow(img)
		window.WaitKey(500)
	}
}
