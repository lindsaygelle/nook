package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// cyranoBirthday represents Cyrano's birthday.
var (
	// cyranoBirthday represents cyrano birthday.
	cyranoBirthday = nook.Birthday{
		Day:   9,
		Month: time.March}
)

// cyranoCode represents Cyrano's unique code.
var (
	// cyranoCode represents cyrano code.
	cyranoCode = nook.Code{
		Value: "ant00"}
)

// Different names for Cyrano in various languages.
var (
	// cyranoAmericanEnglishName represents cyrano american english name.
	cyranoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cyrano"}

	// cyranoCanadianFrenchName represents cyrano canadian french name.
	cyranoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cyrano"}

	// cyranoDutchName represents cyrano dutch name.
	cyranoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyrano"}

	// cyranoFrenchName represents cyrano french name.
	cyranoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cyrano"}

	// cyranoGermanName represents cyrano german name.
	cyranoGermanName = nook.Name{
		Language: language.German,
		Value:    "Theo"}

	// cyranoItalianName represents cyrano italian name.
	cyranoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cirano"}

	// cyranoJapaneseName represents cyrano japanese name.
	cyranoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さくらじま"}

	// cyranoKoreanName represents cyrano korean name.
	cyranoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사지마"}

	// cyranoLatinAmericanSpanishName represents cyrano latin american spanish name.
	cyranoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cirano"}

	// cyranoRussianName represents cyrano russian name.
	cyranoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сирано"}

	// cyranoSimplifiedChineseName represents cyrano simplified chinese name.
	cyranoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阳明"}

	// cyranoSpanishName represents cyrano spanish name.
	cyranoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cirano"}

	// cyranoTraditionalChineseName represents cyrano traditional chinese name.
	cyranoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陽明"}
)

// cyranoName represents Cyrano's name in different languages.
var (
	// cyranoName represents cyrano name.
	cyranoName = nook.Languages{
		language.AmericanEnglish:      cyranoAmericanEnglishName,
		language.CanadianFrench:       cyranoCanadianFrenchName,
		language.Dutch:                cyranoDutchName,
		language.French:               cyranoFrenchName,
		language.German:               cyranoGermanName,
		language.Italian:              cyranoItalianName,
		language.Japanese:             cyranoJapaneseName,
		language.Korean:               cyranoKoreanName,
		language.LatinAmericanSpanish: cyranoLatinAmericanSpanishName,
		language.Russian:              cyranoRussianName,
		language.SimplifiedChinese:    cyranoSimplifiedChineseName,
		language.Spanish:              cyranoSpanishName,
		language.TraditionalChinese:   cyranoTraditionalChineseName}
)

// cyranoCharacter represents Cyrano's character information.
var (
	// cyranoCharacter represents cyrano character.
	cyranoCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: cyranoBirthday,
		Code:     cyranoCode,
		Key:      character.Cyrano,
		Gender:   gender.Male,
		Name:     cyranoName,
		Special:  false}
)

// Different phrases spoken by Cyrano in various languages.
var (
	// cyranoAmericanEnglishPhrase represents cyrano american english phrase.
	cyranoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ah-CHOO"}

	// cyranoCanadianFrenchPhrase represents cyrano canadian french phrase.
	cyranoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ATCHOUM"}

	// cyranoDutchPhrase represents cyrano dutch phrase.
	cyranoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ha-TSJOE"}

	// cyranoFrenchPhrase represents cyrano french phrase.
	cyranoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ATCHOUM"}

	// cyranoGermanPhrase represents cyrano german phrase.
	cyranoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schneuf"}

	// cyranoItalianPhrase represents cyrano italian phrase.
	cyranoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ett-CCIÙ"}

	// cyranoJapanesePhrase represents cyrano japanese phrase.
	cyranoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でごわす"}

	// cyranoKoreanPhrase represents cyrano korean phrase.
	cyranoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "임돠"}

	// cyranoLatinAmericanSpanishPhrase represents cyrano latin american spanish phrase.
	cyranoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "achús"}

	// cyranoRussianPhrase represents cyrano russian phrase.
	cyranoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "апчхи"}

	// cyranoSimplifiedChinesePhrase represents cyrano simplified chinese phrase.
	cyranoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有的"}

	// cyranoSpanishPhrase represents cyrano spanish phrase.
	cyranoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "achús"}

	// cyranoTraditionalChinesePhrase represents cyrano traditional chinese phrase.
	cyranoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有的"}
)

// cyranoPhrase represents Cyrano's phrases in different languages.
var (
	// cyranoPhrase represents cyrano phrase.
	cyranoPhrase = nook.Languages{
		language.AmericanEnglish:      cyranoAmericanEnglishPhrase,
		language.CanadianFrench:       cyranoCanadianFrenchPhrase,
		language.Dutch:                cyranoDutchPhrase,
		language.French:               cyranoFrenchPhrase,
		language.German:               cyranoGermanPhrase,
		language.Italian:              cyranoItalianPhrase,
		language.Japanese:             cyranoJapanesePhrase,
		language.Korean:               cyranoKoreanPhrase,
		language.LatinAmericanSpanish: cyranoLatinAmericanSpanishPhrase,
		language.Russian:              cyranoRussianPhrase,
		language.SimplifiedChinese:    cyranoSimplifiedChinesePhrase,
		language.Spanish:              cyranoSpanishPhrase,
		language.TraditionalChinese:   cyranoTraditionalChinesePhrase}
)

// Cyrano represents the character Cyrano with his complete information.
var (
	// Cyrano represents cyrano.
	Cyrano = nook.Villager{
		Character:   cyranoCharacter,
		Personality: personality.Cranky,
		Phrase:      cyranoPhrase}
)
