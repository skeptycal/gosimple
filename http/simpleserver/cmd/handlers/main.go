package main

import (
	"fmt"
	"net/http"

	"github.com/skeptycal/gosimple/http/simpleserver"
)

const addr = "localhost:8088"

var h = simpleserver.New("sample")

func nilHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sample")
}

func sampleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func sampleHeaders(w http.ResponseWriter, req *http.Request) {
	req.Header.Clone().Write(w)
}

func sampleAbout(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "About")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "request URL: %v\n", req.URL)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, h.String())
}

func getNilHandler(i int) *simpleserver.Handler {
	return simpleserver.NewHandler(
		fmt.Sprintf("/sample/%d", i),
		func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Sample Page Number %d", i) },
	)
}

var sampleHandlerList = []*simpleserver.Handler{
	simpleserver.NewHandler("/", sampleRoot),
	simpleserver.NewHandler("/headers", sampleHeaders),
	simpleserver.NewHandler("/about", sampleAbout),
}

func main() {
	const maxNilHandlers = 20
	start := len(sampleHandlerList)
	for i := 0; i < maxNilHandlers; i++ {
		sampleHandlerList = append(sampleHandlerList, getNilHandler(start+i))
	}
	for i, hand := range sampleHandlerList {
		h.Insert(i, hand)
	}
	h.Register()
	http.ListenAndServe(addr, nil)
}
