package rabbit

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
	// claudeBirthday represents Claude's birthday.
	claudeBirthday = nook.Birthday{
		Day:   3,
		Month: time.December}
)

var (
	// claudeCode represents Claude's unique code ("rbt11").
	claudeCode = nook.Code{
		Value: "rbt11"}
)

var (
	// claudeAmericanEnglishName represents Claude's name in American English.
	claudeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Claude"}

	// claudeCanadianFrenchName represents Claude's name in Canadian French.
	claudeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Claude"}

	// claudeDutchName represents Claude's name in Dutch.
	claudeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Claude"}

	// claudeFrenchName represents Claude's name in French.
	claudeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Claude"}

	// claudeGermanName represents Claude's name in German.
	claudeGermanName = nook.Name{
		Language: language.German,
		Value:    "Claude"}

	// claudeItalianName represents Claude's name in Italian.
	claudeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Claude"}

	// claudeJapaneseName represents Claude's name in Japanese.
	claudeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビネガー"}

	// claudeKoreanName represents Claude's name in Korean.
	claudeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비니거"}

	// claudeLatinAmericanSpanishName represents Claude's name in Latin American Spanish.
	claudeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pablo"}

	// claudeRussianName represents Claude's name in Russian.
	claudeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клод"}

	// claudeSimplifiedChineseName represents Claude's name in Simplified Chinese.
	claudeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿醋"}

	// claudeSpanishName represents Claude's name in Spanish.
	claudeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pablo"}

	// claudeTraditionalChineseName represents Claude's name in Traditional Chinese.
	claudeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿醋"}
)

var (
	// claudeName represents Claude's name in different languages.
	claudeName = nook.Languages{
		language.AmericanEnglish:      claudeAmericanEnglishName,
		language.CanadianFrench:       claudeCanadianFrenchName,
		language.Dutch:                claudeDutchName,
		language.French:               claudeFrenchName,
		language.German:               claudeGermanName,
		language.Italian:              claudeItalianName,
		language.Japanese:             claudeJapaneseName,
		language.Korean:               claudeKoreanName,
		language.LatinAmericanSpanish: claudeLatinAmericanSpanishName,
		language.Russian:              claudeRussianName,
		language.SimplifiedChinese:    claudeSimplifiedChineseName,
		language.Spanish:              claudeSpanishName,
		language.TraditionalChinese:   claudeTraditionalChineseName}
)

var (
	// claudeCharacter represents Claude's character information.
	claudeCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: claudeBirthday,
		Code:     claudeCode,
		Key:      character.Claude,
		Gender:   gender.Male,
		Name:     claudeName,
		Special:  false}
)

var (
	// claudeAmericanEnglishPhrase represents Claude's phrase in American English.
	claudeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hopalong"}

	// claudeCanadianFrenchPhrase represents Claude's phrase in Canadian French.
	claudeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sans rire"}

	// claudeDutchPhrase represents Claude's phrase in Dutch.
	claudeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hopsala"}

	// claudeFrenchPhrase represents Claude's phrase in French.
	claudeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sans rire"}

	// claudeGermanPhrase represents Claude's phrase in German.
	claudeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüpfauf"}

	// claudeItalianPhrase represents Claude's phrase in Italian.
	claudeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hoppela"}

	// claudeJapanesePhrase represents Claude's phrase in Japanese.
	claudeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぶいぶい"}

	// claudeKoreanPhrase represents Claude's phrase in Korean.
	claudeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아으셔"}

	// claudeLatinAmericanSpanishPhrase represents Claude's phrase in Latin American Spanish.
	claudeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hópala"}

	// claudeRussianPhrase represents Claude's phrase in Russian.
	claudeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "скок-поскок"}

	// claudeSimplifiedChinesePhrase represents Claude's phrase in Simplified Chinese.
	claudeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酸酸"}

	// claudeSpanishPhrase represents Claude's phrase in Spanish.
	claudeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hópala"}

	// claudeTraditionalChinesePhrase represents Claude's phrase in Traditional Chinese.
	claudeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "酸酸"}
)

var (
	// claudePhrase represents Claude's phrase in different languages.
	claudePhrase = nook.Languages{
		language.AmericanEnglish:      claudeAmericanEnglishPhrase,
		language.CanadianFrench:       claudeCanadianFrenchPhrase,
		language.Dutch:                claudeDutchPhrase,
		language.French:               claudeFrenchPhrase,
		language.German:               claudeGermanPhrase,
		language.Italian:              claudeItalianPhrase,
		language.Japanese:             claudeJapanesePhrase,
		language.Korean:               claudeKoreanPhrase,
		language.LatinAmericanSpanish: claudeLatinAmericanSpanishPhrase,
		language.Russian:              claudeRussianPhrase,
		language.SimplifiedChinese:    claudeSimplifiedChinesePhrase,
		language.Spanish:              claudeSpanishPhrase,
		language.TraditionalChinese:   claudeTraditionalChinesePhrase}
)

var (
	// Claude represents the character Claude with his complete information.
	Claude = nook.Villager{
		Character:   claudeCharacter,
		Personality: personality.Lazy,
		Phrase:      claudePhrase}
)
