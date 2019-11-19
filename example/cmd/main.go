package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"html/template"
	"database/sql"
	"log"
        "time"
        _ "github.com/go-sql-driver/mysql"
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
	//db related code
	connectionString=user + ":" + password + "@(" + ip + ":" + portno + ")/" + dbname + "?parseTime=true"
	db, err := sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    { // Create a new table
        query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

        if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }
    }

    { // Insert a new user
        username := "johndoe"
        password := "secret"
        createdAt := time.Now()

        result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
        if err != nil {
            log.Fatal(err)
        }

        id, err := result.LastInsertId()
        fmt.Println(id)
    }

    { // Query a single user
        var (
            id        int
            username  string
            password  string
            createdAt time.Time
        )

        query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
        if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
            log.Fatal(err)
        }

        fmt.Println(id, username, password, createdAt)
    }

    { // Query all users
        type user struct {
            id        int
            username  string
            password  string
            createdAt time.Time
        }

        rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var users []user
        for rows.Next() {
            var u user

            err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
            if err != nil {
                log.Fatal(err)
            }
            users = append(users, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%#v", users)
    }

    {
        _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
        if err != nil {
            log.Fatal(err)
        }
    }
	
	
	
	tmpl := template.Must(template.ParseFiles("/opt/app-root/src/github.com/talat-shaheen/openshift-golang-template/example/cmd/forms.html"))
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

