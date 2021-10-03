package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/webview/webview"
)

//go:embed assets/*
var assets embed.FS

func app() {
	root, err := fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(root)))
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func main() {
	go app()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("1431 Game Database")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://127.0.0.1:8080/index.html")
	w.Run()
}
