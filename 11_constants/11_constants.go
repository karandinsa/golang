package main

const someInt = 1         //нетипизированная константа
const typedint int32 = 17 //типизированная константа
const fullName = "Sergey"

//короткий синтаксис
const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	one = iota
	two
	_
	four
)

const (
	_ = iota
	//KB kilobytes
	KB uint64 = 1 << (10 * iota)
	//MB megabytes
	MB
	//GB gigabytes
	GB
	//TB terabytes
	TB
	//PB petabytes
	PB
	//EB exabytes
	EB
)
const (
	flagUserEnabled = 1 << iota
	flagUserVerified
	flagUserPremium
)

func main() {
	pi := 3.14
	println(pi + someInt)
	println(pi + float64(typedint))
	println(KB, MB, GB, PB, EB)
	println(one, two, four)
	println(flagUserEnabled, flagUserVerified, flagUserPremium, 2&flagUserEnabled)
}
