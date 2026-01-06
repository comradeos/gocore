package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	webview "github.com/webview/webview_go"
)

var webFS embed.FS

func main() {
	sub, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(http.ListenAndServe("127.0.0.1:13000", http.FileServer(http.FS(sub))))
	}()

	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("")
	w.SetSize(1920, 1200, webview.HintNone)
	w.Navigate("http://127.0.0.1:13000")
	w.Run()
}
