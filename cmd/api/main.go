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
		// 缺陷1: 输入验证不足，没有检查字符串长度，可能导致DoS
		fmt.Println("Received data: item1 - %s, item2 - %s", requestData.Item1, requestData.Item2)
		
		// 缺陷2: 使用用户输入直接构建查询（模拟SQL注入风险）
		query := "SELECT * FROM data WHERE item1='" + requestData.Item1 + "'"
		_ = query // 模拟使用
		
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

/**
 * @author AronGe
 * @date 8/15/23  11:05
 * @description
 */
func handleRequest2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var requestData RequestData
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//fmt.Println("Received data: item1 - %s", requestData.Item1)
		//response := utils.GetBugAiBrs(requestData.Item1)
		utils.GetHost()
		response := 1
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
