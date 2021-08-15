package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Miroou bien"}

	stitchesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stitchespluche"}

	stitchesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miroou bien"}

	stitchesGermanName = nook.Name{
		Language: language.German,
		Value:    "Berrybrummm"}

	stitchesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Toppettaohilà"}

	stitchesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パッチなのれす"}

	stitchesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패치그런거죠"}

	stitchesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Parchespaguahhh"}

	stitchesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стичесплюш-плюш"}

	stitchesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玩具熊玩玩"}

	stitchesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Parchespeluche"}

	stitchesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玩具熊玩玩"}
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
		Animal:   Bearcub,
		Birthday: stitchesBirthday,
		Code:     stitchesCode,
		Gender:   nook.Male,
		Name:     stitchesName}
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
	Stitches = nook.Villager{
		Character:   stitchesCharacter,
		Personality: nook.Lazy,
		Phrase:      stitchesPhrase}
)
