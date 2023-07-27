package nook

// Villager represents an Animal Crossing character that can be invited to live within the player's town, city, or island.
// Villagers are home owners and can befriend the player, allowing gift and decoration exchanges.
// Each Villager has a unique personality and set of phrases that accompany them.
type Villager struct {
	Character

	// Personality represents the personality type of the Villager.
	Personality Personality

	// Phrase contains the unique phrases associated with the Villager in different languages.
	Phrase Languages
}
