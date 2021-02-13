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

func Loops() {
	// normal for loop
	fmt.Println("Count from 0-10 using normal for loop")
	for i := 0; i <= 10; i++ {
		if i < 10 {
			fmt.Print(i, ", ")
		} else {
			fmt.Println(i)
		}
	}

	// while loop
	fmt.Println("Count from 0-10 using while loop")
	x := 0
	for x <= 10 {
		if x < 10 {
			fmt.Print(x, ", ")
		} else {
			fmt.Println(x)
		}
		x++
	}

	// range loop or foreach
	items := []string{"one", "two"}

	for index, item := range items {
		fmt.Println(index, " -> ", item)
	}
}

// slice
func SliceTypes() {
	names := []string{"gilbert test1", "gilbert test2"}
	friends := []string{"jef", "Plong"}

	result := append(names, "gilbert test3")
	result2 := append(names, friends...)

	fmt.Println(result, " -> ", len(result))
	fmt.Println(result2, " -> ", len(result2))

	for index, name := range result2 {
		fmt.Println(index, " -> ", name)
	}
}

// structs
func Structs() {
	type Player struct {
		name  string
		age   int
		score int
	}
	type Team struct {
		name        string
		description string
		players     []Player
	}

	var philTeam = Team{
		name:        "Phil team org",
		description: "all start players",
		players: []Player{
			{
				name:  "gilbert",
				age:   26,
				score: 1,
			},
			{
				name:  "gilbertwo",
				age:   24,
				score: 10,
			},
		},
	}

	fmt.Println("Team Name: ", philTeam.name)
	fmt.Println("Team Description: ", philTeam.description)
	fmt.Println("Players -> ")

	for index, player := range philTeam.players {
		fmt.Println("[", index+1, "]____________")
		fmt.Println("player name: ", player.name)
		fmt.Println("player age: ", player.name)
		fmt.Println("player score: ", player.name)
	}
}
