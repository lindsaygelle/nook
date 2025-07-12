package cat

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
	// rudyBirthday represents rudy birthday.
	rudyBirthday = nook.Birthday{
		Day:   20,
		Month: time.December}
)

var (
	// rudyCode represents rudy code.
	rudyCode = nook.Code{
		Value: "cat20"}
)

var (
	// rudyAmericanEnglishName represents rudy american english name.
	rudyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rudy"}

	// rudyCanadianFrenchName represents rudy canadian french name.
	rudyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rougepif"}

	// rudyDutchName represents rudy dutch name.
	rudyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rudy"}

	// rudyFrenchName represents rudy french name.
	rudyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rougepif"}

	// rudyGermanName represents rudy german name.
	rudyGermanName = nook.Name{
		Language: language.German,
		Value:    "Heinz"}

	// rudyItalianName represents rudy italian name.
	rudyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gomitolo"}

	// rudyJapaneseName represents rudy japanese name.
	rudyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャス"}

	// rudyKoreanName represents rudy korean name.
	rudyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "찰스"}

	// rudyLatinAmericanSpanishName represents rudy latin american spanish name.
	rudyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rufino"}

	// rudyRussianName represents rudy russian name.
	rudyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Руди"}

	// rudySimplifiedChineseName represents rudy simplified chinese name.
	rudySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茶茶"}

	// rudySpanishName represents rudy spanish name.
	rudySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rufino"}

	// rudyTraditionalChineseName represents rudy traditional chinese name.
	rudyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茶茶"}
)

var (
	// rudyName represents rudy name.
	rudyName = nook.Languages{
		language.AmericanEnglish:      rudyAmericanEnglishName,
		language.CanadianFrench:       rudyCanadianFrenchName,
		language.Dutch:                rudyDutchName,
		language.French:               rudyFrenchName,
		language.German:               rudyGermanName,
		language.Italian:              rudyItalianName,
		language.Japanese:             rudyJapaneseName,
		language.Korean:               rudyKoreanName,
		language.LatinAmericanSpanish: rudyLatinAmericanSpanishName,
		language.Russian:              rudyRussianName,
		language.SimplifiedChinese:    rudySimplifiedChineseName,
		language.Spanish:              rudySpanishName,
		language.TraditionalChinese:   rudyTraditionalChineseName}
)

var (
	// rudyCharacter represents rudy character.
	rudyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: rudyBirthday,
		Code:     rudyCode,
		Key:      character.Rudy,
		Gender:   gender.Male,
		Name:     rudyName,
		Special:  false}
)

var (
	// rudyAmericanEnglishPhrase represents rudy american english phrase.
	rudyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mush"}

	// rudyCanadianFrenchPhrase represents rudy canadian french phrase.
	rudyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouah"}

	// rudyDutchPhrase represents rudy dutch phrase.
	rudyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "robijn"}

	// rudyFrenchPhrase represents rudy french phrase.
	rudyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "alatienne"}

	// rudyGermanPhrase represents rudy german phrase.
	rudyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "katzaaa"}

	// rudyItalianPhrase represents rudy italian phrase.
	rudyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ron ron"}

	// rudyJapanesePhrase represents rudy japanese phrase.
	rudyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とかナ"}

	// rudyKoreanPhrase represents rudy korean phrase.
	rudyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러거나"}

	// rudyLatinAmericanSpanishPhrase represents rudy latin american spanish phrase.
	rudyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cachusco"}

	// rudyRussianPhrase represents rudy russian phrase.
	rudyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шмяк"}

	// rudySimplifiedChinesePhrase represents rudy simplified chinese phrase.
	rudySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "之类的"}

	// rudySpanishPhrase represents rudy spanish phrase.
	rudySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cachusco"}

	// rudyTraditionalChinesePhrase represents rudy traditional chinese phrase.
	rudyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "之類的"}
)

var (
	// rudyPhrase represents rudy phrase.
	rudyPhrase = nook.Languages{
		language.AmericanEnglish:      rudyAmericanEnglishPhrase,
		language.CanadianFrench:       rudyCanadianFrenchPhrase,
		language.Dutch:                rudyDutchPhrase,
		language.French:               rudyFrenchPhrase,
		language.German:               rudyGermanPhrase,
		language.Italian:              rudyItalianPhrase,
		language.Japanese:             rudyJapanesePhrase,
		language.Korean:               rudyKoreanPhrase,
		language.LatinAmericanSpanish: rudyLatinAmericanSpanishPhrase,
		language.Russian:              rudyRussianPhrase,
		language.SimplifiedChinese:    rudySimplifiedChinesePhrase,
		language.Spanish:              rudySpanishPhrase,
		language.TraditionalChinese:   rudyTraditionalChinesePhrase}
)

var (
	// Rudy represents rudy.
	Rudy = nook.Villager{
		Character:   rudyCharacter,
		Personality: personality.Jock,
		Phrase:      rudyPhrase}
)
