package kangaroo

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
	// kittBirthday represents kitt birthday.
	kittBirthday = nook.Birthday{
		Day:   11,
		Month: time.October}
)

var (
	// kittCode represents kitt code.
	kittCode = nook.Code{
		Value: "kgr00"}
)

var (
	// kittAmericanEnglishName represents kitt american english name.
	kittAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kitt"}

	// kittCanadianFrenchName represents kitt canadian french name.
	kittCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Poquette"}

	// kittDutchName represents kitt dutch name.
	kittDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kitt"}

	// kittFrenchName represents kitt french name.
	kittFrenchName = nook.Name{
		Language: language.French,
		Value:    "Poquette"}

	// kittGermanName represents kitt german name.
	kittGermanName = nook.Name{
		Language: language.German,
		Value:    "Kerstin"}

	// kittItalianName represents kitt italian name.
	kittItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kitt"}

	// kittJapaneseName represents kitt japanese name.
	kittJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アップリケ"}

	// kittKoreanName represents kitt korean name.
	kittKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "애플리케"}

	// kittLatinAmericanSpanishName represents kitt latin american spanish name.
	kittLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Antípoda"}

	// kittRussianName represents kitt russian name.
	kittRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китт"}

	// kittSimplifiedChineseName represents kitt simplified chinese name.
	kittSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "缝缝"}

	// kittSpanishName represents kitt spanish name.
	kittSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Antípoda"}

	// kittTraditionalChineseName represents kitt traditional chinese name.
	kittTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "縫縫"}
)

var (
	// kittName represents kitt name.
	kittName = nook.Languages{
		language.AmericanEnglish:      kittAmericanEnglishName,
		language.CanadianFrench:       kittCanadianFrenchName,
		language.Dutch:                kittDutchName,
		language.French:               kittFrenchName,
		language.German:               kittGermanName,
		language.Italian:              kittItalianName,
		language.Japanese:             kittJapaneseName,
		language.Korean:               kittKoreanName,
		language.LatinAmericanSpanish: kittLatinAmericanSpanishName,
		language.Russian:              kittRussianName,
		language.SimplifiedChinese:    kittSimplifiedChineseName,
		language.Spanish:              kittSpanishName,
		language.TraditionalChinese:   kittTraditionalChineseName}
)

var (
	// kittCharacter represents kitt character.
	kittCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: kittBirthday,
		Code:     kittCode,
		Key:      character.Kitt,
		Gender:   gender.Female,
		Name:     kittName,
		Special:  false}
)

var (
	// kittAmericanEnglishPhrase represents kitt american english phrase.
	kittAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "child"}

	// kittCanadianFrenchPhrase represents kitt canadian french phrase.
	kittCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gamin"}

	// kittDutchPhrase represents kitt dutch phrase.
	kittDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kind"}

	// kittFrenchPhrase represents kitt french phrase.
	kittFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gamin"}

	// kittGermanPhrase represents kitt german phrase.
	kittGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kindchen"}

	// kittItalianPhrase represents kitt italian phrase.
	kittItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cucciolo"}

	// kittJapanesePhrase represents kitt japanese phrase.
	kittJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッポン"}

	// kittKoreanPhrase represents kitt korean phrase.
	kittKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "폴짝"}

	// kittLatinAmericanSpanishPhrase represents kitt latin american spanish phrase.
	kittLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chiquichí"}

	// kittRussianPhrase represents kitt russian phrase.
	kittRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "детка"}

	// kittSimplifiedChinesePhrase represents kitt simplified chinese phrase.
	kittSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "补补"}

	// kittSpanishPhrase represents kitt spanish phrase.
	kittSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mi alma"}

	// kittTraditionalChinesePhrase represents kitt traditional chinese phrase.
	kittTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "補補"}
)

var (
	// kittPhrase represents kitt phrase.
	kittPhrase = nook.Languages{
		language.AmericanEnglish:      kittAmericanEnglishPhrase,
		language.CanadianFrench:       kittCanadianFrenchPhrase,
		language.Dutch:                kittDutchPhrase,
		language.French:               kittFrenchPhrase,
		language.German:               kittGermanPhrase,
		language.Italian:              kittItalianPhrase,
		language.Japanese:             kittJapanesePhrase,
		language.Korean:               kittKoreanPhrase,
		language.LatinAmericanSpanish: kittLatinAmericanSpanishPhrase,
		language.Russian:              kittRussianPhrase,
		language.SimplifiedChinese:    kittSimplifiedChinesePhrase,
		language.Spanish:              kittSpanishPhrase,
		language.TraditionalChinese:   kittTraditionalChinesePhrase}
)

var (
	// Kitt represents kitt.
	Kitt = nook.Villager{
		Character:   kittCharacter,
		Personality: personality.Normal,
		Phrase:      kittPhrase}
)
