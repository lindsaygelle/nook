package rabbit

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
	// bunnieBirthday represents Bunnie's birthday.
	bunnieBirthday = nook.Birthday{
		Day:   9,
		Month: time.May}
)

var (
	// bunnieCode represents Bunnie's unique code ("rbt00").
	bunnieCode = nook.Code{
		Value: "rbt00"}
)

var (
	// bunnieAmericanEnglishName represents Bunnie's name in American English.
	bunnieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bunnie"}

	// bunnieCanadianFrenchName represents Bunnie's name in Canadian French.
	bunnieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Clara"}

	// bunnieDutchName represents Bunnie's name in Dutch.
	bunnieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bunnie"}

	// bunnieFrenchName represents Bunnie's name in French.
	bunnieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Clara"}

	// bunnieGermanName represents Bunnie's name in German.
	bunnieGermanName = nook.Name{
		Language: language.German,
		Value:    "Mimmi"}

	// bunnieItalianName represents Bunnie's name in Italian.
	bunnieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bonny"}

	// bunnieJapaneseName represents Bunnie's name in Japanese.
	bunnieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リリアン"}

	// bunnieKoreanName represents Bunnie's name in Korean.
	bunnieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴리안"}

	// bunnieLatinAmericanSpanishName represents Bunnie's name in Latin American Spanish.
	bunnieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Coni"}

	// bunnieRussianName represents Bunnie's name in Russian.
	bunnieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Банни"}

	// bunnieSimplifiedChineseName represents Bunnie's name in Simplified Chinese.
	bunnieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莉莉安"}

	// bunnieSpanishName represents Bunnie's name in Spanish.
	bunnieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Coni"}

	// bunnieTraditionalChineseName represents Bunnie's name in Traditional Chinese.
	bunnieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莉莉安"}
)

var (
	// bunnieName represents Bunnie's name in different languages.
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
	// bunnieCharacter represents Bunnie's character information.
	bunnieCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: bunnieBirthday,
		Code:     bunnieCode,
		Key:      character.Bunnie,
		Gender:   gender.Female,
		Name:     bunnieName,
		Special:  false}
)

var (
	// bunnieAmericanEnglishPhrase represents Bunnie's phrase in American English.
	bunnieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tee-hee"}

	// bunnieCanadianFrenchPhrase represents Bunnie's phrase in Canadian French.
	bunnieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hé hé hé"}

	// bunnieDutchPhrase represents Bunnie's phrase in Dutch.
	bunnieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hihi"}

	// bunnieFrenchPhrase represents Bunnie's phrase in French.
	bunnieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hé hé hé"}

	// bunnieGermanPhrase represents Bunnie's phrase in German.
	bunnieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tiihii"}

	// bunnieItalianPhrase represents Bunnie's phrase in Italian.
	bunnieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tiptap"}

	// bunnieJapanesePhrase represents Bunnie's phrase in Japanese.
	bunnieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "みたいな"}

	// bunnieKoreanPhrase represents Bunnie's phrase in Korean.
	bunnieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇지뭐"}

	// bunnieLatinAmericanSpanishPhrase represents Bunnie's phrase in Latin American Spanish.
	bunnieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oy-oy-oy"}

	// bunnieRussianPhrase represents Bunnie's phrase in Russian.
	bunnieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хи-хи"}

	// bunnieSimplifiedChinesePhrase represents Bunnie's phrase in Simplified Chinese.
	bunnieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好像喔"}

	// bunnieSpanishPhrase represents Bunnie's phrase in Spanish.
	bunnieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oy-oy-oy"}

	// bunnieTraditionalChinesePhrase represents Bunnie's phrase in Traditional Chinese.
	bunnieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好像喔"}
)

var (
	// bunniePhrase represents Bunnie's phrase in different languages.
	bunniePhrase = nook.Languages{
		language.AmericanEnglish:      bunnieAmericanEnglishPhrase,
		language.CanadianFrench:       bunnieCanadianFrenchPhrase,
		language.Dutch:                bunnieDutchPhrase,
		language.French:               bunnieFrenchPhrase,
		language.German:               bunnieGermanPhrase,
		language.Italian:              bunnieItalianPhrase,
		language.Japanese:             bunnieJapanesePhrase,
		language.Korean:               bunnieKoreanPhrase,
		language.LatinAmericanSpanish: bunnieLatinAmericanSpanishPhrase,
		language.Russian:              bunnieRussianPhrase,
		language.SimplifiedChinese:    bunnieSimplifiedChinesePhrase,
		language.Spanish:              bunnieSpanishPhrase,
		language.TraditionalChinese:   bunnieTraditionalChinesePhrase}
)

var (
	// Bunnie represents the character Bunnie with her complete information.
	Bunnie = nook.Villager{
		Character:   bunnieCharacter,
		Personality: personality.Peppy,
		Phrase:      bunniePhrase}
)
