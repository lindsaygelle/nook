package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	quetzalBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	quetzalCode = nook.Code{
		Value: ""}
)

var (
	quetzalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Quetzal"}

	quetzalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	quetzalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	quetzalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Quetzalcouêêêtzal"}

	quetzalGermanName = nook.Name{
		Language: language.German,
		Value:    "EnriqueKRIKRI"}

	quetzalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "QuetzalOCCHIO"}

	quetzalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハチェットゲロッパ"}

	quetzalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	quetzalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	quetzalRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	quetzalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "战斧喳喳"}

	quetzalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Quetzalcuiii"}

	quetzalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	quetzalName = nook.Languages{
		language.AmericanEnglish:      quetzalAmericanEnglishName,
		language.CanadianFrench:       quetzalCanadianFrenchName,
		language.Dutch:                quetzalDutchName,
		language.French:               quetzalFrenchName,
		language.German:               quetzalGermanName,
		language.Italian:              quetzalItalianName,
		language.Japanese:             quetzalJapaneseName,
		language.Korean:               quetzalKoreanName,
		language.LatinAmericanSpanish: quetzalLatinAmericanSpanishName,
		language.Russian:              quetzalRussianName,
		language.SimplifiedChinese:    quetzalSimplifiedChineseName,
		language.Spanish:              quetzalSpanishName,
		language.TraditionalChinese:   quetzalTraditionalChineseName}
)

var (
	quetzalCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: quetzalBirthday,
		Code:     quetzalCode,
		Gender:   nook.Male,
		Name:     quetzalName}
)

var (
	quetzalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	quetzalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	quetzalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	quetzalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "couêêêtzal"}

	quetzalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "KRIKRI"}

	quetzalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "OCCHIO"}

	quetzalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゲロッパ"}

	quetzalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	quetzalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	quetzalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	quetzalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳喳"}

	quetzalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuiii"}

	quetzalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	quetzalPhrase = nook.Languages{
		language.AmericanEnglish:      quetzalAmericanEnglishName,
		language.CanadianFrench:       quetzalCanadianFrenchName,
		language.Dutch:                quetzalDutchName,
		language.French:               quetzalFrenchName,
		language.German:               quetzalGermanName,
		language.Italian:              quetzalItalianName,
		language.Japanese:             quetzalJapaneseName,
		language.Korean:               quetzalKoreanName,
		language.LatinAmericanSpanish: quetzalLatinAmericanSpanishName,
		language.Russian:              quetzalRussianName,
		language.SimplifiedChinese:    quetzalSimplifiedChineseName,
		language.Spanish:              quetzalSpanishName,
		language.TraditionalChinese:   quetzalTraditionalChineseName}
)

var (
	Quetzal = nook.Villager{
		Character:   quetzalCharacter,
		Personality: nook.Jock,
		Phrase:      quetzalPhrase}
)
