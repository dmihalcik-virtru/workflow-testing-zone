package a

import "fmt"

func Heyo(name string) string {
	message := fmt.Sprintf("Hey-o, %v. Welcome from a@v0.0.11!", name)
	return message
}
