package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "saudi aramco"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)
}
