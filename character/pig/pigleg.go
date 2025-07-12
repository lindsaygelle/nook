package pig

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
	// piglegBirthday represents pigleg birthday.
	piglegBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// piglegCode represents pigleg code.
	piglegCode = nook.Code{
		Value: ""}
)

var (
	// piglegAmericanEnglishName represents pigleg american english name.
	piglegAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pigleg"}

	// piglegCanadianFrenchName represents pigleg canadian french name.
	piglegCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Xavier"}

	// piglegDutchName represents pigleg dutch name.
	piglegDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// piglegFrenchName represents pigleg french name.
	piglegFrenchName = nook.Name{
		Language: language.French,
		Value:    "Xavier"}

	// piglegGermanName represents pigleg german name.
	piglegGermanName = nook.Name{
		Language: language.German,
		Value:    "Jonas"}

	// piglegItalianName represents pigleg italian name.
	piglegItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drake"}

	// piglegJapaneseName represents pigleg japanese name.
	piglegJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バチコーメ"}

	// piglegKoreanName represents pigleg korean name.
	piglegKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// piglegLatinAmericanSpanishName represents pigleg latin american spanish name.
	piglegLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jamoncio"}

	// piglegRussianName represents pigleg russian name.
	piglegRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// piglegSimplifiedChineseName represents pigleg simplified chinese name.
	piglegSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// piglegSpanishName represents pigleg spanish name.
	piglegSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jamoncio"}

	// piglegTraditionalChineseName represents pigleg traditional chinese name.
	piglegTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// piglegName represents pigleg name.
	piglegName = nook.Languages{
		language.AmericanEnglish:      piglegAmericanEnglishName,
		language.CanadianFrench:       piglegCanadianFrenchName,
		language.Dutch:                piglegDutchName,
		language.French:               piglegFrenchName,
		language.German:               piglegGermanName,
		language.Italian:              piglegItalianName,
		language.Japanese:             piglegJapaneseName,
		language.Korean:               piglegKoreanName,
		language.LatinAmericanSpanish: piglegLatinAmericanSpanishName,
		language.Russian:              piglegRussianName,
		language.SimplifiedChinese:    piglegSimplifiedChineseName,
		language.Spanish:              piglegSpanishName,
		language.TraditionalChinese:   piglegTraditionalChineseName}
)

var (
	// piglegCharacter represents pigleg character.
	piglegCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: piglegBirthday,
		Code:     piglegCode,
		Key:      character.Pigleg,
		Gender:   gender.Male,
		Name:     piglegName,
		Special:  false}
)

var (
	// piglegAmericanEnglishPhrase represents pigleg american english phrase.
	piglegAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "arrrn"}

	// piglegCanadianFrenchPhrase represents pigleg canadian french phrase.
	piglegCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	// piglegDutchPhrase represents pigleg dutch phrase.
	piglegDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// piglegFrenchPhrase represents pigleg french phrase.
	piglegFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	// piglegGermanPhrase represents pigleg german phrase.
	piglegGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "Unknown"}

	// piglegItalianPhrase represents pigleg italian phrase.
	piglegItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grugno"}

	// piglegJapanesePhrase represents pigleg japanese phrase.
	piglegJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨーソロ"}

	// piglegKoreanPhrase represents pigleg korean phrase.
	piglegKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// piglegLatinAmericanSpanishPhrase represents pigleg latin american spanish phrase.
	piglegLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	// piglegRussianPhrase represents pigleg russian phrase.
	piglegRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// piglegSimplifiedChinesePhrase represents pigleg simplified chinese phrase.
	piglegSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// piglegSpanishPhrase represents pigleg spanish phrase.
	piglegSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	// piglegTraditionalChinesePhrase represents pigleg traditional chinese phrase.
	piglegTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// piglegPhrase represents pigleg phrase.
	piglegPhrase = nook.Languages{
		language.AmericanEnglish:      piglegAmericanEnglishPhrase,
		language.CanadianFrench:       piglegCanadianFrenchPhrase,
		language.Dutch:                piglegDutchPhrase,
		language.French:               piglegFrenchPhrase,
		language.German:               piglegGermanPhrase,
		language.Italian:              piglegItalianPhrase,
		language.Japanese:             piglegJapanesePhrase,
		language.Korean:               piglegKoreanPhrase,
		language.LatinAmericanSpanish: piglegLatinAmericanSpanishPhrase,
		language.Russian:              piglegRussianPhrase,
		language.SimplifiedChinese:    piglegSimplifiedChinesePhrase,
		language.Spanish:              piglegSpanishPhrase,
		language.TraditionalChinese:   piglegTraditionalChinesePhrase}
)

var (
	// Pigleg represents pigleg.
	Pigleg = nook.Villager{
		Character:   piglegCharacter,
		Personality: personality.Jock,
		Phrase:      piglegPhrase}
)
