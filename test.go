//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)

//	func main() {
//		s, s2, s3 := test()
//		fmt.Println(s3)
//		fmt.Println(s2)
//		fmt.Println(s)
//
//		fmt.Println(test())
//		fmt.Println(test1())
//
// }
//func test() (string, string, string) {
//	a := "hello"
//	c := "all"
//	b := "world"
//	return a, b, c
//
//}
//
//func test1() string {
//	return "empty"
//}

//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}

//	func main() {
//		sum := 0
//		for sum < 1000 {
//			sum += 10
//		}
//
//		fmt.Println(sum)
//	}
//
//	func main() {
//		a := 0
//		for a < 1000 {
//			if a == 100 {
//				fmt.Println("a is 100")
//			} else {
//				fmt.Println(fmt.Sprint("a is not 100. a=%d", a))
//			}
//			a++
//		}
//	}
//func main() {
//	a := 0
//	for a < 1000 {
//		i := isTest(a)
//		if i == 1 {
//			fmt.Println("a = 2")
//		} else if i == 2 {
//			fmt.Println("a = 3")
//		} else {
//			fmt.Println(fmt.Sprintf("unknown a. a=%d", a))
//		}
//		a++
//	}
//
//}

//func main() {
//
//	a := 0
//	for a < 1000 {
//		i := isTest(a)
//		switch i {
//		case 1:
//			fmt.Println("a = 2")
//		case 2:
//			fmt.Println("a = 3")
//		default:
//			fmt.Println(fmt.Sprintf("unknown a. a=%d", a))
//		}
//		a++
//	}
//}
//
//func isTest(a int) int {
//	if a == 2 {
//		return 1
//	} else if a == 3 {
//
//		return 2
//	}
//	return 3
//
//}

//
//func main() {
//	defer fmt.Println("wold")
//	fmt.Println("hello")

// }
//func main() {
//	//stack, last in first out  \
//	defer fmt.Println("1")
//	defer fmt.Println("2")
//	defer fmt.Println("3")
//	fmt.Println("hello")
//}
//
//func test56(nume int) {
//	if nume > 5 {
//		fmt.Println(">5")
//
//	} else {
//		fmt.Println("<5")
//
//	}
//
//}

//func main() {
//	pointers()
//}
//
//func pointers() {
//
//	a := "hello world"
//	b := 42
//	fmt.Println(a)
//	fmt.Println(b)
//
//	p := &a
//	fmt.Println(p)
//}

//func main() {
//	pointers()
//}
//
//func pointers() {
//
//	a := "hello world"
//
//	fmt.Println(a)
//
//	p := &a
//	fmt.Println(p)
//	fmt.Println(*p)
//	*p = "oh my"
//	fmt.Println(a)
//
//	b := 42
//	g := &b
//	*g = *g / 2
//
//	fmt.Println(b)
//}

//type Point struct {
//	X int
//	Y int
//	S string
//}
//
//func (p Point) method() {
//	fmt.Println(p.X)
//	fmt.Println(p.Y)
//	fmt.Println(p.S)
//}
//
//func main() {
//	p1 := Point{
//		X: 1,
//		Y: 2,
//		S: "hello",
//	}
//	p2 := &p1
//	p2.method()
//}

//func main() {
//	structs()
//}
//
//func structs() {
//	p1 := Point{
//		X: 1,
//		Y: 2,
//	}
//	fmt.Println(1)
//	fmt.Println(p1.X)
//	p1.X = 123
//	fmt.Println(p1)
//
//	p2 := Point{
//		X: 123,
//	}
//	fmt.Println(p2)
//
//	g := &p1
//	fmt.Println(*g)
//	fmt.Println(g)
//	fmt.Println(&g)
//	fmt.Println(*&g)
//}

//func main() {
//slice and arrays

//animalsArr := [4]string{
//	"dog",
//	"cat",
//	"giraffe",
//	"elephant",
//}
//a := animalsArr[0:2]
//fmt.Println(a)
//b := animalsArr[1:3]
//fmt.Println(b)

//	var a []string = animalsArr[0:2]
//	fmt.Println(a)
//	var b []string = animalsArr[1:3]
//
//	b[0] = "123"
//	fmt.Println(a)
//	fmt.Println(animalsArr)
//
//}

//	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	t := s[:5]
//	fmt.Println(t)
//	tt := s[5:]
//	fmt.Println(tt)
//	ttt := s[0:]
//	fmt.Println(ttt)
//}
//	var snil []int
//	fmt.Println(snil, len(snil), cap(snil))
//	if snil == nil {
//		fmt.Println("slice is nil")
//	}
//
//}

//	letters := []string{"a", "b", "c"}
//	letters[0] = "1"
//	letters = append(letters, "d")
//	letters = append(letters, "e", "f")
//	fmt.Println(letters)
//
//	createSlice := make([]string, 3)
//	fmt.Println(len(createSlice))
//	fmt.Println(cap(createSlice))
//	createSlice[0] = "1"
//	createSlice[1] = "2"
//	createSlice[2] = "3"
//	createSlice = append(createSlice, "4")
//	fmt.Println(createSlice)
//	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
//	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
//
//}

//func main() {
//	//arrays
//	var a [2]string
//	a[0] = "hello"
//	a[1] = "world"
//	// [2<]=errorrrr
//	fmt.Println(a[1])
//
//	numbers := [...]int{1, 2, 3}
//	numbers[2] = 4
//
//}

//func doSomething(callback func(int, int) int, s string) int {
//	result := callback(5, 1)
//	fmt.Println(s)
//	return result
//}
//
//func main() {
//	sumCallback := func(n1, n2 int) int {
//		return n1 + n2
//	}
//	result := doSomething(sumCallback, "plus")
//	fmt.Println(result)
//
//	multipleCallback := func(n1, n2 int) int {
//		return n1 * n2
//	}
//	result = doSomething(multipleCallback, "multiple")
//	fmt.Println(result)
//}

//	func totalPrice(initPrice int) func(int) int {
//		sum := initPrice
//		return func(x int) int {
//			sum += x
//			return sum
//		}
//	}
//
//	func main() {
//		orderPrice := totalPrice(1)
//		fmt.Println(orderPrice(1))
//		fmt.Println(orderPrice(1))
//
//// }
//type Point struct {
//	X, Y int
//}
//
//func movePoint(p Point, x, y int) Point {
//	p.X += x
//	p.Y += y
//	return p
//}
//
//func main() {
//	p := Point{0, 0}
//	fmt.Println(p)
//	fmt.Println(movePoint(p, 3, 4))
//
//}

//func main() {
//	//arr := []string{"a", "b", "c"}
//	//for i, l := range arr {
//	//	fmt.Println(i)
//	//	fmt.Println(l)
//	//}
//
//}

//	type Point struct {
//		X int
//		Y int
//	}
//
//	func (p Point) method() {
//		fmt.Println("call Point method")
//	}
//
// func main() {
//
//	pointsMap := map[string]Point{
//		"b": {
//			Y: 13,
//			X: 15,
//		},
//	}
//
// otherPointsMap := make(map[int]Point)
//
//	pointsMap["a"] = Point{
//		X: 1,
//		Y: 2,
//	}
//
// //fmt.Println(pointsMap)
// //fmt.Println(pointsMap["a"])
// otherPointsMap[1] = Point{1, 2}
// //fmt.Println(otherPointsMap)
// //fmt.Println(otherPointsMap[1])
//
// var oneMorePointsMap map[string]Point
//
//	if oneMorePointsMap == nil {
//		//	fmt.Println("oneMorePointsMap is nil")
//		////} else {
//		////	oneMorePointsMap["a"] = Point{1, 2}
//		oneMorePointsMap = map[string]Point{
//			"a": {1, 2},
//		}
//	}
//
// //fmt.Println(oneMorePointsMap["a"])
// fmt.Println(oneMorePointsMap["a"].X)
// fmt.Println(oneMorePointsMap["a"].Y)
// oneMorePointsMap["a"].method()
//
// key := "a"
// value, ok := oneMorePointsMap[key]
//
//	if ok {
//		fmt.Printf("key= %s exist in map\n", key)
//		fmt.Println(value)
//	} else {
//
//		fmt.Printf("key= %s does exist in map\n", key)
//		fmt.Println(value)
//	}
//
// }
//type structHere struct {
//	N1, N2 int
//}
//
//func (s *structHere) Sum() int {
//	return s.N1 + s.N2
//}
//
//type InterfaceHere interface {
//	Sum() int
//}
//
//func main() {
//var a InterfaceHere
//sh := structHere{1, 2}
//os := otherStruct{2, 3}
//
//a = &sh
////fmt.Println(a.Sum())
//a = os
////fmt.Println(a.Sum())
//var i InterfaceHere = otherStruct{3, 3}
//fmt.Println(i.Sum())
//var os *structHere
//	var i InterfaceHere
//	i = os
//	fmt.Printf("%v, %T)\n", i, i)
//}
//
//type otherStruct struct {
//	A, B int
//}
//
//func (o otherStruct) Sum() int {
//	return o.A + o.B
//
//
//}

//func main() {
//	var a interface{}
//	a = "jelli"
//	fmt.Println(a)
//	fmt.Printf("(%v, %T)\n\n", a, a)
//	a = 42
//	fmt.Println(a)
//	fmt.Printf("(%v, %T)\n\n", a, a)

//}

//func main() {
//var a interface{} = "hello"
//s := a.(string)
//fmt.Println(s)
//
//s, ok := a.(string)
//fmt.Println(s, ok)
//	var a interface{} = "hello"
//
//	switch a.(type) {
//	case int:
//		fmt.Println("a is int")
//	case string:
//		fmt.Println("a is string")
//	default:
//		fmt.Printf("unknown type %T\n", a)
//
//	}
//
//}

//	type AppError struct {
//		Err    error
//		Custom string
//		Field  int
//	}
//
//	func (e *AppError) Error() string {
//		return e.Err.Error()
//	}
//
//	func main() {
//		//err := m()
//		//if err != nil {
//		//}
//		//fmt.Println(err.Error())
//		err := method1()
//		if err != nil {
//			fmt.Println(err)
//			//return
//		}
//		fmt.Println("success")
//	}
//
//	func method1() error {
//		//return &AppError{
//		//	Err:    fmt.Errorf("my error"),
//		//	Custom: "value here",
//		//	Field:  89,
//		//}
//		return method2()
//	}
//
//	func method2() error {
//		return method3()
//
// }
//
//	func method3() error {
//		return fmt.Errorf("some error")
//	}
//
//	func main() {
//		r := strings.NewReader("hello world")
//		arr := make([]byte, 11)
//		for {
//			n, err := r.Read(arr)
//			fmt.Printf("n = %d err = %v arr = %v\n", n, err, arr)
//			fmt.Printf("arr n bytes: %q", arr[:n])
//			if err == io.EOF {
//				break
//			}
//		}

//func main() {
//
//	fmt.Println("my name aslan")
//
//}

//func main() {
//
//	a := 11
//	b := a / 3
//
//	fmt.Println(b)
//
//}

//func main()  {
//
//}

// const (
//
//	i = 8
//	j
//	k
//
// )
//
//	func main() {
//		fmt.Println(i, j, k)
//	}
//func main() {
//	fmt.Print("Мой вес на поверхности Марса равен ")
//	fmt.Print(55.0 * 0.3783) // В результате 20.8065
//	fmt.Print(" килограммам, а мой возраст равен ")
//	fmt.Print(41 * 365 / 687) // В результате 21
//	fmt.Print(" годам.")
//}

//func main() {
//	const lightSpeed = 100800 // км/c
//	var distance = 96300000   // км
//
//	fmt.Println(distance/lightSpeed, "секунд") // В результате 186 секунд
//
//	distance = 401000000
//	fmt.Println(distance/lightSpeed, "секунд") // В результате 1337 секунд
//}

//func main() {

//	const hoursPerDay = 24
//
//	var days = 28
//	var distance = 56000000
//	fmt.Println(distance/(hoursPerDay*days), "км/ч")
//}

//func main() {
//	const hoursPerDay, minutesPerDay = 24, 60
//
//}
//var distance = rand.Intn(345000001) + 56000000
//fmt.Println(distance)

//}
//	const hoursPerDay = 24
//
//	var days = 28
//	var distance = 56000000 // km
//
//	fmt.Println(distance/(days*hoursPerDay), "км/ч")
//}

//	fmt.Println("Вы находитесь в темной пещере.")
//
//	var command = "выйти наружу"
//	var exit = strings.Contains(command, "наружу")
//
//	fmt.Println("Вы покидаете пещеру:", exit) // Выводит: Вы покидаете пещеру: true
//}
//var room = "пещера"
//
//if room == "пещера" {
//	fmt.Println("Вы находитесь в тускло освещенной пещере.")
//} else if room == "вход" {
//	fmt.Println("Здесь есть вход в пещеру и путь на восток.")
//} else if room == "гора" {
//	fmt.Println("Здесь крутой утес. Тропа ведет к подножью горы.")
//} else {
//	fmt.Println("Здесь ничего нет.")
//}

// var yaer = 2000
// var leap = yaer%400 == 0
//
//	{
//		if leap {
//			fmt.Println("vysokostny god")
//			fmt.Println(leap)
//		}
//	}
//package main
//
//func main() {

//var count = 10
//for count > 0 {
//	fmt.Println(count)
//	time.Sleep(time.Second)
//	if rand.Intn(100) == 0 {
//		break
//	}
//	count--
//}
//if count == 0 {
//	fmt.Println("zapusk")
//} else {
//	fmt.Println(" zapusk otmenyetsy")
//}
//	var number = 8
//	for {
//		var n = rand.Intn(10) + 1
//		if n < number {
//			fmt.Printf("%v слишком маленькое число.\n, n")
//
//		} else if n > number {
//			fmt.Printf("%v слишком большое число.\n", n)
//		} else {
//			fmt.Printf("Угадал! %v\n", n)
//			break
//		}
//	}
//}
//var count = 0
//
//for count < 10 { // Начало области видимости
//	var num = rand.Intn(10) + 1
//	fmt.Println(num)
//	count++
//} // Конец области видимости
//var count = 0

//	for count := 10; count > 0; count-- {
//		fmt.Println(count)
//	}
//}
//if num := rand.Intn(3); num == 0 {
//	fmt.Println("Space Adventures")
//} else if num == 1 {
//	fmt.Println("SpaceX")
//} else {
//	fmt.Println("Virgin Galactic")
//}

//	var era = "AD" // переменная era доступна через пакет
//
//	{
//		year := 2018 // переменные era и year находятся в области видимости
//
//		switch month := rand.Intn(12) + 1; month { // переменные era, year и month в области видимости
//		case 2:
//			day := rand.Intn(28) + 1 // новый день
//			fmt.Println(era, year, month, day)
//		case 4, 6, 9, 11:
//			day := rand.Intn(30) + 1
//			fmt.Println(era, year, month, day)
//		default:
//			day := rand.Intn(31) + 1
//			fmt.Println(era, year, month, day)
//		} // month и day за пределами области видимости
//	} // year за пределами области видимости
//}
//var era = "AD"

//	year := 2020
//	month := rand.Intn(12) + 1
//	daysInMonth := 31
//
//	switch month {
//	case 1:
//		daysInMonth = 29
//
//	case 2, 4, 6, 9, 11:
//		daysInMonth = 30
//	}
//
//	day := rand.Intn(daysInMonth) + 1
//	fmt.Println(era, year, month, day)
//}

//var era = "AD"
//
//for count := 0; count < 10; count++ {
//	year := 2020 + rand.Intn(10)
//	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
//	month := rand.Intn(12) + 1
//	daysInMonth := 31
//	switch month {
//	case 2:
//		daysInMonth = 28
//		if leap {
//			daysInMonth = 29
//		}
//	case 4, 6, 9, 11:
//		daysInMonth = 30
//	}
//
//	day := rand.Intn(daysInMonth) + 1
//	fmt.Println(era, year, month, day)
//
//}

//const distance = 62100000 // км
//var speed = rand.Intn(16) + 14
//var hoursPerDay = 24
//var price = rand.Intn(36) + 14
//fmt.Println("days price speed company")
//fmt.Println(distance/speed/hoursPerDay/price, "дней", price, " million $", speed, "км/ч", "SpaceX")

//	fmt.Println(distance/speed/hoursPerDay/price, "дней", price, " million $", speed, "км", "Space Adventures")
//}
//
//fmt.Println(distance/speed/hoursPerDay/price, "дней", price, " million $", speed, "км", "Virgin Galactic")
//third := 0015.1021
//fmt.Println(third)
//fmt.Printf("%.2f", third)

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// Пчела
//type honeyBee struct {
//	name string
//}
//
//func (hb honeyBee) String() string {
//	return hb.name
//}
//
//func (hb honeyBee) move() string {
//	switch rand.Intn(2) {
//	case 0:
//		return "жужжит"
//	default:
//		return "летает и веселится"
//	}
//}
//
//func (hb honeyBee) eat() string {
//	switch rand.Intn(2) {
//	case 0:
//		return "пыльцу"
//	default:
//		return "нектар"
//	}
//}
//
//// Суслик
//type gopher struct {
//	name string
//}
//
//func (g gopher) String() string {
//	return g.name
//}
//
//func (g gopher) move() string {
//	switch rand.Intn(2) {
//	case 0:
//		return "гулят и изучает территорию"
//	default:
//		return "прячется в норку"
//	}
//}
//
//func (g gopher) eat() string {
//	switch rand.Intn(5) {
//	case 0:
//		return "морковку"
//	case 1:
//		return "салат-латук"
//	case 2:
//		return "редиску"
//	case 3:
//		return "кукурузу"
//	default:
//		return "корнеплоды"
//	}
//}
//
//type animal interface {
//	move() string
//	eat() string
//}
//
//func step(a animal) {
//	switch rand.Intn(2) {
//	case 0:
//		fmt.Printf("%v %v.\n", a, a.move())
//	default:
//		fmt.Printf("%v кушает %v.\n", a, a.eat())
//	}
//}
//
//const sunrise, sunset = 8, 18
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	animals := []animal{
//		honeyBee{name: "Шмель Базз"},
//		gopher{name: "Суслик Го"},
//	}
//
//	var sol, hour int
//
//	for {
//		fmt.Printf("%2d:00 ", hour)
//		if hour < sunrise || hour >= sunset {
//			fmt.Println("Животные спят.")
//		} else {
//			i := rand.Intn(len(animals))
//			step(animals[i])
//		}
//
//		time.Sleep(500 * time.Millisecond)
//
//		hour++
//		if hour >= 24 {
//			hour = 0
//			sol++
//			if sol >= 3 {
//				break
//			}
//		}
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	var number1, number2 float64
//	var operator string
//
//	fmt.Print("Enter the first number : ")
//	fmt.Scan(&number1)
//
//	fmt.Print("Enter the second number : ")
//	fmt.Scan(&number2)
//
//	fmt.Print("Enter the operator ( + - * / ) : ")
//	fmt.Scan(&operator)
//
//	switch operator {
//
//	case "+":
//		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1+number2)
//	case "-":
//		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1-number2)
//	case "*":
//		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1*number2)
//	case "/":
//		if number2 == 0.0 {
//			fmt.Println("Divide by Zero situation")
//		} else {
//			fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1/number2)
//		}
//	default:
//		fmt.Println("Invalid operator")
//	}
//}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	arabic := 0
	lastValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < lastValue {
			arabic -= value
		} else {
			arabic += value
		}
		lastValue = value
	}
	return arabic
}

func arabicToRoman(arabic int) string {
	numerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var roman strings.Builder
	for _, numeral := range numerals {
		for arabic >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return roman.String()
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func main() {
	var firstElement, operator, secondElement string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, 1 + 1):")
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Fields(input)
	firstElement = parts[0]
	operator = parts[1]
	secondElement = parts[2]
	flag1 := false
	flag2 := false

	num1, err1 := strconv.Atoi(firstElement)
	if err1 != nil {
		flag1 = true
		num1 = romanToArabic(firstElement)
	}

	num2, err2 := strconv.Atoi(secondElement)
	if err2 != nil {
		flag2 = true
		num2 = romanToArabic(secondElement)
	}
	if flag1 != flag2 {
		fmt.Println("Оба значения должны быть в одинаковом формате")
		return
	}
	result := calculate(num1, num2, operator)

	fmt.Printf("Result (Arabic): %d\n", result)
	fmt.Printf("Result (Roman): %s\n", arabicToRoman(result))
}
