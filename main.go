package main

import (
	_ "embed"

	"github.com/webview/webview"
)

//go:embed static/index.html
var indexHtml string

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("1431 Game Database")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(`data:text/html,` + indexHtml)
	w.Run()
}
