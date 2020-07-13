package structs

//Fields from other packages need to be upper case
type Player struct {
	Name 	string
	Health	int
	Level   int
	Attack  int
	Defense int
}

type Bag struct {
	MyWeapon	Weapon
	MyShield	Shield
}

type Weapon struct {
	IsYielded	 bool
	AttackDamage int
}

type Shield struct {
	IsYielded	 bool
	AttackDamage int
}

var MainPlayer = Player{Name: "Main Name", Health: 10, Level: 0, Attack: 2, Defense: 2}
var ArchyPirate = Player{Name: "Archy Pirate", Health: 10, Level: 0, Attack: 1, Defense: 0}