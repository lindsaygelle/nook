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
	// vladimirBirthday represents vladimir birthday.
	vladimirBirthday = nook.Birthday{
		Day:   2,
		Month: time.August}
)

var (
	// vladimirCode represents vladimir code.
	vladimirCode = nook.Code{
		Value: "cbr06"}
)

var (
	// vladimirAmericanEnglishName represents vladimir american english name.
	vladimirAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vladimir"}

	// vladimirCanadianFrenchName represents vladimir canadian french name.
	vladimirCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vladimir"}

	// vladimirDutchName represents vladimir dutch name.
	vladimirDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vladimir"}

	// vladimirFrenchName represents vladimir french name.
	vladimirFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vladimir"}

	// vladimirGermanName represents vladimir german name.
	vladimirGermanName = nook.Name{
		Language: language.German,
		Value:    "Vladimir"}

	// vladimirItalianName represents vladimir italian name.
	vladimirItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vladimir"}

	// vladimirJapaneseName represents vladimir japanese name.
	vladimirJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガビ"}

	// vladimirKoreanName represents vladimir korean name.
	vladimirKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰비"}

	// vladimirLatinAmericanSpanishName represents vladimir latin american spanish name.
	vladimirLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vladimir"}

	// vladimirRussianName represents vladimir russian name.
	vladimirRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Владимир"}

	// vladimirSimplifiedChineseName represents vladimir simplified chinese name.
	vladimirSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉弼"}

	// vladimirSpanishName represents vladimir spanish name.
	vladimirSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vladimir"}

	// vladimirTraditionalChineseName represents vladimir traditional chinese name.
	vladimirTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉弼"}
)

var (
	// vladimirName represents vladimir name.
	vladimirName = nook.Languages{
		language.AmericanEnglish:      vladimirAmericanEnglishName,
		language.CanadianFrench:       vladimirCanadianFrenchName,
		language.Dutch:                vladimirDutchName,
		language.French:               vladimirFrenchName,
		language.German:               vladimirGermanName,
		language.Italian:              vladimirItalianName,
		language.Japanese:             vladimirJapaneseName,
		language.Korean:               vladimirKoreanName,
		language.LatinAmericanSpanish: vladimirLatinAmericanSpanishName,
		language.Russian:              vladimirRussianName,
		language.SimplifiedChinese:    vladimirSimplifiedChineseName,
		language.Spanish:              vladimirSpanishName,
		language.TraditionalChinese:   vladimirTraditionalChineseName}
)

var (
	// vladimirCharacter represents vladimir character.
	vladimirCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: vladimirBirthday,
		Code:     vladimirCode,
		Key:      character.Vladimir,
		Gender:   gender.Male,
		Name:     vladimirName,
		Special:  false}
)

var (
	// vladimirAmericanEnglishPhrase represents vladimir american english phrase.
	vladimirAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nyet"}

	// vladimirCanadianFrenchPhrase represents vladimir canadian french phrase.
	vladimirCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "niet"}

	// vladimirDutchPhrase represents vladimir dutch phrase.
	vladimirDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "njet"}

	// vladimirFrenchPhrase represents vladimir french phrase.
	vladimirFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "niet"}

	// vladimirGermanPhrase represents vladimir german phrase.
	vladimirGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nastrovje"}

	// vladimirItalianPhrase represents vladimir italian phrase.
	vladimirItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nyet"}

	// vladimirJapanesePhrase represents vladimir japanese phrase.
	vladimirJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やんけ"}

	// vladimirKoreanPhrase represents vladimir korean phrase.
	vladimirKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "참말로"}

	// vladimirLatinAmericanSpanishPhrase represents vladimir latin american spanish phrase.
	vladimirLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nyet"}

	// vladimirRussianPhrase represents vladimir russian phrase.
	vladimirRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ну не-е-ет"}

	// vladimirSimplifiedChinesePhrase represents vladimir simplified chinese phrase.
	vladimirSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唉唷"}

	// vladimirSpanishPhrase represents vladimir spanish phrase.
	vladimirSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nyet"}

	// vladimirTraditionalChinesePhrase represents vladimir traditional chinese phrase.
	vladimirTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唉唷"}
)

var (
	// vladimirPhrase represents vladimir phrase.
	vladimirPhrase = nook.Languages{
		language.AmericanEnglish:      vladimirAmericanEnglishPhrase,
		language.CanadianFrench:       vladimirCanadianFrenchPhrase,
		language.Dutch:                vladimirDutchPhrase,
		language.French:               vladimirFrenchPhrase,
		language.German:               vladimirGermanPhrase,
		language.Italian:              vladimirItalianPhrase,
		language.Japanese:             vladimirJapanesePhrase,
		language.Korean:               vladimirKoreanPhrase,
		language.LatinAmericanSpanish: vladimirLatinAmericanSpanishPhrase,
		language.Russian:              vladimirRussianPhrase,
		language.SimplifiedChinese:    vladimirSimplifiedChinesePhrase,
		language.Spanish:              vladimirSpanishPhrase,
		language.TraditionalChinese:   vladimirTraditionalChinesePhrase}
)

var (
	// Vladimir represents vladimir.
	Vladimir = nook.Villager{
		Character:   vladimirCharacter,
		Personality: personality.Cranky,
		Phrase:      vladimirPhrase}
)
