package eagle

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
	// quetzalBirthday represents quetzal birthday.
	quetzalBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// quetzalCode represents quetzal code.
	quetzalCode = nook.Code{
		Value: ""}
)

var (
	// quetzalAmericanEnglishName represents quetzal american english name.
	quetzalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Quetzal"}

	// quetzalCanadianFrenchName represents quetzal canadian french name.
	quetzalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// quetzalDutchName represents quetzal dutch name.
	quetzalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// quetzalFrenchName represents quetzal french name.
	quetzalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Quetzal"}

	// quetzalGermanName represents quetzal german name.
	quetzalGermanName = nook.Name{
		Language: language.German,
		Value:    "Enrique"}

	// quetzalItalianName represents quetzal italian name.
	quetzalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Quetzal"}

	// quetzalJapaneseName represents quetzal japanese name.
	quetzalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハチェット"}

	// quetzalKoreanName represents quetzal korean name.
	quetzalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// quetzalLatinAmericanSpanishName represents quetzal latin american spanish name.
	quetzalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// quetzalRussianName represents quetzal russian name.
	quetzalRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// quetzalSimplifiedChineseName represents quetzal simplified chinese name.
	quetzalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "战斧"}

	// quetzalSpanishName represents quetzal spanish name.
	quetzalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Quetzal"}

	// quetzalTraditionalChineseName represents quetzal traditional chinese name.
	quetzalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// quetzalName represents quetzal name.
	quetzalName = nook.Languages{
		language.AmericanEnglish:      quetzalAmericanEnglishName,
		language.CanadianFrench:       quetzalCanadianFrenchName,
		language.Dutch:                quetzalDutchName,
		language.French:               quetzalFrenchName,
		language.German:               quetzalGermanName,
		language.Italian:              quetzalItalianName,
		language.Japanese:             quetzalJapaneseName,
		language.Korean:               quetzalKoreanName,
		language.LatinAmericanSpanish: quetzalLatinAmericanSpanishName,
		language.Russian:              quetzalRussianName,
		language.SimplifiedChinese:    quetzalSimplifiedChineseName,
		language.Spanish:              quetzalSpanishName,
		language.TraditionalChinese:   quetzalTraditionalChineseName}
)

var (
	// quetzalCharacter represents quetzal character.
	quetzalCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: quetzalBirthday,
		Code:     quetzalCode,
		Key:      character.Quetzal,
		Gender:   gender.Male,
		Name:     quetzalName,
		Special:  false}
)

var (
	// quetzalAmericanEnglishPhrase represents quetzal american english phrase.
	quetzalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "SKREEE"}

	// quetzalCanadianFrenchPhrase represents quetzal canadian french phrase.
	quetzalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// quetzalDutchPhrase represents quetzal dutch phrase.
	quetzalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// quetzalFrenchPhrase represents quetzal french phrase.
	quetzalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "couêêêtzal"}

	// quetzalGermanPhrase represents quetzal german phrase.
	quetzalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "KRIKRI"}

	// quetzalItalianPhrase represents quetzal italian phrase.
	quetzalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "OCCHIO"}

	// quetzalJapanesePhrase represents quetzal japanese phrase.
	quetzalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゲロッパ"}

	// quetzalKoreanPhrase represents quetzal korean phrase.
	quetzalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// quetzalLatinAmericanSpanishPhrase represents quetzal latin american spanish phrase.
	quetzalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// quetzalRussianPhrase represents quetzal russian phrase.
	quetzalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// quetzalSimplifiedChinesePhrase represents quetzal simplified chinese phrase.
	quetzalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳喳"}

	// quetzalSpanishPhrase represents quetzal spanish phrase.
	quetzalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuiii"}

	// quetzalTraditionalChinesePhrase represents quetzal traditional chinese phrase.
	quetzalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// quetzalPhrase represents quetzal phrase.
	quetzalPhrase = nook.Languages{
		language.AmericanEnglish:      quetzalAmericanEnglishPhrase,
		language.CanadianFrench:       quetzalCanadianFrenchPhrase,
		language.Dutch:                quetzalDutchPhrase,
		language.French:               quetzalFrenchPhrase,
		language.German:               quetzalGermanPhrase,
		language.Italian:              quetzalItalianPhrase,
		language.Japanese:             quetzalJapanesePhrase,
		language.Korean:               quetzalKoreanPhrase,
		language.LatinAmericanSpanish: quetzalLatinAmericanSpanishPhrase,
		language.Russian:              quetzalRussianPhrase,
		language.SimplifiedChinese:    quetzalSimplifiedChinesePhrase,
		language.Spanish:              quetzalSpanishPhrase,
		language.TraditionalChinese:   quetzalTraditionalChinesePhrase}
)

var (
	// Quetzal represents quetzal.
	Quetzal = nook.Villager{
		Character:   quetzalCharacter,
		Personality: personality.Jock,
		Phrase:      quetzalPhrase}
)
