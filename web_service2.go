package main

import (
    "fmt"
    "net/http"
)

// Handler untuk endpoint pertama (/hello)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "FABIANO SALMAEN GANTENGG POLL!!")
}

// Handler untuk endpoint kedua (/info)
func infoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "MANYALA ABANGKUHHH")
}

func main() {
    // Registrasi handler untuk masing-masing endpoint
    http.HandleFunc("/halow", helloHandler)
    http.HandleFunc("/infong", infoHandler)

    // Mulai server web pada port 8080
    fmt.Println("Server berjalan di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

