package main

import (
	"encoding/json"
	"fmt"
)

func stringify(data any) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Cannot stringify data", jsonData)
	}
	return string(jsonData)
}
