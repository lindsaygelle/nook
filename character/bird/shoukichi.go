package bird

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
	// shoukichiBirthday represents shoukichi birthday.
	shoukichiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// shoukichiCode represents shoukichi code.
	shoukichiCode = nook.Code{
		Value: ""}
)

var (
	// shoukichiAmericanEnglishName represents shoukichi american english name.
	shoukichiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shoukichi"}

	// shoukichiCanadianFrenchName represents shoukichi canadian french name.
	shoukichiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// shoukichiDutchName represents shoukichi dutch name.
	shoukichiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// shoukichiFrenchName represents shoukichi french name.
	shoukichiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// shoukichiGermanName represents shoukichi german name.
	shoukichiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// shoukichiItalianName represents shoukichi italian name.
	shoukichiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// shoukichiJapaneseName represents shoukichi japanese name.
	shoukichiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しょうきち"}

	// shoukichiKoreanName represents shoukichi korean name.
	shoukichiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// shoukichiLatinAmericanSpanishName represents shoukichi latin american spanish name.
	shoukichiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// shoukichiRussianName represents shoukichi russian name.
	shoukichiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// shoukichiSimplifiedChineseName represents shoukichi simplified chinese name.
	shoukichiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// shoukichiSpanishName represents shoukichi spanish name.
	shoukichiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// shoukichiTraditionalChineseName represents shoukichi traditional chinese name.
	shoukichiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// shoukichiName represents shoukichi name.
	shoukichiName = nook.Languages{
		language.AmericanEnglish:      shoukichiAmericanEnglishName,
		language.CanadianFrench:       shoukichiCanadianFrenchName,
		language.Dutch:                shoukichiDutchName,
		language.French:               shoukichiFrenchName,
		language.German:               shoukichiGermanName,
		language.Italian:              shoukichiItalianName,
		language.Japanese:             shoukichiJapaneseName,
		language.Korean:               shoukichiKoreanName,
		language.LatinAmericanSpanish: shoukichiLatinAmericanSpanishName,
		language.Russian:              shoukichiRussianName,
		language.SimplifiedChinese:    shoukichiSimplifiedChineseName,
		language.Spanish:              shoukichiSpanishName,
		language.TraditionalChinese:   shoukichiTraditionalChineseName}
)

var (
	// shoukichiCharacter represents shoukichi character.
	shoukichiCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: shoukichiBirthday,
		Code:     shoukichiCode,
		Key:      character.Shoukichi,
		Gender:   gender.Male,
		Name:     shoukichiName,
		Special:  false}
)

var (
	// shoukichiAmericanEnglishPhrase represents shoukichi american english phrase.
	shoukichiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ズン"}

	// shoukichiCanadianFrenchPhrase represents shoukichi canadian french phrase.
	shoukichiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// shoukichiDutchPhrase represents shoukichi dutch phrase.
	shoukichiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// shoukichiFrenchPhrase represents shoukichi french phrase.
	shoukichiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// shoukichiGermanPhrase represents shoukichi german phrase.
	shoukichiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// shoukichiItalianPhrase represents shoukichi italian phrase.
	shoukichiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// shoukichiJapanesePhrase represents shoukichi japanese phrase.
	shoukichiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ズン"}

	// shoukichiKoreanPhrase represents shoukichi korean phrase.
	shoukichiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// shoukichiLatinAmericanSpanishPhrase represents shoukichi latin american spanish phrase.
	shoukichiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// shoukichiRussianPhrase represents shoukichi russian phrase.
	shoukichiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// shoukichiSimplifiedChinesePhrase represents shoukichi simplified chinese phrase.
	shoukichiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// shoukichiSpanishPhrase represents shoukichi spanish phrase.
	shoukichiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// shoukichiTraditionalChinesePhrase represents shoukichi traditional chinese phrase.
	shoukichiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// shoukichiPhrase represents shoukichi phrase.
	shoukichiPhrase = nook.Languages{
		language.AmericanEnglish:      shoukichiAmericanEnglishPhrase,
		language.CanadianFrench:       shoukichiCanadianFrenchPhrase,
		language.Dutch:                shoukichiDutchPhrase,
		language.French:               shoukichiFrenchPhrase,
		language.German:               shoukichiGermanPhrase,
		language.Italian:              shoukichiItalianPhrase,
		language.Japanese:             shoukichiJapanesePhrase,
		language.Korean:               shoukichiKoreanPhrase,
		language.LatinAmericanSpanish: shoukichiLatinAmericanSpanishPhrase,
		language.Russian:              shoukichiRussianPhrase,
		language.SimplifiedChinese:    shoukichiSimplifiedChinesePhrase,
		language.Spanish:              shoukichiSpanishPhrase,
		language.TraditionalChinese:   shoukichiTraditionalChinesePhrase}
)

var (
	// Shoukichi represents shoukichi.
	Shoukichi = nook.Villager{
		Character:   shoukichiCharacter,
		Personality: personality.Jock,
		Phrase:      shoukichiPhrase}
)
