package main

import (
    "html/template"
        "net/http"
        "fmt"
        )

        type ContactDetails struct {
            Email   string
                Subject string
                    Message string
                    }

 func main(){
 fmt.Println("server has started")
 tmpl := template.Must(template.ParseFiles("forms.html"))

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodPost {
      tmpl.Execute(w, nil)
      return
 }

 details := ContactDetails{
     Email:   r.FormValue("email"),
     Subject: r.FormValue("subject"),
     Message: r.FormValue("message"),
 }
                                                                                                                                        _ = details
tmpl.Execute(w, struct{ Success bool }{true})
                                                                                                                                                    })

  http.ListenAndServe(":80", nil)
                                                                                                                                                        }
