package a

import "fmt"

func Heyo(name string) string {
	message := fmt.Sprintf("Hey-o, %v. Welcome from a@v0.0.12!", name)
	return message
}
