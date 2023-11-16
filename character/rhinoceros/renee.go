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
	// reneeBirthday represents Renée's birthday (May 28th).
	reneeBirthday = nook.Birthday{
		Day:   28,
		Month: time.May}
)

var (
	// reneeCode represents Renée's unique code ("rhn08").
	reneeCode = nook.Code{
		Value: "rhn08"}
)

var (
	// reneeAmericanEnglishName represents Renée's name in American English.
	reneeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Renée"}

	// reneeCanadianFrenchName represents Renée's name in Canadian French.
	reneeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rina"}

	// reneeDutchName represents Renée's name in Dutch.
	reneeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Renée"}

	// reneeFrenchName represents Renée's name in French.
	reneeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rina"}

	// reneeGermanName represents Renée's name in German.
	reneeGermanName = nook.Name{
		Language: language.German,
		Value:    "Ilona"}

	// reneeItalianName represents Renée's name in Italian.
	reneeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Renata"}

	// reneeJapaneseName represents Renée's name in Japanese.
	reneeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おさい"}

	// reneeKoreanName represents Renée's name in Korean.
	reneeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뿔님이"}

	// reneeLatinAmericanSpanishName represents Renée's name in Latin American Spanish.
	reneeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rina"}

	// reneeRussianName represents Renée's name in Russian.
	reneeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рене"}

	// reneeSimplifiedChineseName represents Renée's name in Simplified Chinese.
	reneeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "柴姐"}

	// reneeSpanishName represents Renée's name in Spanish.
	reneeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rina"}

	// reneeTraditionalChineseName represents Renée's name in Traditional Chinese.
	reneeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "柴姐"}
)

var (
	// reneeName represents Renée's name in different languages.
	reneeName = nook.Languages{
		language.AmericanEnglish:      reneeAmericanEnglishName,
		language.CanadianFrench:       reneeCanadianFrenchName,
		language.Dutch:                reneeDutchName,
		language.French:               reneeFrenchName,
		language.German:               reneeGermanName,
		language.Italian:              reneeItalianName,
		language.Japanese:             reneeJapaneseName,
		language.Korean:               reneeKoreanName,
		language.LatinAmericanSpanish: reneeLatinAmericanSpanishName,
		language.Russian:              reneeRussianName,
		language.SimplifiedChinese:    reneeSimplifiedChineseName,
		language.Spanish:              reneeSpanishName,
		language.TraditionalChinese:   reneeTraditionalChineseName}
)

var (
	// reneeCharacter represents Renée's character information.
	reneeCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: reneeBirthday,
		Code:     reneeCode,
		Key:      character.Renee,
		Gender:   gender.Female,
		Name:     reneeName,
		Special:  false}
)

var (
	// reneeAmericanEnglishPhrase represents Renée's phrase in American English.
	reneeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yo yo yo"}

	// reneeCanadianFrenchPhrase represents Renée's phrase in Canadian French.
	reneeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "badaboum"}

	// reneeDutchPhrase represents Renée's phrase in Dutch.
	reneeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "yo yo yo"}

	// reneeFrenchPhrase represents Renée's phrase in French.
	reneeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "badaboum"}

	// reneeGermanPhrase represents Renée's phrase in German.
	reneeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bomm"}

	// reneeItalianPhrase represents Renée's phrase in Italian.
	reneeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "stomp"}

	// reneeJapanesePhrase represents Renée's phrase in Japanese.
	reneeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パネェ"}

	// reneeKoreanPhrase represents Renée's phrase in Korean.
	reneeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "킁킁"}

	// reneeLatinAmericanSpanishPhrase represents Renée's phrase in Latin American Spanish.
	reneeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purrum"}

	// reneeRussianPhrase represents Renée's phrase in Russian.
	reneeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "йоу"}

	// reneeSimplifiedChinesePhrase represents Renée's phrase in Simplified Chinese.
	reneeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "超强"}

	// reneeSpanishPhrase represents Renée's phrase in Spanish.
	reneeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "purrum"}

	// reneeTraditionalChinesePhrase represents Renée's phrase in Traditional Chinese.
	reneeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "超強"}
)

var (
	// reneePhrase represents Renée's phrases in different languages.
	reneePhrase = nook.Languages{
		language.AmericanEnglish:      reneeAmericanEnglishPhrase,
		language.CanadianFrench:       reneeCanadianFrenchPhrase,
		language.Dutch:                reneeDutchPhrase,
		language.French:               reneeFrenchPhrase,
		language.German:               reneeGermanPhrase,
		language.Italian:              reneeItalianPhrase,
		language.Japanese:             reneeJapanesePhrase,
		language.Korean:               reneeKoreanPhrase,
		language.LatinAmericanSpanish: reneeLatinAmericanSpanishPhrase,
		language.Russian:              reneeRussianPhrase,
		language.SimplifiedChinese:    reneeSimplifiedChinesePhrase,
		language.Spanish:              reneeSpanishPhrase,
		language.TraditionalChinese:   reneeTraditionalChinesePhrase}
)

var (
	// Renee represents the character Renée with her complete information.
	Renee = nook.Villager{
		Character:   reneeCharacter,
		Personality: personality.BigSister,
		Phrase:      reneePhrase}
)
