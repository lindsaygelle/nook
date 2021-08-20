package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Jojo"}

	hamletDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hamlet"}

	hamletFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jojo"}

	hamletGermanName = nook.Name{
		Language: language.German,
		Value:    "Hamid"}

	hamletItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amleto"}

	hamletJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムスケ"}

	hamletKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄스틴"}

	hamletLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bombo"}

	hamletRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гамлет"}

	hamletSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈姆"}

	hamletSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bombo"}

	hamletTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈姆"}
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
		Animal:   animal.Hamster,
		Birthday: hamletBirthday,
		Code:     hamletCode,
		Key:      character.Hamlet,
		Gender:   gender.Male,
		Name:     hamletName,
		Special:  false}
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
		language.AmericanEnglish:      hamletAmericanEnglishPhrase,
		language.CanadianFrench:       hamletCanadianFrenchPhrase,
		language.Dutch:                hamletDutchPhrase,
		language.French:               hamletFrenchPhrase,
		language.German:               hamletGermanPhrase,
		language.Italian:              hamletItalianPhrase,
		language.Japanese:             hamletJapanesePhrase,
		language.Korean:               hamletKoreanPhrase,
		language.LatinAmericanSpanish: hamletLatinAmericanSpanishPhrase,
		language.Russian:              hamletRussianPhrase,
		language.SimplifiedChinese:    hamletSimplifiedChinesePhrase,
		language.Spanish:              hamletSpanishPhrase,
		language.TraditionalChinese:   hamletTraditionalChinesePhrase}
)

var (
	Hamlet = nook.Villager{
		Character:   hamletCharacter,
		Personality: personality.Jock,
		Phrase:      hamletPhrase}
)
