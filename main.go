package main

import (
	"fmt"
	"github.com/ugaemi/gaemicoin/blockchain"
	"log"
	"net/http"
	"text/template"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home (rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
