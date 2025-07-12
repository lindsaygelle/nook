package elephant

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
	// elinaBirthday represents elina birthday.
	elinaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// elinaCode represents elina code.
	elinaCode = nook.Code{
		Value: ""}
)

var (
	// elinaAmericanEnglishName represents elina american english name.
	elinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elina"}

	// elinaCanadianFrenchName represents elina canadian french name.
	elinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// elinaDutchName represents elina dutch name.
	elinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// elinaFrenchName represents elina french name.
	elinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fanfan"}

	// elinaGermanName represents elina german name.
	elinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Elena"}

	// elinaItalianName represents elina italian name.
	elinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ellie"}

	// elinaJapaneseName represents elina japanese name.
	elinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビンディ"}

	// elinaKoreanName represents elina korean name.
	elinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// elinaLatinAmericanSpanishName represents elina latin american spanish name.
	elinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// elinaRussianName represents elina russian name.
	elinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// elinaSimplifiedChineseName represents elina simplified chinese name.
	elinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// elinaSpanishName represents elina spanish name.
	elinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elina"}

	// elinaTraditionalChineseName represents elina traditional chinese name.
	elinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// elinaName represents elina name.
	elinaName = nook.Languages{
		language.AmericanEnglish:      elinaAmericanEnglishName,
		language.CanadianFrench:       elinaCanadianFrenchName,
		language.Dutch:                elinaDutchName,
		language.French:               elinaFrenchName,
		language.German:               elinaGermanName,
		language.Italian:              elinaItalianName,
		language.Japanese:             elinaJapaneseName,
		language.Korean:               elinaKoreanName,
		language.LatinAmericanSpanish: elinaLatinAmericanSpanishName,
		language.Russian:              elinaRussianName,
		language.SimplifiedChinese:    elinaSimplifiedChineseName,
		language.Spanish:              elinaSpanishName,
		language.TraditionalChinese:   elinaTraditionalChineseName}
)

var (
	// elinaCharacter represents elina character.
	elinaCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: elinaBirthday,
		Code:     elinaCode,
		Key:      character.Elina,
		Gender:   gender.Female,
		Name:     elinaName,
		Special:  false}
)

var (
	// elinaAmericanEnglishPhrase represents elina american english phrase.
	elinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shrimp"}

	// elinaCanadianFrenchPhrase represents elina canadian french phrase.
	elinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// elinaDutchPhrase represents elina dutch phrase.
	elinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// elinaFrenchPhrase represents elina french phrase.
	elinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "garri"}

	// elinaGermanPhrase represents elina german phrase.
	elinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwerg"}

	// elinaItalianPhrase represents elina italian phrase.
	elinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pesciolino"}

	// elinaJapanesePhrase represents elina japanese phrase.
	elinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナマステ"}

	// elinaKoreanPhrase represents elina korean phrase.
	elinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// elinaLatinAmericanSpanishPhrase represents elina latin american spanish phrase.
	elinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// elinaRussianPhrase represents elina russian phrase.
	elinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// elinaSimplifiedChinesePhrase represents elina simplified chinese phrase.
	elinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// elinaSpanishPhrase represents elina spanish phrase.
	elinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "camarón"}

	// elinaTraditionalChinesePhrase represents elina traditional chinese phrase.
	elinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// elinaPhrase represents elina phrase.
	elinaPhrase = nook.Languages{
		language.AmericanEnglish:      elinaAmericanEnglishPhrase,
		language.CanadianFrench:       elinaCanadianFrenchPhrase,
		language.Dutch:                elinaDutchPhrase,
		language.French:               elinaFrenchPhrase,
		language.German:               elinaGermanPhrase,
		language.Italian:              elinaItalianPhrase,
		language.Japanese:             elinaJapanesePhrase,
		language.Korean:               elinaKoreanPhrase,
		language.LatinAmericanSpanish: elinaLatinAmericanSpanishPhrase,
		language.Russian:              elinaRussianPhrase,
		language.SimplifiedChinese:    elinaSimplifiedChinesePhrase,
		language.Spanish:              elinaSpanishPhrase,
		language.TraditionalChinese:   elinaTraditionalChinesePhrase}
)

var (
	// Elina represents elina.
	Elina = nook.Villager{
		Character:   elinaCharacter,
		Personality: personality.Peppy,
		Phrase:      elinaPhrase}
)
