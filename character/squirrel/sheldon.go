package squirrel

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
	sheldonBirthday = nook.Birthday{
		Day:   26,
		Month: time.February}
)

var (
	sheldonCode = nook.Code{
		Value: "squ16"}
)

var (
	sheldonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sheldon"}

	sheldonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Roy"}

	sheldonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sheldon"}

	sheldonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Roy"}

	sheldonGermanName = nook.Name{
		Language: language.German,
		Value:    "Steffen"}

	sheldonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frugolo"}

	sheldonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリス"}

	sheldonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리스"}

	sheldonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Matracas"}

	sheldonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шелдон"}

	sheldonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "克栗斯"}

	sheldonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Matracas"}

	sheldonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "克栗斯"}
)

var (
	sheldonName = nook.Languages{
		language.AmericanEnglish:      sheldonAmericanEnglishName,
		language.CanadianFrench:       sheldonCanadianFrenchName,
		language.Dutch:                sheldonDutchName,
		language.French:               sheldonFrenchName,
		language.German:               sheldonGermanName,
		language.Italian:              sheldonItalianName,
		language.Japanese:             sheldonJapaneseName,
		language.Korean:               sheldonKoreanName,
		language.LatinAmericanSpanish: sheldonLatinAmericanSpanishName,
		language.Russian:              sheldonRussianName,
		language.SimplifiedChinese:    sheldonSimplifiedChineseName,
		language.Spanish:              sheldonSpanishName,
		language.TraditionalChinese:   sheldonTraditionalChineseName}
)

var (
	sheldonCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: sheldonBirthday,
		Code:     sheldonCode,
		Key:      character.Sheldon,
		Gender:   gender.Male,
		Name:     sheldonName,
		Special:  false}
)

var (
	sheldonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cardio"}

	sheldonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pacane"}

	sheldonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "topfit"}

	sheldonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toutouffe"}

	sheldonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoppi"}

	sheldonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flash"}

	sheldonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クリクリ"}

	sheldonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡글땡글"}

	sheldonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zasca"}

	sheldonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "треники"}

	sheldonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鼓栗鼓栗"}

	sheldonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zasca"}

	sheldonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鼓栗鼓栗"}
)

var (
	sheldonPhrase = nook.Languages{
		language.AmericanEnglish:      sheldonAmericanEnglishPhrase,
		language.CanadianFrench:       sheldonCanadianFrenchPhrase,
		language.Dutch:                sheldonDutchPhrase,
		language.French:               sheldonFrenchPhrase,
		language.German:               sheldonGermanPhrase,
		language.Italian:              sheldonItalianPhrase,
		language.Japanese:             sheldonJapanesePhrase,
		language.Korean:               sheldonKoreanPhrase,
		language.LatinAmericanSpanish: sheldonLatinAmericanSpanishPhrase,
		language.Russian:              sheldonRussianPhrase,
		language.SimplifiedChinese:    sheldonSimplifiedChinesePhrase,
		language.Spanish:              sheldonSpanishPhrase,
		language.TraditionalChinese:   sheldonTraditionalChinesePhrase}
)

var (
	Sheldon = nook.Villager{
		Character:   sheldonCharacter,
		Personality: personality.Jock,
		Phrase:      sheldonPhrase}
)
