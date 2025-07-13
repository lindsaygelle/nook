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
	// kaitlinBirthday represents kaitlin birthday.
	kaitlinBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// kaitlinCode represents kaitlin code.
	kaitlinCode = nook.Code{
		Value: "mum"}
)

var (
	// kaitlinAmericanEnglishName represents kaitlin american english name.
	kaitlinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kaitlin"}

	// kaitlinCanadianFrenchName represents kaitlin canadian french name.
	kaitlinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cathou"}

	// kaitlinDutchName represents kaitlin dutch name.
	kaitlinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kaitlin"}

	// kaitlinFrenchName represents kaitlin french name.
	kaitlinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cathou"}

	// kaitlinGermanName represents kaitlin german name.
	kaitlinGermanName = nook.Name{
		Language: language.German,
		Value:    "Kitty"}

	// kaitlinItalianName represents kaitlin italian name.
	kaitlinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Micia"}

	// kaitlinJapaneseName represents kaitlin japanese name.
	kaitlinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おかあさん"}

	// kaitlinKoreanName represents kaitlin korean name.
	kaitlinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엄마"}

	// kaitlinLatinAmericanSpanishName represents kaitlin latin american spanish name.
	kaitlinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Catiana"}

	// kaitlinRussianName represents kaitlin russian name.
	kaitlinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кейтлин"}

	// kaitlinSimplifiedChineseName represents kaitlin simplified chinese name.
	kaitlinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// kaitlinSpanishName represents kaitlin spanish name.
	kaitlinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Catiana"}

	// kaitlinTraditionalChineseName represents kaitlin traditional chinese name.
	kaitlinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// kaitlinName represents kaitlin name.
	kaitlinName = nook.Languages{
		language.AmericanEnglish:      kaitlinAmericanEnglishName,
		language.CanadianFrench:       kaitlinCanadianFrenchName,
		language.Dutch:                kaitlinDutchName,
		language.French:               kaitlinFrenchName,
		language.German:               kaitlinGermanName,
		language.Italian:              kaitlinItalianName,
		language.Japanese:             kaitlinJapaneseName,
		language.Korean:               kaitlinKoreanName,
		language.LatinAmericanSpanish: kaitlinLatinAmericanSpanishName,
		language.Russian:              kaitlinRussianName,
		language.SimplifiedChinese:    kaitlinSimplifiedChineseName,
		language.Spanish:              kaitlinSpanishName,
		language.TraditionalChinese:   kaitlinTraditionalChineseName}
)

var (
	// kaitlinCharacter represents kaitlin character.
	kaitlinCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kaitlinBirthday,
		Code:     kaitlinCode,
		Key:      character.Kaitlin,
		Gender:   gender.Female,
		Name:     kaitlinName,
		Special:  true}
)

var (
	// Kaitlin represents kaitlin.
	Kaitlin = nook.Resident{
		Character: kaitlinCharacter}
)
