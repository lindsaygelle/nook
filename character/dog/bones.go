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
	// bonesBirthday represents bones birthday.
	bonesBirthday = nook.Birthday{
		Day:   4,
		Month: time.August}
)

var (
	// bonesCode represents bones code.
	bonesCode = nook.Code{
		Value: "dog04"}
)

var (
	// bonesAmericanEnglishName represents bones american english name.
	bonesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bones"}

	// bonesCanadianFrenchName represents bones canadian french name.
	bonesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nonos"}

	// bonesDutchName represents bones dutch name.
	bonesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bones"}

	// bonesFrenchName represents bones french name.
	bonesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nonos"}

	// bonesGermanName represents bones german name.
	bonesGermanName = nook.Name{
		Language: language.German,
		Value:    "Strolch"}

	// bonesItalianName represents bones italian name.
	bonesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tobia"}

	// bonesJapaneseName represents bones japanese name.
	bonesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トミ"}

	// bonesKoreanName represents bones korean name.
	bonesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토미"}

	// bonesLatinAmericanSpanishName represents bones latin american spanish name.
	bonesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cándido"}

	// bonesRussianName represents bones russian name.
	bonesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Боунс"}

	// bonesSimplifiedChineseName represents bones simplified chinese name.
	bonesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿富"}

	// bonesSpanishName represents bones spanish name.
	bonesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cándido"}

	// bonesTraditionalChineseName represents bones traditional chinese name.
	bonesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿富"}
)

var (
	// bonesName represents bones name.
	bonesName = nook.Languages{
		language.AmericanEnglish:      bonesAmericanEnglishName,
		language.CanadianFrench:       bonesCanadianFrenchName,
		language.Dutch:                bonesDutchName,
		language.French:               bonesFrenchName,
		language.German:               bonesGermanName,
		language.Italian:              bonesItalianName,
		language.Japanese:             bonesJapaneseName,
		language.Korean:               bonesKoreanName,
		language.LatinAmericanSpanish: bonesLatinAmericanSpanishName,
		language.Russian:              bonesRussianName,
		language.SimplifiedChinese:    bonesSimplifiedChineseName,
		language.Spanish:              bonesSpanishName,
		language.TraditionalChinese:   bonesTraditionalChineseName}
)

var (
	// bonesCharacter represents bones character.
	bonesCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: bonesBirthday,
		Code:     bonesCode,
		Key:      character.Bones,
		Gender:   gender.Male,
		Name:     bonesName,
		Special:  false}
)

var (
	// bonesAmericanEnglishPhrase represents bones american english phrase.
	bonesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yip yip"}

	// bonesCanadianFrenchPhrase represents bones canadian french phrase.
	bonesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yip yip"}

	// bonesDutchPhrase represents bones dutch phrase.
	bonesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuf snuf"}

	// bonesFrenchPhrase represents bones french phrase.
	bonesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yip yip"}

	// bonesGermanPhrase represents bones german phrase.
	bonesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hechel"}

	// bonesItalianPhrase represents bones italian phrase.
	bonesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yip yip"}

	// bonesJapanesePhrase represents bones japanese phrase.
	bonesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ネッ"}

	// bonesKoreanPhrase represents bones korean phrase.
	bonesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "옙"}

	// bonesLatinAmericanSpanishPhrase represents bones latin american spanish phrase.
	bonesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yip yip"}

	// bonesRussianPhrase represents bones russian phrase.
	bonesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв-тяв"}

	// bonesSimplifiedChinesePhrase represents bones simplified chinese phrase.
	bonesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对吧"}

	// bonesSpanishPhrase represents bones spanish phrase.
	bonesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yip yip"}

	// bonesTraditionalChinesePhrase represents bones traditional chinese phrase.
	bonesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對吧"}
)

var (
	// bonesPhrase represents bones phrase.
	bonesPhrase = nook.Languages{
		language.AmericanEnglish:      bonesAmericanEnglishPhrase,
		language.CanadianFrench:       bonesCanadianFrenchPhrase,
		language.Dutch:                bonesDutchPhrase,
		language.French:               bonesFrenchPhrase,
		language.German:               bonesGermanPhrase,
		language.Italian:              bonesItalianPhrase,
		language.Japanese:             bonesJapanesePhrase,
		language.Korean:               bonesKoreanPhrase,
		language.LatinAmericanSpanish: bonesLatinAmericanSpanishPhrase,
		language.Russian:              bonesRussianPhrase,
		language.SimplifiedChinese:    bonesSimplifiedChinesePhrase,
		language.Spanish:              bonesSpanishPhrase,
		language.TraditionalChinese:   bonesTraditionalChinesePhrase}
)

var (
	// Bones represents bones.
	Bones = nook.Villager{
		Character:   bonesCharacter,
		Personality: personality.Lazy,
		Phrase:      bonesPhrase}
)
