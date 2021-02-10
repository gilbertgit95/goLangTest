package testservice

import "fmt"

// variables and datatypes
func Variables() {
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

// conditional statements
func Conditions() {
	var name string
	var askAny string
	var gender string
	var addrName string
	var coffee string
	var taste string

	fmt.Println("Hello, I will be your assistant today.")
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println("Nice to meet you ", name)
	fmt.Println("Can I ask you anything right now?")
	fmt.Println("-> type: [Y] if yes, any if no")
	fmt.Scanln(&askAny)

	if askAny == "Y" {
		fmt.Println("Sorry to ask you this, but what is your gender?")

	genderQ:
		fmt.Println("-> type: [M] if male, [F] if female")
		fmt.Scanln(&gender)

		if gender == "M" {
			addrName = "Mr."

		} else if gender == "F" {
			addrName = "Ms."

		} else {
			fmt.Println("Pardon me ", name, ", because I dont understand, please repeat you answer :-(")
			goto genderQ
		}

		fmt.Println("Do you want any of this coffe?")
		fmt.Println("-> type: caramel, milk, dark, americana")
		fmt.Scanln(&coffee)

		switch coffee {
		case "caramel":
			taste = "yes that a litle bit milky and yummy"
		case "milk":
			taste = "ohw the taste of goodness"
		case "dark":
			taste = "hmmm, that what I want if I had to OT"
		case "americana":
			taste = "too strong for me"
		default:
			taste = "I guese you dont want any of this"

		}

		fmt.Println(taste)
		fmt.Println("Okay ", addrName, " ", name, ", Thank you for your time, thats all.")

	} else {
		fmt.Println("Okay, no problem, take care.")
	}

}