package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cobbBirthday = nook.Birthday{
		Day:   7,
		Month: time.October}
)

var (
	cobbCode = nook.Code{
		Value: "pig08"}
)

var (
	cobbAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cobb"}

	cobbCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Porken"}

	cobbDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cobb"}

	cobbFrenchName = nook.Name{
		Language: language.French,
		Value:    "Porken"}

	cobbGermanName = nook.Name{
		Language: language.German,
		Value:    "Rolo"}

	cobbItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Wurst"}

	cobbJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハカセ"}

	cobbKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "박사"}

	cobbLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Franfuá"}

	cobbRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кобб"}

	cobbSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "博士"}

	cobbSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Franfuá"}

	cobbTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "博士"}
)

var (
	cobbName = nook.Languages{
		language.AmericanEnglish:      cobbAmericanEnglishName,
		language.CanadianFrench:       cobbCanadianFrenchName,
		language.Dutch:                cobbDutchName,
		language.French:               cobbFrenchName,
		language.German:               cobbGermanName,
		language.Italian:              cobbItalianName,
		language.Japanese:             cobbJapaneseName,
		language.Korean:               cobbKoreanName,
		language.LatinAmericanSpanish: cobbLatinAmericanSpanishName,
		language.Russian:              cobbRussianName,
		language.SimplifiedChinese:    cobbSimplifiedChineseName,
		language.Spanish:              cobbSpanishName,
		language.TraditionalChinese:   cobbTraditionalChineseName}
)

var (
	cobbCharacter = nook.Character{
		Animal:   Pig,
		Birthday: cobbBirthday,
		Code:     cobbCode,
		Gender:   nook.Male,
		Name:     cobbName}
)

var (
	cobbAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hot dog"}

	cobbCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "andouille"}

	cobbDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knakker"}

	cobbFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "andouille"}

	cobbGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "würstel"}

	cobbItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "salsicce"}

	cobbJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアール"}

	cobbKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이노라"}

	cobbLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "saperlipop"}

	cobbRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сосиска"}

	cobbSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "公亩"}

	cobbSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "saperlipop"}

	cobbTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "公畝"}
)

var (
	cobbPhrase = nook.Languages{
		language.AmericanEnglish:      cobbAmericanEnglishName,
		language.CanadianFrench:       cobbCanadianFrenchName,
		language.Dutch:                cobbDutchName,
		language.French:               cobbFrenchName,
		language.German:               cobbGermanName,
		language.Italian:              cobbItalianName,
		language.Japanese:             cobbJapaneseName,
		language.Korean:               cobbKoreanName,
		language.LatinAmericanSpanish: cobbLatinAmericanSpanishName,
		language.Russian:              cobbRussianName,
		language.SimplifiedChinese:    cobbSimplifiedChineseName,
		language.Spanish:              cobbSpanishName,
		language.TraditionalChinese:   cobbTraditionalChineseName}
)

var (
	Cobb = nook.Villager{
		Character:   cobbCharacter,
		Personality: nook.Jock,
		Phrase:      cobbPhrase}
)
