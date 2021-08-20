package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	zippertbunnyBirthday = nook.Birthday{
		Day:   11,
		Month: time.March}
)

var (
	zippertbunnyCode = nook.Code{
		Value: "pyn"}
)

var (
	zippertbunnyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zipper T. Bunny"}

	zippertbunnyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Albin"}

	zippertbunnyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zipper"}

	zippertbunnyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Albin"}

	zippertbunnyGermanName = nook.Name{
		Language: language.German,
		Value:    "Ohs"}

	zippertbunnyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ovidio"}

	zippertbunnyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぴょんたろう"}

	zippertbunnyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토빗"}

	zippertbunnyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Coti Conejal"}

	zippertbunnyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Зиппер"}

	zippertbunnySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹦蹦"}

	zippertbunnySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Coti Conejal"}

	zippertbunnyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蹦蹦"}
)

var (
	zippertbunnyName = nook.Languages{
		language.AmericanEnglish:      zippertbunnyAmericanEnglishName,
		language.CanadianFrench:       zippertbunnyCanadianFrenchName,
		language.Dutch:                zippertbunnyDutchName,
		language.French:               zippertbunnyFrenchName,
		language.German:               zippertbunnyGermanName,
		language.Italian:              zippertbunnyItalianName,
		language.Japanese:             zippertbunnyJapaneseName,
		language.Korean:               zippertbunnyKoreanName,
		language.LatinAmericanSpanish: zippertbunnyLatinAmericanSpanishName,
		language.Russian:              zippertbunnyRussianName,
		language.SimplifiedChinese:    zippertbunnySimplifiedChineseName,
		language.Spanish:              zippertbunnySpanishName,
		language.TraditionalChinese:   zippertbunnyTraditionalChineseName}
)

var (
	zippertbunnyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: zippertbunnyBirthday,
		Code:     zippertbunnyCode,
		Key:      character.ZipperTBunny,
		Gender:   gender.Male,
		Name:     zippertbunnyName,
		Special:  true}
)

var (
	ZipperTBunny = nook.Resident{
		Character: zippertbunnyCharacter}
)
