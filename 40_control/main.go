package main

import "fmt"

func main() {
	a := true
	if a {
		fmt.Println("Hello world")
	}
	b := 1
	c := 22
	if (b == 1 && a) || c != 22 {
		fmt.Println("неавное преобразование if b не работает")
	}
	mm := map[string]string{
		"firstName": "Sergey",
		"lastName":  "Karandin",
	}
	if _, ok := mm["firstName"]; ok {
		fmt.Println("Firstname key is exist", ok)
	} else {
		fmt.Println("No firstname key")
	}

	if firstName, ok := mm["firstName"]; !ok {
		fmt.Println("no firstname")
	} else if firstName == "Sergey" {
		fmt.Println("Firstname is Sergey")
	} else {
		fmt.Println("Firstname is not Sergey")
	}

	//Циклы
	for {
		fmt.Println("Бесконечный цикл while(true)")
		break
	}
	sl := []int{3, 4, 5, 6, 7, 8}
	value := 0
	idx := 0
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl[idx]
		idx++
		fmt.Println("while-style loop,idx", idx, "value:", value)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}

	for idx := range sl {
		fmt.Println("range slice by index", idx)
	}
	for idx, val := range sl {
		fmt.Println("range slice by index", idx, val)
	}

}
