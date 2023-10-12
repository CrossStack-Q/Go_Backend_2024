package main

import (
	"fmt"

	decrypt "github.com/CrossStack-Q/Go-Lang_Package_1"
)

func main() {
	encryptedStr := "Kode"
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
