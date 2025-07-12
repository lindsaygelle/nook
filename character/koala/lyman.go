package koala

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
	// lymanBirthday represents lyman birthday.
	lymanBirthday = nook.Birthday{
		Day:   12,
		Month: time.October}
)

var (
	// lymanCode represents lyman code.
	lymanCode = nook.Code{
		Value: "kal09"}
)

var (
	// lymanAmericanEnglishName represents lyman american english name.
	lymanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lyman"}

	// lymanCanadianFrenchName represents lyman canadian french name.
	lymanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kalyptus"}

	// lymanDutchName represents lyman dutch name.
	lymanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lyman"}

	// lymanFrenchName represents lyman french name.
	lymanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kalyptus"}

	// lymanGermanName represents lyman german name.
	lymanGermanName = nook.Name{
		Language: language.German,
		Value:    "Pepe"}

	// lymanItalianName represents lyman italian name.
	lymanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nicola"}

	// lymanJapaneseName represents lyman japanese name.
	lymanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オズモンド"}

	// lymanKoreanName represents lyman korean name.
	lymanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오즈먼드"}

	// lymanLatinAmericanSpanishName represents lyman latin american spanish name.
	lymanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chipi"}

	// lymanRussianName represents lyman russian name.
	lymanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лайман"}

	// lymanSimplifiedChineseName represents lyman simplified chinese name.
	lymanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "敖志明"}

	// lymanSpanishName represents lyman spanish name.
	lymanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chipi"}

	// lymanTraditionalChineseName represents lyman traditional chinese name.
	lymanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "敖志明"}
)

var (
	// lymanName represents lyman name.
	lymanName = nook.Languages{
		language.AmericanEnglish:      lymanAmericanEnglishName,
		language.CanadianFrench:       lymanCanadianFrenchName,
		language.Dutch:                lymanDutchName,
		language.French:               lymanFrenchName,
		language.German:               lymanGermanName,
		language.Italian:              lymanItalianName,
		language.Japanese:             lymanJapaneseName,
		language.Korean:               lymanKoreanName,
		language.LatinAmericanSpanish: lymanLatinAmericanSpanishName,
		language.Russian:              lymanRussianName,
		language.SimplifiedChinese:    lymanSimplifiedChineseName,
		language.Spanish:              lymanSpanishName,
		language.TraditionalChinese:   lymanTraditionalChineseName}
)

var (
	// lymanCharacter represents lyman character.
	lymanCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: lymanBirthday,
		Code:     lymanCode,
		Key:      character.Lyman,
		Gender:   gender.Male,
		Name:     lymanName,
		Special:  false}
)

var (
	// lymanAmericanEnglishPhrase represents lyman american english phrase.
	lymanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chips"}

	// lymanCanadianFrenchPhrase represents lyman canadian french phrase.
	lymanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sandwich"}

	// lymanDutchPhrase represents lyman dutch phrase.
	lymanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "chips"}

	// lymanFrenchPhrase represents lyman french phrase.
	lymanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sandwich"}

	// lymanGermanPhrase represents lyman german phrase.
	lymanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "liptus"}

	// lymanItalianPhrase represents lyman italian phrase.
	lymanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ponzipò"}

	// lymanJapanesePhrase represents lyman japanese phrase.
	lymanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "わ～い"}

	// lymanKoreanPhrase represents lyman korean phrase.
	lymanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우와"}

	// lymanLatinAmericanSpanishPhrase represents lyman latin american spanish phrase.
	lymanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralalá"}

	// lymanRussianPhrase represents lyman russian phrase.
	lymanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ничего"}

	// lymanSimplifiedChinesePhrase represents lyman simplified chinese phrase.
	lymanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇"}

	// lymanSpanishPhrase represents lyman spanish phrase.
	lymanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trololó"}

	// lymanTraditionalChinesePhrase represents lyman traditional chinese phrase.
	lymanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇～"}
)

var (
	// lymanPhrase represents lyman phrase.
	lymanPhrase = nook.Languages{
		language.AmericanEnglish:      lymanAmericanEnglishPhrase,
		language.CanadianFrench:       lymanCanadianFrenchPhrase,
		language.Dutch:                lymanDutchPhrase,
		language.French:               lymanFrenchPhrase,
		language.German:               lymanGermanPhrase,
		language.Italian:              lymanItalianPhrase,
		language.Japanese:             lymanJapanesePhrase,
		language.Korean:               lymanKoreanPhrase,
		language.LatinAmericanSpanish: lymanLatinAmericanSpanishPhrase,
		language.Russian:              lymanRussianPhrase,
		language.SimplifiedChinese:    lymanSimplifiedChinesePhrase,
		language.Spanish:              lymanSpanishPhrase,
		language.TraditionalChinese:   lymanTraditionalChinesePhrase}
)

var (
	// Lyman represents lyman.
	Lyman = nook.Villager{
		Character:   lymanCharacter,
		Personality: personality.Jock,
		Phrase:      lymanPhrase}
)
