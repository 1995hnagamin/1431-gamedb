package main

import (
	"embed"
	"io/fs"
	"net"
	"net/http"

	"github.com/webview/webview"
)

//go:embed assets/*
var assets embed.FS

func startServer() string {
	root, err := fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	hfs := http.FileServer(http.FS(root))
	hfsprefix := "/view/"
	mux.Handle(hfsprefix, http.StripPrefix(hfsprefix, hfs))

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go func() {
		panic(http.Serve(listener, mux))
	}()

	return listener.Addr().String()
}

func main() {
	addr := startServer()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("1431 Game Database")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://" + addr + "/view/")
	w.Run()
}
