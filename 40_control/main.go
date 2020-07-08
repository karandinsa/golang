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
		fmt.Println("range slice by idx-value", idx, val)
	}
	for _, val := range sl {
		fmt.Println("range slice by idx-value", val)
	}

	//операции по map
	for key := range mm {
		fmt.Println("range map by key", key)
	}
	for key, val := range mm {
		fmt.Println("range map by key-val", key, val)
	}
	for _, val := range mm {
		fmt.Println("range map by val", val)
	}

	//операция switch  - потестировать все возможные варианты: есть специфика
	mm["firstName"] = "Sergey"
	mm["flag"] = "Ok"
	switch mm["firstName"] {
	case "Sergey", "Evgeny":
		println("switch - name is Sergey")
		//если кейс выполнился, остальные нижележащие не выполняются.
	case "Petr":
		if mm["flag"] == "Ok" {
			break // принудительный выход из switch
		}
		println("switch - name is Petr")
		fallthrough //переход на следующий вариант
	default:
		println("switch - some other name")
	}
	//как замена множественным if else
	switch {
	case mm["firstName"] == "Sergey":
		println("Switch2 - Sergey")
	case mm["lastName"] == "Karandin":
		println("Switch2 - Karandin")
	}
	//выход из цикла будучи внутри switch
MyLoop: //лэйбл
	for key, val := range mm {
		println("switch in loop", key, val)
		switch {
		case key == "firstName" && val == "Sergey":
			println("switch - break loop here")
			break MyLoop
		}
	}
}
