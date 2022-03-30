package p

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CPULoadHashing(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 2048)
	rand.Read(b)

	for i := 0; i < 100; i++ {
		a := sha512.Sum512(b)
		fmt.Fprintf(ioutil.Discard, "hash: %x\n", a)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cpu load simulation done\n"))
}
