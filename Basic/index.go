package main

import "fmt"

func main() {
	//hello world
	fmt.Println("Hello World")

	//case sensetive
	var val = 1
	var Val = 2
	fmt.Println(val)
	fmt.Println(Val)

	//data types
	fmt.Println(5 + 5)
	fmt.Println("5+5")

	//check data type of variable
	var str = "go"
	var val1 = 5.5
	fmt.Printf("%T\n", str)
	fmt.Printf("%T\n", val1)

	//varibales
	var var2 int
	var2 = 10
	var var3, var4 = "Hello World", 9
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	const x = 90
	fmt.Println(x)

	//arithmetic opeators
	fmt.Println(5 + 5)
	fmt.Println(5 - 6)
	fmt.Println(5 * 5)
	fmt.Println(5 / 5)
	fmt.Println(13 % 5)
	var d = 10
	d++
	fmt.Println(d)

}
