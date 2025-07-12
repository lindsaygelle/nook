package chameleon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// flickBirthday represents flick birthday.
	flickBirthday = nook.Birthday{
		Day:   10,
		Month: time.May}
)

var (
	// flickCode represents flick code.
	flickCode = nook.Code{
		Value: "chy"}
)

var (
	// flickAmericanEnglishName represents flick american english name.
	flickAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flick"}

	// flickCanadianFrenchName represents flick canadian french name.
	flickCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Djason"}

	// flickDutchName represents flick dutch name.
	flickDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flick"}

	// flickFrenchName represents flick french name.
	flickFrenchName = nook.Name{
		Language: language.French,
		Value:    "Djason"}

	// flickGermanName represents flick german name.
	flickGermanName = nook.Name{
		Language: language.German,
		Value:    "Carlson"}

	// flickItalianName represents flick italian name.
	flickItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ivano"}

	// flickJapaneseName represents flick japanese name.
	flickJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レックス"}

	// flickKoreanName represents flick korean name.
	flickKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레온"}

	// flickLatinAmericanSpanishName represents flick latin american spanish name.
	flickLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kamilo"}

	// flickRussianName represents flick russian name.
	flickRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флик"}

	// flickSimplifiedChineseName represents flick simplified chinese name.
	flickSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "龙克斯"}

	// flickSpanishName represents flick spanish name.
	flickSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kamilo"}

	// flickTraditionalChineseName represents flick traditional chinese name.
	flickTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龍克斯"}
)

var (
	// flickName represents flick name.
	flickName = nook.Languages{
		language.AmericanEnglish:      flickAmericanEnglishName,
		language.CanadianFrench:       flickCanadianFrenchName,
		language.Dutch:                flickDutchName,
		language.French:               flickFrenchName,
		language.German:               flickGermanName,
		language.Italian:              flickItalianName,
		language.Japanese:             flickJapaneseName,
		language.Korean:               flickKoreanName,
		language.LatinAmericanSpanish: flickLatinAmericanSpanishName,
		language.Russian:              flickRussianName,
		language.SimplifiedChinese:    flickSimplifiedChineseName,
		language.Spanish:              flickSpanishName,
		language.TraditionalChinese:   flickTraditionalChineseName}
)

var (
	// flickCharacter represents flick character.
	flickCharacter = nook.Character{
		Animal:   animal.Chameleon,
		Birthday: flickBirthday,
		Code:     flickCode,
		Key:      character.Flick,
		Gender:   gender.Male,
		Name:     flickName,
		Special:  true}
)

var (
	// Flick represents flick.
	Flick = nook.Resident{
		Character: flickCharacter}
)
