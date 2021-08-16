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
	avaBirthday = nook.Birthday{
		Day:   28,
		Month: time.April}
)

var (
	avaCode = nook.Code{
		Value: "chn05"}
)

var (
	avaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ava"}

	avaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eva"}

	avaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ava"}

	avaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eva"}

	avaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gisela"}

	avaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ava"}

	avaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドミグラ"}

	avaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에바"}

	avaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ava"}

	avaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ава"}

	avaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "陶米咕"}

	avaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ava"}

	avaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陶米咕"}
)

var (
	avaName = nook.Languages{
		language.AmericanEnglish:      avaAmericanEnglishName,
		language.CanadianFrench:       avaCanadianFrenchName,
		language.Dutch:                avaDutchName,
		language.French:               avaFrenchName,
		language.German:               avaGermanName,
		language.Italian:              avaItalianName,
		language.Japanese:             avaJapaneseName,
		language.Korean:               avaKoreanName,
		language.LatinAmericanSpanish: avaLatinAmericanSpanishName,
		language.Russian:              avaRussianName,
		language.SimplifiedChinese:    avaSimplifiedChineseName,
		language.Spanish:              avaSpanishName,
		language.TraditionalChinese:   avaTraditionalChineseName}
)

var (
	avaCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: avaBirthday,
		Code:     avaCode,
		Gender:   gender.Female,
		Name:     avaName}
)

var (
	avaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "beaker"}

	avaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cô-cômment"}

	avaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snavel"}

	avaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cô-cômment"}

	avaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "biiika"}

	avaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbecsbec"}

	avaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のよぉ"}

	avaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬대요"}

	avaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kikiú"}

	avaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "клювик"}

	avaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是唷"}

	avaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piquito"}

	avaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是唷"}
)

var (
	avaPhrase = nook.Languages{
		language.AmericanEnglish:      avaAmericanEnglishName,
		language.CanadianFrench:       avaCanadianFrenchName,
		language.Dutch:                avaDutchName,
		language.French:               avaFrenchName,
		language.German:               avaGermanName,
		language.Italian:              avaItalianName,
		language.Japanese:             avaJapaneseName,
		language.Korean:               avaKoreanName,
		language.LatinAmericanSpanish: avaLatinAmericanSpanishName,
		language.Russian:              avaRussianName,
		language.SimplifiedChinese:    avaSimplifiedChineseName,
		language.Spanish:              avaSpanishName,
		language.TraditionalChinese:   avaTraditionalChineseName}
)

var (
	Ava = nook.Villager{
		Character:   avaCharacter,
		Personality: personality.Normal,
		Phrase:      avaPhrase}
)
