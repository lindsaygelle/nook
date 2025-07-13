package bird

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
	// otisBirthday represents otis birthday.
	otisBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// otisCode represents otis code.
	otisCode = nook.Code{
		Value: ""}
)

var (
	// otisAmericanEnglishName represents otis american english name.
	otisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Otis"}

	// otisCanadianFrenchName represents otis canadian french name.
	otisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// otisDutchName represents otis dutch name.
	otisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// otisFrenchName represents otis french name.
	otisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ottis"}

	// otisGermanName represents otis german name.
	otisGermanName = nook.Name{
		Language: language.German,
		Value:    "Otto"}

	// otisItalianName represents otis italian name.
	otisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pablo"}

	// otisJapaneseName represents otis japanese name.
	otisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たくあん"}

	// otisKoreanName represents otis korean name.
	otisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// otisLatinAmericanSpanishName represents otis latin american spanish name.
	otisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// otisRussianName represents otis russian name.
	otisRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// otisSimplifiedChineseName represents otis simplified chinese name.
	otisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "长老"}

	// otisSpanishName represents otis spanish name.
	otisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Oto"}

	// otisTraditionalChineseName represents otis traditional chinese name.
	otisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// otisName represents otis name.
	otisName = nook.Languages{
		language.AmericanEnglish:      otisAmericanEnglishName,
		language.CanadianFrench:       otisCanadianFrenchName,
		language.Dutch:                otisDutchName,
		language.French:               otisFrenchName,
		language.German:               otisGermanName,
		language.Italian:              otisItalianName,
		language.Japanese:             otisJapaneseName,
		language.Korean:               otisKoreanName,
		language.LatinAmericanSpanish: otisLatinAmericanSpanishName,
		language.Russian:              otisRussianName,
		language.SimplifiedChinese:    otisSimplifiedChineseName,
		language.Spanish:              otisSpanishName,
		language.TraditionalChinese:   otisTraditionalChineseName}
)

var (
	// otisCharacter represents otis character.
	otisCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: otisBirthday,
		Code:     otisCode,
		Key:      character.Otis,
		Gender:   gender.Male,
		Name:     otisName,
		Special:  false}
)

var (
	// otisAmericanEnglishPhrase represents otis american english phrase.
	otisAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "I s'pose"}

	// otisCanadianFrenchPhrase represents otis canadian french phrase.
	otisCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// otisDutchPhrase represents otis dutch phrase.
	otisDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// otisFrenchPhrase represents otis french phrase.
	otisFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "j'pense"}

	// otisGermanPhrase represents otis german phrase.
	otisGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "denk ich"}

	// otisItalianPhrase represents otis italian phrase.
	otisItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tuitcì"}

	// otisJapanesePhrase represents otis japanese phrase.
	otisJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのじゃ"}

	// otisKoreanPhrase represents otis korean phrase.
	otisKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// otisLatinAmericanSpanishPhrase represents otis latin american spanish phrase.
	otisLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// otisRussianPhrase represents otis russian phrase.
	otisRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// otisSimplifiedChinesePhrase represents otis simplified chinese phrase.
	otisSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "神说"}

	// otisSpanishPhrase represents otis spanish phrase.
	otisSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "digo"}

	// otisTraditionalChinesePhrase represents otis traditional chinese phrase.
	otisTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// otisPhrase represents otis phrase.
	otisPhrase = nook.Languages{
		language.AmericanEnglish:      otisAmericanEnglishPhrase,
		language.CanadianFrench:       otisCanadianFrenchPhrase,
		language.Dutch:                otisDutchPhrase,
		language.French:               otisFrenchPhrase,
		language.German:               otisGermanPhrase,
		language.Italian:              otisItalianPhrase,
		language.Japanese:             otisJapanesePhrase,
		language.Korean:               otisKoreanPhrase,
		language.LatinAmericanSpanish: otisLatinAmericanSpanishPhrase,
		language.Russian:              otisRussianPhrase,
		language.SimplifiedChinese:    otisSimplifiedChinesePhrase,
		language.Spanish:              otisSpanishPhrase,
		language.TraditionalChinese:   otisTraditionalChinesePhrase}
)

var (
	// Otis represents otis.
	Otis = nook.Villager{
		Character:   otisCharacter,
		Personality: personality.Lazy,
		Phrase:      otisPhrase}
)
