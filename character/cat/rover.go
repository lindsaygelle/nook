package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	roverBirthday = nook.Birthday{
		Day:   1,
		Month: time.February}
)

var (
	roverCode = nook.Code{
		Value: "xct"}
)

var (
	roverAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rover"}

	roverCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Charly"}

	roverDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rover"}

	roverFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charly"}

	roverGermanName = nook.Name{
		Language: language.German,
		Value:    "Olli"}

	roverItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Girolamo"}

	roverJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みしらぬネコ"}

	roverKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "낯선고양이"}

	roverLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fran"}

	roverRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ровер"}

	roverSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "陌陌"}

	roverSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fran"}

	roverTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陌陌"}
)

var (
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
	roverCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: roverBirthday,
		Code:     roverCode,
		Gender:   gender.Male,
		Name:     roverName}
)

var (
	Rover = nook.Resident{
		Character: roverCharacter}
)
