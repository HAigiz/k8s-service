package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	appName  string
	bindPort string
)

func init() {
	var ok bool

	if appName, ok = os.LookupEnv("ROOT_APP_NAME"); !ok {
		appName = "ROOT"
	}

	if bindPort, ok = os.LookupEnv("ROOT_BIND_PORT"); !ok {
		bindPort = "8080"
	}
}

func main() {
	log.Printf("Start %s Service on port :%s\n", appName, bindPort)

	http.HandleFunc("/", rootHandler())

	if err := http.ListenAndServe(fmt.Sprintf(":%s", bindPort), nil); err != nil {
		log.Fatalln(err)
	}
}

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t %s\t %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		if _, err := w.Write([]byte("Sirius kube-itog project")); err != nil {
			log.Println(err)
		}
	}
}
