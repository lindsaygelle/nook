package chicken

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
	// hankBirthday represents hank birthday.
	hankBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// hankCode represents hank code.
	hankCode = nook.Code{
		Value: ""}
)

var (
	// hankAmericanEnglishName represents hank american english name.
	hankAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hank"}

	// hankCanadianFrenchName represents hank canadian french name.
	hankCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// hankDutchName represents hank dutch name.
	hankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// hankFrenchName represents hank french name.
	hankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bruce"}

	// hankGermanName represents hank german name.
	hankGermanName = nook.Name{
		Language: language.German,
		Value:    "Galiardo"}

	// hankItalianName represents hank italian name.
	hankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pettore"}

	// hankJapaneseName represents hank japanese name.
	hankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タツた"}

	// hankKoreanName represents hank korean name.
	hankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// hankLatinAmericanSpanishName represents hank latin american spanish name.
	hankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// hankRussianName represents hank russian name.
	hankRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// hankSimplifiedChineseName represents hank simplified chinese name.
	hankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丁丁"}

	// hankSpanishName represents hank spanish name.
	hankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Claudio"}

	// hankTraditionalChineseName represents hank traditional chinese name.
	hankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// hankName represents hank name.
	hankName = nook.Languages{
		language.AmericanEnglish:      hankAmericanEnglishName,
		language.CanadianFrench:       hankCanadianFrenchName,
		language.Dutch:                hankDutchName,
		language.French:               hankFrenchName,
		language.German:               hankGermanName,
		language.Italian:              hankItalianName,
		language.Japanese:             hankJapaneseName,
		language.Korean:               hankKoreanName,
		language.LatinAmericanSpanish: hankLatinAmericanSpanishName,
		language.Russian:              hankRussianName,
		language.SimplifiedChinese:    hankSimplifiedChineseName,
		language.Spanish:              hankSpanishName,
		language.TraditionalChinese:   hankTraditionalChineseName}
)

var (
	// hankCharacter represents hank character.
	hankCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: hankBirthday,
		Code:     hankCode,
		Key:      character.Hank,
		Gender:   gender.Male,
		Name:     hankName,
		Special:  false}
)

var (
	// hankAmericanEnglishPhrase represents hank american english phrase.
	hankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buhk buhk"}

	// hankCanadianFrenchPhrase represents hank canadian french phrase.
	hankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// hankDutchPhrase represents hank dutch phrase.
	hankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// hankFrenchPhrase represents hank french phrase.
	hankFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cocorico"}

	// hankGermanPhrase represents hank german phrase.
	hankGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bakbak"}

	// hankItalianPhrase represents hank italian phrase.
	hankItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boh boh"}

	// hankJapanesePhrase represents hank japanese phrase.
	hankJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃん"}

	// hankKoreanPhrase represents hank korean phrase.
	hankKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// hankLatinAmericanSpanishPhrase represents hank latin american spanish phrase.
	hankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// hankRussianPhrase represents hank russian phrase.
	hankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// hankSimplifiedChinesePhrase represents hank simplified chinese phrase.
	hankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啾啾"}

	// hankSpanishPhrase represents hank spanish phrase.
	hankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cocorico"}

	// hankTraditionalChinesePhrase represents hank traditional chinese phrase.
	hankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// hankPhrase represents hank phrase.
	hankPhrase = nook.Languages{
		language.AmericanEnglish:      hankAmericanEnglishPhrase,
		language.CanadianFrench:       hankCanadianFrenchPhrase,
		language.Dutch:                hankDutchPhrase,
		language.French:               hankFrenchPhrase,
		language.German:               hankGermanPhrase,
		language.Italian:              hankItalianPhrase,
		language.Japanese:             hankJapanesePhrase,
		language.Korean:               hankKoreanPhrase,
		language.LatinAmericanSpanish: hankLatinAmericanSpanishPhrase,
		language.Russian:              hankRussianPhrase,
		language.SimplifiedChinese:    hankSimplifiedChinesePhrase,
		language.Spanish:              hankSpanishPhrase,
		language.TraditionalChinese:   hankTraditionalChinesePhrase}
)

var (
	// Hank represents hank.
	Hank = nook.Villager{
		Character:   hankCharacter,
		Personality: personality.Jock,
		Phrase:      hankPhrase}
)
