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
	stitchesBirthday = nook.Birthday{
		Day:   10,
		Month: time.February}
)

var (
	stitchesCode = nook.Code{
		Value: "cbr05"}
)

var (
	stitchesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stitches"}

	stitchesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miro"}

	stitchesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stitches"}

	stitchesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miro"}

	stitchesGermanName = nook.Name{
		Language: language.German,
		Value:    "Berry"}

	stitchesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Toppetta"}

	stitchesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パッチ"}

	stitchesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패치"}

	stitchesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Parches"}

	stitchesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стичес"}

	stitchesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玩具熊"}

	stitchesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Parches"}

	stitchesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玩具熊"}
)

var (
	stitchesName = nook.Languages{
		language.AmericanEnglish:      stitchesAmericanEnglishName,
		language.CanadianFrench:       stitchesCanadianFrenchName,
		language.Dutch:                stitchesDutchName,
		language.French:               stitchesFrenchName,
		language.German:               stitchesGermanName,
		language.Italian:              stitchesItalianName,
		language.Japanese:             stitchesJapaneseName,
		language.Korean:               stitchesKoreanName,
		language.LatinAmericanSpanish: stitchesLatinAmericanSpanishName,
		language.Russian:              stitchesRussianName,
		language.SimplifiedChinese:    stitchesSimplifiedChineseName,
		language.Spanish:              stitchesSpanishName,
		language.TraditionalChinese:   stitchesTraditionalChineseName}
)

var (
	stitchesCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: stitchesBirthday,
		Code:     stitchesCode,
		Key:      character.Stitches,
		Gender:   gender.Male,
		Name:     stitchesName,
		Special:  false}
)

var (
	stitchesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stuffin'"}

	stitchesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ou bien"}

	stitchesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pluche"}

	stitchesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ou bien"}

	stitchesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brummm"}

	stitchesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ohilà"}

	stitchesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのれす"}

	stitchesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그런거죠"}

	stitchesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "paguahhh"}

	stitchesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "плюш-плюш"}

	stitchesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玩玩"}

	stitchesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peluche"}

	stitchesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玩玩"}
)

var (
	stitchesPhrase = nook.Languages{
		language.AmericanEnglish:      stitchesAmericanEnglishPhrase,
		language.CanadianFrench:       stitchesCanadianFrenchPhrase,
		language.Dutch:                stitchesDutchPhrase,
		language.French:               stitchesFrenchPhrase,
		language.German:               stitchesGermanPhrase,
		language.Italian:              stitchesItalianPhrase,
		language.Japanese:             stitchesJapanesePhrase,
		language.Korean:               stitchesKoreanPhrase,
		language.LatinAmericanSpanish: stitchesLatinAmericanSpanishPhrase,
		language.Russian:              stitchesRussianPhrase,
		language.SimplifiedChinese:    stitchesSimplifiedChinesePhrase,
		language.Spanish:              stitchesSpanishPhrase,
		language.TraditionalChinese:   stitchesTraditionalChinesePhrase}
)

var (
	Stitches = nook.Villager{
		Character:   stitchesCharacter,
		Personality: personality.Lazy,
		Phrase:      stitchesPhrase}
)
