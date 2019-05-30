package awesomeness

import (
	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
)

type AwesomePerson struct {
	name  string
	title string
}

func (a *AwesomePerson) init() {
	a.name = randomdata.SillyName()
	a.title = randomdata.Title(2)
}

func (a *AwesomePerson) String() string {
	return fmt.Sprintf("%s %s", a.title, a.name)
}
func CreateAwesomePerson() (a *AwesomePerson) {
	a = &AwesomePerson{}
	a.init()
	return
}
