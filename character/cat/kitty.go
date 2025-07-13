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
	// kittyBirthday represents kitty birthday.
	kittyBirthday = nook.Birthday{
		Day:   15,
		Month: time.February}
)

var (
	// kittyCode represents kitty code.
	kittyCode = nook.Code{
		Value: "cat14"}
)

var (
	// kittyAmericanEnglishName represents kitty american english name.
	kittyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kitty"}

	// kittyCanadianFrenchName represents kitty canadian french name.
	kittyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kitty"}

	// kittyDutchName represents kitty dutch name.
	kittyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kitty"}

	// kittyFrenchName represents kitty french name.
	kittyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kitty"}

	// kittyGermanName represents kitty german name.
	kittyGermanName = nook.Name{
		Language: language.German,
		Value:    "Karen"}

	// kittyItalianName represents kitty italian name.
	kittyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ester"}

	// kittyJapaneseName represents kitty japanese name.
	kittyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ショコラ"}

	// kittyKoreanName represents kitty korean name.
	kittyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쇼콜라"}

	// kittyLatinAmericanSpanishName represents kitty latin american spanish name.
	kittyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kasandra"}

	// kittyRussianName represents kitty russian name.
	kittyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китти"}

	// kittySimplifiedChineseName represents kitty simplified chinese name.
	kittySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱古莉"}

	// kittySpanishName represents kitty spanish name.
	kittySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kasandra"}

	// kittyTraditionalChineseName represents kitty traditional chinese name.
	kittyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱古莉"}
)

var (
	// kittyName represents kitty name.
	kittyName = nook.Languages{
		language.AmericanEnglish:      kittyAmericanEnglishName,
		language.CanadianFrench:       kittyCanadianFrenchName,
		language.Dutch:                kittyDutchName,
		language.French:               kittyFrenchName,
		language.German:               kittyGermanName,
		language.Italian:              kittyItalianName,
		language.Japanese:             kittyJapaneseName,
		language.Korean:               kittyKoreanName,
		language.LatinAmericanSpanish: kittyLatinAmericanSpanishName,
		language.Russian:              kittyRussianName,
		language.SimplifiedChinese:    kittySimplifiedChineseName,
		language.Spanish:              kittySpanishName,
		language.TraditionalChinese:   kittyTraditionalChineseName}
)

var (
	// kittyCharacter represents kitty character.
	kittyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kittyBirthday,
		Code:     kittyCode,
		Key:      character.Kitty,
		Gender:   gender.Female,
		Name:     kittyName,
		Special:  false}
)

var (
	// kittyAmericanEnglishPhrase represents kitty american english phrase.
	kittyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mrowrr"}

	// kittyCanadianFrenchPhrase represents kitty canadian french phrase.
	kittyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaaaouh"}

	// kittyDutchPhrase represents kitty dutch phrase.
	kittyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "giechel"}

	// kittyFrenchPhrase represents kitty french phrase.
	kittyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaaaouh"}

	// kittyGermanPhrase represents kitty german phrase.
	kittyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miezmiez"}

	// kittyItalianPhrase represents kitty italian phrase.
	kittyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fusolo"}

	// kittyJapanesePhrase represents kitty japanese phrase.
	kittyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フフ"}

	// kittyKoreanPhrase represents kitty korean phrase.
	kittyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "자갸"}

	// kittyLatinAmericanSpanishPhrase represents kitty latin american spanish phrase.
	kittyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miaul"}

	// kittyRussianPhrase represents kitty russian phrase.
	kittyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мр-мр-мр"}

	// kittySimplifiedChinesePhrase represents kitty simplified chinese phrase.
	kittySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古古"}

	// kittySpanishPhrase represents kitty spanish phrase.
	kittySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cascabel"}

	// kittyTraditionalChinesePhrase represents kitty traditional chinese phrase.
	kittyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古古"}
)

var (
	// kittyPhrase represents kitty phrase.
	kittyPhrase = nook.Languages{
		language.AmericanEnglish:      kittyAmericanEnglishPhrase,
		language.CanadianFrench:       kittyCanadianFrenchPhrase,
		language.Dutch:                kittyDutchPhrase,
		language.French:               kittyFrenchPhrase,
		language.German:               kittyGermanPhrase,
		language.Italian:              kittyItalianPhrase,
		language.Japanese:             kittyJapanesePhrase,
		language.Korean:               kittyKoreanPhrase,
		language.LatinAmericanSpanish: kittyLatinAmericanSpanishPhrase,
		language.Russian:              kittyRussianPhrase,
		language.SimplifiedChinese:    kittySimplifiedChinesePhrase,
		language.Spanish:              kittySpanishPhrase,
		language.TraditionalChinese:   kittyTraditionalChinesePhrase}
)

var (
	// Kitty represents kitty.
	Kitty = nook.Villager{
		Character:   kittyCharacter,
		Personality: personality.Snooty,
		Phrase:      kittyPhrase}
)
