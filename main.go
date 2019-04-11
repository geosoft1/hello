package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// https://gist.github.com/denji/12b3a568f092ab951456
// openssl genrsa -out server.key 2048
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
// curl --verbose http://localhost:8080/hello
// curl --verbose --insecure  https://localhost:8090/hello
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	folder, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!\n")
	})
	go func() {
		// allow you to use self signed certificates
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		if err := http.ListenAndServeTLS(":8090", filepath.Join(folder, "server.crt"), filepath.Join(folder, "server.key"), nil); err != nil {
			log.Println(err)
		}
	}()
	http.ListenAndServe(":8080", nil)
}
