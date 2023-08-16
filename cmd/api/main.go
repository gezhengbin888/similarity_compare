package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"similarity_compare/utils"
)

type RequestData struct {
	Item1 string `json:"item1"`
	Item2 string `json:"item2"`
}

/**
 * @author AronGe
 * @date 8/15/23  11:05
 * @description
 */
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var requestData RequestData
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Received data: item1 - %s, item2 - %s", requestData.Item1, requestData.Item2)
		response := utils.GetSimilarity(requestData.Item1, requestData.Item2)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/calc", handleRequest)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
