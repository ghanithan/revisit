package revisit

import (
	"fmt"
)

func Revisit(name string) string {
	greetings := fmt.Sprintf("hello %s", name)
	fmt.Println(greetings)
	return greetings

}
