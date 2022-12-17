package main

import "fmt"

type PhoneBook struct {
	Name   string
	Number int
	Pre    *PhoneBook
	Next   *PhoneBook
}

func newPhoneBook(name string, number int, pre *PhoneBook, next *PhoneBook) PhoneBook {
	return PhoneBook{
		Name:   name,
		Number: number,
		Pre:    pre,
		Next:   next,
	}
}

func main() {
	init := newPhoneBook("pjh", 0, nil, nil)
	one := newPhoneBook("pjh1", 1, &init, nil)
	init.Next = &one
	two := newPhoneBook("pjh1", 2, &one, nil)
	one.Next = &two
	three := newPhoneBook("pjh1", 3, &two, nil)
	two.Next = &three
	curPhone := &init
	for {
		fmt.Println(curPhone.Name)
		fmt.Println(curPhone.Number)
		if curPhone.Next == nil {
			break
		}
		curPhone = curPhone.Next
	}
}
