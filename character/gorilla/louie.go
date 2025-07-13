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
	// louieBirthday represents louie birthday.
	louieBirthday = nook.Birthday{
		Day:   26,
		Month: time.March}
)

var (
	// louieCode represents louie code.
	louieCode = nook.Code{
		Value: "gor04"}
)

var (
	// louieAmericanEnglishName represents louie american english name.
	louieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Louie"}

	// louieCanadianFrenchName represents louie canadian french name.
	louieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Louis"}

	// louieDutchName represents louie dutch name.
	louieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Louie"}

	// louieFrenchName represents louie french name.
	louieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Louis"}

	// louieGermanName represents louie german name.
	louieGermanName = nook.Name{
		Language: language.German,
		Value:    "Ludwig"}

	// louieItalianName represents louie italian name.
	louieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lou"}

	// louieJapaneseName represents louie japanese name.
	louieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マッスル"}

	// louieKoreanName represents louie korean name.
	louieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "머슬"}

	// louieLatinAmericanSpanishName represents louie latin american spanish name.
	louieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lou"}

	// louieRussianName represents louie russian name.
	louieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Луи"}

	// louieSimplifiedChineseName represents louie simplified chinese name.
	louieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "壮壮"}

	// louieSpanishName represents louie spanish name.
	louieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lou"}

	// louieTraditionalChineseName represents louie traditional chinese name.
	louieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "壯壯"}
)

var (
	// louieName represents louie name.
	louieName = nook.Languages{
		language.AmericanEnglish:      louieAmericanEnglishName,
		language.CanadianFrench:       louieCanadianFrenchName,
		language.Dutch:                louieDutchName,
		language.French:               louieFrenchName,
		language.German:               louieGermanName,
		language.Italian:              louieItalianName,
		language.Japanese:             louieJapaneseName,
		language.Korean:               louieKoreanName,
		language.LatinAmericanSpanish: louieLatinAmericanSpanishName,
		language.Russian:              louieRussianName,
		language.SimplifiedChinese:    louieSimplifiedChineseName,
		language.Spanish:              louieSpanishName,
		language.TraditionalChinese:   louieTraditionalChineseName}
)

var (
	// louieCharacter represents louie character.
	louieCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: louieBirthday,
		Code:     louieCode,
		Key:      character.Louie,
		Gender:   gender.Male,
		Name:     louieName,
		Special:  false}
)

var (
	// louieAmericanEnglishPhrase represents louie american english phrase.
	louieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "toots"}

	// louieCanadianFrenchPhrase represents louie canadian french phrase.
	louieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tonneau"}

	// louieDutchPhrase represents louie dutch phrase.
	louieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oe-oe-ah"}

	// louieFrenchPhrase represents louie french phrase.
	louieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tonneau"}

	// louieGermanPhrase represents louie german phrase.
	louieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ugh-ugh"}

	// louieItalianPhrase represents louie italian phrase.
	louieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bum bum"}

	// louieJapanesePhrase represents louie japanese phrase.
	louieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "コング"}

	// louieKoreanPhrase represents louie korean phrase.
	louieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하압"}

	// louieLatinAmericanSpanishPhrase represents louie latin american spanish phrase.
	louieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gorilérico"}

	// louieRussianPhrase represents louie russian phrase.
	louieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-ух-ах"}

	// louieSimplifiedChinesePhrase represents louie simplified chinese phrase.
	louieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "钢钢"}

	// louieSpanishPhrase represents louie spanish phrase.
	louieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gorilérico"}

	// louieTraditionalChinesePhrase represents louie traditional chinese phrase.
	louieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鋼鋼"}
)

var (
	// louiePhrase represents louie phrase.
	louiePhrase = nook.Languages{
		language.AmericanEnglish:      louieAmericanEnglishPhrase,
		language.CanadianFrench:       louieCanadianFrenchPhrase,
		language.Dutch:                louieDutchPhrase,
		language.French:               louieFrenchPhrase,
		language.German:               louieGermanPhrase,
		language.Italian:              louieItalianPhrase,
		language.Japanese:             louieJapanesePhrase,
		language.Korean:               louieKoreanPhrase,
		language.LatinAmericanSpanish: louieLatinAmericanSpanishPhrase,
		language.Russian:              louieRussianPhrase,
		language.SimplifiedChinese:    louieSimplifiedChinesePhrase,
		language.Spanish:              louieSpanishPhrase,
		language.TraditionalChinese:   louieTraditionalChinesePhrase}
)

var (
	// Louie represents louie.
	Louie = nook.Villager{
		Character:   louieCharacter,
		Personality: personality.Jock,
		Phrase:      louiePhrase}
)
