package main

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/load_test", cpuLoadHashing)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8080", loggedRouter)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World :-)"))
}

func cpuLoadHashing(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 2048)
	rand.Read(b)

	for i := 0; i < 100; i++ {
		a := sha512.Sum512(b)
		fmt.Fprintf(ioutil.Discard, "hash: %x\n", a)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cpu load simulation done\n"))
}
