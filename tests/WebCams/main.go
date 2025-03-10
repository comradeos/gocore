package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"time"

	"gocv.io/x/gocv"
)

var (
	cameraIndex = 0
	frameRate   = 20
)

func streamVideo(w http.ResponseWriter, r *http.Request) {

	webcam, err := gocv.OpenVideoCapture(cameraIndex)
	if err != nil {
		http.Error(w, "Error opening video capture", http.StatusInternalServerError)
		return
	}
	defer webcam.Close()


	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			continue
		}

		imgData, err := img.ToImage()
		if err != nil {
			continue
		}

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, imgData, nil); err != nil {
			continue
		}

		fmt.Fprintf(w, "--frame\r\nContent-Type: image/jpeg\r\n\r\n")
		w.Write(buf.Bytes())
		fmt.Fprintf(w, "\r\n")

		time.Sleep(time.Second / time.Duration(frameRate))
	}
}

func main() {
	http.HandleFunc("/video", streamVideo)

	fmt.Println("🚀 Server started at http://localhost:8000/video")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
