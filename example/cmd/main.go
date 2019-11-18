package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"html/template"
)
type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello Modified OpenShift!"
	}

	fmt.Fprintln(w, response)
	fmt.Println("Servicing request.")
}

//func listenAndServe(port string) {
//	fmt.Printf("serving on %s\n", port)
//	err := http.ListenAndServe(":"+port, nil)
//	if err != nil {
//		panic("ListenAndServe: " + err.Error())
//	}
//}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	r := mux.NewRouter()
	//http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	//go listenAndServe(port)

	//port = os.Getenv("SECOND_PORT")
	//if len(port) == 0 {
	//	port = "8888"
	//}
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })
	r.HandleFunc("/",helloHandler)
	r.HandleFunc("/forms", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })
      http.ListenAndServe(":8080", r)
	//go listenAndServe(port)

	//select {}
}

