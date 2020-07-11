package main

import "fmt"

func main() {
	var i int = 10
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1
	var unsignedInt uint = 100500
	var unsignedBigInt uint64 = 1<<64 - 1

	//целые числа
	println("integers:", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	//числа с плавающей точкой
	var p float32 = 3.14
	println("float:", p)

	//булевы значения
	var b bool = true
	println("bool variable:", b)

	// строки
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	//бинарные данные
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)

	//короткое объявление (переменная может быть оъявлена единожды)
	meaningOfLife := 42
	println("Meaning of life is: ", meaningOfLife)

	//приведение типов
	println("float to int conversion", int(p))
	var u1 uint = 17
	var s1 int = 23
	println(int(u1) + s1)
	println("int to string conversion", string(48))

	//комплексные числа
	z := 2 + 3i
	println("complex number:", z)

	//операции со строками
	s11 := "Vasily"
	s22 := "Romanov"
	fullname := s11 + " " + s22
	println("name lenth is: ", fullname, len(fullname))

	escaping := `Hello\r\n
	World`
	println("as-is escaping: ", escaping)

	//значения по умолчанию
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool

	println("default values:", defaultInt, defaultFloat, defaultString, defaultBool)

	//указание нескольких переменных
	var v1, v2 = "v1", "v2"
	println(v1, v2)

	var (
		m0 int    = 12
		m2 string = "string"
		m3        = 23
	)
	println("common declares:", m0, m2, m3)

	x := "7" +
		"6"
	fmt.Println(x)
	return
}
