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
	// sheldonBirthday represents Sheldon's birthday (February 26th).
	sheldonBirthday = nook.Birthday{
		Day:   26,
		Month: time.February}
)

var (
	// sheldonCode represents Sheldon's unique code ("squ16").
	sheldonCode = nook.Code{
		Value: "squ16"}
)

var (
	// sheldonAmericanEnglishName represents Sheldon's name in American English.
	sheldonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sheldon"}

	// sheldonCanadianFrenchName represents Sheldon's name in Canadian French.
	sheldonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Roy"}

	// sheldonDutchName represents Sheldon's name in Dutch.
	sheldonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sheldon"}

	// sheldonFrenchName represents Sheldon's name in French.
	sheldonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Roy"}

	// sheldonGermanName represents Sheldon's name in German.
	sheldonGermanName = nook.Name{
		Language: language.German,
		Value:    "Steffen"}

	// sheldonItalianName represents Sheldon's name in Italian.
	sheldonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frugolo"}

	// sheldonJapaneseName represents Sheldon's name in Japanese.
	sheldonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリス"}

	// sheldonKoreanName represents Sheldon's name in Korean.
	sheldonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리스"}

	// sheldonLatinAmericanSpanishName represents Sheldon's name in Latin American Spanish.
	sheldonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Matracas"}

	// sheldonRussianName represents Sheldon's name in Russian.
	sheldonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шелдон"}

	// sheldonSimplifiedChineseName represents Sheldon's name in Simplified Chinese.
	sheldonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "克栗斯"}

	// sheldonSpanishName represents Sheldon's name in Spanish.
	sheldonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Matracas"}

	// sheldonTraditionalChineseName represents Sheldon's name in Traditional Chinese.
	sheldonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "克栗斯"}
)

var (
	// sheldonName represents Sheldon's name in different languages.
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
	// sheldonCharacter represents Sheldon's character information.
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
	// sheldonAmericanEnglishPhrase represents Sheldon's phrase in American English.
	sheldonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cardio"}

	// sheldonCanadianFrenchPhrase represents Sheldon's phrase in Canadian French.
	sheldonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pacane"}

	// sheldonDutchPhrase represents Sheldon's phrase in Dutch.
	sheldonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "topfit"}

	// sheldonFrenchPhrase represents Sheldon's phrase in French.
	sheldonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toutouffe"}

	// sheldonGermanPhrase represents Sheldon's phrase in German.
	sheldonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoppi"}

	// sheldonItalianPhrase represents Sheldon's phrase in Italian.
	sheldonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flash"}

	// sheldonJapanesePhrase represents Sheldon's phrase in Japanese.
	sheldonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クリクリ"}

	// sheldonKoreanPhrase represents Sheldon's phrase in Korean.
	sheldonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡글땡글"}

	// sheldonLatinAmericanSpanishPhrase represents Sheldon's phrase in Latin American Spanish.
	sheldonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zasca"}

	// sheldonRussianPhrase represents Sheldon's phrase in Russian.
	sheldonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "треники"}

	// sheldonSimplifiedChinesePhrase represents Sheldon's phrase in Simplified Chinese.
	sheldonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鼓栗鼓栗"}

	// sheldonSpanishPhrase represents Sheldon's phrase in Spanish.
	sheldonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zasca"}

	// sheldonTraditionalChinesePhrase represents Sheldon's phrase in Traditional Chinese.
	sheldonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鼓栗鼓栗"}
)

var (
	// sheldonPhrase represents Sheldon's phrases in different languages.
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
	// Sheldon represents the character Sheldon with his complete information.
	Sheldon = nook.Villager{
		Character:   sheldonCharacter,
		Personality: personality.Jock,
		Phrase:      sheldonPhrase}
)
