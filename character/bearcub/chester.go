package bearcub

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
	// chesterBirthday represents chester birthday.
	chesterBirthday = nook.Birthday{
		Day:   6,
		Month: time.August}
)

var (
	// chesterCode represents chester code.
	chesterCode = nook.Code{
		Value: "cbr15"}
)

var (
	// chesterAmericanEnglishName represents chester american english name.
	chesterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chester"}

	// chesterCanadianFrenchName represents chester canadian french name.
	chesterCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Placide"}

	// chesterDutchName represents chester dutch name.
	chesterDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chester"}

	// chesterFrenchName represents chester french name.
	chesterFrenchName = nook.Name{
		Language: language.French,
		Value:    "Placide"}

	// chesterGermanName represents chester german name.
	chesterGermanName = nook.Name{
		Language: language.German,
		Value:    "Eduard"}

	// chesterItalianName represents chester italian name.
	chesterItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Clemente"}

	// chesterJapaneseName represents chester japanese name.
	chesterJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パンタ"}

	// chesterKoreanName represents chester korean name.
	chesterKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "팬타"}

	// chesterLatinAmericanSpanishName represents chester latin american spanish name.
	chesterLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Osunio"}

	// chesterRussianName represents chester russian name.
	chesterRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Честер"}

	// chesterSimplifiedChineseName represents chester simplified chinese name.
	chesterSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胖达"}

	// chesterSpanishName represents chester spanish name.
	chesterSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Osunio"}

	// chesterTraditionalChineseName represents chester traditional chinese name.
	chesterTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "胖達"}
)

var (
	// chesterName represents chester name.
	chesterName = nook.Languages{
		language.AmericanEnglish:      chesterAmericanEnglishName,
		language.CanadianFrench:       chesterCanadianFrenchName,
		language.Dutch:                chesterDutchName,
		language.French:               chesterFrenchName,
		language.German:               chesterGermanName,
		language.Italian:              chesterItalianName,
		language.Japanese:             chesterJapaneseName,
		language.Korean:               chesterKoreanName,
		language.LatinAmericanSpanish: chesterLatinAmericanSpanishName,
		language.Russian:              chesterRussianName,
		language.SimplifiedChinese:    chesterSimplifiedChineseName,
		language.Spanish:              chesterSpanishName,
		language.TraditionalChinese:   chesterTraditionalChineseName}
)

var (
	// chesterCharacter represents chester character.
	chesterCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: chesterBirthday,
		Code:     chesterCode,
		Key:      character.Chester,
		Gender:   gender.Male,
		Name:     chesterName,
		Special:  false}
)

var (
	// chesterAmericanEnglishPhrase represents chester american english phrase.
	chesterAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rookie"}

	// chesterCanadianFrenchPhrase represents chester canadian french phrase.
	chesterCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pla-pla"}

	// chesterDutchPhrase represents chester dutch phrase.
	chesterDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bamboe"}

	// chesterFrenchPhrase represents chester french phrase.
	chesterFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pla-pla"}

	// chesterGermanPhrase represents chester german phrase.
	chesterGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "baaambus"}

	// chesterItalianPhrase represents chester italian phrase.
	chesterItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brupp"}

	// chesterJapanesePhrase represents chester japanese phrase.
	chesterJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だバンブ"}

	// chesterKoreanPhrase represents chester korean phrase.
	chesterKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부끄"}

	// chesterLatinAmericanSpanishPhrase represents chester latin american spanish phrase.
	chesterLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ahivá"}

	// chesterRussianPhrase represents chester russian phrase.
	chesterRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бамбук"}

	// chesterSimplifiedChinesePhrase represents chester simplified chinese phrase.
	chesterSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竹子"}

	// chesterSpanishPhrase represents chester spanish phrase.
	chesterSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ahivá"}

	// chesterTraditionalChinesePhrase represents chester traditional chinese phrase.
	chesterTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竹子"}
)

var (
	// chesterPhrase represents chester phrase.
	chesterPhrase = nook.Languages{
		language.AmericanEnglish:      chesterAmericanEnglishPhrase,
		language.CanadianFrench:       chesterCanadianFrenchPhrase,
		language.Dutch:                chesterDutchPhrase,
		language.French:               chesterFrenchPhrase,
		language.German:               chesterGermanPhrase,
		language.Italian:              chesterItalianPhrase,
		language.Japanese:             chesterJapanesePhrase,
		language.Korean:               chesterKoreanPhrase,
		language.LatinAmericanSpanish: chesterLatinAmericanSpanishPhrase,
		language.Russian:              chesterRussianPhrase,
		language.SimplifiedChinese:    chesterSimplifiedChinesePhrase,
		language.Spanish:              chesterSpanishPhrase,
		language.TraditionalChinese:   chesterTraditionalChinesePhrase}
)

var (
	// Chester represents chester.
	Chester = nook.Villager{
		Character:   chesterCharacter,
		Personality: personality.Lazy,
		Phrase:      chesterPhrase}
)
