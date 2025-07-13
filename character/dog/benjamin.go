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
	// benjaminBirthday represents benjamin birthday.
	benjaminBirthday = nook.Birthday{
		Day:   3,
		Month: time.August}
)

var (
	// benjaminCode represents benjamin code.
	benjaminCode = nook.Code{
		Value: "dog16"}
)

var (
	// benjaminAmericanEnglishName represents benjamin american english name.
	benjaminAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Benjamin"}

	// benjaminCanadianFrenchName represents benjamin canadian french name.
	benjaminCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bernardo"}

	// benjaminDutchName represents benjamin dutch name.
	benjaminDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Benjamin"}

	// benjaminFrenchName represents benjamin french name.
	benjaminFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bernardo"}

	// benjaminGermanName represents benjamin german name.
	benjaminGermanName = nook.Name{
		Language: language.German,
		Value:    "Wastl"}

	// benjaminItalianName represents benjamin italian name.
	benjaminItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bernardo"}

	// benjaminJapaneseName represents benjamin japanese name.
	benjaminJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハチ"}

	// benjaminKoreanName represents benjamin korean name.
	benjaminKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "땡칠이"}

	// benjaminLatinAmericanSpanishName represents benjamin latin american spanish name.
	benjaminLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bernardo"}

	// benjaminRussianName represents benjamin russian name.
	benjaminRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бернардо"}

	// benjaminSimplifiedChineseName represents benjamin simplified chinese name.
	benjaminSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小八"}

	// benjaminSpanishName represents benjamin spanish name.
	benjaminSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bernardo"}

	// benjaminTraditionalChineseName represents benjamin traditional chinese name.
	benjaminTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小八"}
)

var (
	// benjaminName represents benjamin name.
	benjaminName = nook.Languages{
		language.AmericanEnglish:      benjaminAmericanEnglishName,
		language.CanadianFrench:       benjaminCanadianFrenchName,
		language.Dutch:                benjaminDutchName,
		language.French:               benjaminFrenchName,
		language.German:               benjaminGermanName,
		language.Italian:              benjaminItalianName,
		language.Japanese:             benjaminJapaneseName,
		language.Korean:               benjaminKoreanName,
		language.LatinAmericanSpanish: benjaminLatinAmericanSpanishName,
		language.Russian:              benjaminRussianName,
		language.SimplifiedChinese:    benjaminSimplifiedChineseName,
		language.Spanish:              benjaminSpanishName,
		language.TraditionalChinese:   benjaminTraditionalChineseName}
)

var (
	// benjaminCharacter represents benjamin character.
	benjaminCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: benjaminBirthday,
		Code:     benjaminCode,
		Key:      character.Benjamin,
		Gender:   gender.Male,
		Name:     benjaminName,
		Special:  false}
)

var (
	// benjaminAmericanEnglishPhrase represents benjamin american english phrase.
	benjaminAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "alrighty"}

	// benjaminCanadianFrenchPhrase represents benjamin canadian french phrase.
	benjaminCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croquettes"}

	// benjaminDutchPhrase represents benjamin dutch phrase.
	benjaminDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oké dan"}

	// benjaminFrenchPhrase represents benjamin french phrase.
	benjaminFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croquettes"}

	// benjaminGermanPhrase represents benjamin german phrase.
	benjaminGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "winsel"}

	// benjaminItalianPhrase represents benjamin italian phrase.
	benjaminItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bauuau"}

	// benjaminJapanesePhrase represents benjamin japanese phrase.
	benjaminJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ではでは"}

	// benjaminKoreanPhrase represents benjamin korean phrase.
	benjaminKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럼그럼"}

	// benjaminLatinAmericanSpanishPhrase represents benjamin latin american spanish phrase.
	benjaminLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tirirí"}

	// benjaminRussianPhrase represents benjamin russian phrase.
	benjaminRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пор-рядок"}

	// benjaminSimplifiedChinesePhrase represents benjamin simplified chinese phrase.
	benjaminSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "那么那么"}

	// benjaminSpanishPhrase represents benjamin spanish phrase.
	benjaminSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tirirí"}

	// benjaminTraditionalChinesePhrase represents benjamin traditional chinese phrase.
	benjaminTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "那麼那麼"}
)

var (
	// benjaminPhrase represents benjamin phrase.
	benjaminPhrase = nook.Languages{
		language.AmericanEnglish:      benjaminAmericanEnglishPhrase,
		language.CanadianFrench:       benjaminCanadianFrenchPhrase,
		language.Dutch:                benjaminDutchPhrase,
		language.French:               benjaminFrenchPhrase,
		language.German:               benjaminGermanPhrase,
		language.Italian:              benjaminItalianPhrase,
		language.Japanese:             benjaminJapanesePhrase,
		language.Korean:               benjaminKoreanPhrase,
		language.LatinAmericanSpanish: benjaminLatinAmericanSpanishPhrase,
		language.Russian:              benjaminRussianPhrase,
		language.SimplifiedChinese:    benjaminSimplifiedChinesePhrase,
		language.Spanish:              benjaminSpanishPhrase,
		language.TraditionalChinese:   benjaminTraditionalChinesePhrase}
)

var (
	// Benjamin represents benjamin.
	Benjamin = nook.Villager{
		Character:   benjaminCharacter,
		Personality: personality.Lazy,
		Phrase:      benjaminPhrase}
)
