package rhinoceros

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
		Value:    "Cornio"}

	hornsbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hornsby"}

	hornsbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cornio"}

	hornsbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Rüdiger"}

	hornsbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rino"}

	hornsbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みつお"}

	hornsbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뿌람"}

	hornsbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rino"}

	hornsbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хорнсби"}

	hornsbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "光雄"}

	hornsbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rino"}

	hornsbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "光雄"}
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
		Animal:   animal.Rhinoceros,
		Birthday: hornsbyBirthday,
		Code:     hornsbyCode,
		Key:      character.Hornsby,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      hornsbyAmericanEnglishPhrase,
		language.CanadianFrench:       hornsbyCanadianFrenchPhrase,
		language.Dutch:                hornsbyDutchPhrase,
		language.French:               hornsbyFrenchPhrase,
		language.German:               hornsbyGermanPhrase,
		language.Italian:              hornsbyItalianPhrase,
		language.Japanese:             hornsbyJapanesePhrase,
		language.Korean:               hornsbyKoreanPhrase,
		language.LatinAmericanSpanish: hornsbyLatinAmericanSpanishPhrase,
		language.Russian:              hornsbyRussianPhrase,
		language.SimplifiedChinese:    hornsbySimplifiedChinesePhrase,
		language.Spanish:              hornsbySpanishPhrase,
		language.TraditionalChinese:   hornsbyTraditionalChinesePhrase}
)

var (
	Hornsby = nook.Villager{
		Character:   hornsbyCharacter,
		Personality: personality.Lazy,
		Phrase:      hornsbyPhrase}
)
