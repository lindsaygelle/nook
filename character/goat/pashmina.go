package goat

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
	// pashminaBirthday represents pashmina birthday.
	pashminaBirthday = nook.Birthday{
		Day:   26,
		Month: time.December}
)

var (
	// pashminaCode represents pashmina code.
	pashminaCode = nook.Code{
		Value: "goa08"}
)

var (
	// pashminaAmericanEnglishName represents pashmina american english name.
	pashminaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pashmina"}

	// pashminaCanadianFrenchName represents pashmina canadian french name.
	pashminaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chavrina"}

	// pashminaDutchName represents pashmina dutch name.
	pashminaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pashmina"}

	// pashminaFrenchName represents pashmina french name.
	pashminaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chavrina"}

	// pashminaGermanName represents pashmina german name.
	pashminaGermanName = nook.Name{
		Language: language.German,
		Value:    "Pamela"}

	// pashminaItalianName represents pashmina italian name.
	pashminaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pashmina"}

	// pashminaJapaneseName represents pashmina japanese name.
	pashminaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バーバラ"}

	// pashminaKoreanName represents pashmina korean name.
	pashminaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바바라"}

	// pashminaLatinAmericanSpanishName represents pashmina latin american spanish name.
	pashminaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carla"}

	// pashminaRussianName represents pashmina russian name.
	pashminaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пашмина"}

	// pashminaSimplifiedChineseName represents pashmina simplified chinese name.
	pashminaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "芭芭拉"}

	// pashminaSpanishName represents pashmina spanish name.
	pashminaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carla"}

	// pashminaTraditionalChineseName represents pashmina traditional chinese name.
	pashminaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "芭芭拉"}
)

var (
	// pashminaName represents pashmina name.
	pashminaName = nook.Languages{
		language.AmericanEnglish:      pashminaAmericanEnglishName,
		language.CanadianFrench:       pashminaCanadianFrenchName,
		language.Dutch:                pashminaDutchName,
		language.French:               pashminaFrenchName,
		language.German:               pashminaGermanName,
		language.Italian:              pashminaItalianName,
		language.Japanese:             pashminaJapaneseName,
		language.Korean:               pashminaKoreanName,
		language.LatinAmericanSpanish: pashminaLatinAmericanSpanishName,
		language.Russian:              pashminaRussianName,
		language.SimplifiedChinese:    pashminaSimplifiedChineseName,
		language.Spanish:              pashminaSpanishName,
		language.TraditionalChinese:   pashminaTraditionalChineseName}
)

var (
	// pashminaCharacter represents pashmina character.
	pashminaCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: pashminaBirthday,
		Code:     pashminaCode,
		Key:      character.Pashmina,
		Gender:   gender.Female,
		Name:     pashminaName,
		Special:  false}
)

var (
	// pashminaAmericanEnglishPhrase represents pashmina american english phrase.
	pashminaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kidders"}

	// pashminaCanadianFrenchPhrase represents pashmina canadian french phrase.
	pashminaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon moulin"}

	// pashminaDutchPhrase represents pashmina dutch phrase.
	pashminaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bokkig"}

	// pashminaFrenchPhrase represents pashmina french phrase.
	pashminaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon moulin"}

	// pashminaGermanPhrase represents pashmina german phrase.
	pashminaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heckmeck"}

	// pashminaItalianPhrase represents pashmina italian phrase.
	pashminaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "embeeeh"}

	// pashminaJapanesePhrase represents pashmina japanese phrase.
	pashminaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "かな"}

	// pashminaKoreanPhrase represents pashmina korean phrase.
	pashminaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러게"}

	// pashminaLatinAmericanSpanishPhrase represents pashmina latin american spanish phrase.
	pashminaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ne-eh"}

	// pashminaRussianPhrase represents pashmina russian phrase.
	pashminaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "козлятки"}

	// pashminaSimplifiedChinesePhrase represents pashmina simplified chinese phrase.
	pashminaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "或许芭"}

	// pashminaSpanishPhrase represents pashmina spanish phrase.
	pashminaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muy bieeen"}

	// pashminaTraditionalChinesePhrase represents pashmina traditional chinese phrase.
	pashminaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "或許芭"}
)

var (
	// pashminaPhrase represents pashmina phrase.
	pashminaPhrase = nook.Languages{
		language.AmericanEnglish:      pashminaAmericanEnglishPhrase,
		language.CanadianFrench:       pashminaCanadianFrenchPhrase,
		language.Dutch:                pashminaDutchPhrase,
		language.French:               pashminaFrenchPhrase,
		language.German:               pashminaGermanPhrase,
		language.Italian:              pashminaItalianPhrase,
		language.Japanese:             pashminaJapanesePhrase,
		language.Korean:               pashminaKoreanPhrase,
		language.LatinAmericanSpanish: pashminaLatinAmericanSpanishPhrase,
		language.Russian:              pashminaRussianPhrase,
		language.SimplifiedChinese:    pashminaSimplifiedChinesePhrase,
		language.Spanish:              pashminaSpanishPhrase,
		language.TraditionalChinese:   pashminaTraditionalChinesePhrase}
)

var (
	// Pashmina represents pashmina.
	Pashmina = nook.Villager{
		Character:   pashminaCharacter,
		Personality: personality.BigSister,
		Phrase:      pashminaPhrase}
)
