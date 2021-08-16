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
	reneeBirthday = nook.Birthday{
		Day:   28,
		Month: time.May}
)

var (
	reneeCode = nook.Code{
		Value: "rhn08"}
)

var (
	reneeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Renée"}

	reneeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rina"}

	reneeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Renée"}

	reneeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rina"}

	reneeGermanName = nook.Name{
		Language: language.German,
		Value:    "Ilona"}

	reneeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Renata"}

	reneeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おさい"}

	reneeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뿔님이"}

	reneeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rina"}

	reneeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рене"}

	reneeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "柴姐"}

	reneeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rina"}

	reneeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "柴姐"}
)

var (
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
	reneeCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: reneeBirthday,
		Code:     reneeCode,
		Key:      character.Renee,
		Gender:   gender.Female,
		Name:     reneeName}
)

var (
	reneeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yo yo yo"}

	reneeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "badaboum"}

	reneeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "yo yo yo"}

	reneeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "badaboum"}

	reneeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bomm"}

	reneeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "stomp"}

	reneeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パネェ"}

	reneeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "킁킁"}

	reneeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purrum"}

	reneeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "йоу"}

	reneeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "超强"}

	reneeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "purrum"}

	reneeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "超強"}
)

var (
	reneePhrase = nook.Languages{
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
	Renee = nook.Villager{
		Character:   reneeCharacter,
		Personality: personality.BigSister,
		Phrase:      reneePhrase}
)
