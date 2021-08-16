package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Jennifer"}

	reneighDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Reneigh"}

	reneighFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jennifer"}

	reneighGermanName = nook.Name{
		Language: language.German,
		Value:    "Andrea"}

	reneighItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Requina"}

	reneighJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リアーナ"}

	reneighKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리아나"}

	reneighLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corcelia"}

	reneighRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рени"}

	reneighSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷哈娜"}

	reneighSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corcelia"}

	reneighTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷哈娜"}
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
		Animal:   animal.Horse,
		Birthday: reneighBirthday,
		Code:     reneighCode,
		Gender:   gender.Female,
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
		Personality: personality.BigSister,
		Phrase:      reneighPhrase}
)
