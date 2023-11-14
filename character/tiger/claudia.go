package tiger

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
	// claudiaBirthday represents Claudia's birthday (November 22nd).
	claudiaBirthday = nook.Birthday{
		Day:   22,
		Month: time.November}
)

var (
	// claudiaCode represents Claudia's unique code ("tig05").
	claudiaCode = nook.Code{
		Value: "tig05"}
)

var (
	// claudiaAmericanEnglishName represents Claudia's name in American English.
	claudiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Claudia"}

	// claudiaCanadianFrenchName represents Claudia's name in Canadian French.
	claudiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vanina"}

	// claudiaDutchName represents Claudia's name in Dutch.
	claudiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Claudia"}

	// claudiaFrenchName represents Claudia's name in French.
	claudiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vanina"}

	// claudiaGermanName represents Claudia's name in German.
	claudiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lilly"}

	// claudiaItalianName represents Claudia's name in Italian.
	claudiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lilla"}

	// claudiaJapaneseName represents Claudia's name in Japanese.
	claudiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリリン"}

	// claudiaKoreanName represents Claudia's name in Korean.
	claudiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "신디"}

	// claudiaLatinAmericanSpanishName represents Claudia's name in Latin American Spanish.
	claudiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lilu"}

	// claudiaRussianName represents Claudia's name in Russian.
	claudiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клаудия"}

	// claudiaSimplifiedChineseName represents Claudia's name in Simplified Chinese.
	claudiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马丽莲"}

	// claudiaSpanishName represents Claudia's name in Spanish.
	claudiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lilu"}

	// claudiaTraditionalChineseName represents Claudia's name in Traditional Chinese.
	claudiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬麗蓮"}
)

var (
	// claudiaName represents Claudia's name in different languages.
	claudiaName = nook.Languages{
		language.AmericanEnglish:      claudiaAmericanEnglishName,
		language.CanadianFrench:       claudiaCanadianFrenchName,
		language.Dutch:                claudiaDutchName,
		language.French:               claudiaFrenchName,
		language.German:               claudiaGermanName,
		language.Italian:              claudiaItalianName,
		language.Japanese:             claudiaJapaneseName,
		language.Korean:               claudiaKoreanName,
		language.LatinAmericanSpanish: claudiaLatinAmericanSpanishName,
		language.Russian:              claudiaRussianName,
		language.SimplifiedChinese:    claudiaSimplifiedChineseName,
		language.Spanish:              claudiaSpanishName,
		language.TraditionalChinese:   claudiaTraditionalChineseName}
)

var (
	// claudiaCharacter represents Claudia's character information.
	claudiaCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: claudiaBirthday,
		Code:     claudiaCode,
		Key:      character.Claudia,
		Gender:   gender.Female,
		Name:     claudiaName,
		Special:  false}
)

var (
	// claudiaAmericanEnglishPhrase represents Claudia's phrase in American English.
	claudiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ooh la la"}

	// claudiaCanadianFrenchPhrase represents Claudia's phrase in Canadian French.
	claudiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grou grou"}

	// claudiaDutchPhrase represents Claudia's phrase in Dutch.
	claudiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "olala"}

	// claudiaFrenchPhrase represents Claudia's phrase in French.
	claudiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grou grou"}

	// claudiaGermanPhrase represents Claudia's phrase in German.
	claudiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kchhh"}

	// claudiaItalianPhrase represents Claudia's phrase in Italian.
	claudiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "smack"}

	// claudiaJapanesePhrase represents Claudia's phrase in Japanese.
	claudiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アハ～ン"}

	// claudiaKoreanPhrase represents Claudia's phrase in Korean.
	claudiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "잇힝"}

	// claudiaLatinAmericanSpanishPhrase represents Claudia's phrase in Latin American Spanish.
	claudiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mindundi"}

	// claudiaRussianPhrase represents Claudia's phrase in Russian.
	claudiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "о-ля-ля"}

	// claudiaSimplifiedChinesePhrase represents Claudia's phrase in Simplified Chinese.
	claudiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啊哈"}

	// claudiaSpanishPhrase represents Claudia's phrase in Spanish.
	claudiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mindundi"}

	// claudiaTraditionalChinesePhrase represents Claudia's phrase in Traditional Chinese.
	claudiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啊哈"}
)

var (
	// claudiaPhrase represents Claudia's phrase in different languages.
	claudiaPhrase = nook.Languages{
		language.AmericanEnglish:      claudiaAmericanEnglishPhrase,
		language.CanadianFrench:       claudiaCanadianFrenchPhrase,
		language.Dutch:                claudiaDutchPhrase,
		language.French:               claudiaFrenchPhrase,
		language.German:               claudiaGermanPhrase,
		language.Italian:              claudiaItalianPhrase,
		language.Japanese:             claudiaJapanesePhrase,
		language.Korean:               claudiaKoreanPhrase,
		language.LatinAmericanSpanish: claudiaLatinAmericanSpanishPhrase,
		language.Russian:              claudiaRussianPhrase,
		language.SimplifiedChinese:    claudiaSimplifiedChinesePhrase,
		language.Spanish:              claudiaSpanishPhrase,
		language.TraditionalChinese:   claudiaTraditionalChinesePhrase}
)

var (
	// Claudia represents the character Claudia with her complete information.
	Claudia = nook.Villager{
		Character:   claudiaCharacter,
		Personality: personality.Snooty,
		Phrase:      claudiaPhrase}
)
