package dubble_linked_list

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

func addTwoBook() {
	var a int
	for a := 0; a < 3; a++ {
		a = newPhoneBook("pjh", a, nil, nil)

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
	// 두개의 새로운 노드를 추가하는 함수, 모든 노드를 출력하는 함수 별도의 함수로써 만들어서 활용할 수 있게 빼보기
	//가만히 보면 패턴이 있음
	//전역 변수, 다른 변수 추가, golang 더블 링크드 리스트 작성된걸 보고 하는 것
	//수요일이나 목요일까지
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
