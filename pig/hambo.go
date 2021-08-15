package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hamboBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	hamboCode = nook.Code{
		Value: ""}
)

var (
	hamboAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hambo"}

	hamboCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	hamboDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	hamboFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jambo"}

	hamboGermanName = nook.Name{
		Language: language.German,
		Value:    "Silvio"}

	hamboItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampone"}

	hamboJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "はちまき"}

	hamboKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	hamboLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	hamboRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	hamboSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "头巾"}

	hamboSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marcial"}

	hamboTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	hamboName = nook.Languages{
		language.AmericanEnglish:      hamboAmericanEnglishName,
		language.CanadianFrench:       hamboCanadianFrenchName,
		language.Dutch:                hamboDutchName,
		language.French:               hamboFrenchName,
		language.German:               hamboGermanName,
		language.Italian:              hamboItalianName,
		language.Japanese:             hamboJapaneseName,
		language.Korean:               hamboKoreanName,
		language.LatinAmericanSpanish: hamboLatinAmericanSpanishName,
		language.Russian:              hamboRussianName,
		language.SimplifiedChinese:    hamboSimplifiedChineseName,
		language.Spanish:              hamboSpanishName,
		language.TraditionalChinese:   hamboTraditionalChineseName}
)

var (
	hamboCharacter = nook.Character{
		Animal:   Pig,
		Birthday: hamboBirthday,
		Code:     hamboCode,
		Gender:   nook.Male,
		Name:     hamboName}
)

var (
	hamboAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	hamboCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	hamboDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	hamboFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pti boudin"}

	hamboGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hajaaa"}

	hamboItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yo"}

	hamboJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だス"}

	hamboKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	hamboLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	hamboRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	hamboSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈罗"}

	hamboSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "morcillita"}

	hamboTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	hamboPhrase = nook.Languages{
		language.AmericanEnglish:      hamboAmericanEnglishName,
		language.CanadianFrench:       hamboCanadianFrenchName,
		language.Dutch:                hamboDutchName,
		language.French:               hamboFrenchName,
		language.German:               hamboGermanName,
		language.Italian:              hamboItalianName,
		language.Japanese:             hamboJapaneseName,
		language.Korean:               hamboKoreanName,
		language.LatinAmericanSpanish: hamboLatinAmericanSpanishName,
		language.Russian:              hamboRussianName,
		language.SimplifiedChinese:    hamboSimplifiedChineseName,
		language.Spanish:              hamboSpanishName,
		language.TraditionalChinese:   hamboTraditionalChineseName}
)

var (
	Hambo = nook.Villager{
		Character:   hamboCharacter,
		Personality: nook.Jock,
		Phrase:      hamboPhrase}
)
