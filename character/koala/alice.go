package koala

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
	// aliceBirthday represents alice birthday.
	aliceBirthday = nook.Birthday{
		Day:   19,
		Month: time.August}
)

var (
	// aliceCode represents alice code.
	aliceCode = nook.Code{
		Value: "kal01"}
)

var (
	// aliceAmericanEnglishName represents alice american english name.
	aliceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alice"}

	// aliceCanadianFrenchName represents alice canadian french name.
	aliceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Alice"}

	// aliceDutchName represents alice dutch name.
	aliceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alice"}

	// aliceFrenchName represents alice french name.
	aliceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Alice"}

	// aliceGermanName represents alice german name.
	aliceGermanName = nook.Name{
		Language: language.German,
		Value:    "Konny"}

	// aliceItalianName represents alice italian name.
	aliceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alice"}

	// aliceJapaneseName represents alice japanese name.
	aliceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メルボルン"}

	// aliceKoreanName represents alice korean name.
	aliceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멜버른"}

	// aliceLatinAmericanSpanishName represents alice latin american spanish name.
	aliceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zelanda"}

	// aliceRussianName represents alice russian name.
	aliceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Алиса"}

	// aliceSimplifiedChineseName represents alice simplified chinese name.
	aliceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫儿"}

	// aliceSpanishName represents alice spanish name.
	aliceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zelanda"}

	// aliceTraditionalChineseName represents alice traditional chinese name.
	aliceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫兒"}
)

var (
	// aliceName represents alice name.
	aliceName = nook.Languages{
		language.AmericanEnglish:      aliceAmericanEnglishName,
		language.CanadianFrench:       aliceCanadianFrenchName,
		language.Dutch:                aliceDutchName,
		language.French:               aliceFrenchName,
		language.German:               aliceGermanName,
		language.Italian:              aliceItalianName,
		language.Japanese:             aliceJapaneseName,
		language.Korean:               aliceKoreanName,
		language.LatinAmericanSpanish: aliceLatinAmericanSpanishName,
		language.Russian:              aliceRussianName,
		language.SimplifiedChinese:    aliceSimplifiedChineseName,
		language.Spanish:              aliceSpanishName,
		language.TraditionalChinese:   aliceTraditionalChineseName}
)

var (
	// aliceCharacter represents alice character.
	aliceCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: aliceBirthday,
		Code:     aliceCode,
		Key:      character.Alice,
		Gender:   gender.Female,
		Name:     aliceName,
		Special:  false}
)

var (
	// aliceAmericanEnglishPhrase represents alice american english phrase.
	aliceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "guvnor"}

	// aliceCanadianFrenchPhrase represents alice canadian french phrase.
	aliceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chef"}

	// aliceDutchPhrase represents alice dutch phrase.
	aliceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "chef"}

	// aliceFrenchPhrase represents alice french phrase.
	aliceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chef"}

	// aliceGermanPhrase represents alice german phrase.
	aliceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnarchhh"}

	// aliceItalianPhrase represents alice italian phrase.
	aliceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tesorino"}

	// aliceJapanesePhrase represents alice japanese phrase.
	aliceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キラリ"}

	// aliceKoreanPhrase represents alice korean phrase.
	aliceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "반짝"}

	// aliceLatinAmericanSpanishPhrase represents alice latin american spanish phrase.
	aliceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "eucaliup"}

	// aliceRussianPhrase represents alice russian phrase.
	aliceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лидер"}

	// aliceSimplifiedChinesePhrase represents alice simplified chinese phrase.
	aliceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晶亮"}

	// aliceSpanishPhrase represents alice spanish phrase.
	aliceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tú"}

	// aliceTraditionalChinesePhrase represents alice traditional chinese phrase.
	aliceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晶亮"}
)

var (
	// alicePhrase represents alice phrase.
	alicePhrase = nook.Languages{
		language.AmericanEnglish:      aliceAmericanEnglishPhrase,
		language.CanadianFrench:       aliceCanadianFrenchPhrase,
		language.Dutch:                aliceDutchPhrase,
		language.French:               aliceFrenchPhrase,
		language.German:               aliceGermanPhrase,
		language.Italian:              aliceItalianPhrase,
		language.Japanese:             aliceJapanesePhrase,
		language.Korean:               aliceKoreanPhrase,
		language.LatinAmericanSpanish: aliceLatinAmericanSpanishPhrase,
		language.Russian:              aliceRussianPhrase,
		language.SimplifiedChinese:    aliceSimplifiedChinesePhrase,
		language.Spanish:              aliceSpanishPhrase,
		language.TraditionalChinese:   aliceTraditionalChinesePhrase}
)

var (
	// Alice represents alice.
	Alice = nook.Villager{
		Character:   aliceCharacter,
		Personality: personality.Normal,
		Phrase:      alicePhrase}
)
