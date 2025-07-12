package frog

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
	// croqueBirthday represents croque birthday.
	croqueBirthday = nook.Birthday{
		Day:   18,
		Month: time.July}
)

var (
	// croqueCode represents croque code.
	croqueCode = nook.Code{
		Value: "flg17"}
)

var (
	// croqueAmericanEnglishName represents croque american english name.
	croqueAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Croque"}

	// croqueCanadianFrenchName represents croque canadian french name.
	croqueCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Carlos"}

	// croqueDutchName represents croque dutch name.
	croqueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Croque"}

	// croqueFrenchName represents croque french name.
	croqueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Carlos"}

	// croqueGermanName represents croque german name.
	croqueGermanName = nook.Name{
		Language: language.German,
		Value:    "Carlo"}

	// croqueItalianName represents croque italian name.
	croqueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gracido"}

	// croqueJapaneseName represents croque japanese name.
	croqueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タイシ"}

	// croqueKoreanName represents croque korean name.
	croqueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "투투"}

	// croqueLatinAmericanSpanishName represents croque latin american spanish name.
	croqueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ranolfo"}

	// croqueRussianName represents croque russian name.
	croqueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крок"}

	// croqueSimplifiedChineseName represents croque simplified chinese name.
	croqueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太子"}

	// croqueSpanishName represents croque spanish name.
	croqueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ranolfo"}

	// croqueTraditionalChineseName represents croque traditional chinese name.
	croqueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太子"}
)

var (
	// croqueName represents croque name.
	croqueName = nook.Languages{
		language.AmericanEnglish:      croqueAmericanEnglishName,
		language.CanadianFrench:       croqueCanadianFrenchName,
		language.Dutch:                croqueDutchName,
		language.French:               croqueFrenchName,
		language.German:               croqueGermanName,
		language.Italian:              croqueItalianName,
		language.Japanese:             croqueJapaneseName,
		language.Korean:               croqueKoreanName,
		language.LatinAmericanSpanish: croqueLatinAmericanSpanishName,
		language.Russian:              croqueRussianName,
		language.SimplifiedChinese:    croqueSimplifiedChineseName,
		language.Spanish:              croqueSpanishName,
		language.TraditionalChinese:   croqueTraditionalChineseName}
)

var (
	// croqueCharacter represents croque character.
	croqueCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: croqueBirthday,
		Code:     croqueCode,
		Key:      character.Croque,
		Gender:   gender.Male,
		Name:     croqueName,
		Special:  false}
)

var (
	// croqueAmericanEnglishPhrase represents croque american english phrase.
	croqueAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "as if"}

	// croqueCanadianFrenchPhrase represents croque canadian french phrase.
	croqueCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouaoua"}

	// croqueDutchPhrase represents croque dutch phrase.
	croqueDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ja dááág"}

	// croqueFrenchPhrase represents croque french phrase.
	croqueFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crapoupou"}

	// croqueGermanPhrase represents croque german phrase.
	croqueGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schleck"}

	// croqueItalianPhrase represents croque italian phrase.
	croqueItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "umpf"}

	// croqueJapanesePhrase represents croque japanese phrase.
	croqueJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "しからば"}

	// croqueKoreanPhrase represents croque korean phrase.
	croqueKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "툴툴"}

	// croqueLatinAmericanSpanishPhrase represents croque latin american spanish phrase.
	croqueLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yastá"}

	// croqueRussianPhrase represents croque russian phrase.
	croqueRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вроде того"}

	// croqueSimplifiedChinesePhrase represents croque simplified chinese phrase.
	croqueSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "子曰"}

	// croqueSpanishPhrase represents croque spanish phrase.
	croqueSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yastá"}

	// croqueTraditionalChinesePhrase represents croque traditional chinese phrase.
	croqueTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "子曰"}
)

var (
	// croquePhrase represents croque phrase.
	croquePhrase = nook.Languages{
		language.AmericanEnglish:      croqueAmericanEnglishPhrase,
		language.CanadianFrench:       croqueCanadianFrenchPhrase,
		language.Dutch:                croqueDutchPhrase,
		language.French:               croqueFrenchPhrase,
		language.German:               croqueGermanPhrase,
		language.Italian:              croqueItalianPhrase,
		language.Japanese:             croqueJapanesePhrase,
		language.Korean:               croqueKoreanPhrase,
		language.LatinAmericanSpanish: croqueLatinAmericanSpanishPhrase,
		language.Russian:              croqueRussianPhrase,
		language.SimplifiedChinese:    croqueSimplifiedChinesePhrase,
		language.Spanish:              croqueSpanishPhrase,
		language.TraditionalChinese:   croqueTraditionalChinesePhrase}
)

var (
	// Croque represents croque.
	Croque = nook.Villager{
		Character:   croqueCharacter,
		Personality: personality.Cranky,
		Phrase:      croquePhrase}
)
