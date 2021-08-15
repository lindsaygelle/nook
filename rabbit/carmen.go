package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	carmenBirthday = nook.Birthday{
		Day:   6,
		Month: time.January}
)

var (
	carmenCode = nook.Code{
		Value: "rbt16"}
)

var (
	carmenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carmen"}

	carmenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Zoénougat"}

	carmenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Carmennoga"}

	carmenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zoépinou-nou"}

	carmenGermanName = nook.Name{
		Language: language.German,
		Value:    "Hildalangohr"}

	carmenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Connygnuf gnuf"}

	carmenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョコまじで"}

	carmenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "초코진짜진짜"}

	carmenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lindanufinuf"}

	carmenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Карменкарамелька"}

	carmenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧蔻真的"}

	carmenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lindaarcoíris"}

	carmenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧蔻真的"}
)

var (
	carmenName = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishName,
		language.CanadianFrench:       carmenCanadianFrenchName,
		language.Dutch:                carmenDutchName,
		language.French:               carmenFrenchName,
		language.German:               carmenGermanName,
		language.Italian:              carmenItalianName,
		language.Japanese:             carmenJapaneseName,
		language.Korean:               carmenKoreanName,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishName,
		language.Russian:              carmenRussianName,
		language.SimplifiedChinese:    carmenSimplifiedChineseName,
		language.Spanish:              carmenSpanishName,
		language.TraditionalChinese:   carmenTraditionalChineseName}
)

var (
	carmenCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: carmenBirthday,
		Code:     carmenCode,
		Gender:   nook.Female,
		Name:     carmenName}
)

var (
	carmenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nougat"}

	carmenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nougat"}

	carmenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "noga"}

	carmenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pinou-nou"}

	carmenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "langohr"}

	carmenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnuf gnuf"}

	carmenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まじで"}

	carmenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진짜진짜"}

	carmenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nufinuf"}

	carmenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "карамелька"}

	carmenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真的"}

	carmenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "arcoíris"}

	carmenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真的"}
)

var (
	carmenPhrase = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishName,
		language.CanadianFrench:       carmenCanadianFrenchName,
		language.Dutch:                carmenDutchName,
		language.French:               carmenFrenchName,
		language.German:               carmenGermanName,
		language.Italian:              carmenItalianName,
		language.Japanese:             carmenJapaneseName,
		language.Korean:               carmenKoreanName,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishName,
		language.Russian:              carmenRussianName,
		language.SimplifiedChinese:    carmenSimplifiedChineseName,
		language.Spanish:              carmenSpanishName,
		language.TraditionalChinese:   carmenTraditionalChineseName}
)

var (
	Carmen = nook.Villager{
		Character:   carmenCharacter,
		Personality: nook.Peppy,
		Phrase:      carmenPhrase}
)
