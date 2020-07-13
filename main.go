package main

import (
"bufio"
"fmt"
//"math/rand"
"os"
//"time"
"./structs"
"./textfunctions"
"./functions"
//"./crypto"
)

//Alias for shorter print function name
var print   = 	fmt.Print
var println = 	fmt.Println
var printf 	= 	fmt.Printf
var scanf	= 	fmt.Scanf

const highestLevel = 100

//Initialize an array of 5 0's
//0 is still active, 1 has been defeated
var bossesDefeated [5]int

//String Buffer Reader
var reader = bufio.NewReader(os.Stdin)

func initializeLevels() {
	println("What is " + "your name?")
	playerName, _ := reader.ReadString('\n')
	structs.MainPlayer.Name = playerName
}

/*
func levelUp(p player) {
	println("You have leveled up.")
	printf("You're current level is: %d\n", p.level)
	p.level += 1
}

func increaseAttack(p player) {
	p.attack += 1
} 

func increaseDefense(p player) {
	p.defense += 1
}
*/

func main() {
	textfunctions.Startup()
	initializeLevels()
	numberOfTerms := 0
	for {
		textfunctions.Menu()	
		var i int
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			println("Opening Map")
			textfunctions.Map()
		case 2:
			println("You have opened bag")
		case 3:
			println("Checking Stats")
			print("\n")
			println(functions.ParsePlayerStruct(structs.MainPlayer))
		default: 
			println("Quitting the Game")
			os.Exit(3)
			break
		print("\n")
		}	
		numberOfTerms++ // repeated forever
	}
}
