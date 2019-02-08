package main

import (
	"log"
    "fmt"
	"net/http"
    "os"
    "io"
//	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // routing hack :)
        http.ServeFile(w, r, r.URL.Path[1:]+".html")
    }


    if r.Method == "POST" {
        // Parse mutipart maybe not needed because it's done in FormFile?
        // but also I'm still unsure what the args for this should be (MaxMemory)
        err := r.ParseMultipartForm(32)
        // There's probably a better way to deal with these errors
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }   

        // Get file information
        file, header, err := r.FormFile("image")
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        defer file.Close()


        fmt.Fprintf(w, "%v", header.Header)        
        // var file_name string

        // Open write-only file in images folder; Create new file if it doesn't exist
        f, err := os.OpenFile("./images/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }

        // file_name = header.Filename
        defer f.Close()

        io.Copy(f, file) 

    }

    

}

func main() {
	http.HandleFunc("/", indexHandler)
    http.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
