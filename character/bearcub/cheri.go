package bearcub

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
	// cheriBirthday represents cheri birthday.
	cheriBirthday = nook.Birthday{
		Day:   17,
		Month: time.March}
)

var (
	// cheriCode represents cheri code.
	cheriCode = nook.Code{
		Value: "cbr10"}
)

var (
	// cheriAmericanEnglishName represents cheri american english name.
	cheriAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cheri"}

	// cheriCanadianFrenchName represents cheri canadian french name.
	cheriCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosalie"}

	// cheriDutchName represents cheri dutch name.
	cheriDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cheri"}

	// cheriFrenchName represents cheri french name.
	cheriFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosalie"}

	// cheriGermanName represents cheri german name.
	cheriGermanName = nook.Name{
		Language: language.German,
		Value:    "Claudia"}

	// cheriItalianName represents cheri italian name.
	cheriItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bonbon"}

	// cheriJapaneseName represents cheri japanese name.
	cheriJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アセロラ"}

	// cheriKoreanName represents cheri korean name.
	cheriKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아세로라"}

	// cheriLatinAmericanSpanishName represents cheri latin american spanish name.
	cheriLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cerecita"}

	// cheriRussianName represents cheri russian name.
	cheriRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шери"}

	// cheriSimplifiedChineseName represents cheri simplified chinese name.
	cheriSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "樱桃"}

	// cheriSpanishName represents cheri spanish name.
	cheriSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cerecita"}

	// cheriTraditionalChineseName represents cheri traditional chinese name.
	cheriTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "櫻桃"}
)

var (
	// cheriName represents cheri name.
	cheriName = nook.Languages{
		language.AmericanEnglish:      cheriAmericanEnglishName,
		language.CanadianFrench:       cheriCanadianFrenchName,
		language.Dutch:                cheriDutchName,
		language.French:               cheriFrenchName,
		language.German:               cheriGermanName,
		language.Italian:              cheriItalianName,
		language.Japanese:             cheriJapaneseName,
		language.Korean:               cheriKoreanName,
		language.LatinAmericanSpanish: cheriLatinAmericanSpanishName,
		language.Russian:              cheriRussianName,
		language.SimplifiedChinese:    cheriSimplifiedChineseName,
		language.Spanish:              cheriSpanishName,
		language.TraditionalChinese:   cheriTraditionalChineseName}
)

var (
	// cheriCharacter represents cheri character.
	cheriCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: cheriBirthday,
		Code:     cheriCode,
		Key:      character.Cheri,
		Gender:   gender.Female,
		Name:     cheriName,
		Special:  false}
)

var (
	// cheriAmericanEnglishPhrase represents cheri american english phrase.
	cheriAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tralala"}

	// cheriCanadianFrenchPhrase represents cheri canadian french phrase.
	cheriCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tralala"}

	// cheriDutchPhrase represents cheri dutch phrase.
	cheriDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tralala"}

	// cheriFrenchPhrase represents cheri french phrase.
	cheriFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tralala"}

	// cheriGermanPhrase represents cheri german phrase.
	cheriGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tralala"}

	// cheriItalianPhrase represents cheri italian phrase.
	cheriItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trallalà"}

	// cheriJapanesePhrase represents cheri japanese phrase.
	cheriJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんてね"}

	// cheriKoreanPhrase represents cheri korean phrase.
	cheriKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어쩜이래"}

	// cheriLatinAmericanSpanishPhrase represents cheri latin american spanish phrase.
	cheriLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralará"}

	// cheriRussianPhrase represents cheri russian phrase.
	cheriRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тра-ля-ля"}

	// cheriSimplifiedChinesePhrase represents cheri simplified chinese phrase.
	cheriSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开玩笑的"}

	// cheriSpanishPhrase represents cheri spanish phrase.
	cheriSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tralará"}

	// cheriTraditionalChinesePhrase represents cheri traditional chinese phrase.
	cheriTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開玩笑的"}
)

var (
	// cheriPhrase represents cheri phrase.
	cheriPhrase = nook.Languages{
		language.AmericanEnglish:      cheriAmericanEnglishPhrase,
		language.CanadianFrench:       cheriCanadianFrenchPhrase,
		language.Dutch:                cheriDutchPhrase,
		language.French:               cheriFrenchPhrase,
		language.German:               cheriGermanPhrase,
		language.Italian:              cheriItalianPhrase,
		language.Japanese:             cheriJapanesePhrase,
		language.Korean:               cheriKoreanPhrase,
		language.LatinAmericanSpanish: cheriLatinAmericanSpanishPhrase,
		language.Russian:              cheriRussianPhrase,
		language.SimplifiedChinese:    cheriSimplifiedChinesePhrase,
		language.Spanish:              cheriSpanishPhrase,
		language.TraditionalChinese:   cheriTraditionalChinesePhrase}
)

var (
	// Cheri represents cheri.
	Cheri = nook.Villager{
		Character:   cheriCharacter,
		Personality: personality.Peppy,
		Phrase:      cheriPhrase}
)
