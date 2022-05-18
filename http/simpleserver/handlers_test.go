package simpleserver

import (
	"fmt"
	"net/http"
	"testing"
)

// func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))

func sampleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln()
}

func sampleHeaders(w http.ResponseWriter, req *http.Request) {
	req.Header.Clone().Write(w)
}

func sampleAbout(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "About")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "request URL: %v\n", req.URL)
}

// func sampleGraph(w http.ResponseWriter, req *http.Request) {
// 	bar = generateBarChart()
// 	_ = req
// 	bar.Render(w)
// }

var sampleHandlerList = []handler{
	{"/", sampleRoot},
	{"/headers", sampleHeaders},
	{"/about", sampleAbout},
	{"/headers", sampleHeaders},
}

func TestNew(t *testing.T) {
	var sampleHandlers = New("sample")

}
