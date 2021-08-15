package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	blancheBirthday = nook.Birthday{
		Day:   21,
		Month: time.December}
)

var (
	blancheCode = nook.Code{
		Value: "ost08"}
)

var (
	blancheAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blanche"}

	blancheCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sophieviiiite"}

	blancheDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blanchejuist ja"}

	blancheFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sophieviiiite"}

	blancheGermanName = nook.Name{
		Language: language.German,
		Value:    "Christaflaum"}

	blancheItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Emiliabien sur"}

	blancheJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しのぶですのね"}

	blancheKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "신옥그랬구나"}

	blancheLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rocíoigh-igh"}

	blancheRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бланшвот так"}

	blancheSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小偲是嘛"}

	blancheSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rocíogoticas"}

	blancheTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小偲是嘛"}
)

var (
	blancheName = nook.Languages{
		language.AmericanEnglish:      blancheAmericanEnglishName,
		language.CanadianFrench:       blancheCanadianFrenchName,
		language.Dutch:                blancheDutchName,
		language.French:               blancheFrenchName,
		language.German:               blancheGermanName,
		language.Italian:              blancheItalianName,
		language.Japanese:             blancheJapaneseName,
		language.Korean:               blancheKoreanName,
		language.LatinAmericanSpanish: blancheLatinAmericanSpanishName,
		language.Russian:              blancheRussianName,
		language.SimplifiedChinese:    blancheSimplifiedChineseName,
		language.Spanish:              blancheSpanishName,
		language.TraditionalChinese:   blancheTraditionalChineseName}
)

var (
	blancheCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: blancheBirthday,
		Code:     blancheCode,
		Gender:   nook.Female,
		Name:     blancheName}
)

var (
	blancheAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quite so"}

	blancheCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "viiiite"}

	blancheDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "juist ja"}

	blancheFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "viiiite"}

	blancheGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flaum"}

	blancheItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bien sur"}

	blancheJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですのね"}

	blancheKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬구나"}

	blancheLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "igh-igh"}

	blancheRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вот так"}

	blancheSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是嘛"}

	blancheSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "goticas"}

	blancheTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是嘛"}
)

var (
	blanchePhrase = nook.Languages{
		language.AmericanEnglish:      blancheAmericanEnglishName,
		language.CanadianFrench:       blancheCanadianFrenchName,
		language.Dutch:                blancheDutchName,
		language.French:               blancheFrenchName,
		language.German:               blancheGermanName,
		language.Italian:              blancheItalianName,
		language.Japanese:             blancheJapaneseName,
		language.Korean:               blancheKoreanName,
		language.LatinAmericanSpanish: blancheLatinAmericanSpanishName,
		language.Russian:              blancheRussianName,
		language.SimplifiedChinese:    blancheSimplifiedChineseName,
		language.Spanish:              blancheSpanishName,
		language.TraditionalChinese:   blancheTraditionalChineseName}
)

var (
	Blanche = nook.Villager{
		Character:   blancheCharacter,
		Personality: nook.Snooty,
		Phrase:      blanchePhrase}
)
