package main

import "fmt"

func main() {
	var numLines uint
	fmt.Scanf("%d", &numLines)
	for i := uint(1); i < numLines+1; i++ {
		var d uint
		var p string
		fmt.Scanf("%d %s", &d, &p)

		baseDamage := evaluateDamage(p)

		// if damage done is less than or equal shield can take, output 0
		if baseDamage <= d {
			fmt.Printf("Case #%d: 0\n", i)
			continue
		}

		for swaps := uint(0); ; swaps++ {
			p = minDamageProgramOneSwap(p)
			damage := evaluateDamage(p)
			if damage == baseDamage {
				fmt.Printf("Case #%d: Impossible\n", i)
				break
			}
			if damage <= d {
				fmt.Printf("Case #%d: %d\n", i, swaps+1)
				break
			}
			baseDamage = damage
		}
	}
}

func minDamageProgramOneSwap(program string) string {
	minDamage := evaluateDamage(program)
	minDamageProgram := program
	for j := range program[:len(program)-1] {
		swapProgram := swapRunes(program, j, j+1)
		damage := evaluateDamage(swapProgram)
		if damage < minDamage {
			minDamage = damage
			minDamageProgram = swapProgram
		}
	}
	return minDamageProgram
}

func evaluateDamage(program string) uint {
	var totalDamage uint
	power := uint(1)
	for _, runeValue := range program {
		switch runeValue {
		case 'C':
			power *= 2
			break
		case 'S':
			totalDamage += power
			break
		}
	}
	return totalDamage
}

func substring(s string, start int, end int) string { // https://stackoverflow.com/a/38537764
	startStrIdx := 0
	i := 0
	for j := range s {
		if i == start {
			startStrIdx = j
		}
		if i == end {
			return s[startStrIdx:j]
		}
		i++
	}
	return s[startStrIdx:]
}

func swapRunes(input string, first int, second int) string {
	return substring(input, 0, first) + input[second:second] + substring(input, first+1, second) + input[first:first] + input[second+1:]
}
