package horse

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
	// winnieBirthday represents winnie birthday.
	winnieBirthday = nook.Birthday{
		Day:   31,
		Month: time.January}
)

var (
	// winnieCode represents winnie code.
	winnieCode = nook.Code{
		Value: "hrs05"}
)

var (
	// winnieAmericanEnglishName represents winnie american english name.
	winnieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Winnie"}

	// winnieCanadianFrenchName represents winnie canadian french name.
	winnieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anne"}

	// winnieDutchName represents winnie dutch name.
	winnieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Winnie"}

	// winnieFrenchName represents winnie french name.
	winnieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anne"}

	// winnieGermanName represents winnie german name.
	winnieGermanName = nook.Name{
		Language: language.German,
		Value:    "Walli"}

	// winnieItalianName represents winnie italian name.
	winnieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Clara"}

	// winnieJapaneseName represents winnie japanese name.
	winnieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マキバスター"}

	// winnieKoreanName represents winnie korean name.
	winnieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카로틴"}

	// winnieLatinAmericanSpanishName represents winnie latin american spanish name.
	winnieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Soonia"}

	// winnieRussianName represents winnie russian name.
	winnieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Винни"}

	// winnieSimplifiedChineseName represents winnie simplified chinese name.
	winnieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马星星"}

	// winnieSpanishName represents winnie spanish name.
	winnieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Soonia"}

	// winnieTraditionalChineseName represents winnie traditional chinese name.
	winnieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬星星"}
)

var (
	// winnieName represents winnie name.
	winnieName = nook.Languages{
		language.AmericanEnglish:      winnieAmericanEnglishName,
		language.CanadianFrench:       winnieCanadianFrenchName,
		language.Dutch:                winnieDutchName,
		language.French:               winnieFrenchName,
		language.German:               winnieGermanName,
		language.Italian:              winnieItalianName,
		language.Japanese:             winnieJapaneseName,
		language.Korean:               winnieKoreanName,
		language.LatinAmericanSpanish: winnieLatinAmericanSpanishName,
		language.Russian:              winnieRussianName,
		language.SimplifiedChinese:    winnieSimplifiedChineseName,
		language.Spanish:              winnieSpanishName,
		language.TraditionalChinese:   winnieTraditionalChineseName}
)

var (
	// winnieCharacter represents winnie character.
	winnieCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: winnieBirthday,
		Code:     winnieCode,
		Key:      character.Winnie,
		Gender:   gender.Female,
		Name:     winnieName,
		Special:  false}
)

var (
	// winnieAmericanEnglishPhrase represents winnie american english phrase.
	winnieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hay-OK"}

	// winnieCanadianFrenchPhrase represents winnie canadian french phrase.
	winnieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "youp là"}

	// winnieDutchPhrase represents winnie dutch phrase.
	winnieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knol"}

	// winnieFrenchPhrase represents winnie french phrase.
	winnieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "youp là"}

	// winnieGermanPhrase represents winnie german phrase.
	winnieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heu-ja"}

	// winnieItalianPhrase represents winnie italian phrase.
	winnieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "biadina"}

	// winnieJapanesePhrase represents winnie japanese phrase.
	winnieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "みゃあ"}

	// winnieKoreanPhrase represents winnie korean phrase.
	winnieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "당근당근"}

	// winnieLatinAmericanSpanishPhrase represents winnie latin american spanish phrase.
	winnieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "okeimakei"}

	// winnieRussianPhrase represents winnie russian phrase.
	winnieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тык-дык"}

	// winnieSimplifiedChinesePhrase represents winnie simplified chinese phrase.
	winnieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘶嘶"}

	// winnieSpanishPhrase represents winnie spanish phrase.
	winnieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "okeimakei"}

	// winnieTraditionalChinesePhrase represents winnie traditional chinese phrase.
	winnieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘶嘶"}
)

var (
	// winniePhrase represents winnie phrase.
	winniePhrase = nook.Languages{
		language.AmericanEnglish:      winnieAmericanEnglishPhrase,
		language.CanadianFrench:       winnieCanadianFrenchPhrase,
		language.Dutch:                winnieDutchPhrase,
		language.French:               winnieFrenchPhrase,
		language.German:               winnieGermanPhrase,
		language.Italian:              winnieItalianPhrase,
		language.Japanese:             winnieJapanesePhrase,
		language.Korean:               winnieKoreanPhrase,
		language.LatinAmericanSpanish: winnieLatinAmericanSpanishPhrase,
		language.Russian:              winnieRussianPhrase,
		language.SimplifiedChinese:    winnieSimplifiedChinesePhrase,
		language.Spanish:              winnieSpanishPhrase,
		language.TraditionalChinese:   winnieTraditionalChinesePhrase}
)

var (
	// Winnie represents winnie.
	Winnie = nook.Villager{
		Character:   winnieCharacter,
		Personality: personality.Peppy,
		Phrase:      winniePhrase}
)
