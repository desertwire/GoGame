package FirstBossMenu

import "fmt"

//Alias for shorter print function name
var print   = 	fmt.Print
var println = 	fmt.Println
var printf 	= 	fmt.Printf
var scanf	= 	fmt.Scanf

for {

	textfunctions.Menu()	
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		println("You choose to fight the space boss")
		println("To defeat the first boss you must reverse the SHA1 hash of: ")
		crypto.Sha1Checksum("password")
		println("or reverse this game lol")
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