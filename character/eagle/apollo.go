package eagle

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
	// apolloBirthday represents apollo birthday.
	apolloBirthday = nook.Birthday{
		Day:   4,
		Month: time.July}
)

var (
	// apolloCode represents apollo code.
	apolloCode = nook.Code{
		Value: "pbr00"}
)

var (
	// apolloAmericanEnglishName represents apollo american english name.
	apolloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Apollo"}

	// apolloCanadianFrenchName represents apollo canadian french name.
	apolloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Apollon"}

	// apolloDutchName represents apollo dutch name.
	apolloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Apollo"}

	// apolloFrenchName represents apollo french name.
	apolloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Apollon"}

	// apolloGermanName represents apollo german name.
	apolloGermanName = nook.Name{
		Language: language.German,
		Value:    "Apollo"}

	// apolloItalianName represents apollo italian name.
	apolloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Apollo"}

	// apolloJapaneseName represents apollo japanese name.
	apolloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アポロ"}

	// apolloKoreanName represents apollo korean name.
	apolloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아폴로"}

	// apolloLatinAmericanSpanishName represents apollo latin american spanish name.
	apolloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Apolo"}

	// apolloRussianName represents apollo russian name.
	apolloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аполло"}

	// apolloSimplifiedChineseName represents apollo simplified chinese name.
	apolloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿波罗"}

	// apolloSpanishName represents apollo spanish name.
	apolloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Apolo"}

	// apolloTraditionalChineseName represents apollo traditional chinese name.
	apolloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿波羅"}
)

var (
	// apolloName represents apollo name.
	apolloName = nook.Languages{
		language.AmericanEnglish:      apolloAmericanEnglishName,
		language.CanadianFrench:       apolloCanadianFrenchName,
		language.Dutch:                apolloDutchName,
		language.French:               apolloFrenchName,
		language.German:               apolloGermanName,
		language.Italian:              apolloItalianName,
		language.Japanese:             apolloJapaneseName,
		language.Korean:               apolloKoreanName,
		language.LatinAmericanSpanish: apolloLatinAmericanSpanishName,
		language.Russian:              apolloRussianName,
		language.SimplifiedChinese:    apolloSimplifiedChineseName,
		language.Spanish:              apolloSpanishName,
		language.TraditionalChinese:   apolloTraditionalChineseName}
)

var (
	// apolloCharacter represents apollo character.
	apolloCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: apolloBirthday,
		Code:     apolloCode,
		Key:      character.Apollo,
		Gender:   gender.Male,
		Name:     apolloName,
		Special:  false}
)

var (
	// apolloAmericanEnglishPhrase represents apollo american english phrase.
	apolloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pah"}

	// apolloCanadianFrenchPhrase represents apollo canadian french phrase.
	apolloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trraaaa"}

	// apolloDutchPhrase represents apollo dutch phrase.
	apolloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poeh"}

	// apolloFrenchPhrase represents apollo french phrase.
	apolloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trraaaa"}

	// apolloGermanPhrase represents apollo german phrase.
	apolloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahhh"}

	// apolloItalianPhrase represents apollo italian phrase.
	apolloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pah"}

	// apolloJapanesePhrase represents apollo japanese phrase.
	apolloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だワイ"}

	// apolloKoreanPhrase represents apollo korean phrase.
	apolloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다꿇어"}

	// apolloLatinAmericanSpanishPhrase represents apollo latin american spanish phrase.
	apolloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "rapahhh"}

	// apolloRussianPhrase represents apollo russian phrase.
	apolloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пф-ф"}

	// apolloSimplifiedChinesePhrase represents apollo simplified chinese phrase.
	apolloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔咿"}

	// apolloSpanishPhrase represents apollo spanish phrase.
	apolloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rapaz"}

	// apolloTraditionalChinesePhrase represents apollo traditional chinese phrase.
	apolloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔咿"}
)

var (
	// apolloPhrase represents apollo phrase.
	apolloPhrase = nook.Languages{
		language.AmericanEnglish:      apolloAmericanEnglishPhrase,
		language.CanadianFrench:       apolloCanadianFrenchPhrase,
		language.Dutch:                apolloDutchPhrase,
		language.French:               apolloFrenchPhrase,
		language.German:               apolloGermanPhrase,
		language.Italian:              apolloItalianPhrase,
		language.Japanese:             apolloJapanesePhrase,
		language.Korean:               apolloKoreanPhrase,
		language.LatinAmericanSpanish: apolloLatinAmericanSpanishPhrase,
		language.Russian:              apolloRussianPhrase,
		language.SimplifiedChinese:    apolloSimplifiedChinesePhrase,
		language.Spanish:              apolloSpanishPhrase,
		language.TraditionalChinese:   apolloTraditionalChinesePhrase}
)

var (
	// Apollo represents apollo.
	Apollo = nook.Villager{
		Character:   apolloCharacter,
		Personality: personality.Cranky,
		Phrase:      apolloPhrase}
)
