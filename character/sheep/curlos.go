package sheep

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
	// curlosBirthday represents Curlos's birthday (May 8th).
	curlosBirthday = nook.Birthday{
		Day:   8,
		Month: time.May}
)

var (
	// curlosCode represents Curlos's unique code ("shp08").
	curlosCode = nook.Code{
		Value: "shp08"}
)

var (
	// curlosAmericanEnglishName represents Curlos's name in American English.
	curlosAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curlos"}

	// curlosCanadianFrenchName represents Curlos's name in Canadian French.
	curlosCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tonton"}

	// curlosDutchName represents Curlos's name in Dutch.
	curlosDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curlos"}

	// curlosFrenchName represents Curlos's name in French.
	curlosFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tonton"}

	// curlosGermanName represents Curlos's name in German.
	curlosGermanName = nook.Name{
		Language: language.German,
		Value:    "Locke"}

	// curlosItalianName represents Curlos's name in Italian.
	curlosItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Boccolo"}

	// curlosJapaneseName represents Curlos's name in Japanese.
	curlosJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カルロス"}

	// curlosKoreanName represents Curlos's name in Korean.
	curlosKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카를로스"}

	// curlosLatinAmericanSpanishName represents Curlos's name in Latin American Spanish.
	curlosLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carnerio"}

	// curlosRussianName represents Curlos's name in Russian.
	curlosRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керлос"}

	// curlosSimplifiedChineseName represents Curlos's name in Simplified Chinese.
	curlosSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾洛斯"}

	// curlosSpanishName represents Curlos's name in Spanish.
	curlosSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carnerio"}

	// curlosTraditionalChineseName represents Curlos's name in Traditional Chinese.
	curlosTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈洛斯"}
)

var (
	// curlosName represents Curlos's name in different languages.
	curlosName = nook.Languages{
		language.AmericanEnglish:      curlosAmericanEnglishName,
		language.CanadianFrench:       curlosCanadianFrenchName,
		language.Dutch:                curlosDutchName,
		language.French:               curlosFrenchName,
		language.German:               curlosGermanName,
		language.Italian:              curlosItalianName,
		language.Japanese:             curlosJapaneseName,
		language.Korean:               curlosKoreanName,
		language.LatinAmericanSpanish: curlosLatinAmericanSpanishName,
		language.Russian:              curlosRussianName,
		language.SimplifiedChinese:    curlosSimplifiedChineseName,
		language.Spanish:              curlosSpanishName,
		language.TraditionalChinese:   curlosTraditionalChineseName}
)

var (
	// curlosCharacter represents Curlos's character information.
	curlosCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: curlosBirthday,
		Code:     curlosCode,
		Key:      character.Curlos,
		Gender:   gender.Male,
		Name:     curlosName,
		Special:  false}
)

var (
	// curlosAmericanEnglishPhrase represents Curlos's phrase in American English.
	curlosAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shearly"}

	// curlosCanadianFrenchPhrase represents Curlos's phrase in Canadian French.
	curlosCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toison"}

	// curlosDutchPhrase represents Curlos's phrase in Dutch.
	curlosDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "scheerlijk"}

	// curlosFrenchPhrase represents Curlos's phrase in French.
	curlosFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toison"}

	// curlosGermanPhrase represents Curlos's phrase in German.
	curlosGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zottel"}

	// curlosItalianPhrase represents Curlos's phrase in Italian.
	curlosItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hermoso"}

	// curlosJapanesePhrase represents Curlos's phrase in Japanese.
	curlosJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ベイビー"}

	// curlosKoreanPhrase represents Curlos's phrase in Korean.
	curlosKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "자기야"}

	// curlosLatinAmericanSpanishPhrase represents Curlos's phrase in Latin American Spanish.
	curlosLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "laaaná"}

	// curlosRussianPhrase represents Curlos's phrase in Russian.
	curlosRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "завитушка"}

	// curlosSimplifiedChinesePhrase represents Curlos's phrase in Simplified Chinese.
	curlosSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "北鼻"}

	// curlosSpanishPhrase represents Curlos's phrase in Spanish.
	curlosSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lanitas"}

	// curlosTraditionalChinesePhrase represents Curlos's phrase in Traditional Chinese.
	curlosTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "北鼻"}
)

var (
	// curlosPhrase represents Curlos's phrases in different languages.
	curlosPhrase = nook.Languages{
		language.AmericanEnglish:      curlosAmericanEnglishPhrase,
		language.CanadianFrench:       curlosCanadianFrenchPhrase,
		language.Dutch:                curlosDutchPhrase,
		language.French:               curlosFrenchPhrase,
		language.German:               curlosGermanPhrase,
		language.Italian:              curlosItalianPhrase,
		language.Japanese:             curlosJapanesePhrase,
		language.Korean:               curlosKoreanPhrase,
		language.LatinAmericanSpanish: curlosLatinAmericanSpanishPhrase,
		language.Russian:              curlosRussianPhrase,
		language.SimplifiedChinese:    curlosSimplifiedChinesePhrase,
		language.Spanish:              curlosSpanishPhrase,
		language.TraditionalChinese:   curlosTraditionalChinesePhrase}
)

var (
	// Curlos represents the character Curlos with his complete information.
	Curlos = nook.Villager{
		Character:   curlosCharacter,
		Personality: personality.Smug,
		Phrase:      curlosPhrase}
)
