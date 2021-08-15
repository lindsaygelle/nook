package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rizzoBirthday = nook.Birthday{
		Day:   17,
		Month: time.January}
)

var (
	rizzoCode = nook.Code{
		Value: "mus09"}
)

var (
	rizzoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rizzo"}

	rizzoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sourizzinouik"}

	rizzoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rizzopiephoi"}

	rizzoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sourizzinouik"}

	rizzoGermanName = nook.Name{
		Language: language.German,
		Value:    "Rickyfieps"}

	rizzoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rizzosquiii"}

	rizzoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちょろきちがってん"}

	rizzoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "조르쥐힐끔힐끔"}

	rizzoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ratolónñiiiik"}

	rizzoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риццопи-и-и"}

	rizzoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "悄俊蹑蹑"}

	rizzoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ratolónñiiiik"}

	rizzoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "悄俊躡躡"}
)

var (
	rizzoName = nook.Languages{
		language.AmericanEnglish:      rizzoAmericanEnglishName,
		language.CanadianFrench:       rizzoCanadianFrenchName,
		language.Dutch:                rizzoDutchName,
		language.French:               rizzoFrenchName,
		language.German:               rizzoGermanName,
		language.Italian:              rizzoItalianName,
		language.Japanese:             rizzoJapaneseName,
		language.Korean:               rizzoKoreanName,
		language.LatinAmericanSpanish: rizzoLatinAmericanSpanishName,
		language.Russian:              rizzoRussianName,
		language.SimplifiedChinese:    rizzoSimplifiedChineseName,
		language.Spanish:              rizzoSpanishName,
		language.TraditionalChinese:   rizzoTraditionalChineseName}
)

var (
	rizzoCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: rizzoBirthday,
		Code:     rizzoCode,
		Gender:   nook.Male,
		Name:     rizzoName}
)

var (
	rizzoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squee"}

	rizzoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nouik"}

	rizzoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piephoi"}

	rizzoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nouik"}

	rizzoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fieps"}

	rizzoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squiii"}

	rizzoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "がってん"}

	rizzoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "힐끔힐끔"}

	rizzoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiiiik"}

	rizzoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пи-и-и"}

	rizzoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹑蹑"}

	rizzoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñiiiik"}

	rizzoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "躡躡"}
)

var (
	rizzoPhrase = nook.Languages{
		language.AmericanEnglish:      rizzoAmericanEnglishName,
		language.CanadianFrench:       rizzoCanadianFrenchName,
		language.Dutch:                rizzoDutchName,
		language.French:               rizzoFrenchName,
		language.German:               rizzoGermanName,
		language.Italian:              rizzoItalianName,
		language.Japanese:             rizzoJapaneseName,
		language.Korean:               rizzoKoreanName,
		language.LatinAmericanSpanish: rizzoLatinAmericanSpanishName,
		language.Russian:              rizzoRussianName,
		language.SimplifiedChinese:    rizzoSimplifiedChineseName,
		language.Spanish:              rizzoSpanishName,
		language.TraditionalChinese:   rizzoTraditionalChineseName}
)

var (
	Rizzo = nook.Villager{
		Character:   rizzoCharacter,
		Personality: nook.Cranky,
		Phrase:      rizzoPhrase}
)
