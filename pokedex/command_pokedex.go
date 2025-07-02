package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *Config, args []string) error {

	if len(args) != 0 {
		return errors.New("pokemont doesn't take arguments")
	}
	if len(pokedex) == 0 {
		return errors.New("pokedex is empty")
	}
	for _, value := range pokedex {
		fmt.Printf("- %v\n", value)
	}
	return nil
}
