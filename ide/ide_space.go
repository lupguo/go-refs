package ide

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	ID   int
}

func (p *Person) Speak(lang string) (string, error) {
	fmt.Printf(p.Name)
	return "speak" + lang, nil
}

func (p *Person) GetUserName() string {
	return p.Name
}

func afunc() {
	p1 := &Person{}

	p1.Speak("hello/world")
	return
	float, err := strconv.ParseFloat("3.14", 64)
	parseInt, err := strconv.ParseInt("2", 10, 64)
}
