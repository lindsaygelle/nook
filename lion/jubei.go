package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	jubeiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	jubeiFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	jubeiGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	jubeiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	jubeiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュウベエ"}

	jubeiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	jubeiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	jubeiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	jubeiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	jubeiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	jubeiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Lion,
		Birthday: jubeiBirthday,
		Code:     jubeiCode,
		Gender:   nook.Male,
		Name:     jubeiName}
)

var (
	jubeiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	jubeiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	jubeiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	jubeiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	jubeiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	jubeiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	jubeiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いかにも"}

	jubeiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	jubeiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	jubeiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	jubeiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	jubeiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	jubeiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	jubeiPhrase = nook.Languages{
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
	Jubei = nook.Villager{
		Character:   jubeiCharacter,
		Personality: nook.Cranky,
		Phrase:      jubeiPhrase}
)
