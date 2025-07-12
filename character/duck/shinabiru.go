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
	// shinabiruBirthday represents shinabiru birthday.
	shinabiruBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// shinabiruCode represents shinabiru code.
	shinabiruCode = nook.Code{
		Value: ""}
)

var (
	// shinabiruAmericanEnglishName represents shinabiru american english name.
	shinabiruAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shinabiru"}

	// shinabiruCanadianFrenchName represents shinabiru canadian french name.
	shinabiruCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// shinabiruDutchName represents shinabiru dutch name.
	shinabiruDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// shinabiruFrenchName represents shinabiru french name.
	shinabiruFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// shinabiruGermanName represents shinabiru german name.
	shinabiruGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// shinabiruItalianName represents shinabiru italian name.
	shinabiruItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// shinabiruJapaneseName represents shinabiru japanese name.
	shinabiruJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シナビル"}

	// shinabiruKoreanName represents shinabiru korean name.
	shinabiruKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// shinabiruLatinAmericanSpanishName represents shinabiru latin american spanish name.
	shinabiruLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// shinabiruRussianName represents shinabiru russian name.
	shinabiruRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// shinabiruSimplifiedChineseName represents shinabiru simplified chinese name.
	shinabiruSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// shinabiruSpanishName represents shinabiru spanish name.
	shinabiruSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// shinabiruTraditionalChineseName represents shinabiru traditional chinese name.
	shinabiruTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// shinabiruName represents shinabiru name.
	shinabiruName = nook.Languages{
		language.AmericanEnglish:      shinabiruAmericanEnglishName,
		language.CanadianFrench:       shinabiruCanadianFrenchName,
		language.Dutch:                shinabiruDutchName,
		language.French:               shinabiruFrenchName,
		language.German:               shinabiruGermanName,
		language.Italian:              shinabiruItalianName,
		language.Japanese:             shinabiruJapaneseName,
		language.Korean:               shinabiruKoreanName,
		language.LatinAmericanSpanish: shinabiruLatinAmericanSpanishName,
		language.Russian:              shinabiruRussianName,
		language.SimplifiedChinese:    shinabiruSimplifiedChineseName,
		language.Spanish:              shinabiruSpanishName,
		language.TraditionalChinese:   shinabiruTraditionalChineseName}
)

var (
	// shinabiruCharacter represents shinabiru character.
	shinabiruCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: shinabiruBirthday,
		Code:     shinabiruCode,
		Key:      character.Shinabiru,
		Gender:   gender.Male,
		Name:     shinabiruName,
		Special:  false}
)

var (
	// shinabiruAmericanEnglishPhrase represents shinabiru american english phrase.
	shinabiruAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "シナ"}

	// shinabiruCanadianFrenchPhrase represents shinabiru canadian french phrase.
	shinabiruCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// shinabiruDutchPhrase represents shinabiru dutch phrase.
	shinabiruDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// shinabiruFrenchPhrase represents shinabiru french phrase.
	shinabiruFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// shinabiruGermanPhrase represents shinabiru german phrase.
	shinabiruGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// shinabiruItalianPhrase represents shinabiru italian phrase.
	shinabiruItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// shinabiruJapanesePhrase represents shinabiru japanese phrase.
	shinabiruJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "シナ"}

	// shinabiruKoreanPhrase represents shinabiru korean phrase.
	shinabiruKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// shinabiruLatinAmericanSpanishPhrase represents shinabiru latin american spanish phrase.
	shinabiruLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// shinabiruRussianPhrase represents shinabiru russian phrase.
	shinabiruRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// shinabiruSimplifiedChinesePhrase represents shinabiru simplified chinese phrase.
	shinabiruSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// shinabiruSpanishPhrase represents shinabiru spanish phrase.
	shinabiruSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// shinabiruTraditionalChinesePhrase represents shinabiru traditional chinese phrase.
	shinabiruTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// shinabiruPhrase represents shinabiru phrase.
	shinabiruPhrase = nook.Languages{
		language.AmericanEnglish:      shinabiruAmericanEnglishPhrase,
		language.CanadianFrench:       shinabiruCanadianFrenchPhrase,
		language.Dutch:                shinabiruDutchPhrase,
		language.French:               shinabiruFrenchPhrase,
		language.German:               shinabiruGermanPhrase,
		language.Italian:              shinabiruItalianPhrase,
		language.Japanese:             shinabiruJapanesePhrase,
		language.Korean:               shinabiruKoreanPhrase,
		language.LatinAmericanSpanish: shinabiruLatinAmericanSpanishPhrase,
		language.Russian:              shinabiruRussianPhrase,
		language.SimplifiedChinese:    shinabiruSimplifiedChinesePhrase,
		language.Spanish:              shinabiruSpanishPhrase,
		language.TraditionalChinese:   shinabiruTraditionalChinesePhrase}
)

var (
	// Shinabiru represents shinabiru.
	Shinabiru = nook.Villager{
		Character:   shinabiruCharacter,
		Personality: personality.Jock,
		Phrase:      shinabiruPhrase}
)
