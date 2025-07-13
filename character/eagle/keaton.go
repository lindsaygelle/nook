package eagle

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
	// keatonBirthday represents keaton birthday.
	keatonBirthday = nook.Birthday{
		Day:   1,
		Month: time.June}
)

var (
	// keatonCode represents keaton code.
	keatonCode = nook.Code{
		Value: "pbr08"}
)

var (
	// keatonAmericanEnglishName represents keaton american english name.
	keatonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Keaton"}

	// keatonCanadianFrenchName represents keaton canadian french name.
	keatonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Enzo"}

	// keatonDutchName represents keaton dutch name.
	keatonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Keaton"}

	// keatonFrenchName represents keaton french name.
	keatonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Enzo"}

	// keatonGermanName represents keaton german name.
	keatonGermanName = nook.Name{
		Language: language.German,
		Value:    "Kai"}

	// keatonItalianName represents keaton italian name.
	keatonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aquilio"}

	// keatonJapaneseName represents keaton japanese name.
	keatonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランク"}

	// keatonKoreanName represents keaton korean name.
	keatonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랭크"}

	// keatonLatinAmericanSpanishName represents keaton latin american spanish name.
	keatonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lucho"}

	// keatonRussianName represents keaton russian name.
	keatonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китон"}

	// keatonSimplifiedChineseName represents keaton simplified chinese name.
	keatonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "法兰克"}

	// keatonSpanishName represents keaton spanish name.
	keatonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lucho"}

	// keatonTraditionalChineseName represents keaton traditional chinese name.
	keatonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "法蘭克"}
)

var (
	// keatonName represents keaton name.
	keatonName = nook.Languages{
		language.AmericanEnglish:      keatonAmericanEnglishName,
		language.CanadianFrench:       keatonCanadianFrenchName,
		language.Dutch:                keatonDutchName,
		language.French:               keatonFrenchName,
		language.German:               keatonGermanName,
		language.Italian:              keatonItalianName,
		language.Japanese:             keatonJapaneseName,
		language.Korean:               keatonKoreanName,
		language.LatinAmericanSpanish: keatonLatinAmericanSpanishName,
		language.Russian:              keatonRussianName,
		language.SimplifiedChinese:    keatonSimplifiedChineseName,
		language.Spanish:              keatonSpanishName,
		language.TraditionalChinese:   keatonTraditionalChineseName}
)

var (
	// keatonCharacter represents keaton character.
	keatonCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: keatonBirthday,
		Code:     keatonCode,
		Key:      character.Keaton,
		Gender:   gender.Male,
		Name:     keatonName,
		Special:  false}
)

var (
	// keatonAmericanEnglishPhrase represents keaton american english phrase.
	keatonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wingo"}

	// keatonCanadianFrenchPhrase represents keaton canadian french phrase.
	keatonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "royaaaal"}

	// keatonDutchPhrase represents keaton dutch phrase.
	keatonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vleugels"}

	// keatonFrenchPhrase represents keaton french phrase.
	keatonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "royaaaal"}

	// keatonGermanPhrase represents keaton german phrase.
	keatonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kiiiiiih"}

	// keatonItalianPhrase represents keaton italian phrase.
	keatonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baila"}

	// keatonJapanesePhrase represents keaton japanese phrase.
	keatonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でーッス"}

	// keatonKoreanPhrase represents keaton korean phrase.
	keatonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "수리수리"}

	// keatonLatinAmericanSpanishPhrase represents keaton latin american spanish phrase.
	keatonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flopflop"}

	// keatonRussianPhrase represents keaton russian phrase.
	keatonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крыло"}

	// keatonSimplifiedChinesePhrase represents keaton simplified chinese phrase.
	keatonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是"}

	// keatonSpanishPhrase represents keaton spanish phrase.
	keatonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flopflop"}

	// keatonTraditionalChinesePhrase represents keaton traditional chinese phrase.
	keatonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是"}
)

var (
	// keatonPhrase represents keaton phrase.
	keatonPhrase = nook.Languages{
		language.AmericanEnglish:      keatonAmericanEnglishPhrase,
		language.CanadianFrench:       keatonCanadianFrenchPhrase,
		language.Dutch:                keatonDutchPhrase,
		language.French:               keatonFrenchPhrase,
		language.German:               keatonGermanPhrase,
		language.Italian:              keatonItalianPhrase,
		language.Japanese:             keatonJapanesePhrase,
		language.Korean:               keatonKoreanPhrase,
		language.LatinAmericanSpanish: keatonLatinAmericanSpanishPhrase,
		language.Russian:              keatonRussianPhrase,
		language.SimplifiedChinese:    keatonSimplifiedChinesePhrase,
		language.Spanish:              keatonSpanishPhrase,
		language.TraditionalChinese:   keatonTraditionalChinesePhrase}
)

var (
	// Keaton represents keaton.
	Keaton = nook.Villager{
		Character:   keatonCharacter,
		Personality: personality.Smug,
		Phrase:      keatonPhrase}
)
