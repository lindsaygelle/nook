package lion

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
	// budBirthday represents bud birthday.
	budBirthday = nook.Birthday{
		Day:   8,
		Month: time.August}
)

var (
	// budCode represents bud code.
	budCode = nook.Code{
		Value: "lon00"}
)

var (
	// budAmericanEnglishName represents bud american english name.
	budAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bud"}

	// budCanadianFrenchName represents bud canadian french name.
	budCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léonard"}

	// budDutchName represents bud dutch name.
	budDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bud"}

	// budFrenchName represents bud french name.
	budFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léonard"}

	// budGermanName represents bud german name.
	budGermanName = nook.Name{
		Language: language.German,
		Value:    "Dieter"}

	// budItalianName represents bud italian name.
	budItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leo"}

	// budJapaneseName represents bud japanese name.
	budJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グラさん"}

	// budKoreanName represents bud korean name.
	budKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "선글"}

	// budLatinAmericanSpanishName represents bud latin american spanish name.
	budLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Surfleo"}

	// budRussianName represents bud russian name.
	budRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бад"}

	// budSimplifiedChineseName represents bud simplified chinese name.
	budSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫敬"}

	// budSpanishName represents bud spanish name.
	budSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Surfleo"}

	// budTraditionalChineseName represents bud traditional chinese name.
	budTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫敬"}
)

var (
	// budName represents bud name.
	budName = nook.Languages{
		language.AmericanEnglish:      budAmericanEnglishName,
		language.CanadianFrench:       budCanadianFrenchName,
		language.Dutch:                budDutchName,
		language.French:               budFrenchName,
		language.German:               budGermanName,
		language.Italian:              budItalianName,
		language.Japanese:             budJapaneseName,
		language.Korean:               budKoreanName,
		language.LatinAmericanSpanish: budLatinAmericanSpanishName,
		language.Russian:              budRussianName,
		language.SimplifiedChinese:    budSimplifiedChineseName,
		language.Spanish:              budSpanishName,
		language.TraditionalChinese:   budTraditionalChineseName}
)

var (
	// budCharacter represents bud character.
	budCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: budBirthday,
		Code:     budCode,
		Key:      character.Bud,
		Gender:   gender.Male,
		Name:     budName,
		Special:  false}
)

var (
	// budAmericanEnglishPhrase represents bud american english phrase.
	budAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dood"}

	// budCanadianFrenchPhrase represents bud canadian french phrase.
	budCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon man"}

	// budDutchPhrase represents bud dutch phrase.
	budDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "maaaan"}

	// budFrenchPhrase represents bud french phrase.
	budFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "m'pote"}

	// budGermanPhrase represents bud german phrase.
	budGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kumpel"}

	// budItalianPhrase represents bud italian phrase.
	budItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruggi"}

	// budJapanesePhrase represents bud japanese phrase.
	budJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メ～ン"}

	// budKoreanPhrase represents bud korean phrase.
	budKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "요～맨"}

	// budLatinAmericanSpanishPhrase represents bud latin american spanish phrase.
	budLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrrau"}

	// budRussianPhrase represents bud russian phrase.
	budRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "че-е-ел"}

	// budSimplifiedChinesePhrase represents bud simplified chinese phrase.
	budSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "男人"}

	// budSpanishPhrase represents bud spanish phrase.
	budSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sabes"}

	// budTraditionalChinesePhrase represents bud traditional chinese phrase.
	budTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "男人"}
)

var (
	// budPhrase represents bud phrase.
	budPhrase = nook.Languages{
		language.AmericanEnglish:      budAmericanEnglishPhrase,
		language.CanadianFrench:       budCanadianFrenchPhrase,
		language.Dutch:                budDutchPhrase,
		language.French:               budFrenchPhrase,
		language.German:               budGermanPhrase,
		language.Italian:              budItalianPhrase,
		language.Japanese:             budJapanesePhrase,
		language.Korean:               budKoreanPhrase,
		language.LatinAmericanSpanish: budLatinAmericanSpanishPhrase,
		language.Russian:              budRussianPhrase,
		language.SimplifiedChinese:    budSimplifiedChinesePhrase,
		language.Spanish:              budSpanishPhrase,
		language.TraditionalChinese:   budTraditionalChinesePhrase}
)

var (
	// Bud represents bud.
	Bud = nook.Villager{
		Character:   budCharacter,
		Personality: personality.Jock,
		Phrase:      budPhrase}
)
