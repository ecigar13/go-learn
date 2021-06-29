package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := readFile("data.json")
	var storedJson map[string](interface{})
	json.Unmarshal([]byte(data), &storedJson)
	fmt.Println(storedJson)
	json, err := json.MarshalIndent(storedJson, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
