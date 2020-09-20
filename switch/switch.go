package main

import (
	"fmt"
	"runtime"
)

func main() {
	foodType("baby")
	foodType("child")
	foodType("eat")
	os()

}
func foodType(human string) {
	switch human {
	case "baby":
		fmt.Println("liquid food")
	case "child":
		fmt.Println("semi solid")
	default:
		fmt.Println("solid food")
	}

}
func os() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macos")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("no")

	}
}
