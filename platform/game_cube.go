package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// gameCube is the common reference for the Nintendo GameCube platform.
	gameCube = "GameCube"
)

var (
	// gameCubeNameAmericanEnglish represents the Nintendo GameCube platform's name in American English.
	gameCubeNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nintendo GameCube",
	}
)

var (
	// gameCubeName contains the localized names of the Nintendo GameCube platform.
	gameCubeName = nook.Languages{
		language.AmericanEnglish: gameCubeNameAmericanEnglish,
	}
)

var (
	// GameCube represents the Nintendo GameCube platform in the Animal Crossing series.
	GameCube = nook.Platform{
		Key:  nook.Key(gameCube),
		Name: gameCubeName,
	}
)
