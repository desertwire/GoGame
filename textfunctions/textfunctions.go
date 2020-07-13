package textfunctions

import (
	"fmt"
	"os"
	"../structs"
	"../functions"
)

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
	//Bug Main Character health not updating but Archy is
	for {
		//Main Player Attacks Archy
		printf("%s attacks %s\n",  structs.MainPlayer.Name, structs.ArchyPirate.Name)
		//Whenever a method wants to modify the receiver, it must be a pointer to the value;  
		functions.AttackPlayer(&structs.MainPlayer, &structs.ArchyPirate)
		printf("%s health %d\n",  structs.MainPlayer.Name, structs.MainPlayer.Health)
		printf("%s health %s\n",  structs.ArchyPirate.Name, structs.ArchyPirate.Health)
		if structs.ArchyPirate.Health <= 0  {
			println("Archy has been defeated")
			break
		}
		//Archy Attacks Main Player
		printf("%s attacks %s\n",  structs.ArchyPirate.Name, structs.MainPlayer.Name)
		functions.AttackPlayer(&structs.ArchyPirate, &structs.MainPlayer)
		printf("%s health %d\n",  structs.MainPlayer.Name, structs.MainPlayer.Health)
		printf("%s health %d\n",  structs.ArchyPirate.Name, structs.ArchyPirate.Health)
		if structs.MainPlayer.Health <= 0  {
			println("You die")
			os.Exit(3)		
			break
		}
	}
	println("Fight Ended")
	printf("Main Player Health %d", structs.ArchyPirate.Health)
}

func Menu() {
	println("=======================")
	println(" 1. Map")
	println(" 2. Bag")
	println(" 3. Check Stats")
	println("-1. Quit Game")
	println("=======================")
}

func Map() {
	numberOfTerms := 0
	println("=======================")
	println(" 1. Arch Park")
	println(" 2. Thinkpad Cave")
	println(" 3. Gentoo Labs")
	println("-1. Menu")
	println("=======================")
	for {	
		var i int
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			println("Welcome to Arch Park")
			println("Archy the Pirate attacks you")
			FirstBossMenu()
			break
		case 2:
			println("Welcome to Thinkpad Cave")
		case 3:
			println("Welcome to Gentoo Labs")
		default: 
			println("Back to Main Menu")
			break
		}	
		numberOfTerms++ // repeated forever
		break
	}
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
	for {	
		var i int
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			println("Welcome to Arch Park")
			break
		case 2:
			println("You flee away")
			break
		}	
		break
	}
}