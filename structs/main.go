package main

import "fmt"

type ContactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func (p person) data() {
	fmt.Println("result is ", p.firstName)

}
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	//name := person{firstName: "Mukul", lastName: "Mantosh"}
	//fmt.Println(name)
	var name person
	name.firstName = "Mukul"
	name.lastName = "Mantosh"
	name.contact.email = "mukul@gmail.com"
	name.contact.zip = 123
	fmt.Println(name)
	fmt.Printf("%+v", name)
	name.data()

	// Updating using Pointer
	//dataPointer := &name
	//dataPointer.updateName("MM")

	// Pointer Shortcut
	name.updateName("MM")

	fmt.Println(name)
}
