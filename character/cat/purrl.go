package cat

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
	// purrlBirthday represents purrl birthday.
	purrlBirthday = nook.Birthday{
		Day:   29,
		Month: time.May}
)

var (
	// purrlCode represents purrl code.
	purrlCode = nook.Code{
		Value: "cat07"}
)

var (
	// purrlAmericanEnglishName represents purrl american english name.
	purrlAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Purrl"}

	// purrlCanadianFrenchName represents purrl canadian french name.
	purrlCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Perle"}

	// purrlDutchName represents purrl dutch name.
	purrlDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Purrl"}

	// purrlFrenchName represents purrl french name.
	purrlFrenchName = nook.Name{
		Language: language.French,
		Value:    "Perle"}

	// purrlGermanName represents purrl german name.
	purrlGermanName = nook.Name{
		Language: language.German,
		Value:    "Franka"}

	// purrlItalianName represents purrl italian name.
	purrlItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felidia"}

	// purrlJapaneseName represents purrl japanese name.
	purrlJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たま"}

	// purrlKoreanName represents purrl korean name.
	purrlKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "타마"}

	// purrlLatinAmericanSpanishName represents purrl latin american spanish name.
	purrlLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wanda"}

	// purrlRussianName represents purrl russian name.
	purrlRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Перл"}

	// purrlSimplifiedChineseName represents purrl simplified chinese name.
	purrlSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小玉"}

	// purrlSpanishName represents purrl spanish name.
	purrlSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Wanda"}

	// purrlTraditionalChineseName represents purrl traditional chinese name.
	purrlTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小玉"}
)

var (
	// purrlName represents purrl name.
	purrlName = nook.Languages{
		language.AmericanEnglish:      purrlAmericanEnglishName,
		language.CanadianFrench:       purrlCanadianFrenchName,
		language.Dutch:                purrlDutchName,
		language.French:               purrlFrenchName,
		language.German:               purrlGermanName,
		language.Italian:              purrlItalianName,
		language.Japanese:             purrlJapaneseName,
		language.Korean:               purrlKoreanName,
		language.LatinAmericanSpanish: purrlLatinAmericanSpanishName,
		language.Russian:              purrlRussianName,
		language.SimplifiedChinese:    purrlSimplifiedChineseName,
		language.Spanish:              purrlSpanishName,
		language.TraditionalChinese:   purrlTraditionalChineseName}
)

var (
	// purrlCharacter represents purrl character.
	purrlCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: purrlBirthday,
		Code:     purrlCode,
		Key:      character.Purrl,
		Gender:   gender.Female,
		Name:     purrlName,
		Special:  false}
)

var (
	// purrlAmericanEnglishPhrase represents purrl american english phrase.
	purrlAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kitten"}

	// purrlCanadianFrenchPhrase represents purrl canadian french phrase.
	purrlCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chaton"}

	// purrlDutchPhrase represents purrl dutch phrase.
	purrlDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kitten"}

	// purrlFrenchPhrase represents purrl french phrase.
	purrlFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chaton"}

	// purrlGermanPhrase represents purrl german phrase.
	purrlGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kätzchen"}

	// purrlItalianPhrase represents purrl italian phrase.
	purrlItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "micetto"}

	// purrlJapanesePhrase represents purrl japanese phrase.
	purrlJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふんっ"}

	// purrlKoreanPhrase represents purrl korean phrase.
	purrlKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냠"}

	// purrlLatinAmericanSpanishPhrase represents purrl latin american spanish phrase.
	purrlLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michimiu"}

	// purrlRussianPhrase represents purrl russian phrase.
	purrlRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кис-кис"}

	// purrlSimplifiedChinesePhrase represents purrl simplified chinese phrase.
	purrlSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯嗯"}

	// purrlSpanishPhrase represents purrl spanish phrase.
	purrlSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tontaina"}

	// purrlTraditionalChinesePhrase represents purrl traditional chinese phrase.
	purrlTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯嗯"}
)

var (
	// purrlPhrase represents purrl phrase.
	purrlPhrase = nook.Languages{
		language.AmericanEnglish:      purrlAmericanEnglishPhrase,
		language.CanadianFrench:       purrlCanadianFrenchPhrase,
		language.Dutch:                purrlDutchPhrase,
		language.French:               purrlFrenchPhrase,
		language.German:               purrlGermanPhrase,
		language.Italian:              purrlItalianPhrase,
		language.Japanese:             purrlJapanesePhrase,
		language.Korean:               purrlKoreanPhrase,
		language.LatinAmericanSpanish: purrlLatinAmericanSpanishPhrase,
		language.Russian:              purrlRussianPhrase,
		language.SimplifiedChinese:    purrlSimplifiedChinesePhrase,
		language.Spanish:              purrlSpanishPhrase,
		language.TraditionalChinese:   purrlTraditionalChinesePhrase}
)

var (
	// Purrl represents purrl.
	Purrl = nook.Villager{
		Character:   purrlCharacter,
		Personality: personality.Snooty,
		Phrase:      purrlPhrase}
)
