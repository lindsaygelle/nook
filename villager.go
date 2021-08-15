package nook

// Villager is a type of Animal Crossing character that inhabits the players town, city or island.
type Villager struct {
	Character

	Personality Personality
	Phrase      Languages
}
