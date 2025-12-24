package createperson

import "fmt"

type Person struct {
	name    string
	address string
	phone   string
}

func NewPerson(name, address, phone string) Person {
	return Person{
		name:    name,
		address: address,
		phone:   phone,
	}
}

func (p Person) Print() string {
	return fmt.Sprintf("Name: %s\nAddres: %s\nPhone: %s", p.name, p.address, p.phone)
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello %s", p.name)
}

func (p *Person) SetName(newName string) {
	p.name = newName
}
