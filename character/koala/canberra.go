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
	// canberraBirthday represents canberra birthday.
	canberraBirthday = nook.Birthday{
		Day:   14,
		Month: time.May}
)

var (
	// canberraCode represents canberra code.
	canberraCode = nook.Code{
		Value: "kal08"}
)

var (
	// canberraAmericanEnglishName represents canberra american english name.
	canberraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Canberra"}

	// canberraCanadianFrenchName represents canberra canadian french name.
	canberraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kolala"}

	// canberraDutchName represents canberra dutch name.
	canberraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Canberra"}

	// canberraFrenchName represents canberra french name.
	canberraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kolala"}

	// canberraGermanName represents canberra german name.
	canberraGermanName = nook.Name{
		Language: language.German,
		Value:    "Caroline"}

	// canberraItalianName represents canberra italian name.
	canberraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Canberra"}

	// canberraJapaneseName represents canberra japanese name.
	canberraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンベラ"}

	// canberraKoreanName represents canberra korean name.
	canberraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캔버라"}

	// canberraLatinAmericanSpanishName represents canberra latin american spanish name.
	canberraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Canberra"}

	// canberraRussianName represents canberra russian name.
	canberraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Канберра"}

	// canberraSimplifiedChineseName represents canberra simplified chinese name.
	canberraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "简培菈"}

	// canberraSpanishName represents canberra spanish name.
	canberraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Canberra"}

	// canberraTraditionalChineseName represents canberra traditional chinese name.
	canberraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "簡培菈"}
)

var (
	// canberraName represents canberra name.
	canberraName = nook.Languages{
		language.AmericanEnglish:      canberraAmericanEnglishName,
		language.CanadianFrench:       canberraCanadianFrenchName,
		language.Dutch:                canberraDutchName,
		language.French:               canberraFrenchName,
		language.German:               canberraGermanName,
		language.Italian:              canberraItalianName,
		language.Japanese:             canberraJapaneseName,
		language.Korean:               canberraKoreanName,
		language.LatinAmericanSpanish: canberraLatinAmericanSpanishName,
		language.Russian:              canberraRussianName,
		language.SimplifiedChinese:    canberraSimplifiedChineseName,
		language.Spanish:              canberraSpanishName,
		language.TraditionalChinese:   canberraTraditionalChineseName}
)

var (
	// canberraCharacter represents canberra character.
	canberraCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: canberraBirthday,
		Code:     canberraCode,
		Key:      character.Canberra,
		Gender:   gender.Female,
		Name:     canberraName,
		Special:  false}
)

var (
	// canberraAmericanEnglishPhrase represents canberra american english phrase.
	canberraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nuh uh"}

	// canberraCanadianFrenchPhrase represents canberra canadian french phrase.
	canberraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zyva"}

	// canberraDutchPhrase represents canberra dutch phrase.
	canberraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "echnie"}

	// canberraFrenchPhrase represents canberra french phrase.
	canberraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zyva"}

	// canberraGermanPhrase represents canberra german phrase.
	canberraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kolala"}

	// canberraItalianPhrase represents canberra italian phrase.
	canberraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "eurekalì"}

	// canberraJapanesePhrase represents canberra japanese phrase.
	canberraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "え～"}

	// canberraKoreanPhrase represents canberra korean phrase.
	canberraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어어"}

	// canberraLatinAmericanSpanishPhrase represents canberra latin american spanish phrase.
	canberraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "koalalá"}

	// canberraRussianPhrase represents canberra russian phrase.
	canberraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ох-хо"}

	// canberraSimplifiedChinesePhrase represents canberra simplified chinese phrase.
	canberraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咦"}

	// canberraSpanishPhrase represents canberra spanish phrase.
	canberraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kolalá"}

	// canberraTraditionalChinesePhrase represents canberra traditional chinese phrase.
	canberraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咦～"}
)

var (
	// canberraPhrase represents canberra phrase.
	canberraPhrase = nook.Languages{
		language.AmericanEnglish:      canberraAmericanEnglishPhrase,
		language.CanadianFrench:       canberraCanadianFrenchPhrase,
		language.Dutch:                canberraDutchPhrase,
		language.French:               canberraFrenchPhrase,
		language.German:               canberraGermanPhrase,
		language.Italian:              canberraItalianPhrase,
		language.Japanese:             canberraJapanesePhrase,
		language.Korean:               canberraKoreanPhrase,
		language.LatinAmericanSpanish: canberraLatinAmericanSpanishPhrase,
		language.Russian:              canberraRussianPhrase,
		language.SimplifiedChinese:    canberraSimplifiedChinesePhrase,
		language.Spanish:              canberraSpanishPhrase,
		language.TraditionalChinese:   canberraTraditionalChinesePhrase}
)

var (
	// Canberra represents canberra.
	Canberra = nook.Villager{
		Character:   canberraCharacter,
		Personality: personality.BigSister,
		Phrase:      canberraPhrase}
)
