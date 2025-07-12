package duck

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
	// pompomBirthday represents pompom birthday.
	pompomBirthday = nook.Birthday{
		Day:   11,
		Month: time.February}
)

var (
	// pompomCode represents pompom code.
	pompomCode = nook.Code{
		Value: "duk05"}
)

var (
	// pompomAmericanEnglishName represents pompom american english name.
	pompomAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pompom"}

	// pompomCanadianFrenchName represents pompom canadian french name.
	pompomCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pompon"}

	// pompomDutchName represents pompom dutch name.
	pompomDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pompom"}

	// pompomFrenchName represents pompom french name.
	pompomFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pompon"}

	// pompomGermanName represents pompom german name.
	pompomGermanName = nook.Name{
		Language: language.German,
		Value:    "Erika"}

	// pompomItalianName represents pompom italian name.
	pompomItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Erica"}

	// pompomJapaneseName represents pompom japanese name.
	pompomJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりっぺ"}

	// pompomKoreanName represents pompom korean name.
	pompomKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "주디"}

	// pompomLatinAmericanSpanishName represents pompom latin american spanish name.
	pompomLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Flopi"}

	// pompomRussianName represents pompom russian name.
	pompomRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Помпом"}

	// pompomSimplifiedChineseName represents pompom simplified chinese name.
	pompomSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海苔裴"}

	// pompomSpanishName represents pompom spanish name.
	pompomSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Flopi"}

	// pompomTraditionalChineseName represents pompom traditional chinese name.
	pompomTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海苔裴"}
)

var (
	// pompomName represents pompom name.
	pompomName = nook.Languages{
		language.AmericanEnglish:      pompomAmericanEnglishName,
		language.CanadianFrench:       pompomCanadianFrenchName,
		language.Dutch:                pompomDutchName,
		language.French:               pompomFrenchName,
		language.German:               pompomGermanName,
		language.Italian:              pompomItalianName,
		language.Japanese:             pompomJapaneseName,
		language.Korean:               pompomKoreanName,
		language.LatinAmericanSpanish: pompomLatinAmericanSpanishName,
		language.Russian:              pompomRussianName,
		language.SimplifiedChinese:    pompomSimplifiedChineseName,
		language.Spanish:              pompomSpanishName,
		language.TraditionalChinese:   pompomTraditionalChineseName}
)

var (
	// pompomCharacter represents pompom character.
	pompomCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: pompomBirthday,
		Code:     pompomCode,
		Key:      character.Pompom,
		Gender:   gender.Female,
		Name:     pompomName,
		Special:  false}
)

var (
	// pompomAmericanEnglishPhrase represents pompom american english phrase.
	pompomAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rah rah"}

	// pompomCanadianFrenchPhrase represents pompom canadian french phrase.
	pompomCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rah rah"}

	// pompomDutchPhrase represents pompom dutch phrase.
	pompomDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hup hup"}

	// pompomFrenchPhrase represents pompom french phrase.
	pompomFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rah rah"}

	// pompomGermanPhrase represents pompom german phrase.
	pompomGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krah"}

	// pompomItalianPhrase represents pompom italian phrase.
	pompomItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quoquiquà"}

	// pompomJapanesePhrase represents pompom japanese phrase.
	pompomJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっピ"}

	// pompomKoreanPhrase represents pompom korean phrase.
	pompomKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다삐"}

	// pompomLatinAmericanSpanishPhrase represents pompom latin american spanish phrase.
	pompomLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ra-rá"}

	// pompomRussianPhrase represents pompom russian phrase.
	pompomRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крякушки"}

	// pompomSimplifiedChinesePhrase represents pompom simplified chinese phrase.
	pompomSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "裴裴"}

	// pompomSpanishPhrase represents pompom spanish phrase.
	pompomSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sinplumas"}

	// pompomTraditionalChinesePhrase represents pompom traditional chinese phrase.
	pompomTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "裴裴"}
)

var (
	// pompomPhrase represents pompom phrase.
	pompomPhrase = nook.Languages{
		language.AmericanEnglish:      pompomAmericanEnglishPhrase,
		language.CanadianFrench:       pompomCanadianFrenchPhrase,
		language.Dutch:                pompomDutchPhrase,
		language.French:               pompomFrenchPhrase,
		language.German:               pompomGermanPhrase,
		language.Italian:              pompomItalianPhrase,
		language.Japanese:             pompomJapanesePhrase,
		language.Korean:               pompomKoreanPhrase,
		language.LatinAmericanSpanish: pompomLatinAmericanSpanishPhrase,
		language.Russian:              pompomRussianPhrase,
		language.SimplifiedChinese:    pompomSimplifiedChinesePhrase,
		language.Spanish:              pompomSpanishPhrase,
		language.TraditionalChinese:   pompomTraditionalChinesePhrase}
)

var (
	// Pompom represents pompom.
	Pompom = nook.Villager{
		Character:   pompomCharacter,
		Personality: personality.Peppy,
		Phrase:      pompomPhrase}
)
