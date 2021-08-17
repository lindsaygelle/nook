package bull

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
	angusBirthday = nook.Birthday{
		Day:   30,
		Month: time.April}
)

var (
	angusCode = nook.Code{
		Value: "bul00"}
)

var (
	angusAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Angus"}

	angusCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Angus"}

	angusDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Angus"}

	angusFrenchName = nook.Name{
		Language: language.French,
		Value:    "Angus"}

	angusGermanName = nook.Name{
		Language: language.German,
		Value:    "Angus"}

	angusItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angus"}

	angusJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セルバンテス"}

	angusKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "반데스"}

	angusLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aliste"}

	angusRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ангус"}

	angusSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施万德"}

	angusSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aliste"}

	angusTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施萬德"}
)

var (
	angusName = nook.Languages{
		language.AmericanEnglish:      angusAmericanEnglishName,
		language.CanadianFrench:       angusCanadianFrenchName,
		language.Dutch:                angusDutchName,
		language.French:               angusFrenchName,
		language.German:               angusGermanName,
		language.Italian:              angusItalianName,
		language.Japanese:             angusJapaneseName,
		language.Korean:               angusKoreanName,
		language.LatinAmericanSpanish: angusLatinAmericanSpanishName,
		language.Russian:              angusRussianName,
		language.SimplifiedChinese:    angusSimplifiedChineseName,
		language.Spanish:              angusSpanishName,
		language.TraditionalChinese:   angusTraditionalChineseName}
)

var (
	angusCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: angusBirthday,
		Code:     angusCode,
		Key:      character.Angus,
		Gender:   gender.Male,
		Name:     angusName}
)

var (
	angusAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "macmoo"}

	angusCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh meuh"}

	angusDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "denderund"}

	angusFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh meuh"}

	angusGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrruzzl"}

	angusItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "macmoo"}

	angusJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふふ"}

	angusKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "후후"}

	angusLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bufff"}

	angusRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "макму-у-у"}

	angusSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呼呼"}

	angusSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuernos"}

	angusTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呼呼"}
)

var (
	angusPhrase = nook.Languages{
		language.AmericanEnglish:      angusAmericanEnglishPhrase,
		language.CanadianFrench:       angusCanadianFrenchPhrase,
		language.Dutch:                angusDutchPhrase,
		language.French:               angusFrenchPhrase,
		language.German:               angusGermanPhrase,
		language.Italian:              angusItalianPhrase,
		language.Japanese:             angusJapanesePhrase,
		language.Korean:               angusKoreanPhrase,
		language.LatinAmericanSpanish: angusLatinAmericanSpanishPhrase,
		language.Russian:              angusRussianPhrase,
		language.SimplifiedChinese:    angusSimplifiedChinesePhrase,
		language.Spanish:              angusSpanishPhrase,
		language.TraditionalChinese:   angusTraditionalChinesePhrase}
)

var (
	Angus = nook.Villager{
		Character:   angusCharacter,
		Personality: personality.Cranky,
		Phrase:      angusPhrase}
)
