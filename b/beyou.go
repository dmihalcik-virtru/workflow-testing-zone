package b

import "fmt"

func Yourself(s string) string {
	message := fmt.Sprintf("Nice to see you, [%v]. Welcome from b@0.0.1!", s)
	return message
}
