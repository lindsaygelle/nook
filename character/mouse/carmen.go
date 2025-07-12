package mouse

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
	// carmenBirthday represents carmen birthday.
	carmenBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// carmenCode represents carmen code.
	carmenCode = nook.Code{
		Value: ""}
)

var (
	// carmenAmericanEnglishName represents carmen american english name.
	carmenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carmen"}

	// carmenCanadianFrenchName represents carmen canadian french name.
	carmenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// carmenDutchName represents carmen dutch name.
	carmenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// carmenFrenchName represents carmen french name.
	carmenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Carmen"}

	// carmenGermanName represents carmen german name.
	carmenGermanName = nook.Name{
		Language: language.German,
		Value:    "Carmen"}

	// carmenItalianName represents carmen italian name.
	carmenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmen"}

	// carmenJapaneseName represents carmen japanese name.
	carmenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゼリー"}

	// carmenKoreanName represents carmen korean name.
	carmenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// carmenLatinAmericanSpanishName represents carmen latin american spanish name.
	carmenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// carmenRussianName represents carmen russian name.
	carmenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// carmenSimplifiedChineseName represents carmen simplified chinese name.
	carmenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卡门"}

	// carmenSpanishName represents carmen spanish name.
	carmenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmen"}

	// carmenTraditionalChineseName represents carmen traditional chinese name.
	carmenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// carmenName represents carmen name.
	carmenName = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishName,
		language.CanadianFrench:       carmenCanadianFrenchName,
		language.Dutch:                carmenDutchName,
		language.French:               carmenFrenchName,
		language.German:               carmenGermanName,
		language.Italian:              carmenItalianName,
		language.Japanese:             carmenJapaneseName,
		language.Korean:               carmenKoreanName,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishName,
		language.Russian:              carmenRussianName,
		language.SimplifiedChinese:    carmenSimplifiedChineseName,
		language.Spanish:              carmenSpanishName,
		language.TraditionalChinese:   carmenTraditionalChineseName}
)

var (
	// carmenCharacter represents carmen character.
	carmenCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: carmenBirthday,
		Code:     carmenCode,
		Key:      character.Carmen,
		Gender:   gender.Female,
		Name:     carmenName,
		Special:  false}
)

var (
	// carmenAmericanEnglishPhrase represents carmen american english phrase.
	carmenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bumpkin"}

	// carmenCanadianFrenchPhrase represents carmen canadian french phrase.
	carmenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// carmenDutchPhrase represents carmen dutch phrase.
	carmenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// carmenFrenchPhrase represents carmen french phrase.
	carmenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mamour"}

	// carmenGermanPhrase represents carmen german phrase.
	carmenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "äpfelchen"}

	// carmenItalianPhrase represents carmen italian phrase.
	carmenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "buffetto"}

	// carmenJapanesePhrase represents carmen japanese phrase.
	carmenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "んとんと"}

	// carmenKoreanPhrase represents carmen korean phrase.
	carmenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// carmenLatinAmericanSpanishPhrase represents carmen latin american spanish phrase.
	carmenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// carmenRussianPhrase represents carmen russian phrase.
	carmenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// carmenSimplifiedChinesePhrase represents carmen simplified chinese phrase.
	carmenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱哧"}

	// carmenSpanishPhrase represents carmen spanish phrase.
	carmenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bolilla"}

	// carmenTraditionalChinesePhrase represents carmen traditional chinese phrase.
	carmenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// carmenPhrase represents carmen phrase.
	carmenPhrase = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishPhrase,
		language.CanadianFrench:       carmenCanadianFrenchPhrase,
		language.Dutch:                carmenDutchPhrase,
		language.French:               carmenFrenchPhrase,
		language.German:               carmenGermanPhrase,
		language.Italian:              carmenItalianPhrase,
		language.Japanese:             carmenJapanesePhrase,
		language.Korean:               carmenKoreanPhrase,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishPhrase,
		language.Russian:              carmenRussianPhrase,
		language.SimplifiedChinese:    carmenSimplifiedChinesePhrase,
		language.Spanish:              carmenSpanishPhrase,
		language.TraditionalChinese:   carmenTraditionalChinesePhrase}
)

var (
	// Carmen represents carmen.
	Carmen = nook.Villager{
		Character:   carmenCharacter,
		Personality: personality.Snooty,
		Phrase:      carmenPhrase}
)
