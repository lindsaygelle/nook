package chicken

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
	// rhodaBirthday represents rhoda birthday.
	rhodaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// rhodaCode represents rhoda code.
	rhodaCode = nook.Code{
		Value: ""}
)

var (
	// rhodaAmericanEnglishName represents rhoda american english name.
	rhodaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhoda"}

	// rhodaCanadianFrenchName represents rhoda canadian french name.
	rhodaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Arlette"}

	// rhodaDutchName represents rhoda dutch name.
	rhodaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// rhodaFrenchName represents rhoda french name.
	rhodaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Arlette"}

	// rhodaGermanName represents rhoda german name.
	rhodaGermanName = nook.Name{
		Language: language.German,
		Value:    "Helene"}

	// rhodaItalianName represents rhoda italian name.
	rhodaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Petula"}

	// rhodaJapaneseName represents rhoda japanese name.
	rhodaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブイヨン"}

	// rhodaKoreanName represents rhoda korean name.
	rhodaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// rhodaLatinAmericanSpanishName represents rhoda latin american spanish name.
	rhodaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pochola"}

	// rhodaRussianName represents rhoda russian name.
	rhodaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// rhodaSimplifiedChineseName represents rhoda simplified chinese name.
	rhodaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布布"}

	// rhodaSpanishName represents rhoda spanish name.
	rhodaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pochola"}

	// rhodaTraditionalChineseName represents rhoda traditional chinese name.
	rhodaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// rhodaName represents rhoda name.
	rhodaName = nook.Languages{
		language.AmericanEnglish:      rhodaAmericanEnglishName,
		language.CanadianFrench:       rhodaCanadianFrenchName,
		language.Dutch:                rhodaDutchName,
		language.French:               rhodaFrenchName,
		language.German:               rhodaGermanName,
		language.Italian:              rhodaItalianName,
		language.Japanese:             rhodaJapaneseName,
		language.Korean:               rhodaKoreanName,
		language.LatinAmericanSpanish: rhodaLatinAmericanSpanishName,
		language.Russian:              rhodaRussianName,
		language.SimplifiedChinese:    rhodaSimplifiedChineseName,
		language.Spanish:              rhodaSpanishName,
		language.TraditionalChinese:   rhodaTraditionalChineseName}
)

var (
	// rhodaCharacter represents rhoda character.
	rhodaCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: rhodaBirthday,
		Code:     rhodaCode,
		Key:      character.Rhoda,
		Gender:   gender.Female,
		Name:     rhodaName,
		Special:  false}
)

var (
	// rhodaAmericanEnglishPhrase represents rhoda american english phrase.
	rhodaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "clucky"}

	// rhodaCanadianFrenchPhrase represents rhoda canadian french phrase.
	rhodaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	// rhodaDutchPhrase represents rhoda dutch phrase.
	rhodaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// rhodaFrenchPhrase represents rhoda french phrase.
	rhodaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	// rhodaGermanPhrase represents rhoda german phrase.
	rhodaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "glucker"}

	// rhodaItalianPhrase represents rhoda italian phrase.
	rhodaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "skiok"}

	// rhodaJapanesePhrase represents rhoda japanese phrase.
	rhodaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "コッコ"}

	// rhodaKoreanPhrase represents rhoda korean phrase.
	rhodaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// rhodaLatinAmericanSpanishPhrase represents rhoda latin american spanish phrase.
	rhodaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	// rhodaRussianPhrase represents rhoda russian phrase.
	rhodaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// rhodaSimplifiedChinesePhrase represents rhoda simplified chinese phrase.
	rhodaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	// rhodaSpanishPhrase represents rhoda spanish phrase.
	rhodaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	// rhodaTraditionalChinesePhrase represents rhoda traditional chinese phrase.
	rhodaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// rhodaPhrase represents rhoda phrase.
	rhodaPhrase = nook.Languages{
		language.AmericanEnglish:      rhodaAmericanEnglishPhrase,
		language.CanadianFrench:       rhodaCanadianFrenchPhrase,
		language.Dutch:                rhodaDutchPhrase,
		language.French:               rhodaFrenchPhrase,
		language.German:               rhodaGermanPhrase,
		language.Italian:              rhodaItalianPhrase,
		language.Japanese:             rhodaJapanesePhrase,
		language.Korean:               rhodaKoreanPhrase,
		language.LatinAmericanSpanish: rhodaLatinAmericanSpanishPhrase,
		language.Russian:              rhodaRussianPhrase,
		language.SimplifiedChinese:    rhodaSimplifiedChinesePhrase,
		language.Spanish:              rhodaSpanishPhrase,
		language.TraditionalChinese:   rhodaTraditionalChinesePhrase}
)

var (
	// Rhoda represents rhoda.
	Rhoda = nook.Villager{
		Character:   rhodaCharacter,
		Personality: personality.Snooty,
		Phrase:      rhodaPhrase}
)
