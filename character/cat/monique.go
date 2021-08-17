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
	moniqueBirthday = nook.Birthday{
		Day:   30,
		Month: time.September}
)

var (
	moniqueCode = nook.Code{
		Value: "cat11"}
)

var (
	moniqueAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Monique"}

	moniqueCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Monique"}

	moniqueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Monique"}

	moniqueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Monique"}

	moniqueGermanName = nook.Name{
		Language: language.German,
		Value:    "Monique"}

	moniqueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marylin"}

	moniqueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジェーン"}

	moniqueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "제인"}

	moniqueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Monique"}

	moniqueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Моник"}

	moniqueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珍珍"}

	moniqueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Monique"}

	moniqueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珍珍"}
)

var (
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
	moniqueCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: moniqueBirthday,
		Code:     moniqueCode,
		Key:      character.Monique,
		Gender:   gender.Female,
		Name:     moniqueName}
)

var (
	moniqueAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pffffft"}

	moniqueCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pffffft"}

	moniqueDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pffffft"}

	moniqueFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pffffft"}

	moniqueGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mrrrrrrr"}

	moniqueItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miaù"}

	moniqueJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウフーン"}

	moniqueKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우"}

	moniqueLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ronrón"}

	moniqueRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фи-фи"}

	moniqueSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯哼"}

	moniqueSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ronrón"}

	moniqueTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯哼"}
)

var (
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
	Monique = nook.Villager{
		Character:   moniqueCharacter,
		Personality: personality.Snooty,
		Phrase:      moniquePhrase}
)
