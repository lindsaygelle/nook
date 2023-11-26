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
	// francineBirthday represents Francine's birthday.
	francineBirthday = nook.Birthday{
		Day:   22,
		Month: time.January}
)

var (
	// francineCode represents Francine's unique code.
	francineCode = nook.Code{
		Value: "rbt12"}
)

var (
	// francineAmericanEnglishName represents Francine's name in American English.
	francineAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Francine"}

	// francineCanadianFrenchName represents Francine's name in Canadian French.
	francineCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nadine"}

	// francineDutchName represents Francine's name in Dutch.
	francineDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Francine"}

	// francineFrenchName represents Francine's name in French.
	francineFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadine"}

	// francineGermanName represents Francine's name in German.
	francineGermanName = nook.Name{
		Language: language.German,
		Value:    "Manu"}

	// francineItalianName represents Francine's name in Italian.
	francineItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Franca"}

	// francineJapaneseName represents Francine's name in Japanese.
	francineJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランソワ"}

	// francineKoreanName represents Francine's name in Korean.
	francineKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랑소와"}

	// francineLatinAmericanSpanishName represents Francine's name in Latin American Spanish.
	francineLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natacha"}

	// francineRussianName represents Francine's name in Russian.
	francineRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франсин"}

	// francineSimplifiedChineseName represents Francine's name in Simplified Chinese.
	francineSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "法蓝琪"}

	// francineSpanishName represents Francine's name in Spanish.
	francineSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natacha"}

	// francineTraditionalChineseName represents Francine's name in Traditional Chinese.
	francineTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "法藍琪"}
)

var (
	// francineName represents Francine's name in different languages.
	francineName = nook.Languages{
		language.AmericanEnglish:      francineAmericanEnglishName,
		language.CanadianFrench:       francineCanadianFrenchName,
		language.Dutch:                francineDutchName,
		language.French:               francineFrenchName,
		language.German:               francineGermanName,
		language.Italian:              francineItalianName,
		language.Japanese:             francineJapaneseName,
		language.Korean:               francineKoreanName,
		language.LatinAmericanSpanish: francineLatinAmericanSpanishName,
		language.Russian:              francineRussianName,
		language.SimplifiedChinese:    francineSimplifiedChineseName,
		language.Spanish:              francineSpanishName,
		language.TraditionalChinese:   francineTraditionalChineseName}
)

var (
	// francineCharacter represents Francine's character information.
	francineCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: francineBirthday,
		Code:     francineCode,
		Key:      character.Francine,
		Gender:   gender.Female,
		Name:     francineName,
		Special:  false}
)

var (
	// francineAmericanEnglishPhrase represents Francine's phrase in American English.
	francineAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "karat"}

	// francineCanadianFrenchPhrase represents Francine's phrase in Canadian French.
	francineCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "carotte"}

	// francineDutchPhrase represents Francine's phrase in Dutch.
	francineDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wortels"}

	// francineFrenchPhrase represents Francine's phrase in French.
	francineFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "carotte"}

	// francineGermanPhrase represents Francine's phrase in German.
	francineGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hey man"}

	// francineItalianPhrase represents Francine's phrase in Italian.
	francineItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lulalà"}

	// francineJapanesePhrase represents Francine's phrase in Japanese.
	francineJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルララ"}

	// francineKoreanPhrase represents Francine's phrase in Korean.
	francineKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쇼봉"}

	// francineLatinAmericanSpanishPhrase represents Francine's phrase in Latin American Spanish.
	francineLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kaninchen"}

	// francineRussianPhrase represents Francine's phrase in Russian.
	francineRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "морковка"}

	// francineSimplifiedChinesePhrase represents Francine's phrase in Simplified Chinese.
	francineSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噜啦啦"}

	// francineSpanishPhrase represents Francine's phrase in Spanish.
	francineSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kaninchen"}

	// francineTraditionalChinesePhrase represents Francine's phrase in Traditional Chinese.
	francineTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嚕啦啦"}
)

var (
	// francinePhrase represents Francine's phrases in different languages.
	francinePhrase = nook.Languages{
		language.AmericanEnglish:      francineAmericanEnglishPhrase,
		language.CanadianFrench:       francineCanadianFrenchPhrase,
		language.Dutch:                francineDutchPhrase,
		language.French:               francineFrenchPhrase,
		language.German:               francineGermanPhrase,
		language.Italian:              francineItalianPhrase,
		language.Japanese:             francineJapanesePhrase,
		language.Korean:               francineKoreanPhrase,
		language.LatinAmericanSpanish: francineLatinAmericanSpanishPhrase,
		language.Russian:              francineRussianPhrase,
		language.SimplifiedChinese:    francineSimplifiedChinesePhrase,
		language.Spanish:              francineSpanishPhrase,
		language.TraditionalChinese:   francineTraditionalChinesePhrase}
)

var (
	// Francine represents the character Francine with her complete information.
	Francine = nook.Villager{
		Character:   francineCharacter,
		Personality: personality.Snooty,
		Phrase:      francinePhrase}
)
