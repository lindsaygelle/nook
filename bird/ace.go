package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Boumboumma caille"}

	aceGermanName = nook.Name{
		Language: language.German,
		Value:    "Helgeastrein"}

	aceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Assoflapflap"}

	aceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フェザーでヤス"}

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
		Value:    "飞翼高手"}

	aceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Unofigura"}

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
		Animal:   Bird,
		Birthday: aceBirthday,
		Code:     aceCode,
		Gender:   nook.Male,
		Name:     aceName}
)

var (
	aceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	aceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	aceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	aceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	aceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	aceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "高手"}

	aceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "figura"}

	aceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	acePhrase = nook.Languages{
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
	Ace = nook.Villager{
		Character:   aceCharacter,
		Personality: nook.Jock,
		Phrase:      acePhrase}
)
