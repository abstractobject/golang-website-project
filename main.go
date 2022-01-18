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

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
	fmt.Fprintln(w, r.URL.RawPath)
	// switch r.URL.Path {
	// case "/":
	// 	homeHandler(w, r)
	// case "/contact":
	// 	contactHandler(w, r)
	// default:
	// 	//handle page not found error
	// 	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 	//fmt.Fprint(w, "<h1>INVALID URL</h>")
	// 	w.WriteHeader(404)
	// }
}

func main() {
	http.HandleFunc("/", pathHandler)
	//http.HandleFunc("/contact", contactHandler)
	//time.Sleep(3 * time.Second)
	//if you want to make sure your server is actually started
	fmt.Println("Starting the server on :3000...")
	//print isnt necessary, just a dev helper tool to make sure your code is running at least okayish
	http.ListenAndServe(":3000", nil)
}
