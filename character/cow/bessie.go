package cow

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
	// bessieBirthday represents bessie birthday.
	bessieBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// bessieCode represents bessie code.
	bessieCode = nook.Code{
		Value: ""}
)

var (
	// bessieAmericanEnglishName represents bessie american english name.
	bessieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bessie"}

	// bessieCanadianFrenchName represents bessie canadian french name.
	bessieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// bessieDutchName represents bessie dutch name.
	bessieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// bessieFrenchName represents bessie french name.
	bessieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Florette"}

	// bessieGermanName represents bessie german name.
	bessieGermanName = nook.Name{
		Language: language.German,
		Value:    "Karla"}

	// bessieItalianName represents bessie italian name.
	bessieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Betta"}

	// bessieJapaneseName represents bessie japanese name.
	bessieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アモーレ"}

	// bessieKoreanName represents bessie korean name.
	bessieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// bessieLatinAmericanSpanishName represents bessie latin american spanish name.
	bessieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// bessieRussianName represents bessie russian name.
	bessieRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// bessieSimplifiedChineseName represents bessie simplified chinese name.
	bessieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茜茜"}

	// bessieSpanishName represents bessie spanish name.
	bessieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Muuriel"}

	// bessieTraditionalChineseName represents bessie traditional chinese name.
	bessieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bessieName represents bessie name.
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
	// bessieCharacter represents bessie character.
	bessieCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: bessieBirthday,
		Code:     bessieCode,
		Key:      character.Bessie,
		Gender:   gender.Female,
		Name:     bessieName,
		Special:  false}
)

var (
	// bessieAmericanEnglishPhrase represents bessie american english phrase.
	bessieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buttercup"}

	// bessieCanadianFrenchPhrase represents bessie canadian french phrase.
	bessieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// bessieDutchPhrase represents bessie dutch phrase.
	bessieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// bessieFrenchPhrase represents bessie french phrase.
	bessieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh'fin"}

	// bessieGermanPhrase represents bessie german phrase.
	bessieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "butterli"}

	// bessieItalianPhrase represents bessie italian phrase.
	bessieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muuffola"}

	// bessieJapanesePhrase represents bessie japanese phrase.
	bessieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ずら"}

	// bessieKoreanPhrase represents bessie korean phrase.
	bessieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// bessieLatinAmericanSpanishPhrase represents bessie latin american spanish phrase.
	bessieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// bessieRussianPhrase represents bessie russian phrase.
	bessieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// bessieSimplifiedChinesePhrase represents bessie simplified chinese phrase.
	bessieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啧啧"}

	// bessieSpanishPhrase represents bessie spanish phrase.
	bessieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuernis"}

	// bessieTraditionalChinesePhrase represents bessie traditional chinese phrase.
	bessieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bessiePhrase represents bessie phrase.
	bessiePhrase = nook.Languages{
		language.AmericanEnglish:      bessieAmericanEnglishPhrase,
		language.CanadianFrench:       bessieCanadianFrenchPhrase,
		language.Dutch:                bessieDutchPhrase,
		language.French:               bessieFrenchPhrase,
		language.German:               bessieGermanPhrase,
		language.Italian:              bessieItalianPhrase,
		language.Japanese:             bessieJapanesePhrase,
		language.Korean:               bessieKoreanPhrase,
		language.LatinAmericanSpanish: bessieLatinAmericanSpanishPhrase,
		language.Russian:              bessieRussianPhrase,
		language.SimplifiedChinese:    bessieSimplifiedChinesePhrase,
		language.Spanish:              bessieSpanishPhrase,
		language.TraditionalChinese:   bessieTraditionalChinesePhrase}
)

var (
	// Bessie represents bessie.
	Bessie = nook.Villager{
		Character:   bessieCharacter,
		Personality: personality.Normal,
		Phrase:      bessiePhrase}
)
