package main

import (
"bufio"
"fmt"
//"math/rand"
"os"
//"time"
)

type player struct {
    name 	string
	level   int
	attack  int
	defense int
}

//Alias for shorter print function name
var print   = fmt.Print
var println = fmt.Println
var printf 	= fmt.Printf 

//String Buffer Reader
var reader = bufio.NewReader(os.Stdin)

//Initalize On Startup
func startup() {
	print("=======================\n" + 
			"Welcome to the GoGame!\n"	+
			"=======================\n")	
}

func initializeLevels() {
	println("What is your name?")
	playerName, _ := reader.ReadString('\n')
	printf("Your new name is %s\n",  playerName)
	currentPlayer := player{name: playerName, level: 0, attack: 0, defense: 0}
	_ = currentPlayer
	printf("Current Player Struct:\n %+v\n", currentPlayer)
}

func main() {
	//print(startup())
	startup()
	initializeLevels()
	println("=============")
	println("1. Attack")
	println("2. Defend")
	println("3. Bag")
	println("1. Exit")
	println("=============")
    println("Enter text: ")
    text, _ := reader.ReadString('\n')
    println(text)
}
