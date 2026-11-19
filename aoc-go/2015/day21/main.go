package main

import (
	"fmt"
	"iter"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const itemSlots = 4

type ItemSlot uint8

const (
	WEAPON ItemSlot = iota
	ARMOR
	RING1
	RING2
)

type Entity struct {
	hp    int
	dmg   int
	armor int
}

type Item struct {
	cost  int
	dmg   int
	armor int
}

var weapons = [...]Item{{8, 4, 0}, {10, 5, 0}, {25, 6, 0}, {40, 7, 0}, {74, 8, 0}}
var armors = [...]Item{{13, 0, 1}, {31, 0, 2}, {53, 0, 3}, {75, 0, 4}, {102, 0, 5}}
var rings = [...]Item{{25, 1, 0}, {50, 2, 0}, {100, 3, 0}, {20, 0, 1}, {40, 0, 2}, {80, 0, 3}}

func ParseFile(path string) Entity {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	hp, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	dmg, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	armor, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])

	return Entity{hp, dmg, armor}
}

func CheckWinFight(player, enemy Entity) bool {
	playerHit, enemyHit := max(1, player.dmg-enemy.armor), max(1, enemy.dmg-player.armor)
	playerHits := (enemy.hp + playerHit - 1) / playerHit
	enemyHits := (player.hp + enemyHit - 1) / enemyHit
	return playerHits <= enemyHits
}

func ItemCombinations() iter.Seq[[itemSlots]Item] {
	var comb [itemSlots]Item
	var armorChoices [len(armors) + 1]Item
	copy(armorChoices[1:], armors[:])
	var ringChoices [len(rings) + 1]Item
	copy(ringChoices[1:], rings[:])

	return func(yield func([itemSlots]Item) bool) {
		for _, w := range weapons {
			comb[WEAPON] = w
			for _, a := range armorChoices {
				comb[ARMOR] = a
				// 0 rings
				comb[RING1] = ringChoices[0]
				comb[RING2] = ringChoices[0]
				if !yield(comb) {
					return
				}
				// 1-2 rings
				for i, r1 := range ringChoices[:len(ringChoices)-1] {
					comb[RING1] = r1
					for _, r2 := range ringChoices[i+1:] {
						comb[RING2] = r2
						if !yield(comb) {
							return
						}
					}
				}
			}
		}
	}
}

func PartOne(player, enemy Entity) int {
	minGold := math.MaxInt

	for c := range ItemCombinations() {
		var gold, addDmg, addArmor int
		for _, item := range c {
			gold += item.cost
			addDmg += item.dmg
			addArmor += item.armor
		}
		if gold > minGold {
			continue
		}
		p := Entity{
			player.hp,
			player.dmg + addDmg,
			player.armor + addArmor,
		}
		if CheckWinFight(p, enemy) {
			minGold = gold
		}
	}

	return minGold
}

func PartTwo(player, enemy Entity) int {
	maxGold := 0

	for c := range ItemCombinations() {
		var gold, addDmg, addArmor int
		for _, item := range c {
			gold += item.cost
			addDmg += item.dmg
			addArmor += item.armor
		}
		if gold <= maxGold {
			continue
		}
		p := Entity{
			player.hp,
			player.dmg + addDmg,
			player.armor + addArmor,
		}
		if !CheckWinFight(p, enemy) {
			maxGold = gold
		}
	}

	return maxGold
}

func main() {
	enemy := ParseFile("./input.txt")
	player := Entity{100, 0, 0}
	fmt.Println(PartOne(player, enemy))
	fmt.Println(PartTwo(player, enemy))
}
