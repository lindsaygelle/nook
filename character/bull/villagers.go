package bull

import (
	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
)

// Villagers is map of villager instances.
var (
	// Villagers represents villagers.
	Villagers = nook.Villagers{
		character.Angus:  Angus,
		character.Chuck:  Chuck,
		character.Coach:  Coach,
		character.Oxford: Oxford,
		character.Rodeo:  Rodeo,
		character.Stu:    Stu,
		character.TBone:  TBone,
		character.Vic:    Vic,
		character.Weldon: Weldon}
)
