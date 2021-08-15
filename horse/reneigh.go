package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	reneighBirthday = nook.Birthday{
		Day:   4,
		Month: time.June}
)

var (
	reneighCode = nook.Code{
		Value: "hrs16"}
)

var (
	reneighAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Reneigh"}

	reneighCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jennifercrincrin"}

	reneighDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Reneighuhuh"}

	reneighFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jennifercrincrin"}

	reneighGermanName = nook.Name{
		Language: language.German,
		Value:    "Andreatada"}

	reneighItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Requinahiii-op"}

	reneighJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リアーナスン"}

	reneighKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리아나쉬익"}

	reneighLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corceliajuááá"}

	reneighRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ренино-но"}

	reneighSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷哈娜哼"}

	reneighSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corceliaalehop"}

	reneighTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷哈娜哼"}
)

var (
	reneighName = nook.Languages{
		language.AmericanEnglish:      reneighAmericanEnglishName,
		language.CanadianFrench:       reneighCanadianFrenchName,
		language.Dutch:                reneighDutchName,
		language.French:               reneighFrenchName,
		language.German:               reneighGermanName,
		language.Italian:              reneighItalianName,
		language.Japanese:             reneighJapaneseName,
		language.Korean:               reneighKoreanName,
		language.LatinAmericanSpanish: reneighLatinAmericanSpanishName,
		language.Russian:              reneighRussianName,
		language.SimplifiedChinese:    reneighSimplifiedChineseName,
		language.Spanish:              reneighSpanishName,
		language.TraditionalChinese:   reneighTraditionalChineseName}
)

var (
	reneighCharacter = nook.Character{
		Animal:   Horse,
		Birthday: reneighBirthday,
		Code:     reneighCode,
		Gender:   nook.Female,
		Name:     reneighName}
)

var (
	reneighAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ayup yup"}

	reneighCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crincrin"}

	reneighDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "uhuh"}

	reneighFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crincrin"}

	reneighGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tada"}

	reneighItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hiii-op"}

	reneighJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "スン"}

	reneighKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쉬익"}

	reneighLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "juááá"}

	reneighRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "но-но"}

	reneighSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哼"}

	reneighSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "alehop"}

	reneighTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哼"}
)

var (
	reneighPhrase = nook.Languages{
		language.AmericanEnglish:      reneighAmericanEnglishName,
		language.CanadianFrench:       reneighCanadianFrenchName,
		language.Dutch:                reneighDutchName,
		language.French:               reneighFrenchName,
		language.German:               reneighGermanName,
		language.Italian:              reneighItalianName,
		language.Japanese:             reneighJapaneseName,
		language.Korean:               reneighKoreanName,
		language.LatinAmericanSpanish: reneighLatinAmericanSpanishName,
		language.Russian:              reneighRussianName,
		language.SimplifiedChinese:    reneighSimplifiedChineseName,
		language.Spanish:              reneighSpanishName,
		language.TraditionalChinese:   reneighTraditionalChineseName}
)

var (
	Reneigh = nook.Villager{
		Character:   reneighCharacter,
		Personality: nook.BigSister,
		Phrase:      reneighPhrase}
)
