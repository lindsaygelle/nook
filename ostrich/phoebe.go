package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	phoebeBirthday = nook.Birthday{
		Day:   22,
		Month: time.April}
)

var (
	phoebeCode = nook.Code{
		Value: "ost10"}
)

var (
	phoebeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phoebe"}

	phoebeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Faustine"}

	phoebeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phoebe"}

	phoebeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Faustine"}

	phoebeGermanName = nook.Name{
		Language: language.German,
		Value:    "Philippa"}

	phoebeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fenicia"}

	phoebeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒノコ"}

	phoebeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "휘니"}

	phoebeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Avelina"}

	phoebeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фиби"}

	phoebeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "火鸟"}

	phoebeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Avelina"}

	phoebeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "火鳥"}
)

var (
	phoebeName = nook.Languages{
		language.AmericanEnglish:      phoebeAmericanEnglishName,
		language.CanadianFrench:       phoebeCanadianFrenchName,
		language.Dutch:                phoebeDutchName,
		language.French:               phoebeFrenchName,
		language.German:               phoebeGermanName,
		language.Italian:              phoebeItalianName,
		language.Japanese:             phoebeJapaneseName,
		language.Korean:               phoebeKoreanName,
		language.LatinAmericanSpanish: phoebeLatinAmericanSpanishName,
		language.Russian:              phoebeRussianName,
		language.SimplifiedChinese:    phoebeSimplifiedChineseName,
		language.Spanish:              phoebeSpanishName,
		language.TraditionalChinese:   phoebeTraditionalChineseName}
)

var (
	phoebeCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: phoebeBirthday,
		Code:     phoebeCode,
		Gender:   nook.Female,
		Name:     phoebeName}
)

var (
	phoebeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sparky"}

	phoebeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chanmé"}

	phoebeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vonkje"}

	phoebeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chanmé"}

	phoebeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pieks"}

	phoebeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "behmù"}

	phoebeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アツイわ"}

	phoebeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아뜨뜨"}

	phoebeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "equilicuá"}

	phoebeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "искорка"}

	phoebeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好热"}

	phoebeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "equilicuá"}

	phoebeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好熱"}
)

var (
	phoebePhrase = nook.Languages{
		language.AmericanEnglish:      phoebeAmericanEnglishName,
		language.CanadianFrench:       phoebeCanadianFrenchName,
		language.Dutch:                phoebeDutchName,
		language.French:               phoebeFrenchName,
		language.German:               phoebeGermanName,
		language.Italian:              phoebeItalianName,
		language.Japanese:             phoebeJapaneseName,
		language.Korean:               phoebeKoreanName,
		language.LatinAmericanSpanish: phoebeLatinAmericanSpanishName,
		language.Russian:              phoebeRussianName,
		language.SimplifiedChinese:    phoebeSimplifiedChineseName,
		language.Spanish:              phoebeSpanishName,
		language.TraditionalChinese:   phoebeTraditionalChineseName}
)

var (
	Phoebe = nook.Villager{
		Character:   phoebeCharacter,
		Personality: nook.BigSister,
		Phrase:      phoebePhrase}
)
