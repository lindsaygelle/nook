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
	// zippertbunnyBirthday represents zippertbunny birthday.
	zippertbunnyBirthday = nook.Birthday{
		Day:   11,
		Month: time.March}
)

var (
	// zippertbunnyCode represents zippertbunny code.
	zippertbunnyCode = nook.Code{
		Value: "pyn"}
)

var (
	// zippertbunnyAmericanEnglishName represents zippertbunny american english name.
	zippertbunnyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zipper T. Bunny"}

	// zippertbunnyCanadianFrenchName represents zippertbunny canadian french name.
	zippertbunnyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Albin"}

	// zippertbunnyDutchName represents zippertbunny dutch name.
	zippertbunnyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zipper"}

	// zippertbunnyFrenchName represents zippertbunny french name.
	zippertbunnyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Albin"}

	// zippertbunnyGermanName represents zippertbunny german name.
	zippertbunnyGermanName = nook.Name{
		Language: language.German,
		Value:    "Ohs"}

	// zippertbunnyItalianName represents zippertbunny italian name.
	zippertbunnyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ovidio"}

	// zippertbunnyJapaneseName represents zippertbunny japanese name.
	zippertbunnyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぴょんたろう"}

	// zippertbunnyKoreanName represents zippertbunny korean name.
	zippertbunnyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토빗"}

	// zippertbunnyLatinAmericanSpanishName represents zippertbunny latin american spanish name.
	zippertbunnyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Coti Conejal"}

	// zippertbunnyRussianName represents zippertbunny russian name.
	zippertbunnyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Зиппер"}

	// zippertbunnySimplifiedChineseName represents zippertbunny simplified chinese name.
	zippertbunnySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹦蹦"}

	// zippertbunnySpanishName represents zippertbunny spanish name.
	zippertbunnySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Coti Conejal"}

	// zippertbunnyTraditionalChineseName represents zippertbunny traditional chinese name.
	zippertbunnyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蹦蹦"}
)

var (
	// zippertbunnyName represents zippertbunny name.
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
	// zippertbunnyCharacter represents zippertbunny character.
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
	// ZipperTBunny represents zipper t bunny.
	ZipperTBunny = nook.Resident{
		Character: zippertbunnyCharacter}
)
