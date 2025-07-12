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
	// roryBirthday represents rory birthday.
	roryBirthday = nook.Birthday{
		Day:   7,
		Month: time.August}
)

var (
	// roryCode represents rory code.
	roryCode = nook.Code{
		Value: "lon07"}
)

var (
	// roryAmericanEnglishName represents rory american english name.
	roryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rory"}

	// roryCanadianFrenchName represents rory canadian french name.
	roryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hercule"}

	// roryDutchName represents rory dutch name.
	roryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rory"}

	// roryFrenchName represents rory french name.
	roryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hercule"}

	// roryGermanName represents rory german name.
	roryGermanName = nook.Name{
		Language: language.German,
		Value:    "Leon"}

	// roryItalianName represents rory italian name.
	roryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ruggero"}

	// roryJapaneseName represents rory japanese name.
	roryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アーサー"}

	// roryKoreanName represents rory korean name.
	roryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아더"}

	// roryLatinAmericanSpanishName represents rory latin american spanish name.
	roryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rogelio"}

	// roryRussianName represents rory russian name.
	roryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рори"}

	// rorySimplifiedChineseName represents rory simplified chinese name.
	rorySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚瑟"}

	// rorySpanishName represents rory spanish name.
	rorySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rogelio"}

	// roryTraditionalChineseName represents rory traditional chinese name.
	roryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞瑟"}
)

var (
	// roryName represents rory name.
	roryName = nook.Languages{
		language.AmericanEnglish:      roryAmericanEnglishName,
		language.CanadianFrench:       roryCanadianFrenchName,
		language.Dutch:                roryDutchName,
		language.French:               roryFrenchName,
		language.German:               roryGermanName,
		language.Italian:              roryItalianName,
		language.Japanese:             roryJapaneseName,
		language.Korean:               roryKoreanName,
		language.LatinAmericanSpanish: roryLatinAmericanSpanishName,
		language.Russian:              roryRussianName,
		language.SimplifiedChinese:    rorySimplifiedChineseName,
		language.Spanish:              rorySpanishName,
		language.TraditionalChinese:   roryTraditionalChineseName}
)

var (
	// roryCharacter represents rory character.
	roryCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: roryBirthday,
		Code:     roryCode,
		Key:      character.Rory,
		Gender:   gender.Male,
		Name:     roryName,
		Special:  false}
)

var (
	// roryAmericanEnglishPhrase represents rory american english phrase.
	roryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "capital"}

	// roryCanadianFrenchPhrase represents rory canadian french phrase.
	roryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "excellent"}

	// roryDutchPhrase represents rory dutch phrase.
	roryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "koninklijk"}

	// roryFrenchPhrase represents rory french phrase.
	roryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "criniou"}

	// roryGermanPhrase represents rory german phrase.
	roryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "graaah"}

	// roryItalianPhrase represents rory italian phrase.
	roryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "roar"}

	// roryJapanesePhrase represents rory japanese phrase.
	roryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナハッ"}

	// roryKoreanPhrase represents rory korean phrase.
	roryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크릉"}

	// roryLatinAmericanSpanishPhrase represents rory latin american spanish phrase.
	roryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ticotaco"}

	// roryRussianPhrase represents rory russian phrase.
	roryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "капитально"}

	// rorySimplifiedChinesePhrase represents rory simplified chinese phrase.
	rorySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哪哈"}

	// rorySpanishPhrase represents rory spanish phrase.
	rorySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ticotaco"}

	// roryTraditionalChinesePhrase represents rory traditional chinese phrase.
	roryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哪哈"}
)

var (
	// roryPhrase represents rory phrase.
	roryPhrase = nook.Languages{
		language.AmericanEnglish:      roryAmericanEnglishPhrase,
		language.CanadianFrench:       roryCanadianFrenchPhrase,
		language.Dutch:                roryDutchPhrase,
		language.French:               roryFrenchPhrase,
		language.German:               roryGermanPhrase,
		language.Italian:              roryItalianPhrase,
		language.Japanese:             roryJapanesePhrase,
		language.Korean:               roryKoreanPhrase,
		language.LatinAmericanSpanish: roryLatinAmericanSpanishPhrase,
		language.Russian:              roryRussianPhrase,
		language.SimplifiedChinese:    rorySimplifiedChinesePhrase,
		language.Spanish:              rorySpanishPhrase,
		language.TraditionalChinese:   roryTraditionalChinesePhrase}
)

var (
	// Rory represents rory.
	Rory = nook.Villager{
		Character:   roryCharacter,
		Personality: personality.Jock,
		Phrase:      roryPhrase}
)
