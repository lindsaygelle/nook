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
	// wendyBirthday represents Wendy's birthday.
	wendyBirthday = nook.Birthday{
		Day:   15,
		Month: time.August}
)

var (
	// wendyCode represents Wendy's unique code.
	wendyCode = nook.Code{
		Value: "shp09"}
)

var (
	// wendyAmericanEnglishName represents Wendy's name in American English.
	wendyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wendy"}

	// wendyCanadianFrenchName represents Wendy's name in Canadian French.
	wendyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Karen"}

	// wendyDutchName represents Wendy's name in Dutch.
	wendyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wendy"}

	// wendyFrenchName represents Wendy's name in French.
	wendyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Karen"}

	// wendyGermanName represents Wendy's name in German.
	wendyGermanName = nook.Name{
		Language: language.German,
		Value:    "Wolli"}

	// wendyItalianName represents Wendy's name in Italian.
	wendyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agnola"}

	// wendyJapaneseName represents Wendy's name in Japanese.
	wendyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みぞれ"}

	// wendyKoreanName represents Wendy's name in Korean.
	wendyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "눈송이"}

	// wendyLatinAmericanSpanishName represents Wendy's name in Latin American Spanish.
	wendyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Franela"}

	// wendyRussianName represents Wendy's name in Russian.
	wendyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Венди"}

	// wendySimplifiedChineseName represents Wendy's name in Simplified Chinese.
	wendySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪花"}

	// wendySpanishName represents Wendy's name in Spanish.
	wendySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Franela"}

	// wendyTraditionalChineseName represents Wendy's name in Traditional Chinese.
	wendyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪花"}
)

var (
	// wendyName represents Wendy's name in different languages.
	wendyName = nook.Languages{
		language.AmericanEnglish:      wendyAmericanEnglishName,
		language.CanadianFrench:       wendyCanadianFrenchName,
		language.Dutch:                wendyDutchName,
		language.French:               wendyFrenchName,
		language.German:               wendyGermanName,
		language.Italian:              wendyItalianName,
		language.Japanese:             wendyJapaneseName,
		language.Korean:               wendyKoreanName,
		language.LatinAmericanSpanish: wendyLatinAmericanSpanishName,
		language.Russian:              wendyRussianName,
		language.SimplifiedChinese:    wendySimplifiedChineseName,
		language.Spanish:              wendySpanishName,
		language.TraditionalChinese:   wendyTraditionalChineseName}
)

var (
	// wendyCharacter represents Wendy's character information.
	wendyCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: wendyBirthday,
		Code:     wendyCode,
		Key:      character.Wendy,
		Gender:   gender.Female,
		Name:     wendyName,
		Special:  false}
)

var (
	// wendyAmericanEnglishPhrase represents Wendy's phrase in American English.
	wendyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lambkins"}

	// wendyCanadianFrenchPhrase represents Wendy's phrase in Canadian French.
	wendyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bê-bê-bê"}

	// wendyDutchPhrase represents Wendy's phrase in Dutch.
	wendyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lammetje"}

	// wendyFrenchPhrase represents Wendy's phrase in French.
	wendyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bê-bê-bê"}

	// wendyGermanPhrase represents Wendy's phrase in German.
	wendyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wuschel"}

	// wendyItalianPhrase represents Wendy's phrase in Italian.
	wendyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beee-ehi"}

	// wendyJapanesePhrase represents Wendy's phrase in Japanese.
	wendyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "イシシ"}

	// wendyKoreanPhrase represents Wendy's phrase in Korean.
	wendyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꺄르르"}

	// wendyLatinAmericanSpanishPhrase represents Wendy's phrase in Latin American Spanish.
	wendyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bee-biis"}

	// wendyRussianPhrase represents Wendy's phrase in Russian.
	wendyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "овечка"}

	// wendySimplifiedChinesePhrase represents Wendy's phrase in Simplified Chinese.
	wendySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "一丝丝"}

	// wendySpanishPhrase represents Wendy's phrase in Spanish.
	wendySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "joyita"}

	// wendyTraditionalChinesePhrase represents Wendy's phrase in Traditional Chinese.
	wendyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "一絲絲"}
)

var (
	// wendyPhrase represents Wendy's phrases in different languages.
	wendyPhrase = nook.Languages{
		language.AmericanEnglish:      wendyAmericanEnglishPhrase,
		language.CanadianFrench:       wendyCanadianFrenchPhrase,
		language.Dutch:                wendyDutchPhrase,
		language.French:               wendyFrenchPhrase,
		language.German:               wendyGermanPhrase,
		language.Italian:              wendyItalianPhrase,
		language.Japanese:             wendyJapanesePhrase,
		language.Korean:               wendyKoreanPhrase,
		language.LatinAmericanSpanish: wendyLatinAmericanSpanishPhrase,
		language.Russian:              wendyRussianPhrase,
		language.SimplifiedChinese:    wendySimplifiedChinesePhrase,
		language.Spanish:              wendySpanishPhrase,
		language.TraditionalChinese:   wendyTraditionalChinesePhrase}
)

var (
	// Wendy represents the character Wendy with her complete information.
	Wendy = nook.Villager{
		Character:   wendyCharacter,
		Personality: personality.Peppy,
		Phrase:      wendyPhrase}
)
