// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// variables and datatypes
func variables() {
	var age1 int = 25
	var name1 string = "gilbert"
	fmt.Println(name1, ": ", age1)

	age2 := 25
	name2 := "gilbert2"
	fmt.Println(name2, ": ", age2)

	var age3, name3 = 25, "gilbert3"
	fmt.Println(name3, ": ", age3)

	fmt.Println("Enter Your age:")
	fmt.Scanln(&age3)

	fmt.Println("Enter Your name:")
	fmt.Scanln(&name3)

	fmt.Println("So you are ", name3, "and your age is ", age3)
}

func main() {
	fmt.Println("Hello mr. ??")
	variables()
}
