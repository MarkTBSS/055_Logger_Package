package utils

import (
	"encoding/json"
	"fmt"
)

func Debug(data any) {
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Printf("error marshalling data: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

func Output(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshalling data: %v\n", err)
		return nil
	}
	return bytes
}
