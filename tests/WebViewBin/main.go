package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	webview "github.com/webview/webview_go"
)

//go:embed web/*
var webFS embed.FS

func main() {
	// создаём fs только для папки web
	sub, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatal(err)
	}

	// HTTP сервер ИЗ БИНАРЯ
	go func() {
		log.Fatal(http.ListenAndServe("127.0.0.1:8080", http.FileServer(http.FS(sub))))
	}()

	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://127.0.0.1:8080")
	w.Run()
}
