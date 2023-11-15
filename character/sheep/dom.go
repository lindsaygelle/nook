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
	// domBirthday represents Dom's birthday (March 18th).
	domBirthday = nook.Birthday{
		Day:   18,
		Month: time.March}
)

var (
	// domCode represents Dom's unique code ("shp15").
	domCode = nook.Code{
		Value: "shp15"}
)

var (
	// domAmericanEnglishName represents Dom's name in American English.
	domAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dom"}

	// domCanadianFrenchName represents Dom's name in Canadian French.
	domCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bouloche"}

	// domDutchName represents Dom's name in Dutch.
	domDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Donny"}

	// domFrenchName represents Dom's name in French.
	domFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bouloche"}

	// domGermanName represents Dom's name in German.
	domGermanName = nook.Name{
		Language: language.German,
		Value:    "Dominik"}

	// domItalianName represents Dom's name in Italian.
	domItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ovilio"}

	// domJapaneseName represents Dom's name in Japanese.
	domJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちゃちゃまる"}

	// domKoreanName represents Dom's name in Korean.
	domKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "차둘"}

	// domLatinAmericanSpanishName represents Dom's name in Latin American Spanish.
	domLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fibrilio"}

	// domRussianName represents Dom's name in Russian.
	domRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дом"}

	// domSimplifiedChineseName represents Dom's name in Simplified Chinese.
	domSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茶茶丸"}

	// domSpanishName represents Dom's name in Spanish.
	domSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fibrilio"}

	// domTraditionalChineseName represents Dom's name in Traditional Chinese.
	domTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茶茶丸"}
)

var (
	// domName represents Dom's name in different languages.
	domName = nook.Languages{
		language.AmericanEnglish:      domAmericanEnglishName,
		language.CanadianFrench:       domCanadianFrenchName,
		language.Dutch:                domDutchName,
		language.French:               domFrenchName,
		language.German:               domGermanName,
		language.Italian:              domItalianName,
		language.Japanese:             domJapaneseName,
		language.Korean:               domKoreanName,
		language.LatinAmericanSpanish: domLatinAmericanSpanishName,
		language.Russian:              domRussianName,
		language.SimplifiedChinese:    domSimplifiedChineseName,
		language.Spanish:              domSpanishName,
		language.TraditionalChinese:   domTraditionalChineseName}
)

var (
	// domCharacter represents Dom's character information.
	domCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: domBirthday,
		Code:     domCode,
		Key:      character.Dom,
		Gender:   gender.Male,
		Name:     domName,
		Special:  false}
)

var (
	// domAmericanEnglishPhrase represents Dom's phrase in American English.
	domAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "indeedaroo"}

	// domCanadianFrenchPhrase represents Dom's phrase in Canadian French.
	domCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "broutille"}

	// domDutchPhrase represents Dom's phrase in Dutch.
	domDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "precies"}

	// domFrenchPhrase represents Dom's phrase in French.
	domFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "broutille"}

	// domGermanPhrase represents Dom's phrase in German.
	domGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jau"}

	// domItalianPhrase represents Dom's phrase in Italian.
	domItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "veronò"}

	// domJapanesePhrase represents Dom's phrase in Japanese.
	domJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふんふん"}

	// domKoreanPhrase represents Dom's phrase in Korean.
	domKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿜뿜"}

	// domLatinAmericanSpanishPhrase represents Dom's phrase in Latin American Spanish.
	domLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "digoyó"}

	// domRussianPhrase represents Dom's phrase in Russian.
	domRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "точно-точно"}

	// domSimplifiedChinesePhrase represents Dom's phrase in Simplified Chinese.
	domSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇耶"}

	// domSpanishPhrase represents Dom's phrase in Spanish.
	domSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "diiigo yo"}

	// domTraditionalChinesePhrase represents Dom's phrase in Traditional Chinese.
	domTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇耶"}
)

var (
	// domPhrase represents Dom's phrases in different languages.
	domPhrase = nook.Languages{
		language.AmericanEnglish:      domAmericanEnglishPhrase,
		language.CanadianFrench:       domCanadianFrenchPhrase,
		language.Dutch:                domDutchPhrase,
		language.French:               domFrenchPhrase,
		language.German:               domGermanPhrase,
		language.Italian:              domItalianPhrase,
		language.Japanese:             domJapanesePhrase,
		language.Korean:               domKoreanPhrase,
		language.LatinAmericanSpanish: domLatinAmericanSpanishPhrase,
		language.Russian:              domRussianPhrase,
		language.SimplifiedChinese:    domSimplifiedChinesePhrase,
		language.Spanish:              domSpanishPhrase,
		language.TraditionalChinese:   domTraditionalChinesePhrase}
)

var (
	// Dom represents the character Dom with his complete information.
	Dom = nook.Villager{
		Character:   domCharacter,
		Personality: personality.Jock,
		Phrase:      domPhrase}
)
