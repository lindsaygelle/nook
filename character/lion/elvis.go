package lion

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
	// elvisBirthday represents elvis birthday.
	elvisBirthday = nook.Birthday{
		Day:   23,
		Month: time.July}
)

var (
	// elvisCode represents elvis code.
	elvisCode = nook.Code{
		Value: "lon01"}
)

var (
	// elvisAmericanEnglishName represents elvis american english name.
	elvisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elvis"}

	// elvisCanadianFrenchName represents elvis canadian french name.
	elvisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Elvis"}

	// elvisDutchName represents elvis dutch name.
	elvisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elvis"}

	// elvisFrenchName represents elvis french name.
	elvisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Elvis"}

	// elvisGermanName represents elvis german name.
	elvisGermanName = nook.Name{
		Language: language.German,
		Value:    "Leonardo"}

	// elvisItalianName represents elvis italian name.
	elvisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elvis"}

	// elvisJapaneseName represents elvis japanese name.
	elvisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キング"}

	// elvisKoreanName represents elvis korean name.
	elvisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "킹"}

	// elvisLatinAmericanSpanishName represents elvis latin american spanish name.
	elvisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Elvis"}

	// elvisRussianName represents elvis russian name.
	elvisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элвис"}

	// elvisSimplifiedChineseName represents elvis simplified chinese name.
	elvisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皇狮"}

	// elvisSpanishName represents elvis spanish name.
	elvisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elvis"}

	// elvisTraditionalChineseName represents elvis traditional chinese name.
	elvisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皇獅"}
)

var (
	// elvisName represents elvis name.
	elvisName = nook.Languages{
		language.AmericanEnglish:      elvisAmericanEnglishName,
		language.CanadianFrench:       elvisCanadianFrenchName,
		language.Dutch:                elvisDutchName,
		language.French:               elvisFrenchName,
		language.German:               elvisGermanName,
		language.Italian:              elvisItalianName,
		language.Japanese:             elvisJapaneseName,
		language.Korean:               elvisKoreanName,
		language.LatinAmericanSpanish: elvisLatinAmericanSpanishName,
		language.Russian:              elvisRussianName,
		language.SimplifiedChinese:    elvisSimplifiedChineseName,
		language.Spanish:              elvisSpanishName,
		language.TraditionalChinese:   elvisTraditionalChineseName}
)

var (
	// elvisCharacter represents elvis character.
	elvisCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: elvisBirthday,
		Code:     elvisCode,
		Key:      character.Elvis,
		Gender:   gender.Male,
		Name:     elvisName,
		Special:  false}
)

var (
	// elvisAmericanEnglishPhrase represents elvis american english phrase.
	elvisAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "unh-hunh"}

	// elvisCanadianFrenchPhrase represents elvis canadian french phrase.
	elvisCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bébé"}

	// elvisDutchPhrase represents elvis dutch phrase.
	elvisDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "aloha"}

	// elvisFrenchPhrase represents elvis french phrase.
	elvisFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bébé"}

	// elvisGermanPhrase represents elvis german phrase.
	elvisGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grolll"}

	// elvisItalianPhrase represents elvis italian phrase.
	elvisItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "unh-hunh"}

	// elvisJapanesePhrase represents elvis japanese phrase.
	elvisJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ダロガ"}

	// elvisKoreanPhrase represents elvis korean phrase.
	elvisKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "안그냐"}

	// elvisLatinAmericanSpanishPhrase represents elvis latin american spanish phrase.
	elvisLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "groar"}

	// elvisRussianPhrase represents elvis russian phrase.
	elvisRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "буги-вуги"}

	// elvisSimplifiedChinesePhrase represents elvis simplified chinese phrase.
	elvisSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "听懂吧"}

	// elvisSpanishPhrase represents elvis spanish phrase.
	elvisSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "groar"}

	// elvisTraditionalChinesePhrase represents elvis traditional chinese phrase.
	elvisTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聽懂吧"}
)

var (
	// elvisPhrase represents elvis phrase.
	elvisPhrase = nook.Languages{
		language.AmericanEnglish:      elvisAmericanEnglishPhrase,
		language.CanadianFrench:       elvisCanadianFrenchPhrase,
		language.Dutch:                elvisDutchPhrase,
		language.French:               elvisFrenchPhrase,
		language.German:               elvisGermanPhrase,
		language.Italian:              elvisItalianPhrase,
		language.Japanese:             elvisJapanesePhrase,
		language.Korean:               elvisKoreanPhrase,
		language.LatinAmericanSpanish: elvisLatinAmericanSpanishPhrase,
		language.Russian:              elvisRussianPhrase,
		language.SimplifiedChinese:    elvisSimplifiedChinesePhrase,
		language.Spanish:              elvisSpanishPhrase,
		language.TraditionalChinese:   elvisTraditionalChinesePhrase}
)

var (
	// Elvis represents elvis.
	Elvis = nook.Villager{
		Character:   elvisCharacter,
		Personality: personality.Cranky,
		Phrase:      elvisPhrase}
)
