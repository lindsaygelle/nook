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
	// moniqueBirthday represents monique birthday.
	moniqueBirthday = nook.Birthday{
		Day:   30,
		Month: time.September}
)

var (
	// moniqueCode represents monique code.
	moniqueCode = nook.Code{
		Value: "cat11"}
)

var (
	// moniqueAmericanEnglishName represents monique american english name.
	moniqueAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Monique"}

	// moniqueCanadianFrenchName represents monique canadian french name.
	moniqueCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Monique"}

	// moniqueDutchName represents monique dutch name.
	moniqueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Monique"}

	// moniqueFrenchName represents monique french name.
	moniqueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Monique"}

	// moniqueGermanName represents monique german name.
	moniqueGermanName = nook.Name{
		Language: language.German,
		Value:    "Monique"}

	// moniqueItalianName represents monique italian name.
	moniqueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marylin"}

	// moniqueJapaneseName represents monique japanese name.
	moniqueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジェーン"}

	// moniqueKoreanName represents monique korean name.
	moniqueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "제인"}

	// moniqueLatinAmericanSpanishName represents monique latin american spanish name.
	moniqueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Monique"}

	// moniqueRussianName represents monique russian name.
	moniqueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Моник"}

	// moniqueSimplifiedChineseName represents monique simplified chinese name.
	moniqueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珍珍"}

	// moniqueSpanishName represents monique spanish name.
	moniqueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Monique"}

	// moniqueTraditionalChineseName represents monique traditional chinese name.
	moniqueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珍珍"}
)

var (
	// moniqueName represents monique name.
	moniqueName = nook.Languages{
		language.AmericanEnglish:      moniqueAmericanEnglishName,
		language.CanadianFrench:       moniqueCanadianFrenchName,
		language.Dutch:                moniqueDutchName,
		language.French:               moniqueFrenchName,
		language.German:               moniqueGermanName,
		language.Italian:              moniqueItalianName,
		language.Japanese:             moniqueJapaneseName,
		language.Korean:               moniqueKoreanName,
		language.LatinAmericanSpanish: moniqueLatinAmericanSpanishName,
		language.Russian:              moniqueRussianName,
		language.SimplifiedChinese:    moniqueSimplifiedChineseName,
		language.Spanish:              moniqueSpanishName,
		language.TraditionalChinese:   moniqueTraditionalChineseName}
)

var (
	// moniqueCharacter represents monique character.
	moniqueCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: moniqueBirthday,
		Code:     moniqueCode,
		Key:      character.Monique,
		Gender:   gender.Female,
		Name:     moniqueName,
		Special:  false}
)

var (
	// moniqueAmericanEnglishPhrase represents monique american english phrase.
	moniqueAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pffffft"}

	// moniqueCanadianFrenchPhrase represents monique canadian french phrase.
	moniqueCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pffffft"}

	// moniqueDutchPhrase represents monique dutch phrase.
	moniqueDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pffffft"}

	// moniqueFrenchPhrase represents monique french phrase.
	moniqueFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pffffft"}

	// moniqueGermanPhrase represents monique german phrase.
	moniqueGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mrrrrrrr"}

	// moniqueItalianPhrase represents monique italian phrase.
	moniqueItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miaù"}

	// moniqueJapanesePhrase represents monique japanese phrase.
	moniqueJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウフーン"}

	// moniqueKoreanPhrase represents monique korean phrase.
	moniqueKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우"}

	// moniqueLatinAmericanSpanishPhrase represents monique latin american spanish phrase.
	moniqueLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ronrón"}

	// moniqueRussianPhrase represents monique russian phrase.
	moniqueRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фи-фи"}

	// moniqueSimplifiedChinesePhrase represents monique simplified chinese phrase.
	moniqueSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯哼"}

	// moniqueSpanishPhrase represents monique spanish phrase.
	moniqueSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ronrón"}

	// moniqueTraditionalChinesePhrase represents monique traditional chinese phrase.
	moniqueTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯哼"}
)

var (
	// moniquePhrase represents monique phrase.
	moniquePhrase = nook.Languages{
		language.AmericanEnglish:      moniqueAmericanEnglishPhrase,
		language.CanadianFrench:       moniqueCanadianFrenchPhrase,
		language.Dutch:                moniqueDutchPhrase,
		language.French:               moniqueFrenchPhrase,
		language.German:               moniqueGermanPhrase,
		language.Italian:              moniqueItalianPhrase,
		language.Japanese:             moniqueJapanesePhrase,
		language.Korean:               moniqueKoreanPhrase,
		language.LatinAmericanSpanish: moniqueLatinAmericanSpanishPhrase,
		language.Russian:              moniqueRussianPhrase,
		language.SimplifiedChinese:    moniqueSimplifiedChinesePhrase,
		language.Spanish:              moniqueSpanishPhrase,
		language.TraditionalChinese:   moniqueTraditionalChinesePhrase}
)

var (
	// Monique represents monique.
	Monique = nook.Villager{
		Character:   moniqueCharacter,
		Personality: personality.Snooty,
		Phrase:      moniquePhrase}
)
