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
	// boydBirthday represents boyd birthday.
	boydBirthday = nook.Birthday{
		Day:   1,
		Month: time.October}
)

var (
	// boydCode represents boyd code.
	boydCode = nook.Code{
		Value: "gor05"}
)

var (
	// boydAmericanEnglishName represents boyd american english name.
	boydAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boyd"}

	// boydCanadianFrenchName represents boyd canadian french name.
	boydCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Primo"}

	// boydDutchName represents boyd dutch name.
	boydDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boyd"}

	// boydFrenchName represents boyd french name.
	boydFrenchName = nook.Name{
		Language: language.French,
		Value:    "Primo"}

	// boydGermanName represents boyd german name.
	boydGermanName = nook.Name{
		Language: language.German,
		Value:    "Boyd"}

	// boydItalianName represents boyd italian name.
	boydItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brando"}

	// boydJapaneseName represents boyd japanese name.
	boydJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボイド"}

	// boydKoreanName represents boyd korean name.
	boydKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "보이드"}

	// boydLatinAmericanSpanishName represents boyd latin american spanish name.
	boydLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bunga"}

	// boydRussianName represents boyd russian name.
	boydRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бойд"}

	// boydSimplifiedChineseName represents boyd simplified chinese name.
	boydSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "空空"}

	// boydSpanishName represents boyd spanish name.
	boydSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bunga"}

	// boydTraditionalChineseName represents boyd traditional chinese name.
	boydTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "空空"}
)

var (
	// boydName represents boyd name.
	boydName = nook.Languages{
		language.AmericanEnglish:      boydAmericanEnglishName,
		language.CanadianFrench:       boydCanadianFrenchName,
		language.Dutch:                boydDutchName,
		language.French:               boydFrenchName,
		language.German:               boydGermanName,
		language.Italian:              boydItalianName,
		language.Japanese:             boydJapaneseName,
		language.Korean:               boydKoreanName,
		language.LatinAmericanSpanish: boydLatinAmericanSpanishName,
		language.Russian:              boydRussianName,
		language.SimplifiedChinese:    boydSimplifiedChineseName,
		language.Spanish:              boydSpanishName,
		language.TraditionalChinese:   boydTraditionalChineseName}
)

var (
	// boydCharacter represents boyd character.
	boydCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: boydBirthday,
		Code:     boydCode,
		Key:      character.Boyd,
		Gender:   gender.Male,
		Name:     boydName,
		Special:  false}
)

var (
	// boydAmericanEnglishPhrase represents boyd american english phrase.
	boydAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uh-oh"}

	// boydCanadianFrenchPhrase represents boyd canadian french phrase.
	boydCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pelage"}

	// boydDutchPhrase represents boyd dutch phrase.
	boydDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oh-oh"}

	// boydFrenchPhrase represents boyd french phrase.
	boydFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pelage"}

	// boydGermanPhrase represents boyd german phrase.
	boydGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "äffäff"}

	// boydItalianPhrase represents boyd italian phrase.
	boydItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bam bam"}

	// boydJapanesePhrase represents boyd japanese phrase.
	boydJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おうおう"}

	// boydKoreanPhrase represents boyd korean phrase.
	boydKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오우오우"}

	// boydLatinAmericanSpanishPhrase represents boyd latin american spanish phrase.
	boydLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "unga"}

	// boydRussianPhrase represents boyd russian phrase.
	boydRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-ох"}

	// boydSimplifiedChinesePhrase represents boyd simplified chinese phrase.
	boydSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔喔"}

	// boydSpanishPhrase represents boyd spanish phrase.
	boydSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "unga"}

	// boydTraditionalChinesePhrase represents boyd traditional chinese phrase.
	boydTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔喔"}
)

var (
	// boydPhrase represents boyd phrase.
	boydPhrase = nook.Languages{
		language.AmericanEnglish:      boydAmericanEnglishPhrase,
		language.CanadianFrench:       boydCanadianFrenchPhrase,
		language.Dutch:                boydDutchPhrase,
		language.French:               boydFrenchPhrase,
		language.German:               boydGermanPhrase,
		language.Italian:              boydItalianPhrase,
		language.Japanese:             boydJapanesePhrase,
		language.Korean:               boydKoreanPhrase,
		language.LatinAmericanSpanish: boydLatinAmericanSpanishPhrase,
		language.Russian:              boydRussianPhrase,
		language.SimplifiedChinese:    boydSimplifiedChinesePhrase,
		language.Spanish:              boydSpanishPhrase,
		language.TraditionalChinese:   boydTraditionalChinesePhrase}
)

var (
	// Boyd represents boyd.
	Boyd = nook.Villager{
		Character:   boydCharacter,
		Personality: personality.Cranky,
		Phrase:      boydPhrase}
)
