package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	egbertBirthday = nook.Birthday{
		Day:   14,
		Month: time.October}
)

var (
	egbertCode = nook.Code{
		Value: "chn02"}
)

var (
	egbertAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Egbert"}

	egbertCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Herbert"}

	egbertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Egbert"}

	egbertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Herbert"}

	egbertGermanName = nook.Name{
		Language: language.German,
		Value:    "Waldemar"}

	egbertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chicco"}

	egbertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しもやけ"}

	egbertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "김희"}

	egbertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Norberto"}

	egbertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эгберт"}

	egbertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "黄金鸡"}

	egbertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Norberto"}

	egbertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "黃金雞"}
)

var (
	egbertName = nook.Languages{
		language.AmericanEnglish:      egbertAmericanEnglishName,
		language.CanadianFrench:       egbertCanadianFrenchName,
		language.Dutch:                egbertDutchName,
		language.French:               egbertFrenchName,
		language.German:               egbertGermanName,
		language.Italian:              egbertItalianName,
		language.Japanese:             egbertJapaneseName,
		language.Korean:               egbertKoreanName,
		language.LatinAmericanSpanish: egbertLatinAmericanSpanishName,
		language.Russian:              egbertRussianName,
		language.SimplifiedChinese:    egbertSimplifiedChineseName,
		language.Spanish:              egbertSpanishName,
		language.TraditionalChinese:   egbertTraditionalChineseName}
)

var (
	egbertCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: egbertBirthday,
		Code:     egbertCode,
		Gender:   gender.Male,
		Name:     egbertName}
)

var (
	egbertAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "doodle-duh"}

	egbertCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "di di dou"}

	egbertDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kukelejus"}

	egbertFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "di di dou"}

	egbertGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "duudle-du"}

	egbertItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "du du du"}

	egbertJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だヨ"}

	egbertKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜잔"}

	egbertLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kikirikí"}

	egbertRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кукареку"}

	egbertSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "很优"}

	egbertSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kikirikí"}

	egbertTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "很優"}
)

var (
	egbertPhrase = nook.Languages{
		language.AmericanEnglish:      egbertAmericanEnglishName,
		language.CanadianFrench:       egbertCanadianFrenchName,
		language.Dutch:                egbertDutchName,
		language.French:               egbertFrenchName,
		language.German:               egbertGermanName,
		language.Italian:              egbertItalianName,
		language.Japanese:             egbertJapaneseName,
		language.Korean:               egbertKoreanName,
		language.LatinAmericanSpanish: egbertLatinAmericanSpanishName,
		language.Russian:              egbertRussianName,
		language.SimplifiedChinese:    egbertSimplifiedChineseName,
		language.Spanish:              egbertSpanishName,
		language.TraditionalChinese:   egbertTraditionalChineseName}
)

var (
	Egbert = nook.Villager{
		Character:   egbertCharacter,
		Personality: personality.Lazy,
		Phrase:      egbertPhrase}
)
