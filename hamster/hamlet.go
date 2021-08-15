package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hamletBirthday = nook.Birthday{
		Day:   30,
		Month: time.May}
)

var (
	hamletCode = nook.Code{
		Value: "ham00"}
)

var (
	hamletAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hamlet"}

	hamletCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jojochicots"}

	hamletDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hamlethammie"}

	hamletFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jojochicots"}

	hamletGermanName = nook.Name{
		Language: language.German,
		Value:    "Hamidtwirps"}

	hamletItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amletopuff pant"}

	hamletJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムスケハム"}

	hamletKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄스틴햄햄"}

	hamletLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bombochucufá"}

	hamletRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гамлетхомячок"}

	hamletSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈姆火腿"}

	hamletSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bombochucufá"}

	hamletTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈姆火腿"}
)

var (
	hamletName = nook.Languages{
		language.AmericanEnglish:      hamletAmericanEnglishName,
		language.CanadianFrench:       hamletCanadianFrenchName,
		language.Dutch:                hamletDutchName,
		language.French:               hamletFrenchName,
		language.German:               hamletGermanName,
		language.Italian:              hamletItalianName,
		language.Japanese:             hamletJapaneseName,
		language.Korean:               hamletKoreanName,
		language.LatinAmericanSpanish: hamletLatinAmericanSpanishName,
		language.Russian:              hamletRussianName,
		language.SimplifiedChinese:    hamletSimplifiedChineseName,
		language.Spanish:              hamletSpanishName,
		language.TraditionalChinese:   hamletTraditionalChineseName}
)

var (
	hamletCharacter = nook.Character{
		Animal:   Hamster,
		Birthday: hamletBirthday,
		Code:     hamletCode,
		Gender:   nook.Male,
		Name:     hamletName}
)

var (
	hamletAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hammie"}

	hamletCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chicots"}

	hamletDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hammie"}

	hamletFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chicots"}

	hamletGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "twirps"}

	hamletItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "puff pant"}

	hamletJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハム"}

	hamletKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "햄햄"}

	hamletLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chucufá"}

	hamletRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хомячок"}

	hamletSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "火腿"}

	hamletSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chucufá"}

	hamletTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "火腿"}
)

var (
	hamletPhrase = nook.Languages{
		language.AmericanEnglish:      hamletAmericanEnglishName,
		language.CanadianFrench:       hamletCanadianFrenchName,
		language.Dutch:                hamletDutchName,
		language.French:               hamletFrenchName,
		language.German:               hamletGermanName,
		language.Italian:              hamletItalianName,
		language.Japanese:             hamletJapaneseName,
		language.Korean:               hamletKoreanName,
		language.LatinAmericanSpanish: hamletLatinAmericanSpanishName,
		language.Russian:              hamletRussianName,
		language.SimplifiedChinese:    hamletSimplifiedChineseName,
		language.Spanish:              hamletSpanishName,
		language.TraditionalChinese:   hamletTraditionalChineseName}
)

var (
	Hamlet = nook.Villager{
		Character:   hamletCharacter,
		Personality: nook.Jock,
		Phrase:      hamletPhrase}
)
