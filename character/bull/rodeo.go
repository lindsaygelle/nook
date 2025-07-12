package bull

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
	// rodeoBirthday represents rodeo birthday.
	rodeoBirthday = nook.Birthday{
		Day:   29,
		Month: time.October}
)

var (
	// rodeoCode represents rodeo code.
	rodeoCode = nook.Code{
		Value: "bul01"}
)

var (
	// rodeoAmericanEnglishName represents rodeo american english name.
	rodeoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rodeo"}

	// rodeoCanadianFrenchName represents rodeo canadian french name.
	rodeoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Flèche"}

	// rodeoDutchName represents rodeo dutch name.
	rodeoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rodeo"}

	// rodeoFrenchName represents rodeo french name.
	rodeoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Flèche"}

	// rodeoGermanName represents rodeo german name.
	rodeoGermanName = nook.Name{
		Language: language.German,
		Value:    "Toro"}

	// rodeoItalianName represents rodeo italian name.
	rodeoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rodeo"}

	// rodeoJapaneseName represents rodeo japanese name.
	rodeoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロデオ"}

	// rodeoKoreanName represents rodeo korean name.
	rodeoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로데오"}

	// rodeoLatinAmericanSpanishName represents rodeo latin american spanish name.
	rodeoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodeo"}

	// rodeoRussianName represents rodeo russian name.
	rodeoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Родео"}

	// rodeoSimplifiedChineseName represents rodeo simplified chinese name.
	rodeoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗迪欧"}

	// rodeoSpanishName represents rodeo spanish name.
	rodeoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodeo"}

	// rodeoTraditionalChineseName represents rodeo traditional chinese name.
	rodeoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅迪歐"}
)

var (
	// rodeoName represents rodeo name.
	rodeoName = nook.Languages{
		language.AmericanEnglish:      rodeoAmericanEnglishName,
		language.CanadianFrench:       rodeoCanadianFrenchName,
		language.Dutch:                rodeoDutchName,
		language.French:               rodeoFrenchName,
		language.German:               rodeoGermanName,
		language.Italian:              rodeoItalianName,
		language.Japanese:             rodeoJapaneseName,
		language.Korean:               rodeoKoreanName,
		language.LatinAmericanSpanish: rodeoLatinAmericanSpanishName,
		language.Russian:              rodeoRussianName,
		language.SimplifiedChinese:    rodeoSimplifiedChineseName,
		language.Spanish:              rodeoSpanishName,
		language.TraditionalChinese:   rodeoTraditionalChineseName}
)

var (
	// rodeoCharacter represents rodeo character.
	rodeoCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: rodeoBirthday,
		Code:     rodeoCode,
		Key:      character.Rodeo,
		Gender:   gender.Male,
		Name:     rodeoName,
		Special:  false}
)

var (
	// rodeoAmericanEnglishPhrase represents rodeo american english phrase.
	rodeoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chaps"}

	// rodeoCanadianFrenchPhrase represents rodeo canadian french phrase.
	rodeoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "olé"}

	// rodeoDutchPhrase represents rodeo dutch phrase.
	rodeoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cowboy"}

	// rodeoFrenchPhrase represents rodeo french phrase.
	rodeoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crac"}

	// rodeoGermanPhrase represents rodeo german phrase.
	rodeoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oléolé"}

	// rodeoItalianPhrase represents rodeo italian phrase.
	rodeoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oléolé"}

	// rodeoJapanesePhrase represents rodeo japanese phrase.
	rodeoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さすがに"}

	// rodeoKoreanPhrase represents rodeo korean phrase.
	rodeoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "디용"}

	// rodeoLatinAmericanSpanishPhrase represents rodeo latin american spanish phrase.
	rodeoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "olé"}

	// rodeoRussianPhrase represents rodeo russian phrase.
	rodeoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ковбой"}

	// rodeoSimplifiedChinesePhrase represents rodeo simplified chinese phrase.
	rodeoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "了不起"}

	// rodeoSpanishPhrase represents rodeo spanish phrase.
	rodeoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "olé"}

	// rodeoTraditionalChinesePhrase represents rodeo traditional chinese phrase.
	rodeoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "了不起"}
)

var (
	// rodeoPhrase represents rodeo phrase.
	rodeoPhrase = nook.Languages{
		language.AmericanEnglish:      rodeoAmericanEnglishPhrase,
		language.CanadianFrench:       rodeoCanadianFrenchPhrase,
		language.Dutch:                rodeoDutchPhrase,
		language.French:               rodeoFrenchPhrase,
		language.German:               rodeoGermanPhrase,
		language.Italian:              rodeoItalianPhrase,
		language.Japanese:             rodeoJapanesePhrase,
		language.Korean:               rodeoKoreanPhrase,
		language.LatinAmericanSpanish: rodeoLatinAmericanSpanishPhrase,
		language.Russian:              rodeoRussianPhrase,
		language.SimplifiedChinese:    rodeoSimplifiedChinesePhrase,
		language.Spanish:              rodeoSpanishPhrase,
		language.TraditionalChinese:   rodeoTraditionalChinesePhrase}
)

var (
	// Rodeo represents rodeo.
	Rodeo = nook.Villager{
		Character:   rodeoCharacter,
		Personality: personality.Lazy,
		Phrase:      rodeoPhrase}
)
