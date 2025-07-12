package duck

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
	// billBirthday represents bill birthday.
	billBirthday = nook.Birthday{
		Day:   1,
		Month: time.February}
)

var (
	// billCode represents bill code.
	billCode = nook.Code{
		Value: "duk00"}
)

var (
	// billAmericanEnglishName represents bill american english name.
	billAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bill"}

	// billCanadianFrenchName represents bill canadian french name.
	billCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Choco"}

	// billDutchName represents bill dutch name.
	billDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bill"}

	// billFrenchName represents bill french name.
	billFrenchName = nook.Name{
		Language: language.French,
		Value:    "Choco"}

	// billGermanName represents bill german name.
	billGermanName = nook.Name{
		Language: language.German,
		Value:    "Bill"}

	// billItalianName represents bill italian name.
	billItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gino"}

	// billJapaneseName represents bill japanese name.
	billJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピータン"}

	// billKoreanName represents bill korean name.
	billKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "코코아"}

	// billLatinAmericanSpanishName represents bill latin american spanish name.
	billLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paquito"}

	// billRussianName represents bill russian name.
	billRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Билл"}

	// billSimplifiedChineseName represents bill simplified chinese name.
	billSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮蛋"}

	// billSpanishName represents bill spanish name.
	billSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paquito"}

	// billTraditionalChineseName represents bill traditional chinese name.
	billTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮蛋"}
)

var (
	// billName represents bill name.
	billName = nook.Languages{
		language.AmericanEnglish:      billAmericanEnglishName,
		language.CanadianFrench:       billCanadianFrenchName,
		language.Dutch:                billDutchName,
		language.French:               billFrenchName,
		language.German:               billGermanName,
		language.Italian:              billItalianName,
		language.Japanese:             billJapaneseName,
		language.Korean:               billKoreanName,
		language.LatinAmericanSpanish: billLatinAmericanSpanishName,
		language.Russian:              billRussianName,
		language.SimplifiedChinese:    billSimplifiedChineseName,
		language.Spanish:              billSpanishName,
		language.TraditionalChinese:   billTraditionalChineseName}
)

var (
	// billCharacter represents bill character.
	billCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: billBirthday,
		Code:     billCode,
		Key:      character.Bill,
		Gender:   gender.Male,
		Name:     billName,
		Special:  false}
)

var (
	// billAmericanEnglishPhrase represents bill american english phrase.
	billAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quacko"}

	// billCanadianFrenchPhrase represents bill canadian french phrase.
	billCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "choupichou"}

	// billDutchPhrase represents bill dutch phrase.
	billDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaker"}

	// billFrenchPhrase represents bill french phrase.
	billFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "choupichou"}

	// billGermanPhrase represents bill german phrase.
	billGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quako"}

	// billItalianPhrase represents bill italian phrase.
	billItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaquo"}

	// billJapanesePhrase represents bill japanese phrase.
	billJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だね"}

	// billKoreanPhrase represents bill korean phrase.
	billKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "탕"}

	// billLatinAmericanSpanishPhrase represents bill latin american spanish phrase.
	billLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuaco"}

	// billRussianPhrase represents bill russian phrase.
	billRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряко"}

	// billSimplifiedChinesePhrase represents bill simplified chinese phrase.
	billSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蛋是啊"}

	// billSpanishPhrase represents bill spanish phrase.
	billSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuaco"}

	// billTraditionalChinesePhrase represents bill traditional chinese phrase.
	billTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蛋是啊"}
)

var (
	// billPhrase represents bill phrase.
	billPhrase = nook.Languages{
		language.AmericanEnglish:      billAmericanEnglishPhrase,
		language.CanadianFrench:       billCanadianFrenchPhrase,
		language.Dutch:                billDutchPhrase,
		language.French:               billFrenchPhrase,
		language.German:               billGermanPhrase,
		language.Italian:              billItalianPhrase,
		language.Japanese:             billJapanesePhrase,
		language.Korean:               billKoreanPhrase,
		language.LatinAmericanSpanish: billLatinAmericanSpanishPhrase,
		language.Russian:              billRussianPhrase,
		language.SimplifiedChinese:    billSimplifiedChinesePhrase,
		language.Spanish:              billSpanishPhrase,
		language.TraditionalChinese:   billTraditionalChinesePhrase}
)

var (
	// Bill represents bill.
	Bill = nook.Villager{
		Character:   billCharacter,
		Personality: personality.Jock,
		Phrase:      billPhrase}
)
