package pig

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
	curlyBirthday = nook.Birthday{
		Day:   26,
		Month: time.July}
)

var (
	curlyCode = nook.Code{
		Value: "pig00"}
)

var (
	curlyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curly"}

	curlyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tirbou"}

	curlyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curly"}

	curlyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tirbou"}

	curlyGermanName = nook.Name{
		Language: language.German,
		Value:    "Oink"}

	curlyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ricciolo"}

	curlyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムカツ"}

	curlyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄까스"}

	curlyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rufueto"}

	curlyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керли"}

	curlySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胜利"}

	curlySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rufueto"}

	curlyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "勝利"}
)

var (
	curlyName = nook.Languages{
		language.AmericanEnglish:      curlyAmericanEnglishName,
		language.CanadianFrench:       curlyCanadianFrenchName,
		language.Dutch:                curlyDutchName,
		language.French:               curlyFrenchName,
		language.German:               curlyGermanName,
		language.Italian:              curlyItalianName,
		language.Japanese:             curlyJapaneseName,
		language.Korean:               curlyKoreanName,
		language.LatinAmericanSpanish: curlyLatinAmericanSpanishName,
		language.Russian:              curlyRussianName,
		language.SimplifiedChinese:    curlySimplifiedChineseName,
		language.Spanish:              curlySpanishName,
		language.TraditionalChinese:   curlyTraditionalChineseName}
)

var (
	curlyCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: curlyBirthday,
		Code:     curlyCode,
		Key:      character.Curly,
		Gender:   gender.Male,
		Name:     curlyName}
)

var (
	curlyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nyoink"}

	curlyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coink"}

	curlyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hupsaknor"}

	curlyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coink"}

	curlyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oinkoink"}

	curlyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niooink"}

	curlyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どもども"}

	curlyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꿀꿀"}

	curlyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñoink"}

	curlyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк-бряк"}

	curlySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "感谢"}

	curlySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñoink"}

	curlyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "感謝"}
)

var (
	curlyPhrase = nook.Languages{
		language.AmericanEnglish:      curlyAmericanEnglishName,
		language.CanadianFrench:       curlyCanadianFrenchName,
		language.Dutch:                curlyDutchName,
		language.French:               curlyFrenchName,
		language.German:               curlyGermanName,
		language.Italian:              curlyItalianName,
		language.Japanese:             curlyJapaneseName,
		language.Korean:               curlyKoreanName,
		language.LatinAmericanSpanish: curlyLatinAmericanSpanishName,
		language.Russian:              curlyRussianName,
		language.SimplifiedChinese:    curlySimplifiedChineseName,
		language.Spanish:              curlySpanishName,
		language.TraditionalChinese:   curlyTraditionalChineseName}
)

var (
	Curly = nook.Villager{
		Character:   curlyCharacter,
		Personality: personality.Jock,
		Phrase:      curlyPhrase}
)
