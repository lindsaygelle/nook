package chicken

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
	// egbertBirthday represents egbert birthday.
	egbertBirthday = nook.Birthday{
		Day:   14,
		Month: time.October}
)

var (
	// egbertCode represents egbert code.
	egbertCode = nook.Code{
		Value: "chn02"}
)

var (
	// egbertAmericanEnglishName represents egbert american english name.
	egbertAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Egbert"}

	// egbertCanadianFrenchName represents egbert canadian french name.
	egbertCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Herbert"}

	// egbertDutchName represents egbert dutch name.
	egbertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Egbert"}

	// egbertFrenchName represents egbert french name.
	egbertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Herbert"}

	// egbertGermanName represents egbert german name.
	egbertGermanName = nook.Name{
		Language: language.German,
		Value:    "Waldemar"}

	// egbertItalianName represents egbert italian name.
	egbertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chicco"}

	// egbertJapaneseName represents egbert japanese name.
	egbertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しもやけ"}

	// egbertKoreanName represents egbert korean name.
	egbertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "김희"}

	// egbertLatinAmericanSpanishName represents egbert latin american spanish name.
	egbertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Norberto"}

	// egbertRussianName represents egbert russian name.
	egbertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эгберт"}

	// egbertSimplifiedChineseName represents egbert simplified chinese name.
	egbertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "黄金鸡"}

	// egbertSpanishName represents egbert spanish name.
	egbertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Norberto"}

	// egbertTraditionalChineseName represents egbert traditional chinese name.
	egbertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "黃金雞"}
)

var (
	// egbertName represents egbert name.
	egbertName = nook.Languages{
		language.AmericanEnglish:      egbertAmericanEnglishName,
		language.CanadianFrench:       egbertCanadianFrenchName,
		language.Dutch:                egbertDutchName,
		language.French:               egbertFrenchName,
		language.German:               egbertGermanName,
		language.Italian:              egbertItalianName,
		language.Japanese:             egbertJapaneseName,
		language.Korean:               egbertKoreanName,
		language.LatinAmericanSpanish: egbertLatinAmericanSpanishName,
		language.Russian:              egbertRussianName,
		language.SimplifiedChinese:    egbertSimplifiedChineseName,
		language.Spanish:              egbertSpanishName,
		language.TraditionalChinese:   egbertTraditionalChineseName}
)

var (
	// egbertCharacter represents egbert character.
	egbertCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: egbertBirthday,
		Code:     egbertCode,
		Key:      character.Egbert,
		Gender:   gender.Male,
		Name:     egbertName,
		Special:  false}
)

var (
	// egbertAmericanEnglishPhrase represents egbert american english phrase.
	egbertAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "doodle-duh"}

	// egbertCanadianFrenchPhrase represents egbert canadian french phrase.
	egbertCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "di di dou"}

	// egbertDutchPhrase represents egbert dutch phrase.
	egbertDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kukelejus"}

	// egbertFrenchPhrase represents egbert french phrase.
	egbertFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "di di dou"}

	// egbertGermanPhrase represents egbert german phrase.
	egbertGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "duudle-du"}

	// egbertItalianPhrase represents egbert italian phrase.
	egbertItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "du du du"}

	// egbertJapanesePhrase represents egbert japanese phrase.
	egbertJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だヨ"}

	// egbertKoreanPhrase represents egbert korean phrase.
	egbertKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜잔"}

	// egbertLatinAmericanSpanishPhrase represents egbert latin american spanish phrase.
	egbertLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kikirikí"}

	// egbertRussianPhrase represents egbert russian phrase.
	egbertRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кукареку"}

	// egbertSimplifiedChinesePhrase represents egbert simplified chinese phrase.
	egbertSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "很优"}

	// egbertSpanishPhrase represents egbert spanish phrase.
	egbertSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kikirikí"}

	// egbertTraditionalChinesePhrase represents egbert traditional chinese phrase.
	egbertTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "很優"}
)

var (
	// egbertPhrase represents egbert phrase.
	egbertPhrase = nook.Languages{
		language.AmericanEnglish:      egbertAmericanEnglishPhrase,
		language.CanadianFrench:       egbertCanadianFrenchPhrase,
		language.Dutch:                egbertDutchPhrase,
		language.French:               egbertFrenchPhrase,
		language.German:               egbertGermanPhrase,
		language.Italian:              egbertItalianPhrase,
		language.Japanese:             egbertJapanesePhrase,
		language.Korean:               egbertKoreanPhrase,
		language.LatinAmericanSpanish: egbertLatinAmericanSpanishPhrase,
		language.Russian:              egbertRussianPhrase,
		language.SimplifiedChinese:    egbertSimplifiedChinesePhrase,
		language.Spanish:              egbertSpanishPhrase,
		language.TraditionalChinese:   egbertTraditionalChinesePhrase}
)

var (
	// Egbert represents egbert.
	Egbert = nook.Villager{
		Character:   egbertCharacter,
		Personality: personality.Lazy,
		Phrase:      egbertPhrase}
)
