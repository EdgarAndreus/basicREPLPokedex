package main

import (
	"github.com/EdgarAndreus/pokedexcli/pokecache"
)

type Config struct{
	Next string
	Previous string
	cache *pokecache.Cache
}
