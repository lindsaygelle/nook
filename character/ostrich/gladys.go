package ostrich

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
	// gladysBirthday represents gladys birthday.
	gladysBirthday = nook.Birthday{
		Day:   15,
		Month: time.January}
)

var (
	// gladysCode represents gladys code.
	gladysCode = nook.Code{
		Value: "ost01"}
)

var (
	// gladysAmericanEnglishName represents gladys american english name.
	gladysAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gladys"}

	// gladysCanadianFrenchName represents gladys canadian french name.
	gladysCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gladys"}

	// gladysDutchName represents gladys dutch name.
	gladysDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gladys"}

	// gladysFrenchName represents gladys french name.
	gladysFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gladys"}

	// gladysGermanName represents gladys german name.
	gladysGermanName = nook.Name{
		Language: language.German,
		Value:    "Sandra"}

	// gladysItalianName represents gladys italian name.
	gladysItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marta"}

	// gladysJapaneseName represents gladys japanese name.
	gladysJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちとせ"}

	// gladysKoreanName represents gladys korean name.
	gladysKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빅토리아"}

	// gladysLatinAmericanSpanishName represents gladys latin american spanish name.
	gladysLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gladis"}

	// gladysRussianName represents gladys russian name.
	gladysRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Глэдис"}

	// gladysSimplifiedChineseName represents gladys simplified chinese name.
	gladysSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仙仙"}

	// gladysSpanishName represents gladys spanish name.
	gladysSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gladis"}

	// gladysTraditionalChineseName represents gladys traditional chinese name.
	gladysTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仙仙"}
)

var (
	// gladysName represents gladys name.
	gladysName = nook.Languages{
		language.AmericanEnglish:      gladysAmericanEnglishName,
		language.CanadianFrench:       gladysCanadianFrenchName,
		language.Dutch:                gladysDutchName,
		language.French:               gladysFrenchName,
		language.German:               gladysGermanName,
		language.Italian:              gladysItalianName,
		language.Japanese:             gladysJapaneseName,
		language.Korean:               gladysKoreanName,
		language.LatinAmericanSpanish: gladysLatinAmericanSpanishName,
		language.Russian:              gladysRussianName,
		language.SimplifiedChinese:    gladysSimplifiedChineseName,
		language.Spanish:              gladysSpanishName,
		language.TraditionalChinese:   gladysTraditionalChineseName}
)

var (
	// gladysCharacter represents gladys character.
	gladysCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: gladysBirthday,
		Code:     gladysCode,
		Key:      character.Gladys,
		Gender:   gender.Female,
		Name:     gladysName,
		Special:  false}
)

var (
	// gladysAmericanEnglishPhrase represents gladys american english phrase.
	gladysAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stretch"}

	// gladysCanadianFrenchPhrase represents gladys canadian french phrase.
	gladysCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "codêt"}

	// gladysDutchPhrase represents gladys dutch phrase.
	gladysDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "en strek"}

	// gladysFrenchPhrase represents gladys french phrase.
	gladysFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "codêt"}

	// gladysGermanPhrase represents gladys german phrase.
	gladysGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ajajajjj"}

	// gladysItalianPhrase represents gladys italian phrase.
	gladysItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "stretch"}

	// gladysJapanesePhrase represents gladys japanese phrase.
	gladysJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですの"}

	// gladysKoreanPhrase represents gladys korean phrase.
	gladysKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇지라"}

	// gladysLatinAmericanSpanishPhrase represents gladys latin american spanish phrase.
	gladysLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumiplú"}

	// gladysRussianPhrase represents gladys russian phrase.
	gladysRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "растяжка"}

	// gladysSimplifiedChinesePhrase represents gladys simplified chinese phrase.
	gladysSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "千岁"}

	// gladysSpanishPhrase represents gladys spanish phrase.
	gladysSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plumillas"}

	// gladysTraditionalChinesePhrase represents gladys traditional chinese phrase.
	gladysTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "千歲"}
)

var (
	// gladysPhrase represents gladys phrase.
	gladysPhrase = nook.Languages{
		language.AmericanEnglish:      gladysAmericanEnglishPhrase,
		language.CanadianFrench:       gladysCanadianFrenchPhrase,
		language.Dutch:                gladysDutchPhrase,
		language.French:               gladysFrenchPhrase,
		language.German:               gladysGermanPhrase,
		language.Italian:              gladysItalianPhrase,
		language.Japanese:             gladysJapanesePhrase,
		language.Korean:               gladysKoreanPhrase,
		language.LatinAmericanSpanish: gladysLatinAmericanSpanishPhrase,
		language.Russian:              gladysRussianPhrase,
		language.SimplifiedChinese:    gladysSimplifiedChinesePhrase,
		language.Spanish:              gladysSpanishPhrase,
		language.TraditionalChinese:   gladysTraditionalChinesePhrase}
)

var (
	// Gladys represents gladys.
	Gladys = nook.Villager{
		Character:   gladysCharacter,
		Personality: personality.Normal,
		Phrase:      gladysPhrase}
)
