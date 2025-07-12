package monkey

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
	// tammiBirthday represents tammi birthday.
	tammiBirthday = nook.Birthday{
		Day:   23,
		Month: time.April}
)

var (
	// tammiCode represents tammi code.
	tammiCode = nook.Code{
		Value: "mnk03"}
)

var (
	// tammiAmericanEnglishName represents tammi american english name.
	tammiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tammi"}

	// tammiCanadianFrenchName represents tammi canadian french name.
	tammiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lili"}

	// tammiDutchName represents tammi dutch name.
	tammiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tammi"}

	// tammiFrenchName represents tammi french name.
	tammiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lili"}

	// tammiGermanName represents tammi german name.
	tammiGermanName = nook.Name{
		Language: language.German,
		Value:    "Bonni"}

	// tammiItalianName represents tammi italian name.
	tammiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tammi"}

	// tammiJapaneseName represents tammi japanese name.
	tammiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エイプリル"}

	// tammiKoreanName represents tammi korean name.
	tammiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에이프릴"}

	// tammiLatinAmericanSpanishName represents tammi latin american spanish name.
	tammiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tami"}

	// tammiRussianName represents tammi russian name.
	tammiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэмми"}

	// tammiSimplifiedChineseName represents tammi simplified chinese name.
	tammiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "四月"}

	// tammiSpanishName represents tammi spanish name.
	tammiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tami"}

	// tammiTraditionalChineseName represents tammi traditional chinese name.
	tammiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "四月"}
)

var (
	// tammiName represents tammi name.
	tammiName = nook.Languages{
		language.AmericanEnglish:      tammiAmericanEnglishName,
		language.CanadianFrench:       tammiCanadianFrenchName,
		language.Dutch:                tammiDutchName,
		language.French:               tammiFrenchName,
		language.German:               tammiGermanName,
		language.Italian:              tammiItalianName,
		language.Japanese:             tammiJapaneseName,
		language.Korean:               tammiKoreanName,
		language.LatinAmericanSpanish: tammiLatinAmericanSpanishName,
		language.Russian:              tammiRussianName,
		language.SimplifiedChinese:    tammiSimplifiedChineseName,
		language.Spanish:              tammiSpanishName,
		language.TraditionalChinese:   tammiTraditionalChineseName}
)

var (
	// tammiCharacter represents tammi character.
	tammiCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: tammiBirthday,
		Code:     tammiCode,
		Key:      character.Tammi,
		Gender:   gender.Female,
		Name:     tammiName,
		Special:  false}
)

var (
	// tammiAmericanEnglishPhrase represents tammi american english phrase.
	tammiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chimpy"}

	// tammiCanadianFrenchPhrase represents tammi canadian french phrase.
	tammiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "booof"}

	// tammiDutchPhrase represents tammi dutch phrase.
	tammiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "chimpie"}

	// tammiFrenchPhrase represents tammi french phrase.
	tammiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "booof"}

	// tammiGermanPhrase represents tammi german phrase.
	tammiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "agaga"}

	// tammiItalianPhrase represents tammi italian phrase.
	tammiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "scimmietta"}

	// tammiJapanesePhrase represents tammi japanese phrase.
	tammiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワオ"}

	// tammiKoreanPhrase represents tammi korean phrase.
	tammiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "와우"}

	// tammiLatinAmericanSpanishPhrase represents tammi latin american spanish phrase.
	tammiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uki-uki"}

	// tammiRussianPhrase represents tammi russian phrase.
	tammiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шимпатяга"}

	// tammiSimplifiedChinesePhrase represents tammi simplified chinese phrase.
	tammiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇呜"}

	// tammiSpanishPhrase represents tammi spanish phrase.
	tammiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uki-uki"}

	// tammiTraditionalChinesePhrase represents tammi traditional chinese phrase.
	tammiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇嗚"}
)

var (
	// tammiPhrase represents tammi phrase.
	tammiPhrase = nook.Languages{
		language.AmericanEnglish:      tammiAmericanEnglishPhrase,
		language.CanadianFrench:       tammiCanadianFrenchPhrase,
		language.Dutch:                tammiDutchPhrase,
		language.French:               tammiFrenchPhrase,
		language.German:               tammiGermanPhrase,
		language.Italian:              tammiItalianPhrase,
		language.Japanese:             tammiJapanesePhrase,
		language.Korean:               tammiKoreanPhrase,
		language.LatinAmericanSpanish: tammiLatinAmericanSpanishPhrase,
		language.Russian:              tammiRussianPhrase,
		language.SimplifiedChinese:    tammiSimplifiedChinesePhrase,
		language.Spanish:              tammiSpanishPhrase,
		language.TraditionalChinese:   tammiTraditionalChinesePhrase}
)

var (
	// Tammi represents tammi.
	Tammi = nook.Villager{
		Character:   tammiCharacter,
		Personality: personality.Peppy,
		Phrase:      tammiPhrase}
)
