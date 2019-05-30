package main

import (
	"fmt"
	"github.com/eckhard-briedenhann/awesome/awesomeness"
)

func main() {

	var a = awesomeness.CreateAwesomePerson()
	fmt.Printf("Introducing: %s! \n", a.String())

}
