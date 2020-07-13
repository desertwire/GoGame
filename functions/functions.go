package functions

import (
	"fmt"
	"math/rand"
	"strconv"
	"../structs"
)

//Alias for shorter print function name
var print   = 	fmt.Print
var println = 	fmt.Println
var printf 	= 	fmt.Printf
var scanf	= 	fmt.Scanf

//Struct player1 attacks struct player2
//Attack is based off a random number between 0 and Max Attack
func AttackPlayer(TempPlayer1 *structs.Player, TempPlayer2 *structs.Player) {
	TempPlayer2.Health -= rand.Intn(TempPlayer1.Attack)
}

//Return the struct of the player in struct form
func CheckStats(TempPlayer structs.Player) structs.Player {
	return TempPlayer
}

//Convert the struct of the Player into a string
func ParsePlayerStruct(TempPlayer structs.Player) string {
	return 	"=======================" + "\n" + 
			"Name: "    + TempPlayer.Name +
			"Health: "	+ strconv.Itoa(TempPlayer.Health) + " \n" + 
			"Level: "	+ strconv.Itoa(TempPlayer.Level) + 	" \n" + 
			"Attack: "  + strconv.Itoa(TempPlayer.Attack) + " \n" + 
			"Defense: " + strconv.Itoa(TempPlayer.Defense) + "\n" +
			"=======================" + "\n"
}


