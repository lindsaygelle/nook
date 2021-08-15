package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hornsbyBirthday = nook.Birthday{
		Day:   20,
		Month: time.March}
)

var (
	hornsbyCode = nook.Code{
		Value: "rhn04"}
)

var (
	hornsbyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hornsby"}

	hornsbyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cornioeuh"}

	hornsbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hornsbysneuzel"}

	hornsbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cornioeuh"}

	hornsbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Rüdigerschneuzel"}

	hornsbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rinosplush"}

	hornsbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みつおのネン"}

	hornsbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뿌람뿔뿔"}

	hornsbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rinokaplut"}

	hornsbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хорнсбишнобель"}

	hornsbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "光雄念念"}

	hornsbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rinonarigón"}

	hornsbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "光雄念念"}
)

var (
	hornsbyName = nook.Languages{
		language.AmericanEnglish:      hornsbyAmericanEnglishName,
		language.CanadianFrench:       hornsbyCanadianFrenchName,
		language.Dutch:                hornsbyDutchName,
		language.French:               hornsbyFrenchName,
		language.German:               hornsbyGermanName,
		language.Italian:              hornsbyItalianName,
		language.Japanese:             hornsbyJapaneseName,
		language.Korean:               hornsbyKoreanName,
		language.LatinAmericanSpanish: hornsbyLatinAmericanSpanishName,
		language.Russian:              hornsbyRussianName,
		language.SimplifiedChinese:    hornsbySimplifiedChineseName,
		language.Spanish:              hornsbySpanishName,
		language.TraditionalChinese:   hornsbyTraditionalChineseName}
)

var (
	hornsbyCharacter = nook.Character{
		Animal:   Rhinoceros,
		Birthday: hornsbyBirthday,
		Code:     hornsbyCode,
		Gender:   nook.Male,
		Name:     hornsbyName}
)

var (
	hornsbyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schnozzle"}

	hornsbyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "euh"}

	hornsbyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sneuzel"}

	hornsbyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "euh"}

	hornsbyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schneuzel"}

	hornsbyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splush"}

	hornsbyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のネン"}

	hornsbyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿔뿔"}

	hornsbyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kaplut"}

	hornsbyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шнобель"}

	hornsbySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "念念"}

	hornsbySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "narigón"}

	hornsbyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "念念"}
)

var (
	hornsbyPhrase = nook.Languages{
		language.AmericanEnglish:      hornsbyAmericanEnglishName,
		language.CanadianFrench:       hornsbyCanadianFrenchName,
		language.Dutch:                hornsbyDutchName,
		language.French:               hornsbyFrenchName,
		language.German:               hornsbyGermanName,
		language.Italian:              hornsbyItalianName,
		language.Japanese:             hornsbyJapaneseName,
		language.Korean:               hornsbyKoreanName,
		language.LatinAmericanSpanish: hornsbyLatinAmericanSpanishName,
		language.Russian:              hornsbyRussianName,
		language.SimplifiedChinese:    hornsbySimplifiedChineseName,
		language.Spanish:              hornsbySpanishName,
		language.TraditionalChinese:   hornsbyTraditionalChineseName}
)

var (
	Hornsby = nook.Villager{
		Character:   hornsbyCharacter,
		Personality: nook.Lazy,
		Phrase:      hornsbyPhrase}
)
