package functions

import "fmt"

//Alias for shorter print function name
var print   = 	fmt.Print
var println = 	fmt.Println
var printf 	= 	fmt.Printf
var scanf	= 	fmt.Scanf

// This func must be Exported, Capitalized, and comment added.
func Demo() {
    fmt.Println("Second Package Works!")
}

func FirstBossMenu() {
	println("=======================")
	println(" 1. Attack")
	println(" 2. Defend")
	println(" 3. Run")
	println("-1. Exit")
	println("=======================")
}

func Menu() {
	println("=======================")
	println(" 1. Attack")
	println(" 2. Defend")
	println(" 3. Bag")
	println("-1. Exit")
	println("=======================")
}

//Initalize On Startup
func Startup() {
	println("=======================\n" + 
			"Welcome to the GoGame!\n"	+
			"=======================\n")	
}

func FirstBoss() {
	println("=======================")
	println("The Space Boss Randomly Approaches")
	println("Do you choose to accept this fight?")
	println("1. Yes")
	println("2. No")
	println("=======================")
}