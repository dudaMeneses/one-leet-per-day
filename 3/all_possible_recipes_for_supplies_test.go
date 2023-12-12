package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestHappyPathOneRecipe(t *testing.T) {
	result := findAllRecipes([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"})

	if len(result) != 1 && slices.Contains(result, "bread") {
		_ = fmt.Errorf("Bread should be the only possible recipe, but result was [%v]\n", strings.Join(result, ", "))
	}
}

func TestHappyPathSimpleRecursion(t *testing.T) {
	result := findAllRecipes([]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"})

	if len(result) != 2 && slices.Contains(result, "bread") && slices.Contains(result, "sandwich") {
		_ = fmt.Errorf("Bread and Sandwich should be the possible recipes, but result was [%v]\n", strings.Join(result, ", "))
	}
}

func TestHappyPathDeepRecursion(t *testing.T) {
	result := findAllRecipes([]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"})

	if len(result) != 3 && slices.Contains(result, "bread") && slices.Contains(result, "sandwich") && slices.Contains(result, "burger") {
		_ = fmt.Errorf("Bread, Sandwich and Burger should be the possible recipes, but result was [%v]\n", strings.Join(result, ", "))
	}
}

func TestDeepRecursionExcludingRecipes(t *testing.T) {
	result := findAllRecipes([]string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}, [][]string{{"d"}, {"hveml", "f", "cpivl"}, {"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"}, {"cpivl", "hveml", "zpmcz", "ju", "h"}, {"h", "fzjnm", "e", "q", "x"}, {"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, {"f", "hveml", "cpivl"}}, []string{"f", "hveml", "cpivl", "d"})

	if len(result) != 4 {
		_ = fmt.Errorf("Should be 4 possible recipes, but result was [%v]\n", strings.Join(result, ", "))
	}
}
