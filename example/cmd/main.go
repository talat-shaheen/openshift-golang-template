package main
import (
"fmt";
"io";
"net/http")

func writeRootData(w http.ResponseWriter, r *http.Request){
io.WriteString(w, "hi there!\n")
}
func writeData(w http.ResponseWriter, r *http.Request){
io.WriteString(w, "hello\n")
}
func ping(w http.ResponseWriter, r *http.Request){
w.Write([]byte("pong\n"))
}

func main(){
//fmt.Println("1111111111111111")
//http.Handle("/",http.FileServer(http.Dir("/root/www")))
http.HandleFunc("/",writeRootData)
http.HandleFunc("/hello",writeData)
http.HandleFunc("/ping",ping)
http.ListenAndServe(":8000",nil)
}

