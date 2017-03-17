package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
)

const (
	CRLF string = "\r\n"
)

type Controller struct {
}

func (controller Controller) Homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func (controller Controller) About(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	mux := http.NewServeMux()
	var controller Controller
	mux.HandleFunc("/", controller.Homepage)
	mux.HandleFunc("/about", controller.About)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
