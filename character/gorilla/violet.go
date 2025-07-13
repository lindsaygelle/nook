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
	// violetBirthday represents violet birthday.
	violetBirthday = nook.Birthday{
		Day:   1,
		Month: time.September}
)

var (
	// violetCode represents violet code.
	violetCode = nook.Code{
		Value: "gor07"}
)

var (
	// violetAmericanEnglishName represents violet american english name.
	violetAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Violet"}

	// violetCanadianFrenchName represents violet canadian french name.
	violetCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gaëlle"}

	// violetDutchName represents violet dutch name.
	violetDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Violet"}

	// violetFrenchName represents violet french name.
	violetFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaëlle"}

	// violetGermanName represents violet german name.
	violetGermanName = nook.Name{
		Language: language.German,
		Value:    "Konga"}

	// violetItalianName represents violet italian name.
	violetItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Konga"}

	// violetJapaneseName represents violet japanese name.
	violetJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウズメ"}

	// violetKoreanName represents violet korean name.
	violetKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "줌마"}

	// violetLatinAmericanSpanishName represents violet latin american spanish name.
	violetLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Magenta"}

	// violetRussianName represents violet russian name.
	violetRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Конга"}

	// violetSimplifiedChineseName represents violet simplified chinese name.
	violetSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吴紫眉"}

	// violetSpanishName represents violet spanish name.
	violetSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Magenta"}

	// violetTraditionalChineseName represents violet traditional chinese name.
	violetTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吳紫眉"}
)

var (
	// violetName represents violet name.
	violetName = nook.Languages{
		language.AmericanEnglish:      violetAmericanEnglishName,
		language.CanadianFrench:       violetCanadianFrenchName,
		language.Dutch:                violetDutchName,
		language.French:               violetFrenchName,
		language.German:               violetGermanName,
		language.Italian:              violetItalianName,
		language.Japanese:             violetJapaneseName,
		language.Korean:               violetKoreanName,
		language.LatinAmericanSpanish: violetLatinAmericanSpanishName,
		language.Russian:              violetRussianName,
		language.SimplifiedChinese:    violetSimplifiedChineseName,
		language.Spanish:              violetSpanishName,
		language.TraditionalChinese:   violetTraditionalChineseName}
)

var (
	// violetCharacter represents violet character.
	violetCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: violetBirthday,
		Code:     violetCode,
		Key:      character.Violet,
		Gender:   gender.Female,
		Name:     violetName,
		Special:  false}
)

var (
	// violetAmericanEnglishPhrase represents violet american english phrase.
	violetAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "faboom"}

	// violetCanadianFrenchPhrase represents violet canadian french phrase.
	violetCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gong"}

	// violetDutchPhrase represents violet dutch phrase.
	violetDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lieverd"}

	// violetFrenchPhrase represents violet french phrase.
	violetFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gong"}

	// violetGermanPhrase represents violet german phrase.
	violetGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "auauauu"}

	// violetItalianPhrase represents violet italian phrase.
	violetItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tump tump"}

	// violetJapanesePhrase represents violet japanese phrase.
	violetJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アイヤ"}

	// violetKoreanPhrase represents violet korean phrase.
	violetKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "에그머닛"}

	// violetLatinAmericanSpanishPhrase represents violet latin american spanish phrase.
	violetLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-uh"}

	// violetRussianPhrase represents violet russian phrase.
	violetRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сластена"}

	// violetSimplifiedChinesePhrase represents violet simplified chinese phrase.
	violetSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唉呀"}

	// violetSpanishPhrase represents violet spanish phrase.
	violetSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "canapé"}

	// violetTraditionalChinesePhrase represents violet traditional chinese phrase.
	violetTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唉呀"}
)

var (
	// violetPhrase represents violet phrase.
	violetPhrase = nook.Languages{
		language.AmericanEnglish:      violetAmericanEnglishPhrase,
		language.CanadianFrench:       violetCanadianFrenchPhrase,
		language.Dutch:                violetDutchPhrase,
		language.French:               violetFrenchPhrase,
		language.German:               violetGermanPhrase,
		language.Italian:              violetItalianPhrase,
		language.Japanese:             violetJapanesePhrase,
		language.Korean:               violetKoreanPhrase,
		language.LatinAmericanSpanish: violetLatinAmericanSpanishPhrase,
		language.Russian:              violetRussianPhrase,
		language.SimplifiedChinese:    violetSimplifiedChinesePhrase,
		language.Spanish:              violetSpanishPhrase,
		language.TraditionalChinese:   violetTraditionalChinesePhrase}
)

var (
	// Violet represents violet.
	Violet = nook.Villager{
		Character:   violetCharacter,
		Personality: personality.Snooty,
		Phrase:      violetPhrase}
)
