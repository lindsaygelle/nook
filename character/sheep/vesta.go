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
	vestaBirthday = nook.Birthday{
		Day:   16,
		Month: time.April}
)

var (
	vestaCode = nook.Code{
		Value: "shp00"}
)

var (
	vestaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vesta"}

	vestaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hélaine"}

	vestaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vesta"}

	vestaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hélaine"}

	vestaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dolly"}

	vestaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lanella"}

	vestaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メリヤス"}

	vestaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메리어스"}

	vestaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vesta"}

	vestaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Веста"}

	vestaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "绵绵"}

	vestaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vesta"}

	vestaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "綿綿"}
)

var (
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
	vestaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaaffo"}

	vestaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "foulââârd"}

	vestaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bèèè"}

	vestaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pupull"}

	vestaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bäääh"}

	vestaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeeby"}

	vestaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのね"}

	vestaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쩝"}

	vestaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "beeeibi"}

	vestaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бе-е"}

	vestaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就这样呀"}

	vestaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "beeeibi"}

	vestaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就這樣呀"}
)

var (
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
	Vesta = nook.Villager{
		Character:   vestaCharacter,
		Personality: personality.Normal,
		Phrase:      vestaPhrase}
)
