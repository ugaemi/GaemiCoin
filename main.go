package main

import (
	"encoding/json"
	"fmt"
	"github.com/ugaemi/gaemicoin/utils"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL string
	Method string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
