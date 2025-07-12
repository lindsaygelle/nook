package deer

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
	// lopezBirthday represents lopez birthday.
	lopezBirthday = nook.Birthday{
		Day:   20,
		Month: time.August}
)

var (
	// lopezCode represents lopez code.
	lopezCode = nook.Code{
		Value: "der05"}
)

var (
	// lopezAmericanEnglishName represents lopez american english name.
	lopezAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lopez"}

	// lopezCanadianFrenchName represents lopez canadian french name.
	lopezCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jon"}

	// lopezDutchName represents lopez dutch name.
	lopezDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lopez"}

	// lopezFrenchName represents lopez french name.
	lopezFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jon"}

	// lopezGermanName represents lopez german name.
	lopezGermanName = nook.Name{
		Language: language.German,
		Value:    "Eckart"}

	// lopezItalianName represents lopez italian name.
	lopezItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lopez"}

	// lopezJapaneseName represents lopez japanese name.
	lopezJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トムソン"}

	// lopezKoreanName represents lopez korean name.
	lopezKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "톰슨"}

	// lopezLatinAmericanSpanishName represents lopez latin american spanish name.
	lopezLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Agreste"}

	// lopezRussianName represents lopez russian name.
	lopezRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лопес"}

	// lopezSimplifiedChineseName represents lopez simplified chinese name.
	lopezSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "汤姆"}

	// lopezSpanishName represents lopez spanish name.
	lopezSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Agreste"}

	// lopezTraditionalChineseName represents lopez traditional chinese name.
	lopezTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "湯姆"}
)

var (
	// lopezName represents lopez name.
	lopezName = nook.Languages{
		language.AmericanEnglish:      lopezAmericanEnglishName,
		language.CanadianFrench:       lopezCanadianFrenchName,
		language.Dutch:                lopezDutchName,
		language.French:               lopezFrenchName,
		language.German:               lopezGermanName,
		language.Italian:              lopezItalianName,
		language.Japanese:             lopezJapaneseName,
		language.Korean:               lopezKoreanName,
		language.LatinAmericanSpanish: lopezLatinAmericanSpanishName,
		language.Russian:              lopezRussianName,
		language.SimplifiedChinese:    lopezSimplifiedChineseName,
		language.Spanish:              lopezSpanishName,
		language.TraditionalChinese:   lopezTraditionalChineseName}
)

var (
	// lopezCharacter represents lopez character.
	lopezCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: lopezBirthday,
		Code:     lopezCode,
		Key:      character.Lopez,
		Gender:   gender.Male,
		Name:     lopezName,
		Special:  false}
)

var (
	// lopezAmericanEnglishPhrase represents lopez american english phrase.
	lopezAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buckaroo"}

	// lopezCanadianFrenchPhrase represents lopez canadian french phrase.
	lopezCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "you bet"}

	// lopezDutchPhrase represents lopez dutch phrase.
	lopezDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bokje"}

	// lopezFrenchPhrase represents lopez french phrase.
	lopezFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "jaunisse"}

	// lopezGermanPhrase represents lopez german phrase.
	lopezGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bocko"}

	// lopezItalianPhrase represents lopez italian phrase.
	lopezItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sdleng"}

	// lopezJapanesePhrase represents lopez japanese phrase.
	lopezJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "がぜん"}

	// lopezKoreanPhrase represents lopez korean phrase.
	lopezKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엣헴"}

	// lopezLatinAmericanSpanishPhrase represents lopez latin american spanish phrase.
	lopezLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trisca"}

	// lopezRussianPhrase represents lopez russian phrase.
	lopezRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бубух"}

	// lopezSimplifiedChinesePhrase represents lopez simplified chinese phrase.
	lopezSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "变了"}

	// lopezSpanishPhrase represents lopez spanish phrase.
	lopezSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trisca"}

	// lopezTraditionalChinesePhrase represents lopez traditional chinese phrase.
	lopezTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "變了"}
)

var (
	// lopezPhrase represents lopez phrase.
	lopezPhrase = nook.Languages{
		language.AmericanEnglish:      lopezAmericanEnglishPhrase,
		language.CanadianFrench:       lopezCanadianFrenchPhrase,
		language.Dutch:                lopezDutchPhrase,
		language.French:               lopezFrenchPhrase,
		language.German:               lopezGermanPhrase,
		language.Italian:              lopezItalianPhrase,
		language.Japanese:             lopezJapanesePhrase,
		language.Korean:               lopezKoreanPhrase,
		language.LatinAmericanSpanish: lopezLatinAmericanSpanishPhrase,
		language.Russian:              lopezRussianPhrase,
		language.SimplifiedChinese:    lopezSimplifiedChinesePhrase,
		language.Spanish:              lopezSpanishPhrase,
		language.TraditionalChinese:   lopezTraditionalChinesePhrase}
)

var (
	// Lopez represents lopez.
	Lopez = nook.Villager{
		Character:   lopezCharacter,
		Personality: personality.Smug,
		Phrase:      lopezPhrase}
)
