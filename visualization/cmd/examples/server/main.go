package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/skeptycal/gosimple/cli/errorlogger"
)

var log = errorlogger.New()

func init() {
	log.SetLogLevel("TRACE")
	log.Debug("Logging started...")
}

/// Reference: https://stackoverflow.com/questions/53762936/golang-to-create-a-main-layout-webpage

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	log.Debug(strings.ToTitle(tmpl))
	err := templ.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
}

//---------------------------------------Page Handlers----------------------------------//
//Handler for homepage
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "index", &Page{Title: "Home"})
}

//Handler for about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "index", &Page{Title: "About"})
}

//Handler for test Page
func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Test")
		testT, _ := template.ParseFiles("static/test.html")
		testT.Execute(w, nil)
	}
}

func main() {

	//--------------------------------------Routers-------------------------------------//
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/test", testHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//---------------------------------------------------------------------------------//

	//log to file
	// f, err := os.OpenFile("serverlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("Error opening file: %v", err)
	// }
	// defer f.Close()
	// logger := log.New(f, "Logged : ", log.LstdFlags)
	// log.SetOutput(f)

	//start server
	log.Println("Starting server on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}

//-------------------------------------------------------------------------------------//
//Compile templates on start
var templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			log.Info(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	return t
}()

type Page struct {
	Title string
}
