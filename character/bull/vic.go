package bull

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
	// vicBirthday represents vic birthday.
	vicBirthday = nook.Birthday{
		Day:   29,
		Month: time.December}
)

var (
	// vicCode represents vic code.
	vicCode = nook.Code{
		Value: "bul08"}
)

var (
	// vicAmericanEnglishName represents vic american english name.
	vicAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vic"}

	// vicCanadianFrenchName represents vic canadian french name.
	vicCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Toto"}

	// vicDutchName represents vic dutch name.
	vicDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vic"}

	// vicFrenchName represents vic french name.
	vicFrenchName = nook.Name{
		Language: language.French,
		Value:    "Toto"}

	// vicGermanName represents vic german name.
	vicGermanName = nook.Name{
		Language: language.German,
		Value:    "Klaus"}

	// vicItalianName represents vic italian name.
	vicItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Paco"}

	// vicJapaneseName represents vic japanese name.
	vicJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ノルマン"}

	// vicKoreanName represents vic korean name.
	vicKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "노르망"}

	// vicLatinAmericanSpanishName represents vic latin american spanish name.
	vicLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Artorito"}

	// vicRussianName represents vic russian name.
	vicRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вик"}

	// vicSimplifiedChineseName represents vic simplified chinese name.
	vicSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卢尔曼"}

	// vicSpanishName represents vic spanish name.
	vicSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Artorito"}

	// vicTraditionalChineseName represents vic traditional chinese name.
	vicTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "盧爾曼"}
)

var (
	// vicName represents vic name.
	vicName = nook.Languages{
		language.AmericanEnglish:      vicAmericanEnglishName,
		language.CanadianFrench:       vicCanadianFrenchName,
		language.Dutch:                vicDutchName,
		language.French:               vicFrenchName,
		language.German:               vicGermanName,
		language.Italian:              vicItalianName,
		language.Japanese:             vicJapaneseName,
		language.Korean:               vicKoreanName,
		language.LatinAmericanSpanish: vicLatinAmericanSpanishName,
		language.Russian:              vicRussianName,
		language.SimplifiedChinese:    vicSimplifiedChineseName,
		language.Spanish:              vicSpanishName,
		language.TraditionalChinese:   vicTraditionalChineseName}
)

var (
	// vicCharacter represents vic character.
	vicCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: vicBirthday,
		Code:     vicCode,
		Key:      character.Vic,
		Gender:   gender.Male,
		Name:     vicName,
		Special:  false}
)

var (
	// vicAmericanEnglishPhrase represents vic american english phrase.
	vicAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cud"}

	// vicCanadianFrenchPhrase represents vic canadian french phrase.
	vicCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "moumoul"}

	// vicDutchPhrase represents vic dutch phrase.
	vicDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kauw"}

	// vicFrenchPhrase represents vic french phrase.
	vicFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "moumoul"}

	// vicGermanPhrase represents vic german phrase.
	vicGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grumpf"}

	// vicItalianPhrase represents vic italian phrase.
	vicItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trullalà"}

	// vicJapanesePhrase represents vic japanese phrase.
	vicJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "モイ"}

	// vicKoreanPhrase represents vic korean phrase.
	vicKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뭐임"}

	// vicLatinAmericanSpanishPhrase represents vic latin american spanish phrase.
	vicLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brufff"}

	// vicRussianPhrase represents vic russian phrase.
	vicRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "полундра"}

	// vicSimplifiedChinesePhrase represents vic simplified chinese phrase.
	vicSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哞"}

	// vicSpanishPhrase represents vic spanish phrase.
	vicSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muuubién"}

	// vicTraditionalChinesePhrase represents vic traditional chinese phrase.
	vicTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哞"}
)

var (
	// vicPhrase represents vic phrase.
	vicPhrase = nook.Languages{
		language.AmericanEnglish:      vicAmericanEnglishPhrase,
		language.CanadianFrench:       vicCanadianFrenchPhrase,
		language.Dutch:                vicDutchPhrase,
		language.French:               vicFrenchPhrase,
		language.German:               vicGermanPhrase,
		language.Italian:              vicItalianPhrase,
		language.Japanese:             vicJapanesePhrase,
		language.Korean:               vicKoreanPhrase,
		language.LatinAmericanSpanish: vicLatinAmericanSpanishPhrase,
		language.Russian:              vicRussianPhrase,
		language.SimplifiedChinese:    vicSimplifiedChinesePhrase,
		language.Spanish:              vicSpanishPhrase,
		language.TraditionalChinese:   vicTraditionalChinesePhrase}
)

var (
	// Vic represents vic.
	Vic = nook.Villager{
		Character:   vicCharacter,
		Personality: personality.Cranky,
		Phrase:      vicPhrase}
)
