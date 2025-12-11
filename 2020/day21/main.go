package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   []string
}

func parse(input string) ([]Food, map[string]int) {
	var foods []Food
	allIngredients := make(map[string]int)

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, " (contains ")

		ingredients := strings.Fields(parts[0])
		for _, ing := range ingredients {
			allIngredients[ing]++
		}

		var allergens []string
		if len(parts) > 1 {
			allergenStr := strings.TrimSuffix(parts[1], ")")
			allergens = strings.Split(allergenStr, ", ")
		}

		foods = append(foods, Food{ingredients, allergens})
	}

	return foods, allIngredients
}

func getAllAllergens(foods []Food) map[string]bool {
	result := make(map[string]bool)
	for _, food := range foods {
		for _, a := range food.allergens {
			result[a] = true
		}
	}
	return result
}

func couldContain(ingredient, allergen string, foods []Food) bool {
	for _, food := range foods {
		hasAllergen := false
		for _, a := range food.allergens {
			if a == allergen {
				hasAllergen = true
				break
			}
		}
		if !hasAllergen {
			continue
		}
		hasIngredient := false
		for _, ing := range food.ingredients {
			if ing == ingredient {
				hasIngredient = true
				break
			}
		}
		if !hasIngredient {
			return false
		}
	}
	return true
}

func isUsed(ing string, resolved map[string]string) bool {
	for _, v := range resolved {
		if v == ing {
			return true
		}
	}
	return false
}

func part1(foods []Food, allIngredients map[string]int) int {
	canContainAllergen := make(map[string]bool)
	for ing := range allIngredients {
		for allergen := range getAllAllergens(foods) {
			if couldContain(ing, allergen, foods) {
				canContainAllergen[ing] = true
				break
			}
		}
	}

	count := 0
	for ing, appearances := range allIngredients {
		if !canContainAllergen[ing] {
			count += appearances
		}
	}
	return count
}

func part2(foods []Food, allIngredients map[string]int) string {
	allergenToIng := make(map[string]string)
	allAllergens := getAllAllergens(foods)

	for len(allergenToIng) < len(allAllergens) {
		for allergen := range allAllergens {
			if _, resolved := allergenToIng[allergen]; resolved {
				continue
			}
			var candidates []string
			for ing := range allIngredients {
				if isUsed(ing, allergenToIng) {
					continue
				}
				if couldContain(ing, allergen, foods) {
					candidates = append(candidates, ing)
				}
			}
			if len(candidates) == 1 {
				allergenToIng[allergen] = candidates[0]
			}
		}
	}

	var sortedAllergens []string
	for a := range allergenToIng {
		sortedAllergens = append(sortedAllergens, a)
	}
	sort.Strings(sortedAllergens)

	var dangerous []string
	for _, a := range sortedAllergens {
		dangerous = append(dangerous, allergenToIng[a])
	}
	return strings.Join(dangerous, ",")
}

func main() {
	data, _ := os.ReadFile("input.txt")
	foods, allIngredients := parse(string(data))

	fmt.Println("Part 1:", part1(foods, allIngredients))
	fmt.Println("Part 2:", part2(foods, allIngredients))
}
