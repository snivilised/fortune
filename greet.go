package honour

import "fmt"

func Greet(recipient string) string {
	return fmt.Sprintf("greetings %q\n", recipient)
}

func Goodbye(recipient string) {
	fmt.Printf("bye %q\n", recipient)
}
