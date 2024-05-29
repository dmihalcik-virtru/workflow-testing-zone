package main

import (
	"fmt"

	"github.com/dmihalcik-virtru/workflow-testing-zone/a"
	"github.com/dmihalcik-virtru/workflow-testing-zone/b"
)

func main() {
	message := b.Yourself(a.Heyo("Gladys"))
	fmt.Println(message)
}
