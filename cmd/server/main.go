package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":9001", http.FileServer(http.Dir("/Users/john.carnell/work/wasm_linter/assets"))))
}
