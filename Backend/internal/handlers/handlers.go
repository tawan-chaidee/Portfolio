package handlers

import (
    "fmt"
    "net/http"
)

// Home handler function.
func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Homepage")
}


// Info handler function.
func Info(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Info page")
}
