package main

import (
	"GitHub_cicd_test/pets"
	"fmt"
)

func main() {
	var password="123456"

	fmt.Printf("my password %s",password)
	saying := pets.Cat()
	fmt.Println(saying)
}
