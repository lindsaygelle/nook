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
	// bellaBirthday represents bella birthday.
	bellaBirthday = nook.Birthday{
		Day:   28,
		Month: time.December}
)

var (
	// bellaCode represents bella code.
	bellaCode = nook.Code{
		Value: "mus02"}
)

var (
	// bellaAmericanEnglishName represents bella american english name.
	bellaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bella"}

	// bellaCanadianFrenchName represents bella canadian french name.
	bellaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Belle"}

	// bellaDutchName represents bella dutch name.
	bellaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bella"}

	// bellaFrenchName represents bella french name.
	bellaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Belle"}

	// bellaGermanName represents bella german name.
	bellaGermanName = nook.Name{
		Language: language.German,
		Value:    "Susi"}

	// bellaItalianName represents bella italian name.
	bellaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bella"}

	// bellaJapaneseName represents bella japanese name.
	bellaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イザベラ"}

	// bellaKoreanName represents bella korean name.
	bellaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "이자벨"}

	// bellaLatinAmericanSpanishName represents bella latin american spanish name.
	bellaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Prity"}

	// bellaRussianName represents bella russian name.
	bellaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Белла"}

	// bellaSimplifiedChineseName represents bella simplified chinese name.
	bellaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝拉"}

	// bellaSpanishName represents bella spanish name.
	bellaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Prity"}

	// bellaTraditionalChineseName represents bella traditional chinese name.
	bellaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貝拉"}
)

var (
	// bellaName represents bella name.
	bellaName = nook.Languages{
		language.AmericanEnglish:      bellaAmericanEnglishName,
		language.CanadianFrench:       bellaCanadianFrenchName,
		language.Dutch:                bellaDutchName,
		language.French:               bellaFrenchName,
		language.German:               bellaGermanName,
		language.Italian:              bellaItalianName,
		language.Japanese:             bellaJapaneseName,
		language.Korean:               bellaKoreanName,
		language.LatinAmericanSpanish: bellaLatinAmericanSpanishName,
		language.Russian:              bellaRussianName,
		language.SimplifiedChinese:    bellaSimplifiedChineseName,
		language.Spanish:              bellaSpanishName,
		language.TraditionalChinese:   bellaTraditionalChineseName}
)

var (
	// bellaCharacter represents bella character.
	bellaCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: bellaBirthday,
		Code:     bellaCode,
		Key:      character.Bella,
		Gender:   gender.Female,
		Name:     bellaName,
		Special:  false}
)

var (
	// bellaAmericanEnglishPhrase represents bella american english phrase.
	bellaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eeks"}

	// bellaCanadianFrenchPhrase represents bella canadian french phrase.
	bellaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hiiiii"}

	// bellaDutchPhrase represents bella dutch phrase.
	bellaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ieps"}

	// bellaFrenchPhrase represents bella french phrase.
	bellaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hiiiii"}

	// bellaGermanPhrase represents bella german phrase.
	bellaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "visavisa"}

	// bellaItalianPhrase represents bella italian phrase.
	bellaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "eeks"}

	// bellaJapanesePhrase represents bella japanese phrase.
	bellaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ギャハッ"}

	// bellaKoreanPhrase represents bella korean phrase.
	bellaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "캬학"}

	// bellaLatinAmericanSpanishPhrase represents bella latin american spanish phrase.
	bellaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "petisú"}

	// bellaRussianPhrase represents bella russian phrase.
	bellaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пип"}

	// bellaSimplifiedChinesePhrase represents bella simplified chinese phrase.
	bellaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘎哈"}

	// bellaSpanishPhrase represents bella spanish phrase.
	bellaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "petisú"}

	// bellaTraditionalChinesePhrase represents bella traditional chinese phrase.
	bellaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘎哈"}
)

var (
	// bellaPhrase represents bella phrase.
	bellaPhrase = nook.Languages{
		language.AmericanEnglish:      bellaAmericanEnglishPhrase,
		language.CanadianFrench:       bellaCanadianFrenchPhrase,
		language.Dutch:                bellaDutchPhrase,
		language.French:               bellaFrenchPhrase,
		language.German:               bellaGermanPhrase,
		language.Italian:              bellaItalianPhrase,
		language.Japanese:             bellaJapanesePhrase,
		language.Korean:               bellaKoreanPhrase,
		language.LatinAmericanSpanish: bellaLatinAmericanSpanishPhrase,
		language.Russian:              bellaRussianPhrase,
		language.SimplifiedChinese:    bellaSimplifiedChinesePhrase,
		language.Spanish:              bellaSpanishPhrase,
		language.TraditionalChinese:   bellaTraditionalChinesePhrase}
)

var (
	// Bella represents bella.
	Bella = nook.Villager{
		Character:   bellaCharacter,
		Personality: personality.Peppy,
		Phrase:      bellaPhrase}
)
