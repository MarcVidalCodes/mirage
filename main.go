package main 

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"encoding/json"
)

type WebHookPayload struct{
	Action string		`json:"action"`		//open, closed, reopened
	Number int			`json:"number"`		//pr number
	PullRequest struct{
		Head struct{
			Ref string	`json:"ref"`		//branch name
			SHA string	`json:"sha"`		//commit hash
		}`json:"head"`
	}`json:"pull_request"`
	Repository struct{
		Name string 	`json:"name"`		// repo name
	}`json:"repository"`
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func testHealth(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "OK")
}

func handleWebhook(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil{
		http.Error(w, "Error reading body", http.StatusInternalServerError)	
		return
	}
	defer r.Body.Close()

	//Parse
	var payload WebHookPayload
	
	err = json.Unmarshal(body, &payload)
	if err != nil{ 
		http.Error(w, "Could not parse JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Fprintf(w, "[MIRAGE] Received PR #%d on Branch %s: \nAction: %s\n",
		payload.Number,
		payload.PullRequest.Head.Ref,
		payload.Action,
	)

	w.WriteHeader(http.StatusOK)
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", testHealth)
	http.HandleFunc("/webhook", handleWebhook)

	fmt.Println("Server listening on port 8080...")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}