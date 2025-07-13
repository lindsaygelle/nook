package lion

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
	// azizBirthday represents aziz birthday.
	azizBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// azizCode represents aziz code.
	azizCode = nook.Code{
		Value: ""}
)

var (
	// azizAmericanEnglishName represents aziz american english name.
	azizAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aziz"}

	// azizCanadianFrenchName represents aziz canadian french name.
	azizCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// azizDutchName represents aziz dutch name.
	azizDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// azizFrenchName represents aziz french name.
	azizFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aziz"}

	// azizGermanName represents aziz german name.
	azizGermanName = nook.Name{
		Language: language.German,
		Value:    "Daniel"}

	// azizItalianName represents aziz italian name.
	azizItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lionello"}

	// azizJapaneseName represents aziz japanese name.
	azizJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アッサム"}

	// azizKoreanName represents aziz korean name.
	azizKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// azizLatinAmericanSpanishName represents aziz latin american spanish name.
	azizLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// azizRussianName represents aziz russian name.
	azizRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// azizSimplifiedChineseName represents aziz simplified chinese name.
	azizSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿三"}

	// azizSpanishName represents aziz spanish name.
	azizSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leonardo"}

	// azizTraditionalChineseName represents aziz traditional chinese name.
	azizTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// azizName represents aziz name.
	azizName = nook.Languages{
		language.AmericanEnglish:      azizAmericanEnglishName,
		language.CanadianFrench:       azizCanadianFrenchName,
		language.Dutch:                azizDutchName,
		language.French:               azizFrenchName,
		language.German:               azizGermanName,
		language.Italian:              azizItalianName,
		language.Japanese:             azizJapaneseName,
		language.Korean:               azizKoreanName,
		language.LatinAmericanSpanish: azizLatinAmericanSpanishName,
		language.Russian:              azizRussianName,
		language.SimplifiedChinese:    azizSimplifiedChineseName,
		language.Spanish:              azizSpanishName,
		language.TraditionalChinese:   azizTraditionalChineseName}
)

var (
	// azizCharacter represents aziz character.
	azizCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: azizBirthday,
		Code:     azizCode,
		Key:      character.Aziz,
		Gender:   gender.Male,
		Name:     azizName,
		Special:  false}
)

var (
	// azizAmericanEnglishPhrase represents aziz american english phrase.
	azizAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "RAWR"}

	// azizCanadianFrenchPhrase represents aziz canadian french phrase.
	azizCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// azizDutchPhrase represents aziz dutch phrase.
	azizDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// azizFrenchPhrase represents aziz french phrase.
	azizFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zigoto"}

	// azizGermanPhrase represents aziz german phrase.
	azizGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "RRAAHH"}

	// azizItalianPhrase represents aziz italian phrase.
	azizItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ROAR"}

	// azizJapanesePhrase represents aziz japanese phrase.
	azizJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カンジー"}

	// azizKoreanPhrase represents aziz korean phrase.
	azizKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// azizLatinAmericanSpanishPhrase represents aziz latin american spanish phrase.
	azizLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// azizRussianPhrase represents aziz russian phrase.
	azizRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// azizSimplifiedChinesePhrase represents aziz simplified chinese phrase.
	azizSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗尔"}

	// azizSpanishPhrase represents aziz spanish phrase.
	azizSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aborigen"}

	// azizTraditionalChinesePhrase represents aziz traditional chinese phrase.
	azizTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// azizPhrase represents aziz phrase.
	azizPhrase = nook.Languages{
		language.AmericanEnglish:      azizAmericanEnglishPhrase,
		language.CanadianFrench:       azizCanadianFrenchPhrase,
		language.Dutch:                azizDutchPhrase,
		language.French:               azizFrenchPhrase,
		language.German:               azizGermanPhrase,
		language.Italian:              azizItalianPhrase,
		language.Japanese:             azizJapanesePhrase,
		language.Korean:               azizKoreanPhrase,
		language.LatinAmericanSpanish: azizLatinAmericanSpanishPhrase,
		language.Russian:              azizRussianPhrase,
		language.SimplifiedChinese:    azizSimplifiedChinesePhrase,
		language.Spanish:              azizSpanishPhrase,
		language.TraditionalChinese:   azizTraditionalChinesePhrase}
)

var (
	// Aziz represents aziz.
	Aziz = nook.Villager{
		Character:   azizCharacter,
		Personality: personality.Jock,
		Phrase:      azizPhrase}
)
