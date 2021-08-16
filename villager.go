package nook

// Villager is the representation of an Animal Crossing character that can be invited to live within the players town, city or island.
// Characters that are Villagers can be home owners and befriend the player. The player can exchange gifts and decorations
// with any Villager character. All Villagers have personalities and unique phrases that accompany them.
// Examples of characters that are villagers are Tabby the cat, Snooty the wolf and Alfonso the alligator.
type Villager struct {
	Character

	Personality Personality
	Phrase      Languages
}
