package nook

// Character is a composite type that combines various attributes of an Animal Crossing character.
type Character struct {
	// Animal represents the animal type of the character.
	Animal

	// Birthday contains the birthday information of the character.
	Birthday Birthday

	// Code is a unique identifier for the character.
	Code Code

	// Key is the character's canonical package key.
	// Keys are not guaranteed to be globally unique across animal types.
	Key Key

	// Gender represents the gender of the character.
	Gender Gender

	// Name contains the names of the character in different languages.
	Name Languages

	// Special indicates whether the character is special or has a unique role.
	Special bool
}
