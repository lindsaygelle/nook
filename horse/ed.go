package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	edBirthday = nook.Birthday{
		Day:   16,
		Month: time.September}
)

var (
	edCode = nook.Code{
		Value: "hrs06"}
)

var (
	edAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ed"}

	edCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Édouardtagada"}

	edDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Edju"}

	edFrenchName = nook.Name{
		Language: language.French,
		Value:    "Édouardtagada"}

	edGermanName = nook.Name{
		Language: language.German,
		Value:    "Hermanngreenhorn"}

	edItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Codinopirulì"}

	edJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キザノホマレじゃない"}

	edKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "꺼벙게슴츠레"}

	edLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Crinaldojo-joy"}

	edRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эдновичок"}

	edSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马誉不是啦"}

	edSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Crinaldojo-joy"}

	edTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬譽不是啦"}
)

var (
	edName = nook.Languages{
		language.AmericanEnglish:      edAmericanEnglishName,
		language.CanadianFrench:       edCanadianFrenchName,
		language.Dutch:                edDutchName,
		language.French:               edFrenchName,
		language.German:               edGermanName,
		language.Italian:              edItalianName,
		language.Japanese:             edJapaneseName,
		language.Korean:               edKoreanName,
		language.LatinAmericanSpanish: edLatinAmericanSpanishName,
		language.Russian:              edRussianName,
		language.SimplifiedChinese:    edSimplifiedChineseName,
		language.Spanish:              edSpanishName,
		language.TraditionalChinese:   edTraditionalChineseName}
)

var (
	edCharacter = nook.Character{
		Animal:   Horse,
		Birthday: edBirthday,
		Code:     edCode,
		Gender:   nook.Male,
		Name:     edName}
)

var (
	edAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "greenhorn"}

	edCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tagada"}

	edDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ju"}

	edFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tagada"}

	edGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "greenhorn"}

	edItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pirulì"}

	edJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃない"}

	edKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "게슴츠레"}

	edLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jo-joy"}

	edRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "новичок"}

	edSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不是啦"}

	edSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jo-joy"}

	edTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不是啦"}
)

var (
	edPhrase = nook.Languages{
		language.AmericanEnglish:      edAmericanEnglishName,
		language.CanadianFrench:       edCanadianFrenchName,
		language.Dutch:                edDutchName,
		language.French:               edFrenchName,
		language.German:               edGermanName,
		language.Italian:              edItalianName,
		language.Japanese:             edJapaneseName,
		language.Korean:               edKoreanName,
		language.LatinAmericanSpanish: edLatinAmericanSpanishName,
		language.Russian:              edRussianName,
		language.SimplifiedChinese:    edSimplifiedChineseName,
		language.Spanish:              edSpanishName,
		language.TraditionalChinese:   edTraditionalChineseName}
)

var (
	Ed = nook.Villager{
		Character:   edCharacter,
		Personality: nook.Smug,
		Phrase:      edPhrase}
)
