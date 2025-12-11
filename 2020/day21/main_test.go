package main

import "testing"

var input = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`

var foods, allIngredients = parse(input)

func TestPart1(t *testing.T) {
	expected := 5
	actual := part1(foods, allIngredients)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := "mxmxvkd,sqjhc,fvjkl"
	actual := part2(foods, allIngredients)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
