package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	keatonBirthday = nook.Birthday{
		Day:   1,
		Month: time.June}
)

var (
	keatonCode = nook.Code{
		Value: "pbr08"}
)

var (
	keatonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Keaton"}

	keatonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Enzoroyaaaal"}

	keatonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Keatonvleugels"}

	keatonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Enzoroyaaaal"}

	keatonGermanName = nook.Name{
		Language: language.German,
		Value:    "Kaikiiiiiih"}

	keatonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aquiliobaila"}

	keatonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランクでーッス"}

	keatonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랭크수리수리"}

	keatonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Luchoflopflop"}

	keatonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китонкрыло"}

	keatonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "法兰克是"}

	keatonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Luchoflopflop"}

	keatonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "法蘭克是"}
)

var (
	keatonName = nook.Languages{
		language.AmericanEnglish:      keatonAmericanEnglishName,
		language.CanadianFrench:       keatonCanadianFrenchName,
		language.Dutch:                keatonDutchName,
		language.French:               keatonFrenchName,
		language.German:               keatonGermanName,
		language.Italian:              keatonItalianName,
		language.Japanese:             keatonJapaneseName,
		language.Korean:               keatonKoreanName,
		language.LatinAmericanSpanish: keatonLatinAmericanSpanishName,
		language.Russian:              keatonRussianName,
		language.SimplifiedChinese:    keatonSimplifiedChineseName,
		language.Spanish:              keatonSpanishName,
		language.TraditionalChinese:   keatonTraditionalChineseName}
)

var (
	keatonCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: keatonBirthday,
		Code:     keatonCode,
		Gender:   nook.Male,
		Name:     keatonName}
)

var (
	keatonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wingo"}

	keatonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "royaaaal"}

	keatonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vleugels"}

	keatonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "royaaaal"}

	keatonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kiiiiiih"}

	keatonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baila"}

	keatonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でーッス"}

	keatonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "수리수리"}

	keatonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flopflop"}

	keatonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крыло"}

	keatonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是"}

	keatonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flopflop"}

	keatonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是"}
)

var (
	keatonPhrase = nook.Languages{
		language.AmericanEnglish:      keatonAmericanEnglishName,
		language.CanadianFrench:       keatonCanadianFrenchName,
		language.Dutch:                keatonDutchName,
		language.French:               keatonFrenchName,
		language.German:               keatonGermanName,
		language.Italian:              keatonItalianName,
		language.Japanese:             keatonJapaneseName,
		language.Korean:               keatonKoreanName,
		language.LatinAmericanSpanish: keatonLatinAmericanSpanishName,
		language.Russian:              keatonRussianName,
		language.SimplifiedChinese:    keatonSimplifiedChineseName,
		language.Spanish:              keatonSpanishName,
		language.TraditionalChinese:   keatonTraditionalChineseName}
)

var (
	Keaton = nook.Villager{
		Character:   keatonCharacter,
		Personality: nook.Smug,
		Phrase:      keatonPhrase}
)
