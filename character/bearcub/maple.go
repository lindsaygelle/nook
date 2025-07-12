package bearcub

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
	// mapleBirthday represents maple birthday.
	mapleBirthday = nook.Birthday{
		Day:   15,
		Month: time.June}
)

var (
	// mapleCode represents maple code.
	mapleCode = nook.Code{
		Value: "cbr01"}
)

var (
	// mapleAmericanEnglishName represents maple american english name.
	mapleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maple"}

	// mapleCanadianFrenchName represents maple canadian french name.
	mapleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léa"}

	// mapleDutchName represents maple dutch name.
	mapleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maple"}

	// mapleFrenchName represents maple french name.
	mapleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léa"}

	// mapleGermanName represents maple german name.
	mapleGermanName = nook.Name{
		Language: language.German,
		Value:    "Mona"}

	// mapleItalianName represents maple italian name.
	mapleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dulcinea"}

	// mapleJapaneseName represents maple japanese name.
	mapleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メープル"}

	// mapleKoreanName represents maple korean name.
	mapleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메이첼"}

	// mapleLatinAmericanSpanishName represents maple latin american spanish name.
	mapleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dulce"}

	// mapleRussianName represents maple russian name.
	mapleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мейпл"}

	// mapleSimplifiedChineseName represents maple simplified chinese name.
	mapleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小枫"}

	// mapleSpanishName represents maple spanish name.
	mapleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dulce"}

	// mapleTraditionalChineseName represents maple traditional chinese name.
	mapleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小楓"}
)

var (
	// mapleName represents maple name.
	mapleName = nook.Languages{
		language.AmericanEnglish:      mapleAmericanEnglishName,
		language.CanadianFrench:       mapleCanadianFrenchName,
		language.Dutch:                mapleDutchName,
		language.French:               mapleFrenchName,
		language.German:               mapleGermanName,
		language.Italian:              mapleItalianName,
		language.Japanese:             mapleJapaneseName,
		language.Korean:               mapleKoreanName,
		language.LatinAmericanSpanish: mapleLatinAmericanSpanishName,
		language.Russian:              mapleRussianName,
		language.SimplifiedChinese:    mapleSimplifiedChineseName,
		language.Spanish:              mapleSpanishName,
		language.TraditionalChinese:   mapleTraditionalChineseName}
)

var (
	// mapleCharacter represents maple character.
	mapleCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: mapleBirthday,
		Code:     mapleCode,
		Key:      character.Maple,
		Gender:   gender.Female,
		Name:     mapleName,
		Special:  false}
)

var (
	// mapleAmericanEnglishPhrase represents maple american english phrase.
	mapleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honey"}

	// mapleCanadianFrenchPhrase represents maple canadian french phrase.
	mapleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chouchou"}

	// mapleDutchPhrase represents maple dutch phrase.
	mapleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zoeterd"}

	// mapleFrenchPhrase represents maple french phrase.
	mapleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chouchou"}

	// mapleGermanPhrase represents maple german phrase.
	mapleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "darling"}

	// mapleItalianPhrase represents maple italian phrase.
	mapleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miele"}

	// mapleJapanesePhrase represents maple japanese phrase.
	mapleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だベア"}

	// mapleKoreanPhrase represents maple korean phrase.
	mapleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "저기요"}

	// mapleLatinAmericanSpanishPhrase represents maple latin american spanish phrase.
	mapleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mielmiel"}

	// mapleRussianPhrase represents maple russian phrase.
	mapleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лапушка"}

	// mapleSimplifiedChinesePhrase represents maple simplified chinese phrase.
	mapleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊仔"}

	// mapleSpanishPhrase represents maple spanish phrase.
	mapleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caramelito"}

	// mapleTraditionalChinesePhrase represents maple traditional chinese phrase.
	mapleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊仔"}
)

var (
	// maplePhrase represents maple phrase.
	maplePhrase = nook.Languages{
		language.AmericanEnglish:      mapleAmericanEnglishPhrase,
		language.CanadianFrench:       mapleCanadianFrenchPhrase,
		language.Dutch:                mapleDutchPhrase,
		language.French:               mapleFrenchPhrase,
		language.German:               mapleGermanPhrase,
		language.Italian:              mapleItalianPhrase,
		language.Japanese:             mapleJapanesePhrase,
		language.Korean:               mapleKoreanPhrase,
		language.LatinAmericanSpanish: mapleLatinAmericanSpanishPhrase,
		language.Russian:              mapleRussianPhrase,
		language.SimplifiedChinese:    mapleSimplifiedChinesePhrase,
		language.Spanish:              mapleSpanishPhrase,
		language.TraditionalChinese:   mapleTraditionalChinesePhrase}
)

var (
	// Maple represents maple.
	Maple = nook.Villager{
		Character:   mapleCharacter,
		Personality: personality.Normal,
		Phrase:      maplePhrase}
)
