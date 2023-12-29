package main

import (
	"fmt"

	"github.com/ptisma/jenkins-test/greeting"
)

func main() {
	greetingMessage := greeting.GenerateGreeting()
	fmt.Println(greetingMessage)
}
