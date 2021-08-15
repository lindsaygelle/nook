package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sandyBirthday = nook.Birthday{
		Day:   21,
		Month: time.October}
)

var (
	sandyCode = nook.Code{
		Value: "ost02"}
)

var (
	sandyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sandy"}

	sandyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ottieeuh quoi"}

	sandyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sandyraprap"}

	sandyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ottieeuh quoi"}

	sandyGermanName = nook.Name{
		Language: language.German,
		Value:    "Sentaflottilda"}

	sandyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alessiavruum"}

	sandyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ランだチョー"}

	sandyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샌디그렇다조"}

	sandyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Priscilafiuuun"}

	sandyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сэндибыстро"}

	sandySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小岚鸵鸟"}

	sandySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Priscilafiuuun"}

	sandyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小嵐鴕鳥"}
)

var (
	sandyName = nook.Languages{
		language.AmericanEnglish:      sandyAmericanEnglishName,
		language.CanadianFrench:       sandyCanadianFrenchName,
		language.Dutch:                sandyDutchName,
		language.French:               sandyFrenchName,
		language.German:               sandyGermanName,
		language.Italian:              sandyItalianName,
		language.Japanese:             sandyJapaneseName,
		language.Korean:               sandyKoreanName,
		language.LatinAmericanSpanish: sandyLatinAmericanSpanishName,
		language.Russian:              sandyRussianName,
		language.SimplifiedChinese:    sandySimplifiedChineseName,
		language.Spanish:              sandySpanishName,
		language.TraditionalChinese:   sandyTraditionalChineseName}
)

var (
	sandyCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: sandyBirthday,
		Code:     sandyCode,
		Gender:   nook.Female,
		Name:     sandyName}
)

var (
	sandyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "speedy"}

	sandyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "euh quoi"}

	sandyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "raprap"}

	sandyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "euh quoi"}

	sandyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flottilda"}

	sandyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vruum"}

	sandyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だチョー"}

	sandyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇다조"}

	sandyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fiuuun"}

	sandyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "быстро"}

	sandySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鸵鸟"}

	sandySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fiuuun"}

	sandyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鴕鳥"}
)

var (
	sandyPhrase = nook.Languages{
		language.AmericanEnglish:      sandyAmericanEnglishName,
		language.CanadianFrench:       sandyCanadianFrenchName,
		language.Dutch:                sandyDutchName,
		language.French:               sandyFrenchName,
		language.German:               sandyGermanName,
		language.Italian:              sandyItalianName,
		language.Japanese:             sandyJapaneseName,
		language.Korean:               sandyKoreanName,
		language.LatinAmericanSpanish: sandyLatinAmericanSpanishName,
		language.Russian:              sandyRussianName,
		language.SimplifiedChinese:    sandySimplifiedChineseName,
		language.Spanish:              sandySpanishName,
		language.TraditionalChinese:   sandyTraditionalChineseName}
)

var (
	Sandy = nook.Villager{
		Character:   sandyCharacter,
		Personality: nook.Normal,
		Phrase:      sandyPhrase}
)
