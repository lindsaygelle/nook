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
	// doraBirthday represents dora birthday.
	doraBirthday = nook.Birthday{
		Day:   18,
		Month: time.February}
)

var (
	// doraCode represents dora code.
	doraCode = nook.Code{
		Value: "mus00"}
)

var (
	// doraAmericanEnglishName represents dora american english name.
	doraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dora"}

	// doraCanadianFrenchName represents dora canadian french name.
	doraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dora"}

	// doraDutchName represents dora dutch name.
	doraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dora"}

	// doraFrenchName represents dora french name.
	doraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dora"}

	// doraGermanName represents dora german name.
	doraGermanName = nook.Name{
		Language: language.German,
		Value:    "Dora"}

	// doraItalianName represents dora italian name.
	doraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Serena"}

	// doraJapaneseName represents dora japanese name.
	doraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "とめ"}

	// doraKoreanName represents dora korean name.
	doraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티미"}

	// doraLatinAmericanSpanishName represents dora latin american spanish name.
	doraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dori"}

	// doraRussianName represents dora russian name.
	doraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дора"}

	// doraSimplifiedChineseName represents dora simplified chinese name.
	doraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "杜美"}

	// doraSpanishName represents dora spanish name.
	doraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dori"}

	// doraTraditionalChineseName represents dora traditional chinese name.
	doraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "杜美"}
)

var (
	// doraName represents dora name.
	doraName = nook.Languages{
		language.AmericanEnglish:      doraAmericanEnglishName,
		language.CanadianFrench:       doraCanadianFrenchName,
		language.Dutch:                doraDutchName,
		language.French:               doraFrenchName,
		language.German:               doraGermanName,
		language.Italian:              doraItalianName,
		language.Japanese:             doraJapaneseName,
		language.Korean:               doraKoreanName,
		language.LatinAmericanSpanish: doraLatinAmericanSpanishName,
		language.Russian:              doraRussianName,
		language.SimplifiedChinese:    doraSimplifiedChineseName,
		language.Spanish:              doraSpanishName,
		language.TraditionalChinese:   doraTraditionalChineseName}
)

var (
	// doraCharacter represents dora character.
	doraCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: doraBirthday,
		Code:     doraCode,
		Key:      character.Dora,
		Gender:   gender.Female,
		Name:     doraName,
		Special:  false}
)

var (
	// doraAmericanEnglishPhrase represents dora american english phrase.
	doraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squeaky"}

	// doraCanadianFrenchPhrase represents dora canadian french phrase.
	doraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coui-coui"}

	// doraDutchPhrase represents dora dutch phrase.
	doraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pieperig"}

	// doraFrenchPhrase represents dora french phrase.
	doraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tatouille"}

	// doraGermanPhrase represents dora german phrase.
	doraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quieek"}

	// doraItalianPhrase represents dora italian phrase.
	doraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squiksquik"}

	// doraJapanesePhrase represents dora japanese phrase.
	doraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だべ"}

	// doraKoreanPhrase represents dora korean phrase.
	doraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "찍찍"}

	// doraLatinAmericanSpanishPhrase represents dora latin american spanish phrase.
	doraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuiqui"}

	// doraRussianPhrase represents dora russian phrase.
	doraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пип-пип"}

	// doraSimplifiedChinesePhrase represents dora simplified chinese phrase.
	doraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美吧"}

	// doraSpanishPhrase represents dora spanish phrase.
	doraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuiqui"}

	// doraTraditionalChinesePhrase represents dora traditional chinese phrase.
	doraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美吧"}
)

var (
	// doraPhrase represents dora phrase.
	doraPhrase = nook.Languages{
		language.AmericanEnglish:      doraAmericanEnglishPhrase,
		language.CanadianFrench:       doraCanadianFrenchPhrase,
		language.Dutch:                doraDutchPhrase,
		language.French:               doraFrenchPhrase,
		language.German:               doraGermanPhrase,
		language.Italian:              doraItalianPhrase,
		language.Japanese:             doraJapanesePhrase,
		language.Korean:               doraKoreanPhrase,
		language.LatinAmericanSpanish: doraLatinAmericanSpanishPhrase,
		language.Russian:              doraRussianPhrase,
		language.SimplifiedChinese:    doraSimplifiedChinesePhrase,
		language.Spanish:              doraSpanishPhrase,
		language.TraditionalChinese:   doraTraditionalChinesePhrase}
)

var (
	// Dora represents dora.
	Dora = nook.Villager{
		Character:   doraCharacter,
		Personality: personality.Normal,
		Phrase:      doraPhrase}
)
