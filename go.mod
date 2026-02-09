module github.com/vohrr/pokedexcli

go 1.25.6

replace github.com/vohrr/pokeapi v0.0.0 => ./internal/pokeapi

replace github.com/vohrr/pokecache v0.0.0 => ./internal/pokecache

require github.com/vohrr/pokeapi v0.0.0

require github.com/vohrr/pokecache v0.0.0
