package deer

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
	deirdreBirthday = nook.Birthday{
		Day:   4,
		Month: time.May}
)

var (
	deirdreCode = nook.Code{
		Value: "der04"}
)

var (
	deirdreAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deirdre"}

	deirdreCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bichoune"}

	deirdreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Deirdre"}

	deirdreFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bichoune"}

	deirdreGermanName = nook.Name{
		Language: language.German,
		Value:    "Dina"}

	deirdreItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Daria"}

	deirdreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナディア"}

	deirdreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나디아"}

	deirdreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Venada"}

	deirdreRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейрдре"}

	deirdreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娣雅"}

	deirdreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Venada"}

	deirdreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娣雅"}
)

var (
	deirdreName = nook.Languages{
		language.AmericanEnglish:      deirdreAmericanEnglishName,
		language.CanadianFrench:       deirdreCanadianFrenchName,
		language.Dutch:                deirdreDutchName,
		language.French:               deirdreFrenchName,
		language.German:               deirdreGermanName,
		language.Italian:              deirdreItalianName,
		language.Japanese:             deirdreJapaneseName,
		language.Korean:               deirdreKoreanName,
		language.LatinAmericanSpanish: deirdreLatinAmericanSpanishName,
		language.Russian:              deirdreRussianName,
		language.SimplifiedChinese:    deirdreSimplifiedChineseName,
		language.Spanish:              deirdreSpanishName,
		language.TraditionalChinese:   deirdreTraditionalChineseName}
)

var (
	deirdreCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: deirdreBirthday,
		Code:     deirdreCode,
		Key:      character.Deirdre,
		Gender:   gender.Female,
		Name:     deirdreName}
)

var (
	deirdreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whatevs"}

	deirdreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pan"}

	deirdreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nou ja"}

	deirdreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pan"}

	deirdreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tripptripp"}

	deirdreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babam"}

	deirdreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まいっか"}

	deirdreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "됐고"}

	deirdreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uatever"}

	deirdreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "как-то так"}

	deirdreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "随便啦"}

	deirdreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rumiante"}

	deirdreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "隨便啦"}
)

var (
	deirdrePhrase = nook.Languages{
		language.AmericanEnglish:      deirdreAmericanEnglishName,
		language.CanadianFrench:       deirdreCanadianFrenchName,
		language.Dutch:                deirdreDutchName,
		language.French:               deirdreFrenchName,
		language.German:               deirdreGermanName,
		language.Italian:              deirdreItalianName,
		language.Japanese:             deirdreJapaneseName,
		language.Korean:               deirdreKoreanName,
		language.LatinAmericanSpanish: deirdreLatinAmericanSpanishName,
		language.Russian:              deirdreRussianName,
		language.SimplifiedChinese:    deirdreSimplifiedChineseName,
		language.Spanish:              deirdreSpanishName,
		language.TraditionalChinese:   deirdreTraditionalChineseName}
)

var (
	Deirdre = nook.Villager{
		Character:   deirdreCharacter,
		Personality: personality.BigSister,
		Phrase:      deirdrePhrase}
)
