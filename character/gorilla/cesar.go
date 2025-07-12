package gorilla

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
	// cesarBirthday represents cesar birthday.
	cesarBirthday = nook.Birthday{
		Day:   6,
		Month: time.September}
)

var (
	// cesarCode represents cesar code.
	cesarCode = nook.Code{
		Value: "gor00"}
)

var (
	// cesarAmericanEnglishName represents cesar american english name.
	cesarAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cesar"}

	// cesarCanadianFrenchName represents cesar canadian french name.
	cesarCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "César"}

	// cesarDutchName represents cesar dutch name.
	cesarDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cesar"}

	// cesarFrenchName represents cesar french name.
	cesarFrenchName = nook.Name{
		Language: language.French,
		Value:    "César"}

	// cesarGermanName represents cesar german name.
	cesarGermanName = nook.Name{
		Language: language.German,
		Value:    "Alfredo"}

	// cesarItalianName represents cesar italian name.
	cesarItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cesare"}

	// cesarJapaneseName represents cesar japanese name.
	cesarJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アラン"}

	// cesarKoreanName represents cesar korean name.
	cesarKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "앨런"}

	// cesarLatinAmericanSpanishName represents cesar latin american spanish name.
	cesarLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "César"}

	// cesarRussianName represents cesar russian name.
	cesarRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Цезарь"}

	// cesarSimplifiedChineseName represents cesar simplified chinese name.
	cesarSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿朗"}

	// cesarSpanishName represents cesar spanish name.
	cesarSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "César"}

	// cesarTraditionalChineseName represents cesar traditional chinese name.
	cesarTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿朗"}
)

var (
	// cesarName represents cesar name.
	cesarName = nook.Languages{
		language.AmericanEnglish:      cesarAmericanEnglishName,
		language.CanadianFrench:       cesarCanadianFrenchName,
		language.Dutch:                cesarDutchName,
		language.French:               cesarFrenchName,
		language.German:               cesarGermanName,
		language.Italian:              cesarItalianName,
		language.Japanese:             cesarJapaneseName,
		language.Korean:               cesarKoreanName,
		language.LatinAmericanSpanish: cesarLatinAmericanSpanishName,
		language.Russian:              cesarRussianName,
		language.SimplifiedChinese:    cesarSimplifiedChineseName,
		language.Spanish:              cesarSpanishName,
		language.TraditionalChinese:   cesarTraditionalChineseName}
)

var (
	// cesarCharacter represents cesar character.
	cesarCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: cesarBirthday,
		Code:     cesarCode,
		Key:      character.Cesar,
		Gender:   gender.Male,
		Name:     cesarName,
		Special:  false}
)

var (
	// cesarAmericanEnglishPhrase represents cesar american english phrase.
	cesarAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "highness"}

	// cesarCanadianFrenchPhrase represents cesar canadian french phrase.
	cesarCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brutus"}

	// cesarDutchPhrase represents cesar dutch phrase.
	cesarDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hoogheid"}

	// cesarFrenchPhrase represents cesar french phrase.
	cesarFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brutus"}

	// cesarGermanPhrase represents cesar german phrase.
	cesarGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoheit"}

	// cesarItalianPhrase represents cesar italian phrase.
	cesarItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dado"}

	// cesarJapanesePhrase represents cesar japanese phrase.
	cesarJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウホウホ"}

	// cesarKoreanPhrase represents cesar korean phrase.
	cesarKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우갸우갸"}

	// cesarLatinAmericanSpanishPhrase represents cesar latin american spanish phrase.
	cesarLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ejeeem"}

	// cesarRussianPhrase represents cesar russian phrase.
	cesarRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "высочество"}

	// cesarSimplifiedChinesePhrase represents cesar simplified chinese phrase.
	cesarSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拍胸脯"}

	// cesarSpanishPhrase represents cesar spanish phrase.
	cesarSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "monigote"}

	// cesarTraditionalChinesePhrase represents cesar traditional chinese phrase.
	cesarTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "拍胸脯"}
)

var (
	// cesarPhrase represents cesar phrase.
	cesarPhrase = nook.Languages{
		language.AmericanEnglish:      cesarAmericanEnglishPhrase,
		language.CanadianFrench:       cesarCanadianFrenchPhrase,
		language.Dutch:                cesarDutchPhrase,
		language.French:               cesarFrenchPhrase,
		language.German:               cesarGermanPhrase,
		language.Italian:              cesarItalianPhrase,
		language.Japanese:             cesarJapanesePhrase,
		language.Korean:               cesarKoreanPhrase,
		language.LatinAmericanSpanish: cesarLatinAmericanSpanishPhrase,
		language.Russian:              cesarRussianPhrase,
		language.SimplifiedChinese:    cesarSimplifiedChinesePhrase,
		language.Spanish:              cesarSpanishPhrase,
		language.TraditionalChinese:   cesarTraditionalChinesePhrase}
)

var (
	// Cesar represents cesar.
	Cesar = nook.Villager{
		Character:   cesarCharacter,
		Personality: personality.Cranky,
		Phrase:      cesarPhrase}
)
