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
	// champBirthday represents champ birthday.
	champBirthday = nook.Birthday{
		Day:   4,
		Month: time.June}
)

var (
	// champCode represents champ code.
	champCode = nook.Code{
		Value: "mnk00"}
)

var (
	// champAmericanEnglishName represents champ american english name.
	champAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Champ"}

	// champCanadianFrenchName represents champ canadian french name.
	champCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Thibaut"}

	// champDutchName represents champ dutch name.
	champDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// champFrenchName represents champ french name.
	champFrenchName = nook.Name{
		Language: language.French,
		Value:    "Thibaut"}

	// champGermanName represents champ german name.
	champGermanName = nook.Name{
		Language: language.German,
		Value:    "Tschita"}

	// champItalianName represents champ italian name.
	champItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Champ"}

	// champJapaneseName represents champ japanese name.
	champJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さるお"}

	// champKoreanName represents champ korean name.
	champKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽돌이"}

	// champLatinAmericanSpanishName represents champ latin american spanish name.
	champLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mico"}

	// champRussianName represents champ russian name.
	champRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// champSimplifiedChineseName represents champ simplified chinese name.
	champSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// champSpanishName represents champ spanish name.
	champSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mico"}

	// champTraditionalChineseName represents champ traditional chinese name.
	champTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// champName represents champ name.
	champName = nook.Languages{
		language.AmericanEnglish:      champAmericanEnglishName,
		language.CanadianFrench:       champCanadianFrenchName,
		language.Dutch:                champDutchName,
		language.French:               champFrenchName,
		language.German:               champGermanName,
		language.Italian:              champItalianName,
		language.Japanese:             champJapaneseName,
		language.Korean:               champKoreanName,
		language.LatinAmericanSpanish: champLatinAmericanSpanishName,
		language.Russian:              champRussianName,
		language.SimplifiedChinese:    champSimplifiedChineseName,
		language.Spanish:              champSpanishName,
		language.TraditionalChinese:   champTraditionalChineseName}
)

var (
	// champCharacter represents champ character.
	champCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: champBirthday,
		Code:     champCode,
		Key:      character.Champ,
		Gender:   gender.Male,
		Name:     champName,
		Special:  false}
)

var (
	// champAmericanEnglishPhrase represents champ american english phrase.
	champAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "choo CHOO"}

	// champCanadianFrenchPhrase represents champ canadian french phrase.
	champCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tchouTCHOU"}

	// champDutchPhrase represents champ dutch phrase.
	champDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// champFrenchPhrase represents champ french phrase.
	champFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tchouTCHOU"}

	// champGermanPhrase represents champ german phrase.
	champGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "u-u-u"}

	// champItalianPhrase represents champ italian phrase.
	champItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "choo CHOO"}

	// champJapanesePhrase represents champ japanese phrase.
	champJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウッキー"}

	// champKoreanPhrase represents champ korean phrase.
	champKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야아아"}

	// champLatinAmericanSpanishPhrase represents champ latin american spanish phrase.
	champLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-ah-ah"}

	// champRussianPhrase represents champ russian phrase.
	champRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// champSimplifiedChinesePhrase represents champ simplified chinese phrase.
	champSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// champSpanishPhrase represents champ spanish phrase.
	champSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uh-ah-ah"}

	// champTraditionalChinesePhrase represents champ traditional chinese phrase.
	champTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// champPhrase represents champ phrase.
	champPhrase = nook.Languages{
		language.AmericanEnglish:      champAmericanEnglishPhrase,
		language.CanadianFrench:       champCanadianFrenchPhrase,
		language.Dutch:                champDutchPhrase,
		language.French:               champFrenchPhrase,
		language.German:               champGermanPhrase,
		language.Italian:              champItalianPhrase,
		language.Japanese:             champJapanesePhrase,
		language.Korean:               champKoreanPhrase,
		language.LatinAmericanSpanish: champLatinAmericanSpanishPhrase,
		language.Russian:              champRussianPhrase,
		language.SimplifiedChinese:    champSimplifiedChinesePhrase,
		language.Spanish:              champSpanishPhrase,
		language.TraditionalChinese:   champTraditionalChinesePhrase}
)

var (
	// Champ represents champ.
	Champ = nook.Villager{
		Character:   champCharacter,
		Personality: personality.Jock,
		Phrase:      champPhrase}
)
