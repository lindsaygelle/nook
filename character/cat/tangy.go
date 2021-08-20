package cat

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
	tangyBirthday = nook.Birthday{
		Day:   17,
		Month: time.June}
)

var (
	tangyCode = nook.Code{
		Value: "cat05"}
)

var (
	tangyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tangy"}

	tangyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marine"}

	tangyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tangy"}

	tangyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marine"}

	tangyGermanName = nook.Name{
		Language: language.German,
		Value:    "Tanja"}

	tangyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Arancina"}

	tangyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒャクパー"}

	tangyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "백프로"}

	tangyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tricia"}

	tangyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэнджи"}

	tangySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "纯纯"}

	tangySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tricia"}

	tangyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "純純"}
)

var (
	tangyName = nook.Languages{
		language.AmericanEnglish:      tangyAmericanEnglishName,
		language.CanadianFrench:       tangyCanadianFrenchName,
		language.Dutch:                tangyDutchName,
		language.French:               tangyFrenchName,
		language.German:               tangyGermanName,
		language.Italian:              tangyItalianName,
		language.Japanese:             tangyJapaneseName,
		language.Korean:               tangyKoreanName,
		language.LatinAmericanSpanish: tangyLatinAmericanSpanishName,
		language.Russian:              tangyRussianName,
		language.SimplifiedChinese:    tangySimplifiedChineseName,
		language.Spanish:              tangySpanishName,
		language.TraditionalChinese:   tangyTraditionalChineseName}
)

var (
	tangyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: tangyBirthday,
		Code:     tangyCode,
		Key:      character.Tangy,
		Gender:   gender.Female,
		Name:     tangyName,
		Special:  false}
)

var (
	tangyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "reeeeOWR"}

	tangyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "roooaR"}

	tangyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sinas"}

	tangyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "roooaR"}

	tangyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grraul"}

	tangyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miiiiaOU"}

	tangyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "みかん"}

	tangyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "귤귤"}

	tangyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mirriau"}

	tangyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фруктик"}

	tangySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蜜柑"}

	tangySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mandarinas"}

	tangyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蜜柑"}
)

var (
	tangyPhrase = nook.Languages{
		language.AmericanEnglish:      tangyAmericanEnglishPhrase,
		language.CanadianFrench:       tangyCanadianFrenchPhrase,
		language.Dutch:                tangyDutchPhrase,
		language.French:               tangyFrenchPhrase,
		language.German:               tangyGermanPhrase,
		language.Italian:              tangyItalianPhrase,
		language.Japanese:             tangyJapanesePhrase,
		language.Korean:               tangyKoreanPhrase,
		language.LatinAmericanSpanish: tangyLatinAmericanSpanishPhrase,
		language.Russian:              tangyRussianPhrase,
		language.SimplifiedChinese:    tangySimplifiedChinesePhrase,
		language.Spanish:              tangySpanishPhrase,
		language.TraditionalChinese:   tangyTraditionalChinesePhrase}
)

var (
	Tangy = nook.Villager{
		Character:   tangyCharacter,
		Personality: personality.Peppy,
		Phrase:      tangyPhrase}
)
