package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
| METHOD | REQUEST                               | RESPONSE                      |
|--------|---------------------------------------|-------------------------------|
| GET    | `/name/{PARAM}`                       | body: `Hello, PARAM!`         |
| GET    | `/bad`                                | Status: `500`                 |
| POST   | `/data` + Body `PARAM`                | body: `I got message:\nPARAM` |
| POST   | `/headers`+ Headers{"a":"2", "b":"3"} | Header `"a+b": "5"`           |
*/

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", nameRequestHandler)
	router.HandleFunc("/bad", badRequestHandler)
	router.HandleFunc("/data", dataRequestHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", headersRequestHandler).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func nameRequestHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataRequestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "I got message:\n%s", string(data))
}

func headersRequestHandler(w http.ResponseWriter, r *http.Request) {
	aParam := r.Header.Get("a")
	bParam := r.Header.Get("b")
	a, err := strconv.Atoi(aParam)
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(bParam)
	if err != nil {
		log.Fatal(err)
	}
	result := a + b
	w.Header().Set("a+b", strconv.Itoa(result))
}
