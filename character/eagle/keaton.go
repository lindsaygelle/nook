package eagle

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
		Value:    "Enzo"}

	keatonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Keaton"}

	keatonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Enzo"}

	keatonGermanName = nook.Name{
		Language: language.German,
		Value:    "Kai"}

	keatonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aquilio"}

	keatonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランク"}

	keatonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랭크"}

	keatonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lucho"}

	keatonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китон"}

	keatonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "法兰克"}

	keatonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lucho"}

	keatonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "法蘭克"}
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
		Animal:   animal.Eagle,
		Birthday: keatonBirthday,
		Code:     keatonCode,
		Key:      character.Keaton,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      keatonAmericanEnglishPhrase,
		language.CanadianFrench:       keatonCanadianFrenchPhrase,
		language.Dutch:                keatonDutchPhrase,
		language.French:               keatonFrenchPhrase,
		language.German:               keatonGermanPhrase,
		language.Italian:              keatonItalianPhrase,
		language.Japanese:             keatonJapanesePhrase,
		language.Korean:               keatonKoreanPhrase,
		language.LatinAmericanSpanish: keatonLatinAmericanSpanishPhrase,
		language.Russian:              keatonRussianPhrase,
		language.SimplifiedChinese:    keatonSimplifiedChinesePhrase,
		language.Spanish:              keatonSpanishPhrase,
		language.TraditionalChinese:   keatonTraditionalChinesePhrase}
)

var (
	Keaton = nook.Villager{
		Character:   keatonCharacter,
		Personality: personality.Smug,
		Phrase:      keatonPhrase}
)
