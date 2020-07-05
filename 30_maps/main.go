package main

import "fmt"

func main() {
	//[тип ключа] тип данных
	var mm map[string]string
	//var mm2 map[string]map[string]string

	fmt.Println(mm)
	//mm["test"] = "ok"
	//fmt.Println(mm2)

	//полная инициализация
	//var mm2 map[string]string = map[string]string{}
	mm2 := map[string]string{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	//короткая инициализация
	var mm3 = make(map[string]string)
	mm3["firstName"] = "Sergey"
	fmt.Println(mm3)

	//получение значения
	firstName := mm3["firstName"]
	fmt.Println("First name", firstName, len(firstName))

	//если обратиться к несуществующему лючу, отдастся значение по умолчанию
	lastName := mm3["lastName"]
	fmt.Println("Last name", lastName, len(lastName))

	//проверка на то, что значение есть
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is:", lastName, "Exist:", ok)

	//получение тольк признака существования
	_, exist := mm3["firstName"]
	fmt.Println("Firstname exist:", exist)

	//удаление значения
	key := "FirstName1"
	delete(mm3, key)
	_, exist = mm3["firstName"]
	fmt.Println("Firstname exist:", exist)

	mm4 := mm3
	mm4[key] = "test"
	fmt.Println(mm3, mm4)

}
