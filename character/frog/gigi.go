package frog

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
	// gigiBirthday represents gigi birthday.
	gigiBirthday = nook.Birthday{
		Day:   11,
		Month: time.August}
)

var (
	// gigiCode represents gigi code.
	gigiCode = nook.Code{
		Value: "flg16"}
)

var (
	// gigiAmericanEnglishName represents gigi american english name.
	gigiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gigi"}

	// gigiCanadianFrenchName represents gigi canadian french name.
	gigiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gloria"}

	// gigiDutchName represents gigi dutch name.
	gigiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gigi"}

	// gigiFrenchName represents gigi french name.
	gigiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gloria"}

	// gigiGermanName represents gigi german name.
	gigiGermanName = nook.Name{
		Language: language.German,
		Value:    "Violetta"}

	// gigiItalianName represents gigi italian name.
	gigiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Raganà"}

	// gigiJapaneseName represents gigi japanese name.
	gigiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リンダ"}

	// gigiKoreanName represents gigi korean name.
	gigiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "린다"}

	// gigiLatinAmericanSpanishName represents gigi latin american spanish name.
	gigiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cleo"}

	// gigiRussianName represents gigi russian name.
	gigiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джиджи"}

	// gigiSimplifiedChineseName represents gigi simplified chinese name.
	gigiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "林妲"}

	// gigiSpanishName represents gigi spanish name.
	gigiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cleo"}

	// gigiTraditionalChineseName represents gigi traditional chinese name.
	gigiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "林妲"}
)

var (
	// gigiName represents gigi name.
	gigiName = nook.Languages{
		language.AmericanEnglish:      gigiAmericanEnglishName,
		language.CanadianFrench:       gigiCanadianFrenchName,
		language.Dutch:                gigiDutchName,
		language.French:               gigiFrenchName,
		language.German:               gigiGermanName,
		language.Italian:              gigiItalianName,
		language.Japanese:             gigiJapaneseName,
		language.Korean:               gigiKoreanName,
		language.LatinAmericanSpanish: gigiLatinAmericanSpanishName,
		language.Russian:              gigiRussianName,
		language.SimplifiedChinese:    gigiSimplifiedChineseName,
		language.Spanish:              gigiSpanishName,
		language.TraditionalChinese:   gigiTraditionalChineseName}
)

var (
	// gigiCharacter represents gigi character.
	gigiCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: gigiBirthday,
		Code:     gigiCode,
		Key:      character.Gigi,
		Gender:   gender.Female,
		Name:     gigiName,
		Special:  false}
)

var (
	// gigiAmericanEnglishPhrase represents gigi american english phrase.
	gigiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ribbette"}

	// gigiCanadianFrenchPhrase represents gigi canadian french phrase.
	gigiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crôazou"}

	// gigiDutchPhrase represents gigi dutch phrase.
	gigiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaakette"}

	// gigiFrenchPhrase represents gigi french phrase.
	gigiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crôazou"}

	// gigiGermanPhrase represents gigi german phrase.
	gigiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaki"}

	// gigiItalianPhrase represents gigi italian phrase.
	gigiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gurghigù"}

	// gigiJapanesePhrase represents gigi japanese phrase.
	gigiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ベイビィ"}

	// gigiKoreanPhrase represents gigi korean phrase.
	gigiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "베이비"}

	// gigiLatinAmericanSpanishPhrase represents gigi latin american spanish phrase.
	gigiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croá-croá"}

	// gigiRussianPhrase represents gigi russian phrase.
	gigiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квакет"}

	// gigiSimplifiedChinesePhrase represents gigi simplified chinese phrase.
	gigiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宝贝"}

	// gigiSpanishPhrase represents gigi spanish phrase.
	gigiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trebolito"}

	// gigiTraditionalChinesePhrase represents gigi traditional chinese phrase.
	gigiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "寶貝"}
)

var (
	// gigiPhrase represents gigi phrase.
	gigiPhrase = nook.Languages{
		language.AmericanEnglish:      gigiAmericanEnglishPhrase,
		language.CanadianFrench:       gigiCanadianFrenchPhrase,
		language.Dutch:                gigiDutchPhrase,
		language.French:               gigiFrenchPhrase,
		language.German:               gigiGermanPhrase,
		language.Italian:              gigiItalianPhrase,
		language.Japanese:             gigiJapanesePhrase,
		language.Korean:               gigiKoreanPhrase,
		language.LatinAmericanSpanish: gigiLatinAmericanSpanishPhrase,
		language.Russian:              gigiRussianPhrase,
		language.SimplifiedChinese:    gigiSimplifiedChinesePhrase,
		language.Spanish:              gigiSpanishPhrase,
		language.TraditionalChinese:   gigiTraditionalChinesePhrase}
)

var (
	// Gigi represents gigi.
	Gigi = nook.Villager{
		Character:   gigiCharacter,
		Personality: personality.Snooty,
		Phrase:      gigiPhrase}
)
