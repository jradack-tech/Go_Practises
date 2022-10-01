//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
type Player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once
func (player *Player) modifyHealth(health int) {
	player.health = health
	fmt.Println("After modify: Health -", player.health)
}

func (player *Player) modifyEnergy(energy int) {
	player.energy = energy
	fmt.Println("After modify: Energy -", player.energy)
}

func main() {
	player := Player{name: "Brew"}
	fmt.Println("-------- Player Brew's Status ---------")
	fmt.Println("Health -", player.health)
	fmt.Println("Energy -", player.energy)

	player.modifyHealth(98)
	player.modifyEnergy(90)
}
