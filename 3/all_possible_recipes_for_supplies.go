package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Result: %v\n", findAllRecipes([]string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}, [][]string{{"d"}, {"hveml", "f", "cpivl"}, {"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"}, {"cpivl", "hveml", "zpmcz", "ju", "h"}, {"h", "fzjnm", "e", "q", "x"}, {"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, {"f", "hveml", "cpivl"}}, []string{"f", "hveml", "cpivl", "d"}))
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	dependents := make(map[string][]string)
	outdegrees := make(map[string]int)
	for i := 0; i < len(recipes); i++ {
		for _, ingredient := range ingredients[i] {
			dependents[ingredient] = append(dependents[ingredient], recipes[i])
			outdegrees[recipes[i]]++
		}
	}

	recipesMap := make(map[string]struct{})
	for _, recipe := range recipes {
		recipesMap[recipe] = struct{}{}
	}

	var (
		res  []string
		curr string
	)

	for len(supplies) != 0 {
		curr, supplies = supplies[0], supplies[1:]
		if _, ok := recipesMap[curr]; ok {
			res = append(res, curr)
		}
		for _, dependent := range dependents[curr] {
			outdegrees[dependent]--
			if outdegrees[dependent] == 0 {
				supplies = append(supplies, dependent)
			}
		}
	}
	return res
}
