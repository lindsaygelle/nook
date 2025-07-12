package goat

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
	// billyBirthday represents billy birthday.
	billyBirthday = nook.Birthday{
		Day:   25,
		Month: time.March}
)

var (
	// billyCode represents billy code.
	billyCode = nook.Code{
		Value: "goa02"}
)

var (
	// billyAmericanEnglishName represents billy american english name.
	billyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Billy"}

	// billyCanadianFrenchName represents billy canadian french name.
	billyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Seguin"}

	// billyDutchName represents billy dutch name.
	billyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Billy"}

	// billyFrenchName represents billy french name.
	billyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Seguin"}

	// billyGermanName represents billy german name.
	billyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hennes"}

	// billyItalianName represents billy italian name.
	billyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Billy"}

	// billyJapaneseName represents billy japanese name.
	billyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アーシンド"}

	// billyKoreanName represents billy korean name.
	billyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "힘드러"}

	// billyLatinAmericanSpanishName represents billy latin american spanish name.
	billyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brito"}

	// billyRussianName represents billy russian name.
	billyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Билли"}

	// billySimplifiedChineseName represents billy simplified chinese name.
	billySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚星"}

	// billySpanishName represents billy spanish name.
	billySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brito"}

	// billyTraditionalChineseName represents billy traditional chinese name.
	billyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞星"}
)

var (
	// billyName represents billy name.
	billyName = nook.Languages{
		language.AmericanEnglish:      billyAmericanEnglishName,
		language.CanadianFrench:       billyCanadianFrenchName,
		language.Dutch:                billyDutchName,
		language.French:               billyFrenchName,
		language.German:               billyGermanName,
		language.Italian:              billyItalianName,
		language.Japanese:             billyJapaneseName,
		language.Korean:               billyKoreanName,
		language.LatinAmericanSpanish: billyLatinAmericanSpanishName,
		language.Russian:              billyRussianName,
		language.SimplifiedChinese:    billySimplifiedChineseName,
		language.Spanish:              billySpanishName,
		language.TraditionalChinese:   billyTraditionalChineseName}
)

var (
	// billyCharacter represents billy character.
	billyCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: billyBirthday,
		Code:     billyCode,
		Key:      character.Billy,
		Gender:   gender.Male,
		Name:     billyName,
		Special:  false}
)

var (
	// billyAmericanEnglishPhrase represents billy american english phrase.
	billyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dagnaabit"}

	// billyCanadianFrenchPhrase represents billy canadian french phrase.
	billyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bêê quoi"}

	// billyDutchPhrase represents billy dutch phrase.
	billyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nôh"}

	// billyFrenchPhrase represents billy french phrase.
	billyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "eh bââlèze"}

	// billyGermanPhrase represents billy german phrase.
	billyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "o-o-ja-a-a"}

	// billyItalianPhrase represents billy italian phrase.
	billyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sorbeeelli"}

	// billyJapanesePhrase represents billy japanese phrase.
	billyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よのう"}

	// billyKoreanPhrase represents billy korean phrase.
	billyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알아줘"}

	// billyLatinAmericanSpanishPhrase represents billy latin american spanish phrase.
	billyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "beerree"}

	// billyRussianPhrase represents billy russian phrase.
	billyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "такие-сякие"}

	// billySimplifiedChinesePhrase represents billy simplified chinese phrase.
	billySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷喏"}

	// billySpanishPhrase represents billy spanish phrase.
	billySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fenómeno"}

	// billyTraditionalChinesePhrase represents billy traditional chinese phrase.
	billyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷喏"}
)

var (
	// billyPhrase represents billy phrase.
	billyPhrase = nook.Languages{
		language.AmericanEnglish:      billyAmericanEnglishPhrase,
		language.CanadianFrench:       billyCanadianFrenchPhrase,
		language.Dutch:                billyDutchPhrase,
		language.French:               billyFrenchPhrase,
		language.German:               billyGermanPhrase,
		language.Italian:              billyItalianPhrase,
		language.Japanese:             billyJapanesePhrase,
		language.Korean:               billyKoreanPhrase,
		language.LatinAmericanSpanish: billyLatinAmericanSpanishPhrase,
		language.Russian:              billyRussianPhrase,
		language.SimplifiedChinese:    billySimplifiedChinesePhrase,
		language.Spanish:              billySpanishPhrase,
		language.TraditionalChinese:   billyTraditionalChinesePhrase}
)

var (
	// Billy represents billy.
	Billy = nook.Villager{
		Character:   billyCharacter,
		Personality: personality.Jock,
		Phrase:      billyPhrase}
)
