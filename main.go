package main

import (
"bufio"
"fmt"
//"math/rand"
"os"
//"time"
"./functions"
)

type player struct {
    name 	string
	level   int
	attack  int
	defense int
}

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

var currentPlayer = player{name: "", level: 0, attack: 0, defense: 0}
	//_ = currentPlayer

func initializeLevels() {
	println("What is " + "your name?")
	playerName, _ := reader.ReadString('\n')
	printf("Your new name is %s\n",  playerName)
	//printf("Current Player Struct:\n %+v\n", currentPlayer)
	currentPlayer.name = playerName
}

func levelUp(p player) {
	p.level += 1
}

func increaseAttack(p player) {
	p.attack += 1
}

func increaseDefense(p player) {
	p.defense += 1
}

func main() {
	functions.Startup()
	initializeLevels()
	numberOfTerms := 0
	for {
		functions.Menu()
		var i int
		fmt.Scanf("%d", &i)
		println(i)
		functions.FirstBoss()
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			fmt.Println("You choose to fight the space boss")
			functions.FirstBossMenu()
			if currentPlayer.level > 1 {
				println("Boss Defeated")
			} else {
				println("You've been defeated")
				os.Exit(3)
			}
		case 2:
			fmt.Println("You decide to flee")
		case 3:
			fmt.Println("three")
		default: 
			println("Quitting the Game")
			os.Exit(3)
			break
		}	
		numberOfTerms++ // repeated forever
		}
	}
