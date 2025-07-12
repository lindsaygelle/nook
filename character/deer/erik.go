package deer

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
	// erikBirthday represents erik birthday.
	erikBirthday = nook.Birthday{
		Day:   27,
		Month: time.July}
)

var (
	// erikCode represents erik code.
	erikCode = nook.Code{
		Value: "der09"}
)

var (
	// erikAmericanEnglishName represents erik american english name.
	erikAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Erik"}

	// erikCanadianFrenchName represents erik canadian french name.
	erikCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Abraham"}

	// erikDutchName represents erik dutch name.
	erikDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Erik"}

	// erikFrenchName represents erik french name.
	erikFrenchName = nook.Name{
		Language: language.French,
		Value:    "Abraham"}

	// erikGermanName represents erik german name.
	erikGermanName = nook.Name{
		Language: language.German,
		Value:    "Erik"}

	// erikItalianName represents erik italian name.
	erikItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cervasio"}

	// erikJapaneseName represents erik japanese name.
	erikJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャック"}

	// erikKoreanName represents erik korean name.
	erikKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "자끄"}

	// erikLatinAmericanSpanishName represents erik latin american spanish name.
	erikLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cervasio"}

	// erikRussianName represents erik russian name.
	erikRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эрик"}

	// erikSimplifiedChineseName represents erik simplified chinese name.
	erikSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "查克"}

	// erikSpanishName represents erik spanish name.
	erikSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cervasio"}

	// erikTraditionalChineseName represents erik traditional chinese name.
	erikTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "查克"}
)

var (
	// erikName represents erik name.
	erikName = nook.Languages{
		language.AmericanEnglish:      erikAmericanEnglishName,
		language.CanadianFrench:       erikCanadianFrenchName,
		language.Dutch:                erikDutchName,
		language.French:               erikFrenchName,
		language.German:               erikGermanName,
		language.Italian:              erikItalianName,
		language.Japanese:             erikJapaneseName,
		language.Korean:               erikKoreanName,
		language.LatinAmericanSpanish: erikLatinAmericanSpanishName,
		language.Russian:              erikRussianName,
		language.SimplifiedChinese:    erikSimplifiedChineseName,
		language.Spanish:              erikSpanishName,
		language.TraditionalChinese:   erikTraditionalChineseName}
)

var (
	// erikCharacter represents erik character.
	erikCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: erikBirthday,
		Code:     erikCode,
		Key:      character.Erik,
		Gender:   gender.Male,
		Name:     erikName,
		Special:  false}
)

var (
	// erikAmericanEnglishPhrase represents erik american english phrase.
	erikAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chow down"}

	// erikCanadianFrenchPhrase represents erik canadian french phrase.
	erikCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "breuh"}

	// erikDutchPhrase represents erik dutch phrase.
	erikDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schorsie"}

	// erikFrenchPhrase represents erik french phrase.
	erikFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "breuh"}

	// erikGermanPhrase represents erik german phrase.
	erikGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hirschl"}

	// erikItalianPhrase represents erik italian phrase.
	erikItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "relax"}

	// erikJapanesePhrase represents erik japanese phrase.
	erikJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "しかじか"}

	// erikKoreanPhrase represents erik korean phrase.
	erikKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뚜비뚜바"}

	// erikLatinAmericanSpanishPhrase represents erik latin american spanish phrase.
	erikLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñume-ñume"}

	// erikRussianPhrase represents erik russian phrase.
	erikRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ем мох"}

	// erikSimplifiedChinesePhrase represents erik simplified chinese phrase.
	erikSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹿鹿"}

	// erikSpanishPhrase represents erik spanish phrase.
	erikSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñume-ñume"}

	// erikTraditionalChinesePhrase represents erik traditional chinese phrase.
	erikTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鹿鹿"}
)

var (
	// erikPhrase represents erik phrase.
	erikPhrase = nook.Languages{
		language.AmericanEnglish:      erikAmericanEnglishPhrase,
		language.CanadianFrench:       erikCanadianFrenchPhrase,
		language.Dutch:                erikDutchPhrase,
		language.French:               erikFrenchPhrase,
		language.German:               erikGermanPhrase,
		language.Italian:              erikItalianPhrase,
		language.Japanese:             erikJapanesePhrase,
		language.Korean:               erikKoreanPhrase,
		language.LatinAmericanSpanish: erikLatinAmericanSpanishPhrase,
		language.Russian:              erikRussianPhrase,
		language.SimplifiedChinese:    erikSimplifiedChinesePhrase,
		language.Spanish:              erikSpanishPhrase,
		language.TraditionalChinese:   erikTraditionalChinesePhrase}
)

var (
	// Erik represents erik.
	Erik = nook.Villager{
		Character:   erikCharacter,
		Personality: personality.Lazy,
		Phrase:      erikPhrase}
)
