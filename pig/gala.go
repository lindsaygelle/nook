package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	galaBirthday = nook.Birthday{
		Day:   5,
		Month: time.March}
)

var (
	galaCode = nook.Code{
		Value: "pig13"}
)

var (
	galaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gala"}

	galaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Camille"}

	galaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gala"}

	galaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Camille"}

	galaGermanName = nook.Name{
		Language: language.German,
		Value:    "Oinka"}

	galaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lisetta"}

	galaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ためこ"}

	galaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "꽃지"}

	galaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marita"}

	galaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гала"}

	galaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小芽"}

	galaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marita"}

	galaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小芽"}
)

var (
	galaName = nook.Languages{
		language.AmericanEnglish:      galaAmericanEnglishName,
		language.CanadianFrench:       galaCanadianFrenchName,
		language.Dutch:                galaDutchName,
		language.French:               galaFrenchName,
		language.German:               galaGermanName,
		language.Italian:              galaItalianName,
		language.Japanese:             galaJapaneseName,
		language.Korean:               galaKoreanName,
		language.LatinAmericanSpanish: galaLatinAmericanSpanishName,
		language.Russian:              galaRussianName,
		language.SimplifiedChinese:    galaSimplifiedChineseName,
		language.Spanish:              galaSpanishName,
		language.TraditionalChinese:   galaTraditionalChineseName}
)

var (
	galaCharacter = nook.Character{
		Animal:   Pig,
		Birthday: galaBirthday,
		Code:     galaCode,
		Gender:   nook.Female,
		Name:     galaName}
)

var (
	galaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snortie"}

	galaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tigroin"}

	galaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snufferd"}

	galaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tigroin"}

	galaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lächel"}

	galaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruffeffè"}

	galaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちゃりん"}

	galaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡그랑"}

	galaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chanchi"}

	galaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюк"}

	galaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "当啷"}

	galaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chanchi"}

	galaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噹啷"}
)

var (
	galaPhrase = nook.Languages{
		language.AmericanEnglish:      galaAmericanEnglishName,
		language.CanadianFrench:       galaCanadianFrenchName,
		language.Dutch:                galaDutchName,
		language.French:               galaFrenchName,
		language.German:               galaGermanName,
		language.Italian:              galaItalianName,
		language.Japanese:             galaJapaneseName,
		language.Korean:               galaKoreanName,
		language.LatinAmericanSpanish: galaLatinAmericanSpanishName,
		language.Russian:              galaRussianName,
		language.SimplifiedChinese:    galaSimplifiedChineseName,
		language.Spanish:              galaSpanishName,
		language.TraditionalChinese:   galaTraditionalChineseName}
)

var (
	Gala = nook.Villager{
		Character:   galaCharacter,
		Personality: nook.Normal,
		Phrase:      galaPhrase}
)
