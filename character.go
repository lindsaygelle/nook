package nook

// Character is the common information shared between all Animal Crossing characters, regardless of their role.
type Character struct {
	Animal

	Birthday Birthday
	Code     Code
	Gender   Gender
	Name     Languages
}
