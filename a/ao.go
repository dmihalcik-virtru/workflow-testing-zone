package a

import "fmt"

func Heyo(name string) string {
	message := fmt.Sprintf("Hey-o, %v. Welcome from a@v0.1.0!", name)
	return message
}
