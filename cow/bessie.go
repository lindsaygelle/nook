package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bessieBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	bessieCode = nook.Code{
		Value: ""}
)

var (
	bessieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bessie"}

	bessieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	bessieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	bessieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Florettemeuh'fin"}

	bessieGermanName = nook.Name{
		Language: language.German,
		Value:    "Karlabutterli"}

	bessieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bettamuuffola"}

	bessieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アモーレずら"}

	bessieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	bessieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	bessieRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	bessieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茜茜啧啧"}

	bessieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Muurielcuernis"}

	bessieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	bessieName = nook.Languages{
		language.AmericanEnglish:      bessieAmericanEnglishName,
		language.CanadianFrench:       bessieCanadianFrenchName,
		language.Dutch:                bessieDutchName,
		language.French:               bessieFrenchName,
		language.German:               bessieGermanName,
		language.Italian:              bessieItalianName,
		language.Japanese:             bessieJapaneseName,
		language.Korean:               bessieKoreanName,
		language.LatinAmericanSpanish: bessieLatinAmericanSpanishName,
		language.Russian:              bessieRussianName,
		language.SimplifiedChinese:    bessieSimplifiedChineseName,
		language.Spanish:              bessieSpanishName,
		language.TraditionalChinese:   bessieTraditionalChineseName}
)

var (
	bessieCharacter = nook.Character{
		Animal:   Cow,
		Birthday: bessieBirthday,
		Code:     bessieCode,
		Gender:   nook.Female,
		Name:     bessieName}
)

var (
	bessieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	bessieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	bessieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	bessieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh'fin"}

	bessieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "butterli"}

	bessieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muuffola"}

	bessieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ずら"}

	bessieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	bessieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	bessieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	bessieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啧啧"}

	bessieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuernis"}

	bessieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	bessiePhrase = nook.Languages{
		language.AmericanEnglish:      bessieAmericanEnglishName,
		language.CanadianFrench:       bessieCanadianFrenchName,
		language.Dutch:                bessieDutchName,
		language.French:               bessieFrenchName,
		language.German:               bessieGermanName,
		language.Italian:              bessieItalianName,
		language.Japanese:             bessieJapaneseName,
		language.Korean:               bessieKoreanName,
		language.LatinAmericanSpanish: bessieLatinAmericanSpanishName,
		language.Russian:              bessieRussianName,
		language.SimplifiedChinese:    bessieSimplifiedChineseName,
		language.Spanish:              bessieSpanishName,
		language.TraditionalChinese:   bessieTraditionalChineseName}
)

var (
	Bessie = nook.Villager{
		Character:   bessieCharacter,
		Personality: nook.Normal,
		Phrase:      bessiePhrase}
)
