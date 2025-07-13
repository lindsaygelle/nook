package chicken

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
	// kenBirthday represents ken birthday.
	kenBirthday = nook.Birthday{
		Day:   23,
		Month: time.December}
)

var (
	// kenCode represents ken code.
	kenCode = nook.Code{
		Value: "chn13"}
)

var (
	// kenAmericanEnglishName represents ken american english name.
	kenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ken"}

	// kenCanadianFrenchName represents ken canadian french name.
	kenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ken"}

	// kenDutchName represents ken dutch name.
	kenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ken"}

	// kenFrenchName represents ken french name.
	kenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ken"}

	// kenGermanName represents ken german name.
	kenGermanName = nook.Name{
		Language: language.German,
		Value:    "Hannes"}

	// kenItalianName represents ken italian name.
	kenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aligalli"}

	// kenJapaneseName represents ken japanese name.
	kenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロベエ"}

	// kenKoreanName represents ken korean name.
	kenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오골"}

	// kenLatinAmericanSpanishName represents ken latin american spanish name.
	kenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pichón"}

	// kenRussianName represents ken russian name.
	kenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кен"}

	// kenSimplifiedChineseName represents ken simplified chinese name.
	kenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "乌骨鸡"}

	// kenSpanishName represents ken spanish name.
	kenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pichón"}

	// kenTraditionalChineseName represents ken traditional chinese name.
	kenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "烏骨雞"}
)

var (
	// kenName represents ken name.
	kenName = nook.Languages{
		language.AmericanEnglish:      kenAmericanEnglishName,
		language.CanadianFrench:       kenCanadianFrenchName,
		language.Dutch:                kenDutchName,
		language.French:               kenFrenchName,
		language.German:               kenGermanName,
		language.Italian:              kenItalianName,
		language.Japanese:             kenJapaneseName,
		language.Korean:               kenKoreanName,
		language.LatinAmericanSpanish: kenLatinAmericanSpanishName,
		language.Russian:              kenRussianName,
		language.SimplifiedChinese:    kenSimplifiedChineseName,
		language.Spanish:              kenSpanishName,
		language.TraditionalChinese:   kenTraditionalChineseName}
)

var (
	// kenCharacter represents ken character.
	kenCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: kenBirthday,
		Code:     kenCode,
		Key:      character.Ken,
		Gender:   gender.Male,
		Name:     kenName,
		Special:  false}
)

var (
	// kenAmericanEnglishPhrase represents ken american english phrase.
	kenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "no doubt"}

	// kenCanadianFrenchPhrase represents ken canadian french phrase.
	kenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chicken"}

	// kenDutchPhrase represents ken dutch phrase.
	kenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eiteraard"}

	// kenFrenchPhrase represents ken french phrase.
	kenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pôt pôt"}

	// kenGermanPhrase represents ken german phrase.
	kenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gockgock"}

	// kenItalianPhrase represents ken italian phrase.
	kenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quiriqui"}

	// kenJapanesePhrase represents ken japanese phrase.
	kenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "コッケイ"}

	// kenKoreanPhrase represents ken korean phrase.
	kenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "깜깜"}

	// kenLatinAmericanSpanishPhrase represents ken latin american spanish phrase.
	kenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kirrí"}

	// kenRussianPhrase represents ken russian phrase.
	kenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "это точно"}

	// kenSimplifiedChinesePhrase represents ken simplified chinese phrase.
	kenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "骨鸡"}

	// kenSpanishPhrase represents ken spanish phrase.
	kenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kirrí"}

	// kenTraditionalChinesePhrase represents ken traditional chinese phrase.
	kenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "骨雞"}
)

var (
	// kenPhrase represents ken phrase.
	kenPhrase = nook.Languages{
		language.AmericanEnglish:      kenAmericanEnglishPhrase,
		language.CanadianFrench:       kenCanadianFrenchPhrase,
		language.Dutch:                kenDutchPhrase,
		language.French:               kenFrenchPhrase,
		language.German:               kenGermanPhrase,
		language.Italian:              kenItalianPhrase,
		language.Japanese:             kenJapanesePhrase,
		language.Korean:               kenKoreanPhrase,
		language.LatinAmericanSpanish: kenLatinAmericanSpanishPhrase,
		language.Russian:              kenRussianPhrase,
		language.SimplifiedChinese:    kenSimplifiedChinesePhrase,
		language.Spanish:              kenSpanishPhrase,
		language.TraditionalChinese:   kenTraditionalChinesePhrase}
)

var (
	// Ken represents ken.
	Ken = nook.Villager{
		Character:   kenCharacter,
		Personality: personality.Smug,
		Phrase:      kenPhrase}
)
