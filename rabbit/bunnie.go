package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bunnieBirthday = nook.Birthday{
		Day:   9,
		Month: time.May}
)

var (
	bunnieCode = nook.Code{
		Value: "rbt00"}
)

var (
	bunnieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bunnie"}

	bunnieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Clara"}

	bunnieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bunnie"}

	bunnieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Clara"}

	bunnieGermanName = nook.Name{
		Language: language.German,
		Value:    "Mimmi"}

	bunnieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bonny"}

	bunnieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リリアン"}

	bunnieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴리안"}

	bunnieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Coni"}

	bunnieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Банни"}

	bunnieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莉莉安"}

	bunnieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Coni"}

	bunnieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莉莉安"}
)

var (
	bunnieName = nook.Languages{
		language.AmericanEnglish:      bunnieAmericanEnglishName,
		language.CanadianFrench:       bunnieCanadianFrenchName,
		language.Dutch:                bunnieDutchName,
		language.French:               bunnieFrenchName,
		language.German:               bunnieGermanName,
		language.Italian:              bunnieItalianName,
		language.Japanese:             bunnieJapaneseName,
		language.Korean:               bunnieKoreanName,
		language.LatinAmericanSpanish: bunnieLatinAmericanSpanishName,
		language.Russian:              bunnieRussianName,
		language.SimplifiedChinese:    bunnieSimplifiedChineseName,
		language.Spanish:              bunnieSpanishName,
		language.TraditionalChinese:   bunnieTraditionalChineseName}
)

var (
	bunnieCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: bunnieBirthday,
		Code:     bunnieCode,
		Gender:   nook.Female,
		Name:     bunnieName}
)

var (
	bunnieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tee-hee"}

	bunnieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hé hé hé"}

	bunnieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hihi"}

	bunnieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hé hé hé"}

	bunnieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tiihii"}

	bunnieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tiptap"}

	bunnieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "みたいな"}

	bunnieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇지뭐"}

	bunnieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oy-oy-oy"}

	bunnieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хи-хи"}

	bunnieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好像喔"}

	bunnieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oy-oy-oy"}

	bunnieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好像喔"}
)

var (
	bunniePhrase = nook.Languages{
		language.AmericanEnglish:      bunnieAmericanEnglishName,
		language.CanadianFrench:       bunnieCanadianFrenchName,
		language.Dutch:                bunnieDutchName,
		language.French:               bunnieFrenchName,
		language.German:               bunnieGermanName,
		language.Italian:              bunnieItalianName,
		language.Japanese:             bunnieJapaneseName,
		language.Korean:               bunnieKoreanName,
		language.LatinAmericanSpanish: bunnieLatinAmericanSpanishName,
		language.Russian:              bunnieRussianName,
		language.SimplifiedChinese:    bunnieSimplifiedChineseName,
		language.Spanish:              bunnieSpanishName,
		language.TraditionalChinese:   bunnieTraditionalChineseName}
)

var (
	Bunnie = nook.Villager{
		Character:   bunnieCharacter,
		Personality: nook.Peppy,
		Phrase:      bunniePhrase}
)
