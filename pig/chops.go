package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chopsBirthday = nook.Birthday{
		Day:   13,
		Month: time.October}
)

var (
	chopsCode = nook.Code{
		Value: "pig14"}
)

var (
	chopsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chops"}

	chopsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aaronlardon"}

	chopsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chopsmoddervet"}

	chopsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aaronlardon"}

	chopsGermanName = nook.Name{
		Language: language.German,
		Value:    "Clemensborsti"}

	chopsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Linooinc"}

	chopsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トンファンだトン"}

	chopsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "돈후앙떼끼에로"}

	chopsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Solomínzoink"}

	chopsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чопсхряк-с"}

	chopsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豚皇豚"}

	chopsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Solomínzoink"}

	chopsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豚皇豚"}
)

var (
	chopsName = nook.Languages{
		language.AmericanEnglish:      chopsAmericanEnglishName,
		language.CanadianFrench:       chopsCanadianFrenchName,
		language.Dutch:                chopsDutchName,
		language.French:               chopsFrenchName,
		language.German:               chopsGermanName,
		language.Italian:              chopsItalianName,
		language.Japanese:             chopsJapaneseName,
		language.Korean:               chopsKoreanName,
		language.LatinAmericanSpanish: chopsLatinAmericanSpanishName,
		language.Russian:              chopsRussianName,
		language.SimplifiedChinese:    chopsSimplifiedChineseName,
		language.Spanish:              chopsSpanishName,
		language.TraditionalChinese:   chopsTraditionalChineseName}
)

var (
	chopsCharacter = nook.Character{
		Animal:   Pig,
		Birthday: chopsBirthday,
		Code:     chopsCode,
		Gender:   nook.Male,
		Name:     chopsName}
)

var (
	chopsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zoink"}

	chopsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lardon"}

	chopsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "moddervet"}

	chopsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lardon"}

	chopsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "borsti"}

	chopsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oinc"}

	chopsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だトン"}

	chopsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "떼끼에로"}

	chopsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zoink"}

	chopsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк-с"}

	chopsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豚"}

	chopsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zoink"}

	chopsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豚"}
)

var (
	chopsPhrase = nook.Languages{
		language.AmericanEnglish:      chopsAmericanEnglishName,
		language.CanadianFrench:       chopsCanadianFrenchName,
		language.Dutch:                chopsDutchName,
		language.French:               chopsFrenchName,
		language.German:               chopsGermanName,
		language.Italian:              chopsItalianName,
		language.Japanese:             chopsJapaneseName,
		language.Korean:               chopsKoreanName,
		language.LatinAmericanSpanish: chopsLatinAmericanSpanishName,
		language.Russian:              chopsRussianName,
		language.SimplifiedChinese:    chopsSimplifiedChineseName,
		language.Spanish:              chopsSpanishName,
		language.TraditionalChinese:   chopsTraditionalChineseName}
)

var (
	Chops = nook.Villager{
		Character:   chopsCharacter,
		Personality: nook.Smug,
		Phrase:      chopsPhrase}
)
