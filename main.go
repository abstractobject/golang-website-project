package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in contact with me, email at <a href=\"mailto:hexadecagon@gmail.com\">hexadecagon@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>Q: is this a test question?></li>
		<li>A: yes it is!></li>
	</ul>
`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		//handle page not found error
// 		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 		//fmt.Fprint(w, "<h1>INVALID URL</h>")
// 		w.WriteHeader(404)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		//handle page not found error
		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//fmt.Fprint(w, "<h1>INVALID URL</h>")
		w.WriteHeader(404)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	//print isnt necessary, just a dev helper tool to make sure your code is running at least okayish
	http.ListenAndServe(":3000", router)
}
