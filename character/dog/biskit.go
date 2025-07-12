package dog

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
	// biskitBirthday represents biskit birthday.
	biskitBirthday = nook.Birthday{
		Day:   13,
		Month: time.May}
)

var (
	// biskitCode represents biskit code.
	biskitCode = nook.Code{
		Value: "dog03"}
)

var (
	// biskitAmericanEnglishName represents biskit american english name.
	biskitAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Biskit"}

	// biskitCanadianFrenchName represents biskit canadian french name.
	biskitCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crocket"}

	// biskitDutchName represents biskit dutch name.
	biskitDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Biskit"}

	// biskitFrenchName represents biskit french name.
	biskitFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crocket"}

	// biskitGermanName represents biskit german name.
	biskitGermanName = nook.Name{
		Language: language.German,
		Value:    "Keks"}

	// biskitItalianName represents biskit italian name.
	biskitItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Buffetto"}

	// biskitJapaneseName represents biskit japanese name.
	biskitJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロビン"}

	// biskitKoreanName represents biskit korean name.
	biskitKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로빈"}

	// biskitLatinAmericanSpanishName represents biskit latin american spanish name.
	biskitLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amnesio"}

	// biskitRussianName represents biskit russian name.
	biskitRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бискит"}

	// biskitSimplifiedChineseName represents biskit simplified chinese name.
	biskitSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗宾"}

	// biskitSpanishName represents biskit spanish name.
	biskitSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amnesio"}

	// biskitTraditionalChineseName represents biskit traditional chinese name.
	biskitTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅賓"}
)

var (
	// biskitName represents biskit name.
	biskitName = nook.Languages{
		language.AmericanEnglish:      biskitAmericanEnglishName,
		language.CanadianFrench:       biskitCanadianFrenchName,
		language.Dutch:                biskitDutchName,
		language.French:               biskitFrenchName,
		language.German:               biskitGermanName,
		language.Italian:              biskitItalianName,
		language.Japanese:             biskitJapaneseName,
		language.Korean:               biskitKoreanName,
		language.LatinAmericanSpanish: biskitLatinAmericanSpanishName,
		language.Russian:              biskitRussianName,
		language.SimplifiedChinese:    biskitSimplifiedChineseName,
		language.Spanish:              biskitSpanishName,
		language.TraditionalChinese:   biskitTraditionalChineseName}
)

var (
	// biskitCharacter represents biskit character.
	biskitCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: biskitBirthday,
		Code:     biskitCode,
		Key:      character.Biskit,
		Gender:   gender.Male,
		Name:     biskitName,
		Special:  false}
)

var (
	// biskitAmericanEnglishPhrase represents biskit american english phrase.
	biskitAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dawg"}

	// biskitCanadianFrenchPhrase represents biskit canadian french phrase.
	biskitCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "clair"}

	// biskitDutchPhrase represents biskit dutch phrase.
	biskitDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "makker"}

	// biskitFrenchPhrase represents biskit french phrase.
	biskitFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "clair"}

	// biskitGermanPhrase represents biskit german phrase.
	biskitGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kläff"}

	// biskitItalianPhrase represents biskit italian phrase.
	biskitItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrranco"}

	// biskitJapanesePhrase represents biskit japanese phrase.
	biskitJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だイヌ"}

	// biskitKoreanPhrase represents biskit korean phrase.
	biskitKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "멍멍"}

	// biskitLatinAmericanSpanishPhrase represents biskit latin american spanish phrase.
	biskitLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guau"}

	// biskitRussianPhrase represents biskit russian phrase.
	biskitRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пр-риятяв"}

	// biskitSimplifiedChinesePhrase represents biskit simplified chinese phrase.
	biskitSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狗狗"}

	// biskitSpanishPhrase represents biskit spanish phrase.
	biskitSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guau"}

	// biskitTraditionalChinesePhrase represents biskit traditional chinese phrase.
	biskitTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "狗狗"}
)

var (
	// biskitPhrase represents biskit phrase.
	biskitPhrase = nook.Languages{
		language.AmericanEnglish:      biskitAmericanEnglishPhrase,
		language.CanadianFrench:       biskitCanadianFrenchPhrase,
		language.Dutch:                biskitDutchPhrase,
		language.French:               biskitFrenchPhrase,
		language.German:               biskitGermanPhrase,
		language.Italian:              biskitItalianPhrase,
		language.Japanese:             biskitJapanesePhrase,
		language.Korean:               biskitKoreanPhrase,
		language.LatinAmericanSpanish: biskitLatinAmericanSpanishPhrase,
		language.Russian:              biskitRussianPhrase,
		language.SimplifiedChinese:    biskitSimplifiedChinesePhrase,
		language.Spanish:              biskitSpanishPhrase,
		language.TraditionalChinese:   biskitTraditionalChinesePhrase}
)

var (
	// Biskit represents biskit.
	Biskit = nook.Villager{
		Character:   biskitCharacter,
		Personality: personality.Lazy,
		Phrase:      biskitPhrase}
)
