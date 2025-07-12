package dog

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
	// megumiBirthday represents megumi birthday.
	megumiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// megumiCode represents megumi code.
	megumiCode = nook.Code{
		Value: ""}
)

var (
	// megumiAmericanEnglishName represents megumi american english name.
	megumiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Megumi"}

	// megumiCanadianFrenchName represents megumi canadian french name.
	megumiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// megumiDutchName represents megumi dutch name.
	megumiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// megumiFrenchName represents megumi french name.
	megumiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// megumiGermanName represents megumi german name.
	megumiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// megumiItalianName represents megumi italian name.
	megumiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// megumiJapaneseName represents megumi japanese name.
	megumiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メグミ"}

	// megumiKoreanName represents megumi korean name.
	megumiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// megumiLatinAmericanSpanishName represents megumi latin american spanish name.
	megumiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// megumiRussianName represents megumi russian name.
	megumiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// megumiSimplifiedChineseName represents megumi simplified chinese name.
	megumiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// megumiSpanishName represents megumi spanish name.
	megumiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// megumiTraditionalChineseName represents megumi traditional chinese name.
	megumiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// megumiName represents megumi name.
	megumiName = nook.Languages{
		language.AmericanEnglish:      megumiAmericanEnglishName,
		language.CanadianFrench:       megumiCanadianFrenchName,
		language.Dutch:                megumiDutchName,
		language.French:               megumiFrenchName,
		language.German:               megumiGermanName,
		language.Italian:              megumiItalianName,
		language.Japanese:             megumiJapaneseName,
		language.Korean:               megumiKoreanName,
		language.LatinAmericanSpanish: megumiLatinAmericanSpanishName,
		language.Russian:              megumiRussianName,
		language.SimplifiedChinese:    megumiSimplifiedChineseName,
		language.Spanish:              megumiSpanishName,
		language.TraditionalChinese:   megumiTraditionalChineseName}
)

var (
	// megumiCharacter represents megumi character.
	megumiCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: megumiBirthday,
		Code:     megumiCode,
		Key:      character.Megumi,
		Gender:   gender.Female,
		Name:     megumiName,
		Special:  false}
)

var (
	// megumiAmericanEnglishPhrase represents megumi american english phrase.
	megumiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "きゅん"}

	// megumiCanadianFrenchPhrase represents megumi canadian french phrase.
	megumiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// megumiDutchPhrase represents megumi dutch phrase.
	megumiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// megumiFrenchPhrase represents megumi french phrase.
	megumiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// megumiGermanPhrase represents megumi german phrase.
	megumiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// megumiItalianPhrase represents megumi italian phrase.
	megumiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// megumiJapanesePhrase represents megumi japanese phrase.
	megumiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゅん"}

	// megumiKoreanPhrase represents megumi korean phrase.
	megumiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// megumiLatinAmericanSpanishPhrase represents megumi latin american spanish phrase.
	megumiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// megumiRussianPhrase represents megumi russian phrase.
	megumiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// megumiSimplifiedChinesePhrase represents megumi simplified chinese phrase.
	megumiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// megumiSpanishPhrase represents megumi spanish phrase.
	megumiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// megumiTraditionalChinesePhrase represents megumi traditional chinese phrase.
	megumiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// megumiPhrase represents megumi phrase.
	megumiPhrase = nook.Languages{
		language.AmericanEnglish:      megumiAmericanEnglishPhrase,
		language.CanadianFrench:       megumiCanadianFrenchPhrase,
		language.Dutch:                megumiDutchPhrase,
		language.French:               megumiFrenchPhrase,
		language.German:               megumiGermanPhrase,
		language.Italian:              megumiItalianPhrase,
		language.Japanese:             megumiJapanesePhrase,
		language.Korean:               megumiKoreanPhrase,
		language.LatinAmericanSpanish: megumiLatinAmericanSpanishPhrase,
		language.Russian:              megumiRussianPhrase,
		language.SimplifiedChinese:    megumiSimplifiedChinesePhrase,
		language.Spanish:              megumiSpanishPhrase,
		language.TraditionalChinese:   megumiTraditionalChinesePhrase}
)

var (
	// Megumi represents megumi.
	Megumi = nook.Villager{
		Character:   megumiCharacter,
		Personality: personality.Peppy,
		Phrase:      megumiPhrase}
)
