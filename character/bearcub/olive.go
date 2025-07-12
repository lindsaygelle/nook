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
	// oliveBirthday represents olive birthday.
	oliveBirthday = nook.Birthday{
		Day:   12,
		Month: time.July}
)

var (
	// oliveCode represents olive code.
	oliveCode = nook.Code{
		Value: "cbr09"}
)

var (
	// oliveAmericanEnglishName represents olive american english name.
	oliveAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Olive"}

	// oliveCanadianFrenchName represents olive canadian french name.
	oliveCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grisa"}

	// oliveDutchName represents olive dutch name.
	oliveDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olive"}

	// oliveFrenchName represents olive french name.
	oliveFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grisa"}

	// oliveGermanName represents olive german name.
	oliveGermanName = nook.Name{
		Language: language.German,
		Value:    "Linda"}

	// oliveItalianName represents olive italian name.
	oliveItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Olivia"}

	// oliveJapaneseName represents olive japanese name.
	oliveJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピッコロ"}

	// oliveKoreanName represents olive korean name.
	oliveKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리브"}

	// oliveLatinAmericanSpanishName represents olive latin american spanish name.
	oliveLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Osalina"}

	// oliveRussianName represents olive russian name.
	oliveRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Олив"}

	// oliveSimplifiedChineseName represents olive simplified chinese name.
	oliveSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毕克萝"}

	// oliveSpanishName represents olive spanish name.
	oliveSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Osalina"}

	// oliveTraditionalChineseName represents olive traditional chinese name.
	oliveTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "畢克蘿"}
)

var (
	// oliveName represents olive name.
	oliveName = nook.Languages{
		language.AmericanEnglish:      oliveAmericanEnglishName,
		language.CanadianFrench:       oliveCanadianFrenchName,
		language.Dutch:                oliveDutchName,
		language.French:               oliveFrenchName,
		language.German:               oliveGermanName,
		language.Italian:              oliveItalianName,
		language.Japanese:             oliveJapaneseName,
		language.Korean:               oliveKoreanName,
		language.LatinAmericanSpanish: oliveLatinAmericanSpanishName,
		language.Russian:              oliveRussianName,
		language.SimplifiedChinese:    oliveSimplifiedChineseName,
		language.Spanish:              oliveSpanishName,
		language.TraditionalChinese:   oliveTraditionalChineseName}
)

var (
	// oliveCharacter represents olive character.
	oliveCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: oliveBirthday,
		Code:     oliveCode,
		Key:      character.Olive,
		Gender:   gender.Female,
		Name:     oliveName,
		Special:  false}
)

var (
	// oliveAmericanEnglishPhrase represents olive american english phrase.
	oliveAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sweet pea"}

	// oliveCanadianFrenchPhrase represents olive canadian french phrase.
	oliveCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trognon"}

	// oliveDutchPhrase represents olive dutch phrase.
	oliveDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "honingkelkje"}

	// oliveFrenchPhrase represents olive french phrase.
	oliveFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trognon"}

	// oliveGermanPhrase represents olive german phrase.
	oliveGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "honigtopf"}

	// oliveItalianPhrase represents olive italian phrase.
	oliveItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fagiolino"}

	// oliveJapanesePhrase represents olive japanese phrase.
	oliveJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マグ"}

	// oliveKoreanPhrase represents olive korean phrase.
	oliveKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오잉"}

	// oliveLatinAmericanSpanishPhrase represents olive latin american spanish phrase.
	oliveLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "panipof"}

	// oliveRussianPhrase represents olive russian phrase.
	oliveRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "милашка"}

	// oliveSimplifiedChinesePhrase represents olive simplified chinese phrase.
	oliveSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马克杯"}

	// oliveSpanishPhrase represents olive spanish phrase.
	oliveSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "panalito"}

	// oliveTraditionalChinesePhrase represents olive traditional chinese phrase.
	oliveTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬克杯"}
)

var (
	// olivePhrase represents olive phrase.
	olivePhrase = nook.Languages{
		language.AmericanEnglish:      oliveAmericanEnglishPhrase,
		language.CanadianFrench:       oliveCanadianFrenchPhrase,
		language.Dutch:                oliveDutchPhrase,
		language.French:               oliveFrenchPhrase,
		language.German:               oliveGermanPhrase,
		language.Italian:              oliveItalianPhrase,
		language.Japanese:             oliveJapanesePhrase,
		language.Korean:               oliveKoreanPhrase,
		language.LatinAmericanSpanish: oliveLatinAmericanSpanishPhrase,
		language.Russian:              oliveRussianPhrase,
		language.SimplifiedChinese:    oliveSimplifiedChinesePhrase,
		language.Spanish:              oliveSpanishPhrase,
		language.TraditionalChinese:   oliveTraditionalChinesePhrase}
)

var (
	// Olive represents olive.
	Olive = nook.Villager{
		Character:   oliveCharacter,
		Personality: personality.Normal,
		Phrase:      olivePhrase}
)
