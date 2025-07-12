package bull

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
	// chuckBirthday represents chuck birthday.
	chuckBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// chuckCode represents chuck code.
	chuckCode = nook.Code{
		Value: ""}
)

var (
	// chuckAmericanEnglishName represents chuck american english name.
	chuckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chuck"}

	// chuckCanadianFrenchName represents chuck canadian french name.
	chuckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// chuckDutchName represents chuck dutch name.
	chuckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// chuckFrenchName represents chuck french name.
	chuckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léon"}

	// chuckGermanName represents chuck german name.
	chuckGermanName = nook.Name{
		Language: language.German,
		Value:    "Paco"}

	// chuckItalianName represents chuck italian name.
	chuckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bullo"}

	// chuckJapaneseName represents chuck japanese name.
	chuckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビフテキ"}

	// chuckKoreanName represents chuck korean name.
	chuckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// chuckLatinAmericanSpanishName represents chuck latin american spanish name.
	chuckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// chuckRussianName represents chuck russian name.
	chuckRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// chuckSimplifiedChineseName represents chuck simplified chinese name.
	chuckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拉面"}

	// chuckSpanishName represents chuck spanish name.
	chuckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tauricio"}

	// chuckTraditionalChineseName represents chuck traditional chinese name.
	chuckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// chuckName represents chuck name.
	chuckName = nook.Languages{
		language.AmericanEnglish:      chuckAmericanEnglishName,
		language.CanadianFrench:       chuckCanadianFrenchName,
		language.Dutch:                chuckDutchName,
		language.French:               chuckFrenchName,
		language.German:               chuckGermanName,
		language.Italian:              chuckItalianName,
		language.Japanese:             chuckJapaneseName,
		language.Korean:               chuckKoreanName,
		language.LatinAmericanSpanish: chuckLatinAmericanSpanishName,
		language.Russian:              chuckRussianName,
		language.SimplifiedChinese:    chuckSimplifiedChineseName,
		language.Spanish:              chuckSpanishName,
		language.TraditionalChinese:   chuckTraditionalChineseName}
)

var (
	// chuckCharacter represents chuck character.
	chuckCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: chuckBirthday,
		Code:     chuckCode,
		Key:      character.Chuck,
		Gender:   gender.Male,
		Name:     chuckName,
		Special:  false}
)

var (
	// chuckAmericanEnglishPhrase represents chuck american english phrase.
	chuckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "jerky"}

	// chuckCanadianFrenchPhrase represents chuck canadian french phrase.
	chuckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// chuckDutchPhrase represents chuck dutch phrase.
	chuckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// chuckFrenchPhrase represents chuck french phrase.
	chuckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dugong"}

	// chuckGermanPhrase represents chuck german phrase.
	chuckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "daddel"}

	// chuckItalianPhrase represents chuck italian phrase.
	chuckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mummamia"}

	// chuckJapanesePhrase represents chuck japanese phrase.
	chuckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってんだ"}

	// chuckKoreanPhrase represents chuck korean phrase.
	chuckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// chuckLatinAmericanSpanishPhrase represents chuck latin american spanish phrase.
	chuckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// chuckRussianPhrase represents chuck russian phrase.
	chuckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// chuckSimplifiedChinesePhrase represents chuck simplified chinese phrase.
	chuckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伙计"}

	// chuckSpanishPhrase represents chuck spanish phrase.
	chuckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cobarde"}

	// chuckTraditionalChinesePhrase represents chuck traditional chinese phrase.
	chuckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// chuckPhrase represents chuck phrase.
	chuckPhrase = nook.Languages{
		language.AmericanEnglish:      chuckAmericanEnglishPhrase,
		language.CanadianFrench:       chuckCanadianFrenchPhrase,
		language.Dutch:                chuckDutchPhrase,
		language.French:               chuckFrenchPhrase,
		language.German:               chuckGermanPhrase,
		language.Italian:              chuckItalianPhrase,
		language.Japanese:             chuckJapanesePhrase,
		language.Korean:               chuckKoreanPhrase,
		language.LatinAmericanSpanish: chuckLatinAmericanSpanishPhrase,
		language.Russian:              chuckRussianPhrase,
		language.SimplifiedChinese:    chuckSimplifiedChinesePhrase,
		language.Spanish:              chuckSpanishPhrase,
		language.TraditionalChinese:   chuckTraditionalChinesePhrase}
)

var (
	// Chuck represents chuck.
	Chuck = nook.Villager{
		Character:   chuckCharacter,
		Personality: personality.Cranky,
		Phrase:      chuckPhrase}
)
