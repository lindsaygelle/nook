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
	jubeiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	jubeiCode = nook.Code{
		Value: ""}
)

var (
	jubeiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jūbei"}

	jubeiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	jubeiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	jubeiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	jubeiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	jubeiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	jubeiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュウベエ"}

	jubeiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	jubeiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	jubeiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	jubeiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	jubeiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	jubeiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	jubeiName = nook.Languages{
		language.AmericanEnglish:      jubeiAmericanEnglishName,
		language.CanadianFrench:       jubeiCanadianFrenchName,
		language.Dutch:                jubeiDutchName,
		language.French:               jubeiFrenchName,
		language.German:               jubeiGermanName,
		language.Italian:              jubeiItalianName,
		language.Japanese:             jubeiJapaneseName,
		language.Korean:               jubeiKoreanName,
		language.LatinAmericanSpanish: jubeiLatinAmericanSpanishName,
		language.Russian:              jubeiRussianName,
		language.SimplifiedChinese:    jubeiSimplifiedChineseName,
		language.Spanish:              jubeiSpanishName,
		language.TraditionalChinese:   jubeiTraditionalChineseName}
)

var (
	jubeiCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: jubeiBirthday,
		Code:     jubeiCode,
		Key:      character.Jubei,
		Gender:   gender.Male,
		Name:     jubeiName,
		Special:  false}
)

var (
	jubeiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "いかにも"}

	jubeiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	jubeiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	jubeiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	jubeiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	jubeiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	jubeiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いかにも"}

	jubeiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	jubeiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	jubeiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	jubeiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	jubeiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	jubeiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	jubeiPhrase = nook.Languages{
		language.AmericanEnglish:      jubeiAmericanEnglishPhrase,
		language.CanadianFrench:       jubeiCanadianFrenchPhrase,
		language.Dutch:                jubeiDutchPhrase,
		language.French:               jubeiFrenchPhrase,
		language.German:               jubeiGermanPhrase,
		language.Italian:              jubeiItalianPhrase,
		language.Japanese:             jubeiJapanesePhrase,
		language.Korean:               jubeiKoreanPhrase,
		language.LatinAmericanSpanish: jubeiLatinAmericanSpanishPhrase,
		language.Russian:              jubeiRussianPhrase,
		language.SimplifiedChinese:    jubeiSimplifiedChinesePhrase,
		language.Spanish:              jubeiSpanishPhrase,
		language.TraditionalChinese:   jubeiTraditionalChinesePhrase}
)

var (
	Jubei = nook.Villager{
		Character:   jubeiCharacter,
		Personality: personality.Cranky,
		Phrase:      jubeiPhrase}
)
