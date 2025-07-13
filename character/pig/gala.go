package pig

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
	// galaBirthday represents gala birthday.
	galaBirthday = nook.Birthday{
		Day:   5,
		Month: time.March}
)

var (
	// galaCode represents gala code.
	galaCode = nook.Code{
		Value: "pig13"}
)

var (
	// galaAmericanEnglishName represents gala american english name.
	galaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gala"}

	// galaCanadianFrenchName represents gala canadian french name.
	galaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Camille"}

	// galaDutchName represents gala dutch name.
	galaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gala"}

	// galaFrenchName represents gala french name.
	galaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Camille"}

	// galaGermanName represents gala german name.
	galaGermanName = nook.Name{
		Language: language.German,
		Value:    "Oinka"}

	// galaItalianName represents gala italian name.
	galaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lisetta"}

	// galaJapaneseName represents gala japanese name.
	galaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ためこ"}

	// galaKoreanName represents gala korean name.
	galaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "꽃지"}

	// galaLatinAmericanSpanishName represents gala latin american spanish name.
	galaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marita"}

	// galaRussianName represents gala russian name.
	galaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гала"}

	// galaSimplifiedChineseName represents gala simplified chinese name.
	galaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小芽"}

	// galaSpanishName represents gala spanish name.
	galaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marita"}

	// galaTraditionalChineseName represents gala traditional chinese name.
	galaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小芽"}
)

var (
	// galaName represents gala name.
	galaName = nook.Languages{
		language.AmericanEnglish:      galaAmericanEnglishName,
		language.CanadianFrench:       galaCanadianFrenchName,
		language.Dutch:                galaDutchName,
		language.French:               galaFrenchName,
		language.German:               galaGermanName,
		language.Italian:              galaItalianName,
		language.Japanese:             galaJapaneseName,
		language.Korean:               galaKoreanName,
		language.LatinAmericanSpanish: galaLatinAmericanSpanishName,
		language.Russian:              galaRussianName,
		language.SimplifiedChinese:    galaSimplifiedChineseName,
		language.Spanish:              galaSpanishName,
		language.TraditionalChinese:   galaTraditionalChineseName}
)

var (
	// galaCharacter represents gala character.
	galaCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: galaBirthday,
		Code:     galaCode,
		Key:      character.Gala,
		Gender:   gender.Female,
		Name:     galaName,
		Special:  false}
)

var (
	// galaAmericanEnglishPhrase represents gala american english phrase.
	galaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snortie"}

	// galaCanadianFrenchPhrase represents gala canadian french phrase.
	galaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tigroin"}

	// galaDutchPhrase represents gala dutch phrase.
	galaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snufferd"}

	// galaFrenchPhrase represents gala french phrase.
	galaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tigroin"}

	// galaGermanPhrase represents gala german phrase.
	galaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lächel"}

	// galaItalianPhrase represents gala italian phrase.
	galaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruffeffè"}

	// galaJapanesePhrase represents gala japanese phrase.
	galaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちゃりん"}

	// galaKoreanPhrase represents gala korean phrase.
	galaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡그랑"}

	// galaLatinAmericanSpanishPhrase represents gala latin american spanish phrase.
	galaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chanchi"}

	// galaRussianPhrase represents gala russian phrase.
	galaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюк"}

	// galaSimplifiedChinesePhrase represents gala simplified chinese phrase.
	galaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "当啷"}

	// galaSpanishPhrase represents gala spanish phrase.
	galaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chanchi"}

	// galaTraditionalChinesePhrase represents gala traditional chinese phrase.
	galaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噹啷"}
)

var (
	// galaPhrase represents gala phrase.
	galaPhrase = nook.Languages{
		language.AmericanEnglish:      galaAmericanEnglishPhrase,
		language.CanadianFrench:       galaCanadianFrenchPhrase,
		language.Dutch:                galaDutchPhrase,
		language.French:               galaFrenchPhrase,
		language.German:               galaGermanPhrase,
		language.Italian:              galaItalianPhrase,
		language.Japanese:             galaJapanesePhrase,
		language.Korean:               galaKoreanPhrase,
		language.LatinAmericanSpanish: galaLatinAmericanSpanishPhrase,
		language.Russian:              galaRussianPhrase,
		language.SimplifiedChinese:    galaSimplifiedChinesePhrase,
		language.Spanish:              galaSpanishPhrase,
		language.TraditionalChinese:   galaTraditionalChinesePhrase}
)

var (
	// Gala represents gala.
	Gala = nook.Villager{
		Character:   galaCharacter,
		Personality: personality.Normal,
		Phrase:      galaPhrase}
)
