package rhinoceros

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
	tiaraBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	tiaraCode = nook.Code{
		Value: ""}
)

var (
	tiaraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tiara"}

	tiaraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	tiaraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	tiaraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tiara"}

	tiaraGermanName = nook.Name{
		Language: language.German,
		Value:    "Roswitha"}

	tiaraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cornelia"}

	tiaraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ワイハ"}

	tiaraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	tiaraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	tiaraRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	tiaraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "锋锋"}

	tiaraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tiara"}

	tiaraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	tiaraName = nook.Languages{
		language.AmericanEnglish:      tiaraAmericanEnglishName,
		language.CanadianFrench:       tiaraCanadianFrenchName,
		language.Dutch:                tiaraDutchName,
		language.French:               tiaraFrenchName,
		language.German:               tiaraGermanName,
		language.Italian:              tiaraItalianName,
		language.Japanese:             tiaraJapaneseName,
		language.Korean:               tiaraKoreanName,
		language.LatinAmericanSpanish: tiaraLatinAmericanSpanishName,
		language.Russian:              tiaraRussianName,
		language.SimplifiedChinese:    tiaraSimplifiedChineseName,
		language.Spanish:              tiaraSpanishName,
		language.TraditionalChinese:   tiaraTraditionalChineseName}
)

var (
	tiaraCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: tiaraBirthday,
		Code:     tiaraCode,
		Key:      character.Tiara,
		Gender:   gender.Female,
		Name:     tiaraName,
		Special:  false}
)

var (
	tiaraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lovey"}

	tiaraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	tiaraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	tiaraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon hippy"}

	tiaraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "herzchen"}

	tiaraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dolcezza"}

	tiaraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てゆーか"}

	tiaraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	tiaraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	tiaraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	tiaraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀嗨"}

	tiaraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "amorcito"}

	tiaraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	tiaraPhrase = nook.Languages{
		language.AmericanEnglish:      tiaraAmericanEnglishPhrase,
		language.CanadianFrench:       tiaraCanadianFrenchPhrase,
		language.Dutch:                tiaraDutchPhrase,
		language.French:               tiaraFrenchPhrase,
		language.German:               tiaraGermanPhrase,
		language.Italian:              tiaraItalianPhrase,
		language.Japanese:             tiaraJapanesePhrase,
		language.Korean:               tiaraKoreanPhrase,
		language.LatinAmericanSpanish: tiaraLatinAmericanSpanishPhrase,
		language.Russian:              tiaraRussianPhrase,
		language.SimplifiedChinese:    tiaraSimplifiedChinesePhrase,
		language.Spanish:              tiaraSpanishPhrase,
		language.TraditionalChinese:   tiaraTraditionalChinesePhrase}
)

var (
	Tiara = nook.Villager{
		Character:   tiaraCharacter,
		Personality: personality.Snooty,
		Phrase:      tiaraPhrase}
)
