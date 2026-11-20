package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Entity struct {
	hp            int
	dmg           int
	armor         int
	mana          int
	statusEffects [statusEffectCount]StatusEffect
}

func (e *Entity) TickStatusEffects() {
	for i := range e.statusEffects {
		sf := &e.statusEffects[i]
		if sf.turns <= 0 {
			continue
		}

		switch sf.effectType {
		case DAMAGE:
			e.hp -= sf.value
		case MANA:
			e.mana += sf.value
		}

		sf.turns--
	}
}

func (e Entity) Armor() int {
	if sf := e.statusEffects[ARMOR]; sf.turns > 0 {
		return e.armor + sf.value
	}
	return e.armor
}

func (e *Entity) Attack(target *Entity) {
	target.hp -= max(1, e.dmg-target.Armor())
}

func (e *Entity) CastSpell(s Spell, target *Entity) bool {
	if e.mana < s.manacost {
		return false
	}

	if s.statusEffect != nil {
		idx := s.statusEffect.effectType
		if target.statusEffects[idx].turns > 0 {
			return false
		}
		target.statusEffects[idx] = *s.statusEffect
	}

	e.mana -= s.manacost
	target.hp -= s.dmg
	e.hp += s.heal

	return true
}

type Spell struct {
	name         string
	manacost     int
	dmg          int
	heal         int
	statusEffect *StatusEffect
}

var spells = [...]Spell{
	{"Magic Missle", 53, 4, 0, nil},
	{"Drain", 73, 2, 2, nil},
	{"Shield", 113, 0, 0, &StatusEffect{ARMOR, 7, 6}},
	{"Poison", 173, 0, 0, &StatusEffect{DAMAGE, 3, 6}},
	{"Recharge", 229, 0, 0, &StatusEffect{MANA, 101, 5}},
}

type StatusEffectType uint8

const (
	DAMAGE StatusEffectType = iota
	MANA
	ARMOR
	statusEffectCount
)

type StatusEffect struct {
	effectType StatusEffectType
	value      int
	turns      int
}

func ParseFile(path string) Entity {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	hp, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	dmg, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])

	return Entity{hp: hp, dmg: dmg}
}

func PartOne(player, enemy Entity) int {
	minMana := math.MaxInt

	var dfs func(p, e Entity, manaSpent int)
	dfs = func(p, e Entity, manaSpent int) {
		if manaSpent >= minMana {
			return
		}

		// -- Player turn --
		p.TickStatusEffects()
		e.TickStatusEffects()
		// Boss died by StatusEffect on Player's turn
		if e.hp <= 0 {
			minMana = min(manaSpent, minMana)
			return
		}

		casted := false
		for _, spell := range spells {
			pCopy, eCopy := p, e
			var target *Entity
			switch spell.name {
			case "Recharge", "Shield":
				target = &pCopy
			default:
				target = &eCopy
			}
			if !pCopy.CastSpell(spell, target) {
				continue
			}
			casted = true
			newManaSpent := manaSpent + spell.manacost
			// Boss died by Spell damage on Player's turn
			if eCopy.hp <= 0 {
				minMana = min(newManaSpent, minMana)
				continue
			}
			// -- Boss turn --
			pCopy.TickStatusEffects()
			eCopy.TickStatusEffects()
			// Boss died by StatusEffect on Own turn
			if eCopy.hp <= 0 {
				minMana = min(newManaSpent, minMana)
				continue
			}
			eCopy.Attack(&pCopy)
			// Player died
			if pCopy.hp <= 0 {
				continue
			}

			dfs(pCopy, eCopy, newManaSpent)
		}

		// Unable to cast any Spell
		if !casted {
			return
		}
	}

	dfs(player, enemy, 0)

	return minMana
}

func PartTwo(player, enemy Entity) int {
	minMana := math.MaxInt

	var dfs func(p, e Entity, manaSpent int)
	dfs = func(p, e Entity, manaSpent int) {
		if manaSpent >= minMana {
			return
		}

		// -- Player turn --
		p.hp--
		if p.hp <= 0 {
			return
		}
		p.TickStatusEffects()
		e.TickStatusEffects()
		// Boss died by StatusEffect on Player's turn
		if e.hp <= 0 {
			minMana = min(manaSpent, minMana)
			return
		}

		casted := false
		for _, spell := range spells {
			pCopy, eCopy := p, e
			var target *Entity
			switch spell.name {
			case "Recharge", "Shield":
				target = &pCopy
			default:
				target = &eCopy
			}
			if !pCopy.CastSpell(spell, target) {
				continue
			}
			casted = true
			newManaSpent := manaSpent + spell.manacost
			// Boss died by Spell damage on Player's turn
			if eCopy.hp <= 0 {
				minMana = min(newManaSpent, minMana)
				continue
			}
			// -- Boss turn --
			pCopy.TickStatusEffects()
			eCopy.TickStatusEffects()
			// Boss died by StatusEffect on Own turn
			if eCopy.hp <= 0 {
				minMana = min(newManaSpent, minMana)
				continue
			}
			eCopy.Attack(&pCopy)
			// Player died
			if pCopy.hp <= 0 {
				continue
			}

			dfs(pCopy, eCopy, newManaSpent)
		}

		// Unable to cast any Spell
		if !casted {
			return
		}
	}

	dfs(player, enemy, 0)

	return minMana
}

func main() {
	enemy := ParseFile("./input.txt")
	player := Entity{hp: 50, mana: 500}
	fmt.Println(PartOne(player, enemy))
	fmt.Println(PartTwo(player, enemy))
}
