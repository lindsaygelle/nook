package gorilla

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
	// rillaBirthday represents rilla birthday.
	rillaBirthday = nook.Birthday{
		Day:   1,
		Month: time.November}
)

var (
	// rillaCode represents rilla code.
	rillaCode = nook.Code{
		Value: "gor11"}
)

var (
	// rillaAmericanEnglishName represents rilla american english name.
	rillaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rilla"}

	// rillaCanadianFrenchName represents rilla canadian french name.
	rillaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rilla"}

	// rillaDutchName represents rilla dutch name.
	rillaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rilla"}

	// rillaFrenchName represents rilla french name.
	rillaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rilla"}

	// rillaGermanName represents rilla german name.
	rillaGermanName = nook.Name{
		Language: language.German,
		Value:    "Rilla"}

	// rillaItalianName represents rilla italian name.
	rillaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rilla"}

	// rillaJapaneseName represents rilla japanese name.
	rillaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リラ"}

	// rillaKoreanName represents rilla korean name.
	rillaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴라"}

	// rillaLatinAmericanSpanishName represents rilla latin american spanish name.
	rillaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rila"}

	// rillaRussianName represents rilla russian name.
	rillaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рилла"}

	// rillaSimplifiedChineseName represents rilla simplified chinese name.
	rillaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莉拉"}

	// rillaSpanishName represents rilla spanish name.
	rillaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rila"}

	// rillaTraditionalChineseName represents rilla traditional chinese name.
	rillaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莉拉"}
)

var (
	// rillaName represents rilla name.
	rillaName = nook.Languages{
		language.AmericanEnglish:      rillaAmericanEnglishName,
		language.CanadianFrench:       rillaCanadianFrenchName,
		language.Dutch:                rillaDutchName,
		language.French:               rillaFrenchName,
		language.German:               rillaGermanName,
		language.Italian:              rillaItalianName,
		language.Japanese:             rillaJapaneseName,
		language.Korean:               rillaKoreanName,
		language.LatinAmericanSpanish: rillaLatinAmericanSpanishName,
		language.Russian:              rillaRussianName,
		language.SimplifiedChinese:    rillaSimplifiedChineseName,
		language.Spanish:              rillaSpanishName,
		language.TraditionalChinese:   rillaTraditionalChineseName}
)

var (
	// rillaCharacter represents rilla character.
	rillaCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: rillaBirthday,
		Code:     rillaCode,
		Key:      character.Rilla,
		Gender:   gender.Female,
		Name:     rillaName,
		Special:  false}
)

var (
	// rillaAmericanEnglishPhrase represents rilla american english phrase.
	rillaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hello"}

	// rillaCanadianFrenchPhrase represents rilla canadian french phrase.
	rillaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rorille"}

	// rillaDutchPhrase represents rilla dutch phrase.
	rillaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hello"}

	// rillaFrenchPhrase represents rilla french phrase.
	rillaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rorille"}

	// rillaGermanPhrase represents rilla german phrase.
	rillaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kittykong"}

	// rillaItalianPhrase represents rilla italian phrase.
	rillaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arrruuu"}

	// rillaJapanesePhrase represents rilla japanese phrase.
	rillaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハロー"}

	// rillaKoreanPhrase represents rilla korean phrase.
	rillaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "헬로"}

	// rillaLatinAmericanSpanishPhrase represents rilla latin american spanish phrase.
	rillaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kittybún"}

	// rillaRussianPhrase represents rilla russian phrase.
	rillaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "алло"}

	// rillaSimplifiedChinesePhrase represents rilla simplified chinese phrase.
	rillaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈罗"}

	// rillaSpanishPhrase represents rilla spanish phrase.
	rillaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kittybún"}

	// rillaTraditionalChinesePhrase represents rilla traditional chinese phrase.
	rillaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈囉"}
)

var (
	// rillaPhrase represents rilla phrase.
	rillaPhrase = nook.Languages{
		language.AmericanEnglish:      rillaAmericanEnglishPhrase,
		language.CanadianFrench:       rillaCanadianFrenchPhrase,
		language.Dutch:                rillaDutchPhrase,
		language.French:               rillaFrenchPhrase,
		language.German:               rillaGermanPhrase,
		language.Italian:              rillaItalianPhrase,
		language.Japanese:             rillaJapanesePhrase,
		language.Korean:               rillaKoreanPhrase,
		language.LatinAmericanSpanish: rillaLatinAmericanSpanishPhrase,
		language.Russian:              rillaRussianPhrase,
		language.SimplifiedChinese:    rillaSimplifiedChinesePhrase,
		language.Spanish:              rillaSpanishPhrase,
		language.TraditionalChinese:   rillaTraditionalChinesePhrase}
)

var (
	// Rilla represents rilla.
	Rilla = nook.Villager{
		Character:   rillaCharacter,
		Personality: personality.Peppy,
		Phrase:      rillaPhrase}
)
