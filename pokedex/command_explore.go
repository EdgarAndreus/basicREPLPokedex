package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)
func commandExplore(config *Config, args []string) error{
	if len(args) < 1 {
		return fmt.Errorf("please provide an area name")
	}

	var pokemon PokemonEncounter
	fullUrl := baseUrl + "/location-area/" + args[0]
	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pok := range pokemon.PokemonEncounters{
		fmt.Printf("- %v\n", pok.Pokemon.Name)
	}
	return nil
}