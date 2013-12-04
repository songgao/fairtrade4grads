package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Signature struct {
	Name       string `json:"name"`
	AUUsername string `json:"au_username"`
}

var fDev bool
var fLaddr string

func main() {
	flag.Parse()
	init_DB()
	fmt.Println(http.ListenAndServe(fLaddr, buildMux()))
}

func buildMux() *http.ServeMux {
	frontend, err := getFrontendPath()
	if err != nil {
		panic(fmt.Sprintf("Error getting frontend path using go command: %s\n", err))
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/sign", handleSign)
	mux.HandleFunc("/api/count", handleCount)
	mux.Handle("/", http.FileServer(http.Dir(frontend)))
	return mux
}

func handleSign(rsp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rsp.WriteHeader(http.StatusBadRequest)
	} else {
		var sig *Signature
		json.NewDecoder(req.Body).Decode(&sig)
		newSignature(sig)
		rsp.WriteHeader(http.StatusOK)
	}
}

func handleCount(rsp http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		rsp.WriteHeader(http.StatusBadRequest)
	} else {
		count, err := getCount()
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
		} else {
			io.WriteString(rsp, strconv.Itoa(count))
		}
	}
}
