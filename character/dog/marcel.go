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
	// marcelBirthday represents marcel birthday.
	marcelBirthday = nook.Birthday{
		Day:   31,
		Month: time.December}
)

var (
	// marcelCode represents marcel code.
	marcelCode = nook.Code{
		Value: "dog15"}
)

var (
	// marcelAmericanEnglishName represents marcel american english name.
	marcelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marcel"}

	// marcelCanadianFrenchName represents marcel canadian french name.
	marcelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ismaël"}

	// marcelDutchName represents marcel dutch name.
	marcelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marcel"}

	// marcelFrenchName represents marcel french name.
	marcelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ismaël"}

	// marcelGermanName represents marcel german name.
	marcelGermanName = nook.Name{
		Language: language.German,
		Value:    "Ronaldo"}

	// marcelItalianName represents marcel italian name.
	marcelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giosuè"}

	// marcelJapaneseName represents marcel japanese name.
	marcelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "もんじゃ"}

	// marcelKoreanName represents marcel korean name.
	marcelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에드워드"}

	// marcelLatinAmericanSpanishName represents marcel latin american spanish name.
	marcelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eto"}

	// marcelRussianName represents marcel russian name.
	marcelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марсель"}

	// marcelSimplifiedChineseName represents marcel simplified chinese name.
	marcelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文字烧"}

	// marcelSpanishName represents marcel spanish name.
	marcelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eto"}

	// marcelTraditionalChineseName represents marcel traditional chinese name.
	marcelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文字燒"}
)

var (
	// marcelName represents marcel name.
	marcelName = nook.Languages{
		language.AmericanEnglish:      marcelAmericanEnglishName,
		language.CanadianFrench:       marcelCanadianFrenchName,
		language.Dutch:                marcelDutchName,
		language.French:               marcelFrenchName,
		language.German:               marcelGermanName,
		language.Italian:              marcelItalianName,
		language.Japanese:             marcelJapaneseName,
		language.Korean:               marcelKoreanName,
		language.LatinAmericanSpanish: marcelLatinAmericanSpanishName,
		language.Russian:              marcelRussianName,
		language.SimplifiedChinese:    marcelSimplifiedChineseName,
		language.Spanish:              marcelSpanishName,
		language.TraditionalChinese:   marcelTraditionalChineseName}
)

var (
	// marcelCharacter represents marcel character.
	marcelCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: marcelBirthday,
		Code:     marcelCode,
		Key:      character.Marcel,
		Gender:   gender.Male,
		Name:     marcelName,
		Special:  false}
)

var (
	// marcelAmericanEnglishPhrase represents marcel american english phrase.
	marcelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "non"}

	// marcelCanadianFrenchPhrase represents marcel canadian french phrase.
	marcelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaouf"}

	// marcelDutchPhrase represents marcel dutch phrase.
	marcelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mais non"}

	// marcelFrenchPhrase represents marcel french phrase.
	marcelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaouf"}

	// marcelGermanPhrase represents marcel german phrase.
	marcelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tusch"}

	// marcelItalianPhrase represents marcel italian phrase.
	marcelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snif snif"}

	// marcelJapanesePhrase represents marcel japanese phrase.
	marcelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんじゃ"}

	// marcelKoreanPhrase represents marcel korean phrase.
	marcelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하옵소서"}

	// marcelLatinAmericanSpanishPhrase represents marcel latin american spanish phrase.
	marcelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nenonó"}

	// marcelRussianPhrase represents marcel russian phrase.
	marcelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-уи"}

	// marcelSimplifiedChinesePhrase represents marcel simplified chinese phrase.
	marcelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "什么字"}

	// marcelSpanishPhrase represents marcel spanish phrase.
	marcelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "colmillos"}

	// marcelTraditionalChinesePhrase represents marcel traditional chinese phrase.
	marcelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "什麼字"}
)

var (
	// marcelPhrase represents marcel phrase.
	marcelPhrase = nook.Languages{
		language.AmericanEnglish:      marcelAmericanEnglishPhrase,
		language.CanadianFrench:       marcelCanadianFrenchPhrase,
		language.Dutch:                marcelDutchPhrase,
		language.French:               marcelFrenchPhrase,
		language.German:               marcelGermanPhrase,
		language.Italian:              marcelItalianPhrase,
		language.Japanese:             marcelJapanesePhrase,
		language.Korean:               marcelKoreanPhrase,
		language.LatinAmericanSpanish: marcelLatinAmericanSpanishPhrase,
		language.Russian:              marcelRussianPhrase,
		language.SimplifiedChinese:    marcelSimplifiedChinesePhrase,
		language.Spanish:              marcelSpanishPhrase,
		language.TraditionalChinese:   marcelTraditionalChinesePhrase}
)

var (
	// Marcel represents marcel.
	Marcel = nook.Villager{
		Character:   marcelCharacter,
		Personality: personality.Lazy,
		Phrase:      marcelPhrase}
)
