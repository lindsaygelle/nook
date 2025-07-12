package duck

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
	// scootBirthday represents scoot birthday.
	scootBirthday = nook.Birthday{
		Day:   13,
		Month: time.June}
)

var (
	// scootCode represents scoot code.
	scootCode = nook.Code{
		Value: "duk10"}
)

var (
	// scootAmericanEnglishName represents scoot american english name.
	scootAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Scoot"}

	// scootCanadianFrenchName represents scoot canadian french name.
	scootCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Scooter"}

	// scootDutchName represents scoot dutch name.
	scootDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Scoot"}

	// scootFrenchName represents scoot french name.
	scootFrenchName = nook.Name{
		Language: language.French,
		Value:    "Scooter"}

	// scootGermanName represents scoot german name.
	scootGermanName = nook.Name{
		Language: language.German,
		Value:    "Helmut"}

	// scootItalianName represents scoot italian name.
	scootItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Guizzo"}

	// scootJapaneseName represents scoot japanese name.
	scootJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マモル"}

	// scootKoreanName represents scoot korean name.
	scootKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "지키미"}

	// scootLatinAmericanSpanishName represents scoot latin american spanish name.
	scootLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chema"}

	// scootRussianName represents scoot russian name.
	scootRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Скут"}

	// scootSimplifiedChineseName represents scoot simplified chinese name.
	scootSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿守"}

	// scootSpanishName represents scoot spanish name.
	scootSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chema"}

	// scootTraditionalChineseName represents scoot traditional chinese name.
	scootTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿守"}
)

var (
	// scootName represents scoot name.
	scootName = nook.Languages{
		language.AmericanEnglish:      scootAmericanEnglishName,
		language.CanadianFrench:       scootCanadianFrenchName,
		language.Dutch:                scootDutchName,
		language.French:               scootFrenchName,
		language.German:               scootGermanName,
		language.Italian:              scootItalianName,
		language.Japanese:             scootJapaneseName,
		language.Korean:               scootKoreanName,
		language.LatinAmericanSpanish: scootLatinAmericanSpanishName,
		language.Russian:              scootRussianName,
		language.SimplifiedChinese:    scootSimplifiedChineseName,
		language.Spanish:              scootSpanishName,
		language.TraditionalChinese:   scootTraditionalChineseName}
)

var (
	// scootCharacter represents scoot character.
	scootCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: scootBirthday,
		Code:     scootCode,
		Key:      character.Scoot,
		Gender:   gender.Male,
		Name:     scootName,
		Special:  false}
)

var (
	// scootAmericanEnglishPhrase represents scoot american english phrase.
	scootAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zip zoom"}

	// scootCanadianFrenchPhrase represents scoot canadian french phrase.
	scootCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zak zak"}

	// scootDutchPhrase represents scoot dutch phrase.
	scootDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "broem"}

	// scootFrenchPhrase represents scoot french phrase.
	scootFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zak zak"}

	// scootGermanPhrase represents scoot german phrase.
	scootGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flitzer"}

	// scootItalianPhrase represents scoot italian phrase.
	scootItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zip zoom"}

	// scootJapanesePhrase represents scoot japanese phrase.
	scootJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グワッ"}

	// scootKoreanPhrase represents scoot korean phrase.
	scootKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꾸왁"}

	// scootLatinAmericanSpanishPhrase represents scoot latin american spanish phrase.
	scootLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zapizum"}

	// scootRussianPhrase represents scoot russian phrase.
	scootRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брум-брум"}

	// scootSimplifiedChinesePhrase represents scoot simplified chinese phrase.
	scootSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呱呱"}

	// scootSpanishPhrase represents scoot spanish phrase.
	scootSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zapizum"}

	// scootTraditionalChinesePhrase represents scoot traditional chinese phrase.
	scootTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呱呱"}
)

var (
	// scootPhrase represents scoot phrase.
	scootPhrase = nook.Languages{
		language.AmericanEnglish:      scootAmericanEnglishPhrase,
		language.CanadianFrench:       scootCanadianFrenchPhrase,
		language.Dutch:                scootDutchPhrase,
		language.French:               scootFrenchPhrase,
		language.German:               scootGermanPhrase,
		language.Italian:              scootItalianPhrase,
		language.Japanese:             scootJapanesePhrase,
		language.Korean:               scootKoreanPhrase,
		language.LatinAmericanSpanish: scootLatinAmericanSpanishPhrase,
		language.Russian:              scootRussianPhrase,
		language.SimplifiedChinese:    scootSimplifiedChinesePhrase,
		language.Spanish:              scootSpanishPhrase,
		language.TraditionalChinese:   scootTraditionalChinesePhrase}
)

var (
	// Scoot represents scoot.
	Scoot = nook.Villager{
		Character:   scootCharacter,
		Personality: personality.Jock,
		Phrase:      scootPhrase}
)
