package main

import (
	"fmt"
	"os"

	"github.com/Mr-Bruno/goencryption/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Too few arguments.")
		return
	}
	id := []byte(os.Args[1])
	payload := []byte(os.Args[2])

	mc := client.MyClient{}
	key, _ := mc.Store(id, payload)
	originalData, _ := mc.Retrieve(id, key)

	fmt.Println(string(originalData))
}
