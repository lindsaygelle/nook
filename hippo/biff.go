package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	biffBirthday = nook.Birthday{
		Day:   29,
		Month: time.March}
)

var (
	biffCode = nook.Code{
		Value: "hip04"}
)

var (
	biffAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Biff"}

	biffCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Biff"}

	biffDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Biff"}

	biffFrenchName = nook.Name{
		Language: language.French,
		Value:    "Biff"}

	biffGermanName = nook.Name{
		Language: language.German,
		Value:    "Norbert"}

	biffItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pippo"}

	biffJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガブリエル"}

	biffKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가브리엘"}

	biffLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pipo"}

	biffRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бифф"}

	biffSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾宝为"}

	biffSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pipo"}

	biffTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈寶為"}
)

var (
	biffName = nook.Languages{
		language.AmericanEnglish:      biffAmericanEnglishName,
		language.CanadianFrench:       biffCanadianFrenchName,
		language.Dutch:                biffDutchName,
		language.French:               biffFrenchName,
		language.German:               biffGermanName,
		language.Italian:              biffItalianName,
		language.Japanese:             biffJapaneseName,
		language.Korean:               biffKoreanName,
		language.LatinAmericanSpanish: biffLatinAmericanSpanishName,
		language.Russian:              biffRussianName,
		language.SimplifiedChinese:    biffSimplifiedChineseName,
		language.Spanish:              biffSpanishName,
		language.TraditionalChinese:   biffTraditionalChineseName}
)

var (
	biffCharacter = nook.Character{
		Animal:   Hippo,
		Birthday: biffBirthday,
		Code:     biffCode,
		Gender:   nook.Male,
		Name:     biffName}
)

var (
	biffAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squirt"}

	biffCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hips hips"}

	biffDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "spetter"}

	biffFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hips hips"}

	biffGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "skwiirt"}

	biffItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squirt"}

	biffJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃけん"}

	biffKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오도독"}

	biffLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hipo-pipo"}

	biffRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малек"}

	biffSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "因为"}

	biffSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hipo-pipo"}

	biffTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "因為"}
)

var (
	biffPhrase = nook.Languages{
		language.AmericanEnglish:      biffAmericanEnglishName,
		language.CanadianFrench:       biffCanadianFrenchName,
		language.Dutch:                biffDutchName,
		language.French:               biffFrenchName,
		language.German:               biffGermanName,
		language.Italian:              biffItalianName,
		language.Japanese:             biffJapaneseName,
		language.Korean:               biffKoreanName,
		language.LatinAmericanSpanish: biffLatinAmericanSpanishName,
		language.Russian:              biffRussianName,
		language.SimplifiedChinese:    biffSimplifiedChineseName,
		language.Spanish:              biffSpanishName,
		language.TraditionalChinese:   biffTraditionalChineseName}
)

var (
	Biff = nook.Villager{
		Character:   biffCharacter,
		Personality: nook.Jock,
		Phrase:      biffPhrase}
)
