package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// roverBirthday represents rover birthday.
	roverBirthday = nook.Birthday{
		Day:   1,
		Month: time.February}
)

var (
	// roverCode represents rover code.
	roverCode = nook.Code{
		Value: "xct"}
)

var (
	// roverAmericanEnglishName represents rover american english name.
	roverAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rover"}

	// roverCanadianFrenchName represents rover canadian french name.
	roverCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Charly"}

	// roverDutchName represents rover dutch name.
	roverDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rover"}

	// roverFrenchName represents rover french name.
	roverFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charly"}

	// roverGermanName represents rover german name.
	roverGermanName = nook.Name{
		Language: language.German,
		Value:    "Olli"}

	// roverItalianName represents rover italian name.
	roverItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Girolamo"}

	// roverJapaneseName represents rover japanese name.
	roverJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みしらぬネコ"}

	// roverKoreanName represents rover korean name.
	roverKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "낯선고양이"}

	// roverLatinAmericanSpanishName represents rover latin american spanish name.
	roverLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fran"}

	// roverRussianName represents rover russian name.
	roverRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ровер"}

	// roverSimplifiedChineseName represents rover simplified chinese name.
	roverSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "陌陌"}

	// roverSpanishName represents rover spanish name.
	roverSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fran"}

	// roverTraditionalChineseName represents rover traditional chinese name.
	roverTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陌陌"}
)

var (
	// roverName represents rover name.
	roverName = nook.Languages{
		language.AmericanEnglish:      roverAmericanEnglishName,
		language.CanadianFrench:       roverCanadianFrenchName,
		language.Dutch:                roverDutchName,
		language.French:               roverFrenchName,
		language.German:               roverGermanName,
		language.Italian:              roverItalianName,
		language.Japanese:             roverJapaneseName,
		language.Korean:               roverKoreanName,
		language.LatinAmericanSpanish: roverLatinAmericanSpanishName,
		language.Russian:              roverRussianName,
		language.SimplifiedChinese:    roverSimplifiedChineseName,
		language.Spanish:              roverSpanishName,
		language.TraditionalChinese:   roverTraditionalChineseName}
)

var (
	// roverCharacter represents rover character.
	roverCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: roverBirthday,
		Code:     roverCode,
		Key:      character.Rover,
		Gender:   gender.Male,
		Name:     roverName,
		Special:  true}
)

var (
	// Rover represents rover.
	Rover = nook.Resident{
		Character: roverCharacter}
)
