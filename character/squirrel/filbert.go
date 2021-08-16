package squirrel

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
		Value:    "Filibert"}

	filbertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Filbert"}

	filbertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Filibert"}

	filbertGermanName = nook.Name{
		Language: language.German,
		Value:    "Felix"}

	filbertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Scriccio"}

	filbertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リッキー"}

	filbertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리키"}

	filbertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Filberto"}

	filbertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Филберт"}

	filbertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "瑞奇"}

	filbertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Filberto"}

	filbertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑞奇"}
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
		Animal:   animal.Squirrel,
		Birthday: filbertBirthday,
		Code:     filbertCode,
		Key:      character.Filbert,
		Gender:   gender.Male,
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
		Personality: personality.Lazy,
		Phrase:      filbertPhrase}
)
