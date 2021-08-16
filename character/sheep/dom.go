package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	domBirthday = nook.Birthday{
		Day:   18,
		Month: time.March}
)

var (
	domCode = nook.Code{
		Value: "shp15"}
)

var (
	domAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dom"}

	domCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bouloche"}

	domDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Donny"}

	domFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bouloche"}

	domGermanName = nook.Name{
		Language: language.German,
		Value:    "Dominik"}

	domItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ovilio"}

	domJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちゃちゃまる"}

	domKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "차둘"}

	domLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fibrilio"}

	domRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дом"}

	domSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茶茶丸"}

	domSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fibrilio"}

	domTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茶茶丸"}
)

var (
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
	domCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: domBirthday,
		Code:     domCode,
		Gender:   gender.Male,
		Name:     domName}
)

var (
	domAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "indeedaroo"}

	domCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "broutille"}

	domDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "precies"}

	domFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "broutille"}

	domGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jau"}

	domItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "veronò"}

	domJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふんふん"}

	domKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿜뿜"}

	domLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "digoyó"}

	domRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "точно-точно"}

	domSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇耶"}

	domSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "diiigo yo"}

	domTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇耶"}
)

var (
	domPhrase = nook.Languages{
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
	Dom = nook.Villager{
		Character:   domCharacter,
		Personality: personality.Jock,
		Phrase:      domPhrase}
)
