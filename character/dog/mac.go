package dog

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
	// macBirthday represents mac birthday.
	macBirthday = nook.Birthday{
		Day:   11,
		Month: time.November}
)

var (
	// macCode represents mac code.
	macCode = nook.Code{
		Value: "dog14"}
)

var (
	// macAmericanEnglishName represents mac american english name.
	macAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mac"}

	// macCanadianFrenchName represents mac canadian french name.
	macCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Brutus"}

	// macDutchName represents mac dutch name.
	macDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mac"}

	// macFrenchName represents mac french name.
	macFrenchName = nook.Name{
		Language: language.French,
		Value:    "Brutus"}

	// macGermanName represents mac german name.
	macGermanName = nook.Name{
		Language: language.German,
		Value:    "Wuffi"}

	// macItalianName represents mac italian name.
	macItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teo"}

	// macJapaneseName represents mac japanese name.
	macJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャンプ"}

	// macKoreanName represents mac korean name.
	macKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "챔프"}

	// macLatinAmericanSpanishName represents mac latin american spanish name.
	macLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pit"}

	// macRussianName represents mac russian name.
	macRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мак"}

	// macSimplifiedChineseName represents mac simplified chinese name.
	macSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "金牌"}

	// macSpanishName represents mac spanish name.
	macSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pit"}

	// macTraditionalChineseName represents mac traditional chinese name.
	macTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "金牌"}
)

var (
	// macName represents mac name.
	macName = nook.Languages{
		language.AmericanEnglish:      macAmericanEnglishName,
		language.CanadianFrench:       macCanadianFrenchName,
		language.Dutch:                macDutchName,
		language.French:               macFrenchName,
		language.German:               macGermanName,
		language.Italian:              macItalianName,
		language.Japanese:             macJapaneseName,
		language.Korean:               macKoreanName,
		language.LatinAmericanSpanish: macLatinAmericanSpanishName,
		language.Russian:              macRussianName,
		language.SimplifiedChinese:    macSimplifiedChineseName,
		language.Spanish:              macSpanishName,
		language.TraditionalChinese:   macTraditionalChineseName}
)

var (
	// macCharacter represents mac character.
	macCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: macBirthday,
		Code:     macCode,
		Key:      character.Mac,
		Gender:   gender.Male,
		Name:     macName,
		Special:  false}
)

var (
	// macAmericanEnglishPhrase represents mac american english phrase.
	macAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woo woof"}

	// macCanadianFrenchPhrase represents mac canadian french phrase.
	macCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "babines"}

	// macDutchPhrase represents mac dutch phrase.
	macDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "woewoef"}

	// macFrenchPhrase represents mac french phrase.
	macFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "babines"}

	// macGermanPhrase represents mac german phrase.
	macGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bellbell"}

	// macItalianPhrase represents mac italian phrase.
	macItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arfidenti"}

	// macJapanesePhrase represents mac japanese phrase.
	macJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "モグ"}

	// macKoreanPhrase represents mac korean phrase.
	macKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우물우물"}

	// macLatinAmericanSpanishPhrase represents mac latin american spanish phrase.
	macLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "frusky"}

	// macRussianPhrase represents mac russian phrase.
	macRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-у-гав"}

	// macSimplifiedChinesePhrase represents mac simplified chinese phrase.
	macSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "摘"}

	// macSpanishPhrase represents mac spanish phrase.
	macSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "frusky"}

	// macTraditionalChinesePhrase represents mac traditional chinese phrase.
	macTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "摘"}
)

var (
	// macPhrase represents mac phrase.
	macPhrase = nook.Languages{
		language.AmericanEnglish:      macAmericanEnglishPhrase,
		language.CanadianFrench:       macCanadianFrenchPhrase,
		language.Dutch:                macDutchPhrase,
		language.French:               macFrenchPhrase,
		language.German:               macGermanPhrase,
		language.Italian:              macItalianPhrase,
		language.Japanese:             macJapanesePhrase,
		language.Korean:               macKoreanPhrase,
		language.LatinAmericanSpanish: macLatinAmericanSpanishPhrase,
		language.Russian:              macRussianPhrase,
		language.SimplifiedChinese:    macSimplifiedChinesePhrase,
		language.Spanish:              macSpanishPhrase,
		language.TraditionalChinese:   macTraditionalChinesePhrase}
)

var (
	// Mac represents mac.
	Mac = nook.Villager{
		Character:   macCharacter,
		Personality: personality.Jock,
		Phrase:      macPhrase}
)
