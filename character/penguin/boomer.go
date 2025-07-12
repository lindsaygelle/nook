package penguin

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
	// boomerBirthday represents boomer birthday.
	boomerBirthday = nook.Birthday{
		Day:   7,
		Month: time.February}
)

var (
	// boomerCode represents boomer code.
	boomerCode = nook.Code{
		Value: "pgn10"}
)

var (
	// boomerAmericanEnglishName represents boomer american english name.
	boomerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boomer"}

	// boomerCanadianFrenchName represents boomer canadian french name.
	boomerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ethan"}

	// boomerDutchName represents boomer dutch name.
	boomerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boomer"}

	// boomerFrenchName represents boomer french name.
	boomerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ethan"}

	// boomerGermanName represents boomer german name.
	boomerGermanName = nook.Name{
		Language: language.German,
		Value:    "Max"}

	// boomerItalianName represents boomer italian name.
	boomerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Icaro"}

	// boomerJapaneseName represents boomer japanese name.
	boomerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ショーイ"}

	// boomerKoreanName represents boomer korean name.
	boomerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "팽기"}

	// boomerLatinAmericanSpanishName represents boomer latin american spanish name.
	boomerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Serafino"}

	// boomerRussianName represents boomer russian name.
	boomerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бумер"}

	// boomerSimplifiedChineseName represents boomer simplified chinese name.
	boomerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "秀益"}

	// boomerSpanishName represents boomer spanish name.
	boomerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Serafino"}

	// boomerTraditionalChineseName represents boomer traditional chinese name.
	boomerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "秀益"}
)

var (
	// boomerName represents boomer name.
	boomerName = nook.Languages{
		language.AmericanEnglish:      boomerAmericanEnglishName,
		language.CanadianFrench:       boomerCanadianFrenchName,
		language.Dutch:                boomerDutchName,
		language.French:               boomerFrenchName,
		language.German:               boomerGermanName,
		language.Italian:              boomerItalianName,
		language.Japanese:             boomerJapaneseName,
		language.Korean:               boomerKoreanName,
		language.LatinAmericanSpanish: boomerLatinAmericanSpanishName,
		language.Russian:              boomerRussianName,
		language.SimplifiedChinese:    boomerSimplifiedChineseName,
		language.Spanish:              boomerSpanishName,
		language.TraditionalChinese:   boomerTraditionalChineseName}
)

var (
	// boomerCharacter represents boomer character.
	boomerCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: boomerBirthday,
		Code:     boomerCode,
		Key:      character.Boomer,
		Gender:   gender.Male,
		Name:     boomerName,
		Special:  false}
)

var (
	// boomerAmericanEnglishPhrase represents boomer american english phrase.
	boomerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "human"}

	// boomerCanadianFrenchPhrase represents boomer canadian french phrase.
	boomerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "humanoïde"}

	// boomerDutchPhrase represents boomer dutch phrase.
	boomerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mens"}

	// boomerFrenchPhrase represents boomer french phrase.
	boomerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "boudiou"}

	// boomerGermanPhrase represents boomer german phrase.
	boomerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "menschlein"}

	// boomerItalianPhrase represents boomer italian phrase.
	boomerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "glacius"}

	// boomerJapanesePhrase represents boomer japanese phrase.
	boomerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ツーツツ"}

	// boomerKoreanPhrase represents boomer korean phrase.
	boomerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "에헴헴"}

	// boomerLatinAmericanSpanishPhrase represents boomer latin american spanish phrase.
	boomerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glaciux"}

	// boomerRussianPhrase represents boomer russian phrase.
	boomerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "человек"}

	// boomerSimplifiedChinesePhrase represents boomer simplified chinese phrase.
	boomerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "轧轧"}

	// boomerSpanishPhrase represents boomer spanish phrase.
	boomerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "humanoide"}

	// boomerTraditionalChinesePhrase represents boomer traditional chinese phrase.
	boomerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "軋軋"}
)

var (
	// boomerPhrase represents boomer phrase.
	boomerPhrase = nook.Languages{
		language.AmericanEnglish:      boomerAmericanEnglishPhrase,
		language.CanadianFrench:       boomerCanadianFrenchPhrase,
		language.Dutch:                boomerDutchPhrase,
		language.French:               boomerFrenchPhrase,
		language.German:               boomerGermanPhrase,
		language.Italian:              boomerItalianPhrase,
		language.Japanese:             boomerJapanesePhrase,
		language.Korean:               boomerKoreanPhrase,
		language.LatinAmericanSpanish: boomerLatinAmericanSpanishPhrase,
		language.Russian:              boomerRussianPhrase,
		language.SimplifiedChinese:    boomerSimplifiedChinesePhrase,
		language.Spanish:              boomerSpanishPhrase,
		language.TraditionalChinese:   boomerTraditionalChinesePhrase}
)

var (
	// Boomer represents boomer.
	Boomer = nook.Villager{
		Character:   boomerCharacter,
		Personality: personality.Lazy,
		Phrase:      boomerPhrase}
)
