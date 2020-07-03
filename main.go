package main

import (
"bufio"
"fmt"
//"math/rand"
"os"
//"time"
"./structs"
"./textfunctions"
"./crypto"
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

//Import struct from another package
var testplayer = structs.Player{Name: "", Level: 0, Attack: 0, Defense: 0}

func initializeLevels() {
	println("What is " + "your name?")
	playerName, _ := reader.ReadString('\n')
	printf("Your new name is %s\n",  playerName)
	//printf("Current Player Struct:\n %+v\n", currentPlayer)
	currentPlayer.name = playerName
}

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

func main() {
	crypto.Sha1Checksum("Test")
	textfunctions.Startup()
	initializeLevels()
	numberOfTerms := 0
	for {

		textfunctions.Menu()
		var i int
		fmt.Scanf("%d", &i)
		println(i)
		textfunctions.FirstBoss()
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			println("You choose to fight the space boss")
			textfunctions.FirstBossMenu()
			if currentPlayer.level > 1 {
				println("Boss Defeated")
			} else {
				println("You've been defeated")
				os.Exit(3)
			}
		case 2:
			println("You decide to flee")
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
