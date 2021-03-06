package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type (

	MyResponse struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
)

func handlePost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %s\n", r.Form.Get("name"))
	fname := "default"
	if val := r.Form.Get("name"); val != ""{
		fname = val
	}

	w.Write([]byte(fmt.Sprintf("fname: %s\n", fname)))
}
func handleTest(w http.ResponseWriter, r *http.Request){

	fname := "Roy"
	if names, ok := r.URL.Query()["fname"]; ok {
		fname = names[0]
	}

	respObj := &MyResponse{FirstName:fname, LastName:"Pearl"}
	respSer, err := json.Marshal(respObj)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respSer)
}

func main(){

	http.HandleFunc("/test", handleTest)
	http.HandleFunc("/testPost", handlePost)

	http.ListenAndServe(":8080", nil)
}