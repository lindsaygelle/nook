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
		Value:    "N/A"}

	bessieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	bessieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Florette"}

	bessieGermanName = nook.Name{
		Language: language.German,
		Value:    "Karla"}

	bessieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Betta"}

	bessieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アモーレ"}

	bessieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	bessieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	bessieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	bessieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茜茜"}

	bessieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Muuriel"}

	bessieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
