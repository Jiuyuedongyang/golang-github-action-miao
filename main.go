package main

import (
	"GitHub_cicd_test/pets"
	"fmt"
)
var (
	password="123456"
)
func main() {
	fmt.Printf("my password %s",password)
	saying := pets.Cat()
	fmt.Println(saying)
}
