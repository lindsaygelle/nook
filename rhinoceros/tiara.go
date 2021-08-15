package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	tiaraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	tiaraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	tiaraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	tiaraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "锋锋"}

	tiaraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tiara"}

	tiaraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Rhinoceros,
		Birthday: tiaraBirthday,
		Code:     tiaraCode,
		Gender:   nook.Female,
		Name:     tiaraName}
)

var (
	tiaraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	tiaraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	tiaraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	tiaraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	tiaraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	tiaraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀嗨"}

	tiaraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "amorcito"}

	tiaraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	tiaraPhrase = nook.Languages{
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
	Tiara = nook.Villager{
		Character:   tiaraCharacter,
		Personality: nook.Snooty,
		Phrase:      tiaraPhrase}
)
