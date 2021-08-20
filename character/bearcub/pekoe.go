package bearcub

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
	pekoeBirthday = nook.Birthday{
		Day:   18,
		Month: time.May}
)

var (
	pekoeCode = nook.Code{
		Value: "cbr14"}
)

var (
	pekoeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pekoe"}

	pekoeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pauline"}

	pekoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pekoe"}

	pekoeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pauline"}

	pekoeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sandrine"}

	pekoeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mina"}

	pekoeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャスミン"}

	pekoeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "재스민"}

	pekoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vera"}

	pekoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пеко"}

	pekoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佳敏"}

	pekoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vera"}

	pekoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佳敏"}
)

var (
	pekoeName = nook.Languages{
		language.AmericanEnglish:      pekoeAmericanEnglishName,
		language.CanadianFrench:       pekoeCanadianFrenchName,
		language.Dutch:                pekoeDutchName,
		language.French:               pekoeFrenchName,
		language.German:               pekoeGermanName,
		language.Italian:              pekoeItalianName,
		language.Japanese:             pekoeJapaneseName,
		language.Korean:               pekoeKoreanName,
		language.LatinAmericanSpanish: pekoeLatinAmericanSpanishName,
		language.Russian:              pekoeRussianName,
		language.SimplifiedChinese:    pekoeSimplifiedChineseName,
		language.Spanish:              pekoeSpanishName,
		language.TraditionalChinese:   pekoeTraditionalChineseName}
)

var (
	pekoeCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: pekoeBirthday,
		Code:     pekoeCode,
		Key:      character.Pekoe,
		Gender:   gender.Female,
		Name:     pekoeName,
		Special:  false}
)

var (
	pekoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bud"}

	pekoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "linou"}

	pekoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bloempje"}

	pekoeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "linou"}

	pekoeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flüster"}

	pekoeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "purrrimba"}

	pekoeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チャ"}

	pekoeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "띵호아"}

	pekoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pitiminí"}

	pekoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цветочек"}

	pekoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佳佳"}

	pekoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pitiminí"}

	pekoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佳佳"}
)

var (
	pekoePhrase = nook.Languages{
		language.AmericanEnglish:      pekoeAmericanEnglishPhrase,
		language.CanadianFrench:       pekoeCanadianFrenchPhrase,
		language.Dutch:                pekoeDutchPhrase,
		language.French:               pekoeFrenchPhrase,
		language.German:               pekoeGermanPhrase,
		language.Italian:              pekoeItalianPhrase,
		language.Japanese:             pekoeJapanesePhrase,
		language.Korean:               pekoeKoreanPhrase,
		language.LatinAmericanSpanish: pekoeLatinAmericanSpanishPhrase,
		language.Russian:              pekoeRussianPhrase,
		language.SimplifiedChinese:    pekoeSimplifiedChinesePhrase,
		language.Spanish:              pekoeSpanishPhrase,
		language.TraditionalChinese:   pekoeTraditionalChinesePhrase}
)

var (
	Pekoe = nook.Villager{
		Character:   pekoeCharacter,
		Personality: personality.Normal,
		Phrase:      pekoePhrase}
)
