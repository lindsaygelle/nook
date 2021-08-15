package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	filbertBirthday = nook.Birthday{
		Day:   3,
		Month: time.June}
)

var (
	filbertCode = nook.Code{
		Value: "squ02"}
)

var (
	filbertAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Filbert"}

	filbertCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Filibertgogo"}

	filbertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Filbertvent"}

	filbertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Filibertgogo"}

	filbertGermanName = nook.Name{
		Language: language.German,
		Value:    "Felixmanno"}

	filbertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Scricciognic gnic"}

	filbertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リッキーでしゅ"}

	filbertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리키예용"}

	filbertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Filbertoguirlache"}

	filbertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Филбертдружище"}

	filbertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "瑞奇是哦"}

	filbertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Filbertoguirlache"}

	filbertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑞奇是囉"}
)

var (
	filbertName = nook.Languages{
		language.AmericanEnglish:      filbertAmericanEnglishName,
		language.CanadianFrench:       filbertCanadianFrenchName,
		language.Dutch:                filbertDutchName,
		language.French:               filbertFrenchName,
		language.German:               filbertGermanName,
		language.Italian:              filbertItalianName,
		language.Japanese:             filbertJapaneseName,
		language.Korean:               filbertKoreanName,
		language.LatinAmericanSpanish: filbertLatinAmericanSpanishName,
		language.Russian:              filbertRussianName,
		language.SimplifiedChinese:    filbertSimplifiedChineseName,
		language.Spanish:              filbertSpanishName,
		language.TraditionalChinese:   filbertTraditionalChineseName}
)

var (
	filbertCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: filbertBirthday,
		Code:     filbertCode,
		Gender:   nook.Male,
		Name:     filbertName}
)

var (
	filbertAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bucko"}

	filbertCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gogo"}

	filbertDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vent"}

	filbertFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gogo"}

	filbertGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "manno"}

	filbertItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnic gnic"}

	filbertJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしゅ"}

	filbertKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예용"}

	filbertLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guirlache"}

	filbertRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дружище"}

	filbertSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是哦"}

	filbertSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guirlache"}

	filbertTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是囉"}
)

var (
	filbertPhrase = nook.Languages{
		language.AmericanEnglish:      filbertAmericanEnglishName,
		language.CanadianFrench:       filbertCanadianFrenchName,
		language.Dutch:                filbertDutchName,
		language.French:               filbertFrenchName,
		language.German:               filbertGermanName,
		language.Italian:              filbertItalianName,
		language.Japanese:             filbertJapaneseName,
		language.Korean:               filbertKoreanName,
		language.LatinAmericanSpanish: filbertLatinAmericanSpanishName,
		language.Russian:              filbertRussianName,
		language.SimplifiedChinese:    filbertSimplifiedChineseName,
		language.Spanish:              filbertSpanishName,
		language.TraditionalChinese:   filbertTraditionalChineseName}
)

var (
	Filbert = nook.Villager{
		Character:   filbertCharacter,
		Personality: nook.Lazy,
		Phrase:      filbertPhrase}
)
