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
	// hornsbyBirthday represents Hornsby's birthday (March 20th).
	hornsbyBirthday = nook.Birthday{
		Day:   20,
		Month: time.March}
)

var (
	// hornsbyCode represents Hornsby's unique code ("rhn04").
	hornsbyCode = nook.Code{
		Value: "rhn04"}
)

var (
	// hornsbyAmericanEnglishName represents Hornsby's name in American English.
	hornsbyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hornsby"}

	// hornsbyCanadianFrenchName represents Hornsby's name in Canadian French.
	hornsbyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cornio"}

	// hornsbyDutchName represents Hornsby's name in Dutch.
	hornsbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hornsby"}

	// hornsbyFrenchName represents Hornsby's name in French.
	hornsbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cornio"}

	// hornsbyGermanName represents Hornsby's name in German.
	hornsbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Rüdiger"}

	// hornsbyItalianName represents Hornsby's name in Italian.
	hornsbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rino"}

	// hornsbyJapaneseName represents Hornsby's name in Japanese.
	hornsbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みつお"}

	// hornsbyKoreanName represents Hornsby's name in Korean.
	hornsbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뿌람"}

	// hornsbyLatinAmericanSpanishName represents Hornsby's name in Latin American Spanish.
	hornsbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rino"}

	// hornsbyRussianName represents Hornsby's name in Russian.
	hornsbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хорнсби"}

	// hornsbySimplifiedChineseName represents Hornsby's name in Simplified Chinese.
	hornsbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "光雄"}

	// hornsbySpanishName represents Hornsby's name in Spanish.
	hornsbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rino"}

	// hornsbyTraditionalChineseName represents Hornsby's name in Traditional Chinese.
	hornsbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "光雄"}
)

var (
	// hornsbyName represents Hornsby's name in different languages.
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
	// hornsbyCharacter represents Hornsby's character information.
	hornsbyCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: hornsbyBirthday,
		Code:     hornsbyCode,
		Key:      character.Hornsby,
		Gender:   gender.Male,
		Name:     hornsbyName,
		Special:  false}
)

var (
	// hornsbyAmericanEnglishPhrase represents Hornsby's phrase in American English.
	hornsbyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schnozzle"}

	// hornsbyCanadianFrenchPhrase represents Hornsby's phrase in Canadian French.
	hornsbyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "euh"}

	// hornsbyDutchPhrase represents Hornsby's phrase in Dutch.
	hornsbyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sneuzel"}

	// hornsbyFrenchPhrase represents Hornsby's phrase in French.
	hornsbyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "euh"}

	// hornsbyGermanPhrase represents Hornsby's phrase in German.
	hornsbyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schneuzel"}

	// hornsbyItalianPhrase represents Hornsby's phrase in Italian.
	hornsbyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splush"}

	// hornsbyJapanesePhrase represents Hornsby's phrase in Japanese.
	hornsbyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のネン"}

	// hornsbyKoreanPhrase represents Hornsby's phrase in Korean.
	hornsbyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿔뿔"}

	// hornsbyLatinAmericanSpanishPhrase represents Hornsby's phrase in Latin American Spanish.
	hornsbyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kaplut"}

	// hornsbyRussianPhrase represents Hornsby's phrase in Russian.
	hornsbyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шнобель"}

	// hornsbySimplifiedChinesePhrase represents Hornsby's phrase in Simplified Chinese.
	hornsbySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "念念"}

	// hornsbySpanishPhrase represents Hornsby's phrase in Spanish.
	hornsbySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "narigón"}

	// hornsbyTraditionalChinesePhrase represents Hornsby's phrase in Traditional Chinese.
	hornsbyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "念念"}
)

var (
	// hornsbyPhrase represents Hornsby's phrases in different languages.
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
	// Hornsby represents the character Hornsby with his complete information.
	Hornsby = nook.Villager{
		Character:   hornsbyCharacter,
		Personality: personality.Lazy,
		Phrase:      hornsbyPhrase}
)
