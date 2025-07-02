package main

import (
	//"math/rand"
	"fmt"
	"net/http"
	"encoding/json"
)

func commandCatch(config *Config, args []string) error {
	fullUrl := baseUrl + "/pokemon/" + args[0]
	var pokemon Pokemon
	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	pokedex[pokemon.Name] = pokemon
	return nil
}
