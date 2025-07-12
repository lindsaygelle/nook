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
	// deliBirthday represents deli birthday.
	deliBirthday = nook.Birthday{
		Day:   24,
		Month: time.May}
)

var (
	// deliCode represents deli code.
	deliCode = nook.Code{
		Value: "mnk08"}
)

var (
	// deliAmericanEnglishName represents deli american english name.
	deliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deli"}

	// deliCanadianFrenchName represents deli canadian french name.
	deliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Magogo"}

	// deliDutchName represents deli dutch name.
	deliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Deli"}

	// deliFrenchName represents deli french name.
	deliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Magogo"}

	// deliGermanName represents deli german name.
	deliGermanName = nook.Name{
		Language: language.German,
		Value:    "Anton"}

	// deliItalianName represents deli italian name.
	deliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bobo"}

	// deliJapaneseName represents deli japanese name.
	deliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デリー"}

	// deliKoreanName represents deli korean name.
	deliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "델리"}

	// deliLatinAmericanSpanishName represents deli latin american spanish name.
	deliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cana"}

	// deliRussianName represents deli russian name.
	deliRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дели"}

	// deliSimplifiedChineseName represents deli simplified chinese name.
	deliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "德里"}

	// deliSpanishName represents deli spanish name.
	deliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cana"}

	// deliTraditionalChineseName represents deli traditional chinese name.
	deliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "德里"}
)

var (
	// deliName represents deli name.
	deliName = nook.Languages{
		language.AmericanEnglish:      deliAmericanEnglishName,
		language.CanadianFrench:       deliCanadianFrenchName,
		language.Dutch:                deliDutchName,
		language.French:               deliFrenchName,
		language.German:               deliGermanName,
		language.Italian:              deliItalianName,
		language.Japanese:             deliJapaneseName,
		language.Korean:               deliKoreanName,
		language.LatinAmericanSpanish: deliLatinAmericanSpanishName,
		language.Russian:              deliRussianName,
		language.SimplifiedChinese:    deliSimplifiedChineseName,
		language.Spanish:              deliSpanishName,
		language.TraditionalChinese:   deliTraditionalChineseName}
)

var (
	// deliCharacter represents deli character.
	deliCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: deliBirthday,
		Code:     deliCode,
		Key:      character.Deli,
		Gender:   gender.Male,
		Name:     deliName,
		Special:  false}
)

var (
	// deliAmericanEnglishPhrase represents deli american english phrase.
	deliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "monch"}

	// deliCanadianFrenchPhrase represents deli canadian french phrase.
	deliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "liaaaane"}

	// deliDutchPhrase represents deli dutch phrase.
	deliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smak"}

	// deliFrenchPhrase represents deli french phrase.
	deliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "liaaaane"}

	// deliGermanPhrase represents deli german phrase.
	deliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miekmiek"}

	// deliItalianPhrase represents deli italian phrase.
	deliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "panzé"}

	// deliJapanesePhrase represents deli japanese phrase.
	deliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぜ"}

	// deliKoreanPhrase represents deli korean phrase.
	deliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흠냐리"}

	// deliLatinAmericanSpanishPhrase represents deli latin american spanish phrase.
	deliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñim-ñim"}

	// deliRussianPhrase represents deli russian phrase.
	deliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чавк-чавк"}

	// deliSimplifiedChinesePhrase represents deli simplified chinese phrase.
	deliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "去吧"}

	// deliSpanishPhrase represents deli spanish phrase.
	deliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "millón"}

	// deliTraditionalChinesePhrase represents deli traditional chinese phrase.
	deliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "去吧"}
)

var (
	// deliPhrase represents deli phrase.
	deliPhrase = nook.Languages{
		language.AmericanEnglish:      deliAmericanEnglishPhrase,
		language.CanadianFrench:       deliCanadianFrenchPhrase,
		language.Dutch:                deliDutchPhrase,
		language.French:               deliFrenchPhrase,
		language.German:               deliGermanPhrase,
		language.Italian:              deliItalianPhrase,
		language.Japanese:             deliJapanesePhrase,
		language.Korean:               deliKoreanPhrase,
		language.LatinAmericanSpanish: deliLatinAmericanSpanishPhrase,
		language.Russian:              deliRussianPhrase,
		language.SimplifiedChinese:    deliSimplifiedChinesePhrase,
		language.Spanish:              deliSpanishPhrase,
		language.TraditionalChinese:   deliTraditionalChinesePhrase}
)

var (
	// Deli represents deli.
	Deli = nook.Villager{
		Character:   deliCharacter,
		Personality: personality.Lazy,
		Phrase:      deliPhrase}
)
