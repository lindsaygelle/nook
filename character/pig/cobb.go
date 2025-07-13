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
	// cobbBirthday represents cobb birthday.
	cobbBirthday = nook.Birthday{
		Day:   7,
		Month: time.October}
)

var (
	// cobbCode represents cobb code.
	cobbCode = nook.Code{
		Value: "pig08"}
)

var (
	// cobbAmericanEnglishName represents cobb american english name.
	cobbAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cobb"}

	// cobbCanadianFrenchName represents cobb canadian french name.
	cobbCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Porken"}

	// cobbDutchName represents cobb dutch name.
	cobbDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cobb"}

	// cobbFrenchName represents cobb french name.
	cobbFrenchName = nook.Name{
		Language: language.French,
		Value:    "Porken"}

	// cobbGermanName represents cobb german name.
	cobbGermanName = nook.Name{
		Language: language.German,
		Value:    "Rolo"}

	// cobbItalianName represents cobb italian name.
	cobbItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Wurst"}

	// cobbJapaneseName represents cobb japanese name.
	cobbJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハカセ"}

	// cobbKoreanName represents cobb korean name.
	cobbKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "박사"}

	// cobbLatinAmericanSpanishName represents cobb latin american spanish name.
	cobbLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Franfuá"}

	// cobbRussianName represents cobb russian name.
	cobbRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кобб"}

	// cobbSimplifiedChineseName represents cobb simplified chinese name.
	cobbSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "博士"}

	// cobbSpanishName represents cobb spanish name.
	cobbSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Franfuá"}

	// cobbTraditionalChineseName represents cobb traditional chinese name.
	cobbTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "博士"}
)

var (
	// cobbName represents cobb name.
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
	// cobbCharacter represents cobb character.
	cobbCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: cobbBirthday,
		Code:     cobbCode,
		Key:      character.Cobb,
		Gender:   gender.Male,
		Name:     cobbName,
		Special:  false}
)

var (
	// cobbAmericanEnglishPhrase represents cobb american english phrase.
	cobbAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hot dog"}

	// cobbCanadianFrenchPhrase represents cobb canadian french phrase.
	cobbCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "andouille"}

	// cobbDutchPhrase represents cobb dutch phrase.
	cobbDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knakker"}

	// cobbFrenchPhrase represents cobb french phrase.
	cobbFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "andouille"}

	// cobbGermanPhrase represents cobb german phrase.
	cobbGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "würstel"}

	// cobbItalianPhrase represents cobb italian phrase.
	cobbItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "salsicce"}

	// cobbJapanesePhrase represents cobb japanese phrase.
	cobbJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアール"}

	// cobbKoreanPhrase represents cobb korean phrase.
	cobbKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이노라"}

	// cobbLatinAmericanSpanishPhrase represents cobb latin american spanish phrase.
	cobbLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "saperlipop"}

	// cobbRussianPhrase represents cobb russian phrase.
	cobbRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сосиска"}

	// cobbSimplifiedChinesePhrase represents cobb simplified chinese phrase.
	cobbSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "公亩"}

	// cobbSpanishPhrase represents cobb spanish phrase.
	cobbSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "saperlipop"}

	// cobbTraditionalChinesePhrase represents cobb traditional chinese phrase.
	cobbTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "公畝"}
)

var (
	// cobbPhrase represents cobb phrase.
	cobbPhrase = nook.Languages{
		language.AmericanEnglish:      cobbAmericanEnglishPhrase,
		language.CanadianFrench:       cobbCanadianFrenchPhrase,
		language.Dutch:                cobbDutchPhrase,
		language.French:               cobbFrenchPhrase,
		language.German:               cobbGermanPhrase,
		language.Italian:              cobbItalianPhrase,
		language.Japanese:             cobbJapanesePhrase,
		language.Korean:               cobbKoreanPhrase,
		language.LatinAmericanSpanish: cobbLatinAmericanSpanishPhrase,
		language.Russian:              cobbRussianPhrase,
		language.SimplifiedChinese:    cobbSimplifiedChinesePhrase,
		language.Spanish:              cobbSpanishPhrase,
		language.TraditionalChinese:   cobbTraditionalChinesePhrase}
)

var (
	// Cobb represents cobb.
	Cobb = nook.Villager{
		Character:   cobbCharacter,
		Personality: personality.Jock,
		Phrase:      cobbPhrase}
)
