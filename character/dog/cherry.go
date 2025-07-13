package dog

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
	// cherryBirthday represents cherry birthday.
	cherryBirthday = nook.Birthday{
		Day:   11,
		Month: time.May}
)

var (
	// cherryCode represents cherry code.
	cherryCode = nook.Code{
		Value: "dog17"}
)

var (
	// cherryAmericanEnglishName represents cherry american english name.
	cherryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cherry"}

	// cherryCanadianFrenchName represents cherry canadian french name.
	cherryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anna"}

	// cherryDutchName represents cherry dutch name.
	cherryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cherry"}

	// cherryFrenchName represents cherry french name.
	cherryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anna"}

	// cherryGermanName represents cherry german name.
	cherryGermanName = nook.Name{
		Language: language.German,
		Value:    "Bella"}

	// cherryItalianName represents cherry italian name.
	cherryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amarena"}

	// cherryJapaneseName represents cherry japanese name.
	cherryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハンナ"}

	// cherryKoreanName represents cherry korean name.
	cherryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "한나"}

	// cherryLatinAmericanSpanishName represents cherry latin american spanish name.
	cherryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Luna"}

	// cherryRussianName represents cherry russian name.
	cherryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Черри"}

	// cherrySimplifiedChineseName represents cherry simplified chinese name.
	cherrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "花娜"}

	// cherrySpanishName represents cherry spanish name.
	cherrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Luna"}

	// cherryTraditionalChineseName represents cherry traditional chinese name.
	cherryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "花娜"}
)

var (
	// cherryName represents cherry name.
	cherryName = nook.Languages{
		language.AmericanEnglish:      cherryAmericanEnglishName,
		language.CanadianFrench:       cherryCanadianFrenchName,
		language.Dutch:                cherryDutchName,
		language.French:               cherryFrenchName,
		language.German:               cherryGermanName,
		language.Italian:              cherryItalianName,
		language.Japanese:             cherryJapaneseName,
		language.Korean:               cherryKoreanName,
		language.LatinAmericanSpanish: cherryLatinAmericanSpanishName,
		language.Russian:              cherryRussianName,
		language.SimplifiedChinese:    cherrySimplifiedChineseName,
		language.Spanish:              cherrySpanishName,
		language.TraditionalChinese:   cherryTraditionalChineseName}
)

var (
	// cherryCharacter represents cherry character.
	cherryCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: cherryBirthday,
		Code:     cherryCode,
		Key:      character.Cherry,
		Gender:   gender.Female,
		Name:     cherryName,
		Special:  false}
)

var (
	// cherryAmericanEnglishPhrase represents cherry american english phrase.
	cherryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "what what"}

	// cherryCanadianFrenchPhrase represents cherry canadian french phrase.
	cherryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "niniche"}

	// cherryDutchPhrase represents cherry dutch phrase.
	cherryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wat jij"}

	// cherryFrenchPhrase represents cherry french phrase.
	cherryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "niniche"}

	// cherryGermanPhrase represents cherry german phrase.
	cherryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wedel"}

	// cherryItalianPhrase represents cherry italian phrase.
	cherryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arf arf"}

	// cherryJapanesePhrase represents cherry japanese phrase.
	cherryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だってネ"}

	// cherryKoreanPhrase represents cherry korean phrase.
	cherryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그냥저냥"}

	// cherryLatinAmericanSpanishPhrase represents cherry latin american spanish phrase.
	cherryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guauchi"}

	// cherryRussianPhrase represents cherry russian phrase.
	cherryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "что-что"}

	// cherrySimplifiedChinesePhrase represents cherry simplified chinese phrase.
	cherrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "听说娜"}

	// cherrySpanishPhrase represents cherry spanish phrase.
	cherrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guauchi"}

	// cherryTraditionalChinesePhrase represents cherry traditional chinese phrase.
	cherryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聽說娜"}
)

var (
	// cherryPhrase represents cherry phrase.
	cherryPhrase = nook.Languages{
		language.AmericanEnglish:      cherryAmericanEnglishPhrase,
		language.CanadianFrench:       cherryCanadianFrenchPhrase,
		language.Dutch:                cherryDutchPhrase,
		language.French:               cherryFrenchPhrase,
		language.German:               cherryGermanPhrase,
		language.Italian:              cherryItalianPhrase,
		language.Japanese:             cherryJapanesePhrase,
		language.Korean:               cherryKoreanPhrase,
		language.LatinAmericanSpanish: cherryLatinAmericanSpanishPhrase,
		language.Russian:              cherryRussianPhrase,
		language.SimplifiedChinese:    cherrySimplifiedChinesePhrase,
		language.Spanish:              cherrySpanishPhrase,
		language.TraditionalChinese:   cherryTraditionalChinesePhrase}
)

var (
	// Cherry represents cherry.
	Cherry = nook.Villager{
		Character:   cherryCharacter,
		Personality: personality.BigSister,
		Phrase:      cherryPhrase}
)
