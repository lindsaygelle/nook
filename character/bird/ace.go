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
	aceBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	aceCode = nook.Code{
		Value: ""}
)

var (
	aceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ace"}

	aceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	aceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	aceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Boumboum"}

	aceGermanName = nook.Name{
		Language: language.German,
		Value:    "Helge"}

	aceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Asso"}

	aceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フェザー"}

	aceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	aceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	aceRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	aceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "飞翼"}

	aceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Uno"}

	aceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	aceCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: aceBirthday,
		Code:     aceCode,
		Key:      character.Ace,
		Gender:   gender.Male,
		Name:     aceName}
)

var (
	aceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ace"}

	aceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	aceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	aceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma caille"}

	aceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "astrein"}

	aceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flapflap"}

	aceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヤス"}

	aceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	aceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	aceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	aceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "高手"}

	aceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "figura"}

	aceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	Ace = nook.Villager{
		Character:   aceCharacter,
		Personality: personality.Jock,
		Phrase:      acePhrase}
)
