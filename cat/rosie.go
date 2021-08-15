package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rosieBirthday = nook.Birthday{
		Day:   27,
		Month: time.February}
)

var (
	rosieCode = nook.Code{
		Value: "cat02"}
)

var (
	rosieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rosie"}

	rosieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosieflûte"}

	rosieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rosiegekkie"}

	rosieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosieflûte"}

	rosieGermanName = nook.Name{
		Language: language.German,
		Value:    "Sophieflöt"}

	rosieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grinfiatontolon"}

	rosieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブーケチェキ"}

	rosieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "부케헤이"}

	rosieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mininamiaaau"}

	rosieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Розирыбка"}

	rosieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彭花看看"}

	rosieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mininamiaaau"}

	rosieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彭花看看"}
)

var (
	rosieName = nook.Languages{
		language.AmericanEnglish:      rosieAmericanEnglishName,
		language.CanadianFrench:       rosieCanadianFrenchName,
		language.Dutch:                rosieDutchName,
		language.French:               rosieFrenchName,
		language.German:               rosieGermanName,
		language.Italian:              rosieItalianName,
		language.Japanese:             rosieJapaneseName,
		language.Korean:               rosieKoreanName,
		language.LatinAmericanSpanish: rosieLatinAmericanSpanishName,
		language.Russian:              rosieRussianName,
		language.SimplifiedChinese:    rosieSimplifiedChineseName,
		language.Spanish:              rosieSpanishName,
		language.TraditionalChinese:   rosieTraditionalChineseName}
)

var (
	rosieCharacter = nook.Character{
		Animal:   Cat,
		Birthday: rosieBirthday,
		Code:     rosieCode,
		Gender:   nook.Female,
		Name:     rosieName}
)

var (
	rosieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "silly"}

	rosieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flûte"}

	rosieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gekkie"}

	rosieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flûte"}

	rosieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flöt"}

	rosieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tontolon"}

	rosieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チェキ"}

	rosieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "헤이"}

	rosieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miaaau"}

	rosieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "рыбка"}

	rosieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "看看"}

	rosieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miaaau"}

	rosieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "看看"}
)

var (
	rosiePhrase = nook.Languages{
		language.AmericanEnglish:      rosieAmericanEnglishName,
		language.CanadianFrench:       rosieCanadianFrenchName,
		language.Dutch:                rosieDutchName,
		language.French:               rosieFrenchName,
		language.German:               rosieGermanName,
		language.Italian:              rosieItalianName,
		language.Japanese:             rosieJapaneseName,
		language.Korean:               rosieKoreanName,
		language.LatinAmericanSpanish: rosieLatinAmericanSpanishName,
		language.Russian:              rosieRussianName,
		language.SimplifiedChinese:    rosieSimplifiedChineseName,
		language.Spanish:              rosieSpanishName,
		language.TraditionalChinese:   rosieTraditionalChineseName}
)

var (
	Rosie = nook.Villager{
		Character:   rosieCharacter,
		Personality: nook.Peppy,
		Phrase:      rosiePhrase}
)
