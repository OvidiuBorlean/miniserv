package main

import (
        "fmt"
        "net/http"
        "io"
        "errors"
        "os"
)


func getHeader(w http.ResponseWriter, r *http.Request) {
     ua := r.Header.Get("")
     //fmt.Fprintf(w, "User agent: %s\n", ua)
     for k, v := range r.Header {
            fmt.Println(k,v)
           fmt.Fprintf(w, "%q, %q\n", k, v)
      }
     fmt.Printf(ua)
}
func getIP(w http.ResponseWriter, r *http.Request) {
     //fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
      fmt.Println(r.RemoteAddr)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
     fmt.Printf("got / request\n")
     io.WriteString(w, "This is request\n")
}

func getHealthz(w http.ResponseWriter, r *http.Request) {
     fmt.Printf(" got /healthz request ")
     io.WriteString(w, "This is a Health")
}

func main() {
     fmt.Println("Starting Server")
     http.HandleFunc("/", getRoot)
     http.HandleFunc("/header", getHeader)
     http.HandleFunc("/healthz", getHealthz)
     http.HandleFunc("/ip", getIP)
     err :=  http.ListenAndServe(":3333", nil)
     if errors.Is(err, http.ErrServerClosed) {
                fmt.Printf("server closed\n")
        } else if err != nil {
                fmt.Printf("error starting server: %s\n", err)
                os.Exit(1)
        }

}
