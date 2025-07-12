package bird

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
	// aceBirthday represents ace birthday.
	aceBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// aceCode represents ace code.
	aceCode = nook.Code{
		Value: ""}
)

var (
	// aceAmericanEnglishName represents ace american english name.
	aceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ace"}

	// aceCanadianFrenchName represents ace canadian french name.
	aceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// aceDutchName represents ace dutch name.
	aceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// aceFrenchName represents ace french name.
	aceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Boumboum"}

	// aceGermanName represents ace german name.
	aceGermanName = nook.Name{
		Language: language.German,
		Value:    "Helge"}

	// aceItalianName represents ace italian name.
	aceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Asso"}

	// aceJapaneseName represents ace japanese name.
	aceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フェザー"}

	// aceKoreanName represents ace korean name.
	aceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// aceLatinAmericanSpanishName represents ace latin american spanish name.
	aceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// aceRussianName represents ace russian name.
	aceRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// aceSimplifiedChineseName represents ace simplified chinese name.
	aceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "飞翼"}

	// aceSpanishName represents ace spanish name.
	aceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Uno"}

	// aceTraditionalChineseName represents ace traditional chinese name.
	aceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// aceName represents ace name.
	aceName = nook.Languages{
		language.AmericanEnglish:      aceAmericanEnglishName,
		language.CanadianFrench:       aceCanadianFrenchName,
		language.Dutch:                aceDutchName,
		language.French:               aceFrenchName,
		language.German:               aceGermanName,
		language.Italian:              aceItalianName,
		language.Japanese:             aceJapaneseName,
		language.Korean:               aceKoreanName,
		language.LatinAmericanSpanish: aceLatinAmericanSpanishName,
		language.Russian:              aceRussianName,
		language.SimplifiedChinese:    aceSimplifiedChineseName,
		language.Spanish:              aceSpanishName,
		language.TraditionalChinese:   aceTraditionalChineseName}
)

var (
	// aceCharacter represents ace character.
	aceCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: aceBirthday,
		Code:     aceCode,
		Key:      character.Ace,
		Gender:   gender.Male,
		Name:     aceName,
		Special:  false}
)

var (
	// aceAmericanEnglishPhrase represents ace american english phrase.
	aceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ace"}

	// aceCanadianFrenchPhrase represents ace canadian french phrase.
	aceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// aceDutchPhrase represents ace dutch phrase.
	aceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// aceFrenchPhrase represents ace french phrase.
	aceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma caille"}

	// aceGermanPhrase represents ace german phrase.
	aceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "astrein"}

	// aceItalianPhrase represents ace italian phrase.
	aceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flapflap"}

	// aceJapanesePhrase represents ace japanese phrase.
	aceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヤス"}

	// aceKoreanPhrase represents ace korean phrase.
	aceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// aceLatinAmericanSpanishPhrase represents ace latin american spanish phrase.
	aceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// aceRussianPhrase represents ace russian phrase.
	aceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// aceSimplifiedChinesePhrase represents ace simplified chinese phrase.
	aceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "高手"}

	// aceSpanishPhrase represents ace spanish phrase.
	aceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "figura"}

	// aceTraditionalChinesePhrase represents ace traditional chinese phrase.
	aceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// acePhrase represents ace phrase.
	acePhrase = nook.Languages{
		language.AmericanEnglish:      aceAmericanEnglishPhrase,
		language.CanadianFrench:       aceCanadianFrenchPhrase,
		language.Dutch:                aceDutchPhrase,
		language.French:               aceFrenchPhrase,
		language.German:               aceGermanPhrase,
		language.Italian:              aceItalianPhrase,
		language.Japanese:             aceJapanesePhrase,
		language.Korean:               aceKoreanPhrase,
		language.LatinAmericanSpanish: aceLatinAmericanSpanishPhrase,
		language.Russian:              aceRussianPhrase,
		language.SimplifiedChinese:    aceSimplifiedChinesePhrase,
		language.Spanish:              aceSpanishPhrase,
		language.TraditionalChinese:   aceTraditionalChinesePhrase}
)

var (
	// Ace represents ace.
	Ace = nook.Villager{
		Character:   aceCharacter,
		Personality: personality.Jock,
		Phrase:      acePhrase}
)
