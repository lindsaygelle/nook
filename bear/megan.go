package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	meganBirthday = nook.Birthday{
		Day:   13,
		Month: time.March}
)

var (
	meganCode = nook.Code{
		Value: "bea15"}
)

var (
	meganAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Megan"}

	meganCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Candyberlingot"}

	meganDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Meganlolly"}

	meganFrenchName = nook.Name{
		Language: language.French,
		Value:    "Candyberlingot"}

	meganGermanName = nook.Name{
		Language: language.German,
		Value:    "Dagmarlollipop"}

	meganItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dolciniasucré"}

	meganJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンディぺろ"}

	meganKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캔디낼름"}

	meganLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carmelatutifruti"}

	meganRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мегансладенько"}

	meganSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "糖果舔舔"}

	meganSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmelatutifruti"}

	meganTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "糖果舔舔"}
)

var (
	meganName = nook.Languages{
		language.AmericanEnglish:      meganAmericanEnglishName,
		language.CanadianFrench:       meganCanadianFrenchName,
		language.Dutch:                meganDutchName,
		language.French:               meganFrenchName,
		language.German:               meganGermanName,
		language.Italian:              meganItalianName,
		language.Japanese:             meganJapaneseName,
		language.Korean:               meganKoreanName,
		language.LatinAmericanSpanish: meganLatinAmericanSpanishName,
		language.Russian:              meganRussianName,
		language.SimplifiedChinese:    meganSimplifiedChineseName,
		language.Spanish:              meganSpanishName,
		language.TraditionalChinese:   meganTraditionalChineseName}
)

var (
	meganCharacter = nook.Character{
		Animal:   Bear,
		Birthday: meganBirthday,
		Code:     meganCode,
		Gender:   nook.Female,
		Name:     meganName}
)

var (
	meganAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sundae"}

	meganCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "berlingot"}

	meganDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lolly"}

	meganFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "berlingot"}

	meganGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lollipop"}

	meganItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sucré"}

	meganJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぺろ"}

	meganKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "낼름"}

	meganLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tutifruti"}

	meganRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сладенько"}

	meganSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舔舔"}

	meganSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tutifruti"}

	meganTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舔舔"}
)

var (
	meganPhrase = nook.Languages{
		language.AmericanEnglish:      meganAmericanEnglishName,
		language.CanadianFrench:       meganCanadianFrenchName,
		language.Dutch:                meganDutchName,
		language.French:               meganFrenchName,
		language.German:               meganGermanName,
		language.Italian:              meganItalianName,
		language.Japanese:             meganJapaneseName,
		language.Korean:               meganKoreanName,
		language.LatinAmericanSpanish: meganLatinAmericanSpanishName,
		language.Russian:              meganRussianName,
		language.SimplifiedChinese:    meganSimplifiedChineseName,
		language.Spanish:              meganSpanishName,
		language.TraditionalChinese:   meganTraditionalChineseName}
)

var (
	Megan = nook.Villager{
		Character:   meganCharacter,
		Personality: nook.Normal,
		Phrase:      meganPhrase}
)
