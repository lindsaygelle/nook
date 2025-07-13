package mouse

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
	// limbergBirthday represents limberg birthday.
	limbergBirthday = nook.Birthday{
		Day:   17,
		Month: time.October}
)

var (
	// limbergCode represents limberg code.
	limbergCode = nook.Code{
		Value: "mus01"}
)

var (
	// limbergAmericanEnglishName represents limberg american english name.
	limbergAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Limberg"}

	// limbergCanadianFrenchName represents limberg canadian french name.
	limbergCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gruyère"}

	// limbergDutchName represents limberg dutch name.
	limbergDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Limberg"}

	// limbergFrenchName represents limberg french name.
	limbergFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gruyère"}

	// limbergGermanName represents limberg german name.
	limbergGermanName = nook.Name{
		Language: language.German,
		Value:    "Rafael"}

	// limbergItalianName represents limberg italian name.
	limbergItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosico"}

	// limbergJapaneseName represents limberg japanese name.
	limbergJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "らっきょ"}

	// limbergKoreanName represents limberg korean name.
	limbergKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "단무지"}

	// limbergLatinAmericanSpanishName represents limberg latin american spanish name.
	limbergLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camember"}

	// limbergRussianName represents limberg russian name.
	limbergRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лимберг"}

	// limbergSimplifiedChineseName represents limberg simplified chinese name.
	limbergSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "韭菜"}

	// limbergSpanishName represents limberg spanish name.
	limbergSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camember"}

	// limbergTraditionalChineseName represents limberg traditional chinese name.
	limbergTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "韭菜"}
)

var (
	// limbergName represents limberg name.
	limbergName = nook.Languages{
		language.AmericanEnglish:      limbergAmericanEnglishName,
		language.CanadianFrench:       limbergCanadianFrenchName,
		language.Dutch:                limbergDutchName,
		language.French:               limbergFrenchName,
		language.German:               limbergGermanName,
		language.Italian:              limbergItalianName,
		language.Japanese:             limbergJapaneseName,
		language.Korean:               limbergKoreanName,
		language.LatinAmericanSpanish: limbergLatinAmericanSpanishName,
		language.Russian:              limbergRussianName,
		language.SimplifiedChinese:    limbergSimplifiedChineseName,
		language.Spanish:              limbergSpanishName,
		language.TraditionalChinese:   limbergTraditionalChineseName}
)

var (
	// limbergCharacter represents limberg character.
	limbergCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: limbergBirthday,
		Code:     limbergCode,
		Key:      character.Limberg,
		Gender:   gender.Male,
		Name:     limbergName,
		Special:  false}
)

var (
	// limbergAmericanEnglishPhrase represents limberg american english phrase.
	limbergAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squinky"}

	// limbergCanadianFrenchPhrase represents limberg canadian french phrase.
	limbergCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "scriiich"}

	// limbergDutchPhrase represents limberg dutch phrase.
	limbergDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepsels"}

	// limbergFrenchPhrase represents limberg french phrase.
	limbergFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "scriiich"}

	// limbergGermanPhrase represents limberg german phrase.
	limbergGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pfiffikus"}

	// limbergItalianPhrase represents limberg italian phrase.
	limbergItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squick"}

	// limbergJapanesePhrase represents limberg japanese phrase.
	limbergJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てやんで"}

	// limbergKoreanPhrase represents limberg korean phrase.
	limbergKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "긍가벼"}

	// limbergLatinAmericanSpanishPhrase represents limberg latin american spanish phrase.
	limbergLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "scroch"}

	// limbergRussianPhrase represents limberg russian phrase.
	limbergRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фью-фью"}

	// limbergSimplifiedChinesePhrase represents limberg simplified chinese phrase.
	limbergSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "非也"}

	// limbergSpanishPhrase represents limberg spanish phrase.
	limbergSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caramba"}

	// limbergTraditionalChinesePhrase represents limberg traditional chinese phrase.
	limbergTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "非也"}
)

var (
	// limbergPhrase represents limberg phrase.
	limbergPhrase = nook.Languages{
		language.AmericanEnglish:      limbergAmericanEnglishPhrase,
		language.CanadianFrench:       limbergCanadianFrenchPhrase,
		language.Dutch:                limbergDutchPhrase,
		language.French:               limbergFrenchPhrase,
		language.German:               limbergGermanPhrase,
		language.Italian:              limbergItalianPhrase,
		language.Japanese:             limbergJapanesePhrase,
		language.Korean:               limbergKoreanPhrase,
		language.LatinAmericanSpanish: limbergLatinAmericanSpanishPhrase,
		language.Russian:              limbergRussianPhrase,
		language.SimplifiedChinese:    limbergSimplifiedChinesePhrase,
		language.Spanish:              limbergSpanishPhrase,
		language.TraditionalChinese:   limbergTraditionalChinesePhrase}
)

var (
	// Limberg represents limberg.
	Limberg = nook.Villager{
		Character:   limbergCharacter,
		Personality: personality.Cranky,
		Phrase:      limbergPhrase}
)
