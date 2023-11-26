package wolf

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
	// chiefBirthday represents Chief's birthday.
	chiefBirthday = nook.Birthday{
		Day:   19,
		Month: time.December}
)

var (
	// chiefCode represents Chief's unique code.
	chiefCode = nook.Code{
		Value: "wol00"}
)

var (
	// chiefAmericanEnglishName represents Chief's name in American English.
	chiefAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chief"}

	// chiefCanadianFrenchName represents Chief's name in Canadian French.
	chiefCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chef"}

	// chiefDutchName represents Chief's name in Dutch.
	chiefDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chief"}

	// chiefFrenchName represents Chief's name in French.
	chiefFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chef"}

	// chiefGermanName represents Chief's name in German.
	chiefGermanName = nook.Name{
		Language: language.German,
		Value:    "Sascha"}

	// chiefItalianName represents Chief's name in Italian.
	chiefItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Artiglio"}

	// chiefJapaneseName represents Chief's name in Japanese.
	chiefJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チーフ"}

	// chiefKoreanName represents Chief's name in Korean.
	chiefKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대장"}

	// chiefLatinAmericanSpanishName represents Chief's name in Latin American Spanish.
	chiefLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zoilo"}

	// chiefRussianName represents Chief's name in Russian.
	chiefRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чиф"}

	// chiefSimplifiedChineseName represents Chief's name in Simplified Chinese.
	chiefSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施傅"}

	// chiefSpanishName represents Chief's name in Spanish.
	chiefSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zoilo"}

	// chiefTraditionalChineseName represents Chief's name in Traditional Chinese.
	chiefTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施傅"}
)

var (
	// chiefName represents Chief's name in different languages.
	chiefName = nook.Languages{
		language.AmericanEnglish:      chiefAmericanEnglishName,
		language.CanadianFrench:       chiefCanadianFrenchName,
		language.Dutch:                chiefDutchName,
		language.French:               chiefFrenchName,
		language.German:               chiefGermanName,
		language.Italian:              chiefItalianName,
		language.Japanese:             chiefJapaneseName,
		language.Korean:               chiefKoreanName,
		language.LatinAmericanSpanish: chiefLatinAmericanSpanishName,
		language.Russian:              chiefRussianName,
		language.SimplifiedChinese:    chiefSimplifiedChineseName,
		language.Spanish:              chiefSpanishName,
		language.TraditionalChinese:   chiefTraditionalChineseName}
)

var (
	// chiefCharacter represents Chief's character information.
	chiefCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: chiefBirthday,
		Code:     chiefCode,
		Key:      character.Chief,
		Gender:   gender.Male,
		Name:     chiefName,
		Special:  false}
)

var (
	// chiefAmericanEnglishPhrase represents Chief's phrase in American English.
	chiefAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "harrumph"}

	// chiefCanadianFrenchPhrase represents Chief's phrase in Canadian French.
	chiefCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mmmph"}

	// chiefDutchPhrase represents Chief's phrase in Dutch.
	chiefDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "huil"}

	// chiefFrenchPhrase represents Chief's phrase in French.
	chiefFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mmmph"}

	// chiefGermanPhrase represents Chief's phrase in German.
	chiefGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "harrumff"}

	// chiefItalianPhrase represents Chief's phrase in Italian.
	chiefItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arrrgh"}

	// chiefJapanesePhrase represents Chief's phrase in Japanese.
	chiefJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やんか"}

	// chiefKoreanPhrase represents Chief's phrase in Korean.
	chiefKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇잖여"}

	// chiefLatinAmericanSpanishPhrase represents Chief's phrase in Latin American Spanish.
	chiefLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ejem"}

	// chiefRussianPhrase represents Chief's phrase in Russian.
	chiefRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фыр-р-рф"}

	// chiefSimplifiedChinesePhrase represents Chief's phrase in Simplified Chinese.
	chiefSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不是嘛"}

	// chiefSpanishPhrase represents Chief's phrase in Spanish.
	chiefSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ejem"}

	// chiefTraditionalChinesePhrase represents Chief's phrase in Traditional Chinese.
	chiefTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不是嘛"}
)

var (
	// chiefPhrase represents Chief's phrase in different languages.
	chiefPhrase = nook.Languages{
		language.AmericanEnglish:      chiefAmericanEnglishPhrase,
		language.CanadianFrench:       chiefCanadianFrenchPhrase,
		language.Dutch:                chiefDutchPhrase,
		language.French:               chiefFrenchPhrase,
		language.German:               chiefGermanPhrase,
		language.Italian:              chiefItalianPhrase,
		language.Japanese:             chiefJapanesePhrase,
		language.Korean:               chiefKoreanPhrase,
		language.LatinAmericanSpanish: chiefLatinAmericanSpanishPhrase,
		language.Russian:              chiefRussianPhrase,
		language.SimplifiedChinese:    chiefSimplifiedChinesePhrase,
		language.Spanish:              chiefSpanishPhrase,
		language.TraditionalChinese:   chiefTraditionalChinesePhrase}
)

var (
	// Chief represents the character Chief with his complete information.
	Chief = nook.Villager{
		Character:   chiefCharacter,
		Personality: personality.Cranky,
		Phrase:      chiefPhrase}
)
