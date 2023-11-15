package sheep

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
	// vestaBirthday represents Vesta's birthday.
	vestaBirthday = nook.Birthday{
		Day:   16,
		Month: time.April}
)

var (
	// vestaCode represents Vesta's unique code.
	vestaCode = nook.Code{
		Value: "shp00"}
)

var (
	// vestaAmericanEnglishName represents Vesta's name in American English.
	vestaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vesta"}

	// vestaCanadianFrenchName represents Vesta's name in Canadian French.
	vestaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hélaine"}

	// vestaDutchName represents Vesta's name in Dutch.
	vestaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vesta"}

	// vestaFrenchName represents Vesta's name in French.
	vestaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hélaine"}

	// vestaGermanName represents Vesta's name in German.
	vestaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dolly"}

	// vestaItalianName represents Vesta's name in Italian.
	vestaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lanella"}

	// vestaJapaneseName represents Vesta's name in Japanese.
	vestaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メリヤス"}

	// vestaKoreanName represents Vesta's name in Korean.
	vestaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메리어스"}

	// vestaLatinAmericanSpanishName represents Vesta's name in Latin American Spanish.
	vestaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vesta"}

	// vestaRussianName represents Vesta's name in Russian.
	vestaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Веста"}

	// vestaSimplifiedChineseName represents Vesta's name in Simplified Chinese.
	vestaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "绵绵"}

	// vestaSpanishName represents Vesta's name in Spanish.
	vestaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vesta"}

	// vestaTraditionalChineseName represents Vesta's name in Traditional Chinese.
	vestaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "綿綿"}
)

var (
	// vestaName represents Vesta's name in different languages.
	vestaName = nook.Languages{
		language.AmericanEnglish:      vestaAmericanEnglishName,
		language.CanadianFrench:       vestaCanadianFrenchName,
		language.Dutch:                vestaDutchName,
		language.French:               vestaFrenchName,
		language.German:               vestaGermanName,
		language.Italian:              vestaItalianName,
		language.Japanese:             vestaJapaneseName,
		language.Korean:               vestaKoreanName,
		language.LatinAmericanSpanish: vestaLatinAmericanSpanishName,
		language.Russian:              vestaRussianName,
		language.SimplifiedChinese:    vestaSimplifiedChineseName,
		language.Spanish:              vestaSpanishName,
		language.TraditionalChinese:   vestaTraditionalChineseName}
)

var (
	// vestaCharacter represents Vesta's character information.
	vestaCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: vestaBirthday,
		Code:     vestaCode,
		Key:      character.Vesta,
		Gender:   gender.Female,
		Name:     vestaName,
		Special:  false}
)

var (
	// vestaAmericanEnglishPhrase represents Vesta's phrase in American English.
	vestaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaaffo"}

	// vestaCanadianFrenchPhrase represents Vesta's phrase in Canadian French.
	vestaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "foulââârd"}

	// vestaDutchPhrase represents Vesta's phrase in Dutch.
	vestaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bèèè"}

	// vestaFrenchPhrase represents Vesta's phrase in French.
	vestaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pupull"}

	// vestaGermanPhrase represents Vesta's phrase in German.
	vestaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bäääh"}

	// vestaItalianPhrase represents Vesta's phrase in Italian.
	vestaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeeby"}

	// vestaJapanesePhrase represents Vesta's phrase in Japanese.
	vestaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのね"}

	// vestaKoreanPhrase represents Vesta's phrase in Korean.
	vestaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쩝"}

	// vestaLatinAmericanSpanishPhrase represents Vesta's phrase in Latin American Spanish.
	vestaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "beeeibi"}

	// vestaRussianPhrase represents Vesta's phrase in Russian.
	vestaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бе-е"}

	// vestaSimplifiedChinesePhrase represents Vesta's phrase in Simplified Chinese.
	vestaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就这样呀"}

	// vestaSpanishPhrase represents Vesta's phrase in Spanish.
	vestaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "beeeibi"}

	// vestaTraditionalChinesePhrase represents Vesta's phrase in Traditional Chinese.
	vestaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就這樣呀"}
)

var (
	// vestaPhrase represents Vesta's phrases in different languages.
	vestaPhrase = nook.Languages{
		language.AmericanEnglish:      vestaAmericanEnglishPhrase,
		language.CanadianFrench:       vestaCanadianFrenchPhrase,
		language.Dutch:                vestaDutchPhrase,
		language.French:               vestaFrenchPhrase,
		language.German:               vestaGermanPhrase,
		language.Italian:              vestaItalianPhrase,
		language.Japanese:             vestaJapanesePhrase,
		language.Korean:               vestaKoreanPhrase,
		language.LatinAmericanSpanish: vestaLatinAmericanSpanishPhrase,
		language.Russian:              vestaRussianPhrase,
		language.SimplifiedChinese:    vestaSimplifiedChinesePhrase,
		language.Spanish:              vestaSpanishPhrase,
		language.TraditionalChinese:   vestaTraditionalChinesePhrase}
)

var (
	// Vesta represents the character Vesta with her complete information.
	Vesta = nook.Villager{
		Character:   vestaCharacter,
		Personality: personality.Normal,
		Phrase:      vestaPhrase}
)
