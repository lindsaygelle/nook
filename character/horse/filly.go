package horse

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
	// fillyBirthday represents filly birthday.
	fillyBirthday = nook.Birthday{
		Day:   11,
		Month: time.July}
)

var (
	// fillyCode represents filly code.
	fillyCode = nook.Code{
		Value: "hrs14"}
)

var (
	// fillyAmericanEnglishName represents filly american english name.
	fillyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Filly"}

	// fillyCanadianFrenchName represents filly canadian french name.
	fillyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Palomina"}

	// fillyDutchName represents filly dutch name.
	fillyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// fillyFrenchName represents filly french name.
	fillyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Palomina"}

	// fillyGermanName represents filly german name.
	fillyGermanName = nook.Name{
		Language: language.German,
		Value:    "Beate"}

	// fillyItalianName represents filly italian name.
	fillyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Filippa"}

	// fillyJapaneseName represents filly japanese name.
	fillyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "7ごう"}

	// fillyKoreanName represents filly korean name.
	fillyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "7호"}

	// fillyLatinAmericanSpanishName represents filly latin american spanish name.
	fillyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Séptima"}

	// fillyRussianName represents filly russian name.
	fillyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// fillySimplifiedChineseName represents filly simplified chinese name.
	fillySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// fillySpanishName represents filly spanish name.
	fillySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Séptima"}

	// fillyTraditionalChineseName represents filly traditional chinese name.
	fillyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// fillyName represents filly name.
	fillyName = nook.Languages{
		language.AmericanEnglish:      fillyAmericanEnglishName,
		language.CanadianFrench:       fillyCanadianFrenchName,
		language.Dutch:                fillyDutchName,
		language.French:               fillyFrenchName,
		language.German:               fillyGermanName,
		language.Italian:              fillyItalianName,
		language.Japanese:             fillyJapaneseName,
		language.Korean:               fillyKoreanName,
		language.LatinAmericanSpanish: fillyLatinAmericanSpanishName,
		language.Russian:              fillyRussianName,
		language.SimplifiedChinese:    fillySimplifiedChineseName,
		language.Spanish:              fillySpanishName,
		language.TraditionalChinese:   fillyTraditionalChineseName}
)

var (
	// fillyCharacter represents filly character.
	fillyCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: fillyBirthday,
		Code:     fillyCode,
		Key:      character.Filly,
		Gender:   gender.Female,
		Name:     fillyName,
		Special:  false}
)

var (
	// fillyAmericanEnglishPhrase represents filly american english phrase.
	fillyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hah"}

	// fillyCanadianFrenchPhrase represents filly canadian french phrase.
	fillyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "papaye"}

	// fillyDutchPhrase represents filly dutch phrase.
	fillyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// fillyFrenchPhrase represents filly french phrase.
	fillyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "papaye"}

	// fillyGermanPhrase represents filly german phrase.
	fillyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hühott"}

	// fillyItalianPhrase represents filly italian phrase.
	fillyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "galop"}

	// fillyJapanesePhrase represents filly japanese phrase.
	fillyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "はぁっ！"}

	// fillyKoreanPhrase represents filly korean phrase.
	fillyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "타앗"}

	// fillyLatinAmericanSpanishPhrase represents filly latin american spanish phrase.
	fillyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "galope"}

	// fillyRussianPhrase represents filly russian phrase.
	fillyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// fillySimplifiedChinesePhrase represents filly simplified chinese phrase.
	fillySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// fillySpanishPhrase represents filly spanish phrase.
	fillySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galope"}

	// fillyTraditionalChinesePhrase represents filly traditional chinese phrase.
	fillyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// fillyPhrase represents filly phrase.
	fillyPhrase = nook.Languages{
		language.AmericanEnglish:      fillyAmericanEnglishPhrase,
		language.CanadianFrench:       fillyCanadianFrenchPhrase,
		language.Dutch:                fillyDutchPhrase,
		language.French:               fillyFrenchPhrase,
		language.German:               fillyGermanPhrase,
		language.Italian:              fillyItalianPhrase,
		language.Japanese:             fillyJapanesePhrase,
		language.Korean:               fillyKoreanPhrase,
		language.LatinAmericanSpanish: fillyLatinAmericanSpanishPhrase,
		language.Russian:              fillyRussianPhrase,
		language.SimplifiedChinese:    fillySimplifiedChinesePhrase,
		language.Spanish:              fillySpanishPhrase,
		language.TraditionalChinese:   fillyTraditionalChinesePhrase}
)

var (
	// Filly represents filly.
	Filly = nook.Villager{
		Character:   fillyCharacter,
		Personality: personality.Normal,
		Phrase:      fillyPhrase}
)
