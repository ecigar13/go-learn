package main

import (
	"encoding/json"
	"fmt"
	"os"
)



func main() {
	data := readFile("data.json")
	
	//fmt.Printf("%s", data)
	var storedJson interface{}
	json.Unmarshal([]byte(data), &storedJson)
	//fmt.Printf("%s", storedJson)
	searchResult, err := dfsKey(storedJson, "longitude")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(searchResult)

	

}
