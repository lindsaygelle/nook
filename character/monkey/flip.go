package monkey

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
	flipBirthday = nook.Birthday{
		Day:   21,
		Month: time.November}
)

var (
	flipCode = nook.Code{
		Value: "mnk06"}
)

var (
	flipAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flip"}

	flipCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rudy"}

	flipDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flip"}

	flipFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rudy"}

	flipGermanName = nook.Name{
		Language: language.German,
		Value:    "Pippo"}

	flipItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Benny"}

	flipJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さすけ"}

	flipKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "원승"}

	flipLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Monet"}

	flipRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флип"}

	flipSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "天佑"}

	flipSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Monet"}

	flipTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "天佑"}
)

var (
	flipName = nook.Languages{
		language.AmericanEnglish:      flipAmericanEnglishName,
		language.CanadianFrench:       flipCanadianFrenchName,
		language.Dutch:                flipDutchName,
		language.French:               flipFrenchName,
		language.German:               flipGermanName,
		language.Italian:              flipItalianName,
		language.Japanese:             flipJapaneseName,
		language.Korean:               flipKoreanName,
		language.LatinAmericanSpanish: flipLatinAmericanSpanishName,
		language.Russian:              flipRussianName,
		language.SimplifiedChinese:    flipSimplifiedChineseName,
		language.Spanish:              flipSpanishName,
		language.TraditionalChinese:   flipTraditionalChineseName}
)

var (
	flipCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: flipBirthday,
		Code:     flipCode,
		Key:      character.Flip,
		Gender:   gender.Male,
		Name:     flipName}
)

var (
	flipAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rerack"}

	flipCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh hisse"}

	flipDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oe-⁠a-⁠a"}

	flipFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh hisse"}

	flipGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kikiki"}

	flipItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clap clap"}

	flipJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どっこい"}

	flipKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "빠샤"}

	flipLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "afuá"}

	flipRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "у-у-у"}

	flipSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "差不多"}

	flipSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "afuá"}

	flipTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "差不多"}
)

var (
	flipPhrase = nook.Languages{
		language.AmericanEnglish:      flipAmericanEnglishPhrase,
		language.CanadianFrench:       flipCanadianFrenchPhrase,
		language.Dutch:                flipDutchPhrase,
		language.French:               flipFrenchPhrase,
		language.German:               flipGermanPhrase,
		language.Italian:              flipItalianPhrase,
		language.Japanese:             flipJapanesePhrase,
		language.Korean:               flipKoreanPhrase,
		language.LatinAmericanSpanish: flipLatinAmericanSpanishPhrase,
		language.Russian:              flipRussianPhrase,
		language.SimplifiedChinese:    flipSimplifiedChinesePhrase,
		language.Spanish:              flipSpanishPhrase,
		language.TraditionalChinese:   flipTraditionalChinesePhrase}
)

var (
	Flip = nook.Villager{
		Character:   flipCharacter,
		Personality: personality.Jock,
		Phrase:      flipPhrase}
)
