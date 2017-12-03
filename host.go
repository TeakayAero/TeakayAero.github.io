package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Version struct {
	Version string `json:version`
}

func versionRequest(w http.ResponseWriter, r *http.Request){
        v := Version{Version: "1.0.0"}
	fmt.Println(v)
        fmt.Println("version request received")
	json.NewEncoder(w).Encode(v)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/version", versionRequest)
	http.ListenAndServe(":3000", nil)
}
