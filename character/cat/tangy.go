package cat

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
	// tangyBirthday represents tangy birthday.
	tangyBirthday = nook.Birthday{
		Day:   17,
		Month: time.June}
)

var (
	// tangyCode represents tangy code.
	tangyCode = nook.Code{
		Value: "cat05"}
)

var (
	// tangyAmericanEnglishName represents tangy american english name.
	tangyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tangy"}

	// tangyCanadianFrenchName represents tangy canadian french name.
	tangyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marine"}

	// tangyDutchName represents tangy dutch name.
	tangyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tangy"}

	// tangyFrenchName represents tangy french name.
	tangyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marine"}

	// tangyGermanName represents tangy german name.
	tangyGermanName = nook.Name{
		Language: language.German,
		Value:    "Tanja"}

	// tangyItalianName represents tangy italian name.
	tangyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Arancina"}

	// tangyJapaneseName represents tangy japanese name.
	tangyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒャクパー"}

	// tangyKoreanName represents tangy korean name.
	tangyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "백프로"}

	// tangyLatinAmericanSpanishName represents tangy latin american spanish name.
	tangyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tricia"}

	// tangyRussianName represents tangy russian name.
	tangyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэнджи"}

	// tangySimplifiedChineseName represents tangy simplified chinese name.
	tangySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "纯纯"}

	// tangySpanishName represents tangy spanish name.
	tangySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tricia"}

	// tangyTraditionalChineseName represents tangy traditional chinese name.
	tangyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "純純"}
)

var (
	// tangyName represents tangy name.
	tangyName = nook.Languages{
		language.AmericanEnglish:      tangyAmericanEnglishName,
		language.CanadianFrench:       tangyCanadianFrenchName,
		language.Dutch:                tangyDutchName,
		language.French:               tangyFrenchName,
		language.German:               tangyGermanName,
		language.Italian:              tangyItalianName,
		language.Japanese:             tangyJapaneseName,
		language.Korean:               tangyKoreanName,
		language.LatinAmericanSpanish: tangyLatinAmericanSpanishName,
		language.Russian:              tangyRussianName,
		language.SimplifiedChinese:    tangySimplifiedChineseName,
		language.Spanish:              tangySpanishName,
		language.TraditionalChinese:   tangyTraditionalChineseName}
)

var (
	// tangyCharacter represents tangy character.
	tangyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: tangyBirthday,
		Code:     tangyCode,
		Key:      character.Tangy,
		Gender:   gender.Female,
		Name:     tangyName,
		Special:  false}
)

var (
	// tangyAmericanEnglishPhrase represents tangy american english phrase.
	tangyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "reeeeOWR"}

	// tangyCanadianFrenchPhrase represents tangy canadian french phrase.
	tangyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "roooaR"}

	// tangyDutchPhrase represents tangy dutch phrase.
	tangyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sinas"}

	// tangyFrenchPhrase represents tangy french phrase.
	tangyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "roooaR"}

	// tangyGermanPhrase represents tangy german phrase.
	tangyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grraul"}

	// tangyItalianPhrase represents tangy italian phrase.
	tangyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miiiiaOU"}

	// tangyJapanesePhrase represents tangy japanese phrase.
	tangyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "みかん"}

	// tangyKoreanPhrase represents tangy korean phrase.
	tangyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "귤귤"}

	// tangyLatinAmericanSpanishPhrase represents tangy latin american spanish phrase.
	tangyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mirriau"}

	// tangyRussianPhrase represents tangy russian phrase.
	tangyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фруктик"}

	// tangySimplifiedChinesePhrase represents tangy simplified chinese phrase.
	tangySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蜜柑"}

	// tangySpanishPhrase represents tangy spanish phrase.
	tangySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mandarinas"}

	// tangyTraditionalChinesePhrase represents tangy traditional chinese phrase.
	tangyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蜜柑"}
)

var (
	// tangyPhrase represents tangy phrase.
	tangyPhrase = nook.Languages{
		language.AmericanEnglish:      tangyAmericanEnglishPhrase,
		language.CanadianFrench:       tangyCanadianFrenchPhrase,
		language.Dutch:                tangyDutchPhrase,
		language.French:               tangyFrenchPhrase,
		language.German:               tangyGermanPhrase,
		language.Italian:              tangyItalianPhrase,
		language.Japanese:             tangyJapanesePhrase,
		language.Korean:               tangyKoreanPhrase,
		language.LatinAmericanSpanish: tangyLatinAmericanSpanishPhrase,
		language.Russian:              tangyRussianPhrase,
		language.SimplifiedChinese:    tangySimplifiedChinesePhrase,
		language.Spanish:              tangySpanishPhrase,
		language.TraditionalChinese:   tangyTraditionalChinesePhrase}
)

var (
	// Tangy represents tangy.
	Tangy = nook.Villager{
		Character:   tangyCharacter,
		Personality: personality.Peppy,
		Phrase:      tangyPhrase}
)
