package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kaitlinBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	kaitlinCode = nook.Code{
		Value: "mum"}
)

var (
	kaitlinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kaitlin"}

	kaitlinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cathou"}

	kaitlinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kaitlin"}

	kaitlinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cathou"}

	kaitlinGermanName = nook.Name{
		Language: language.German,
		Value:    "Kitty"}

	kaitlinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Micia"}

	kaitlinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おかあさん"}

	kaitlinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엄마"}

	kaitlinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Catiana"}

	kaitlinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кейтлин"}

	kaitlinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	kaitlinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Catiana"}

	kaitlinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
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
	kaitlinCharacter = nook.Character{
		Animal:   Cat,
		Birthday: kaitlinBirthday,
		Code:     kaitlinCode,
		Gender:   nook.Female,
		Name:     kaitlinName}
)

var (
	Kaitlin = nook.Resident{
		Character: kaitlinCharacter}
)
