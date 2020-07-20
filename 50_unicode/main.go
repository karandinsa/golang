package main

import "fmt"

func main() {
	fmt.Println("I`m ready to write Unicode")
	var symbol rune = 'a'
	var autoSymbol = 'a'
	unicodeSymbol := 'ὥ'
	unicodeSymbolByNumber := '\u2318'
	println(symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)

	str1 := "Привет, мир!"
	fmt.Println("ru: ", str1, len(str1))
	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	str2 := "你好世界"
	fmt.Println("cn:", str2, len(str2))
	for index, runeValue := range str2 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
		//обращение непосредственно к байтовому значению
		println(str2[1])
	}

}
