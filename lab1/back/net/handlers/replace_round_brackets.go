package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type SortWordsRequest struct {
	Message string `json:"message"`
}

func ReplaceRoundBrackets(w http.ResponseWriter, r *http.Request) {
	var req SortWordsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error parsing the request:", err)
		http.Error(w, "Error parsing the request", http.StatusBadRequest)
		return
	}

	words := strings.Split(req.Message, " ")
	sort.Strings(words)

	var response = SortWordsRequest{
		Message: strings.Join(words, " "),
	}

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling the response:", err)
		http.Error(w, "Error marshalling the response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error marshalling the response:", err)
		http.Error(w, "Error marshalling the response", http.StatusBadGateway)
		return
	}
}
