package dubble_linked_list_1

import "fmt"

type Array struct {
	PhoneBook []*PhoneBook
}

type PhoneBook struct {
	Name   string
	Number int
	Pre    *PhoneBook
	Next   *PhoneBook
}

var arr Array

func initArray() *Array {
	arr = Array{}
	return &arr
}

func (a *Array) AddBook(name string, number int, pre *PhoneBook, next *PhoneBook) {
	book := PhoneBook{
		Name:   name,
		Number: number,
		Pre:    pre,
		Next:   next,
	}
	a.PhoneBook = append(a.PhoneBook, &book)
}

func (a *Array) AddTwo() *Array {
	rootArr := initArray()
	for a := 0; a < 3; a++ {
		rootArr.AddBook("pjh", a, nil, nil)
		if a >= 1 {
			b := a - 1
			rootArr.PhoneBook[b].Next = rootArr.PhoneBook[a]
			rootArr.PhoneBook[a].Pre = rootArr.PhoneBook[b]
		}
	}
	return rootArr
}

func (a *Array) ShowNode() {
	for i := 0; i < 3; i++ {
		fmt.Println(arr.PhoneBook[i].Name, arr.PhoneBook[i].Number)
		if arr.PhoneBook[i].Next == nil {
			return
		} else {
			fmt.Println(arr.PhoneBook[i].Next.Name, arr.PhoneBook[i].Next.Number)
		}
	}
}

func main() {
	arr.AddTwo()
	arr.ShowNode()
}
