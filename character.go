package nook

// Character is a representation of the common attributes for Animal Crossing characters.
type Character struct {
	Animal

	Birthday Birthday
	Code     Code
	Key      Key
	Gender   Gender
	Name     Languages
	Special  bool
}
