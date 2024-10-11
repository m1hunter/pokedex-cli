package main

import (
	"github.com/m1hunter/pokedex-cli/internal/commands"
	"github.com/m1hunter/pokedex-cli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
		CaughtPokemon: map[string]pokeapi.Pokemon{},
	}

	commands.StartRepl(cfg)
}
