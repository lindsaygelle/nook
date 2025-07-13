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
	// merryBirthday represents merry birthday.
	merryBirthday = nook.Birthday{
		Day:   29,
		Month: time.June}
)

var (
	// merryCode represents merry code.
	merryCode = nook.Code{
		Value: "cat16"}
)

var (
	// merryAmericanEnglishName represents merry american english name.
	merryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Merry"}

	// merryCanadianFrenchName represents merry canadian french name.
	merryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Suzy"}

	// merryDutchName represents merry dutch name.
	merryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Merry"}

	// merryFrenchName represents merry french name.
	merryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Suzy"}

	// merryGermanName represents merry german name.
	merryGermanName = nook.Name{
		Language: language.German,
		Value:    "Mischka"}

	// merryItalianName represents merry italian name.
	merryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Katy"}

	// merryJapaneseName represents merry japanese name.
	merryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さっち"}

	// merryKoreanName represents merry korean name.
	merryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유네찌"}

	// merryLatinAmericanSpanishName represents merry latin american spanish name.
	merryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Susi"}

	// merryRussianName represents merry russian name.
	merryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мерри"}

	// merrySimplifiedChineseName represents merry simplified chinese name.
	merrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎莎"}

	// merrySpanishName represents merry spanish name.
	merrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Susi"}

	// merryTraditionalChineseName represents merry traditional chinese name.
	merryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎莎"}
)

var (
	// merryName represents merry name.
	merryName = nook.Languages{
		language.AmericanEnglish:      merryAmericanEnglishName,
		language.CanadianFrench:       merryCanadianFrenchName,
		language.Dutch:                merryDutchName,
		language.French:               merryFrenchName,
		language.German:               merryGermanName,
		language.Italian:              merryItalianName,
		language.Japanese:             merryJapaneseName,
		language.Korean:               merryKoreanName,
		language.LatinAmericanSpanish: merryLatinAmericanSpanishName,
		language.Russian:              merryRussianName,
		language.SimplifiedChinese:    merrySimplifiedChineseName,
		language.Spanish:              merrySpanishName,
		language.TraditionalChinese:   merryTraditionalChineseName}
)

var (
	// merryCharacter represents merry character.
	merryCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: merryBirthday,
		Code:     merryCode,
		Key:      character.Merry,
		Gender:   gender.Female,
		Name:     merryName,
		Special:  false}
)

var (
	// merryAmericanEnglishPhrase represents merry american english phrase.
	merryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mweee"}

	// merryCanadianFrenchPhrase represents merry canadian french phrase.
	merryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miou-mi"}

	// merryDutchPhrase represents merry dutch phrase.
	merryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mi-jippie"}

	// merryFrenchPhrase represents merry french phrase.
	merryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miou-mi"}

	// merryGermanPhrase represents merry german phrase.
	merryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "prrrrr"}

	// merryItalianPhrase represents merry italian phrase.
	merryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pffffft"}

	// merryJapanesePhrase represents merry japanese phrase.
	merryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃん"}

	// merryKoreanPhrase represents merry korean phrase.
	merryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냥냥"}

	// merryLatinAmericanSpanishPhrase represents merry latin american spanish phrase.
	merryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michifú"}

	// merryRussianPhrase represents merry russian phrase.
	merryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мяуи-и-и"}

	// merrySimplifiedChinesePhrase represents merry simplified chinese phrase.
	merrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喵嗯"}

	// merrySpanishPhrase represents merry spanish phrase.
	merrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "michifú"}

	// merryTraditionalChinesePhrase represents merry traditional chinese phrase.
	merryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喵嗯"}
)

var (
	// merryPhrase represents merry phrase.
	merryPhrase = nook.Languages{
		language.AmericanEnglish:      merryAmericanEnglishPhrase,
		language.CanadianFrench:       merryCanadianFrenchPhrase,
		language.Dutch:                merryDutchPhrase,
		language.French:               merryFrenchPhrase,
		language.German:               merryGermanPhrase,
		language.Italian:              merryItalianPhrase,
		language.Japanese:             merryJapanesePhrase,
		language.Korean:               merryKoreanPhrase,
		language.LatinAmericanSpanish: merryLatinAmericanSpanishPhrase,
		language.Russian:              merryRussianPhrase,
		language.SimplifiedChinese:    merrySimplifiedChinesePhrase,
		language.Spanish:              merrySpanishPhrase,
		language.TraditionalChinese:   merryTraditionalChinesePhrase}
)

var (
	// Merry represents merry.
	Merry = nook.Villager{
		Character:   merryCharacter,
		Personality: personality.Peppy,
		Phrase:      merryPhrase}
)
