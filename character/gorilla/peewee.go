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
	// peeweeBirthday represents peewee birthday.
	peeweeBirthday = nook.Birthday{
		Day:   11,
		Month: time.September}
)

var (
	// peeweeCode represents peewee code.
	peeweeCode = nook.Code{
		Value: "gor01"}
)

var (
	// peeweeAmericanEnglishName represents peewee american english name.
	peeweeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peewee"}

	// peeweeCanadianFrenchName represents peewee canadian french name.
	peeweeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gogo"}

	// peeweeDutchName represents peewee dutch name.
	peeweeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peewee"}

	// peeweeFrenchName represents peewee french name.
	peeweeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gogo"}

	// peeweeGermanName represents peewee german name.
	peeweeGermanName = nook.Name{
		Language: language.German,
		Value:    "Manfred"}

	// peeweeItalianName represents peewee italian name.
	peeweeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kong"}

	// peeweeJapaneseName represents peewee japanese name.
	peeweeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダンベル"}

	// peeweeKoreanName represents peewee korean name.
	peeweeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덤벨"}

	// peeweeLatinAmericanSpanishName represents peewee latin american spanish name.
	peeweeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lotar"}

	// peeweeRussianName represents peewee russian name.
	peeweeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пиви"}

	// peeweeSimplifiedChineseName represents peewee simplified chinese name.
	peeweeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哑铃"}

	// peeweeSpanishName represents peewee spanish name.
	peeweeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lotar"}

	// peeweeTraditionalChineseName represents peewee traditional chinese name.
	peeweeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啞鈴"}
)

var (
	// peeweeName represents peewee name.
	peeweeName = nook.Languages{
		language.AmericanEnglish:      peeweeAmericanEnglishName,
		language.CanadianFrench:       peeweeCanadianFrenchName,
		language.Dutch:                peeweeDutchName,
		language.French:               peeweeFrenchName,
		language.German:               peeweeGermanName,
		language.Italian:              peeweeItalianName,
		language.Japanese:             peeweeJapaneseName,
		language.Korean:               peeweeKoreanName,
		language.LatinAmericanSpanish: peeweeLatinAmericanSpanishName,
		language.Russian:              peeweeRussianName,
		language.SimplifiedChinese:    peeweeSimplifiedChineseName,
		language.Spanish:              peeweeSpanishName,
		language.TraditionalChinese:   peeweeTraditionalChineseName}
)

var (
	// peeweeCharacter represents peewee character.
	peeweeCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: peeweeBirthday,
		Code:     peeweeCode,
		Key:      character.Peewee,
		Gender:   gender.Male,
		Name:     peeweeName,
		Special:  false}
)

var (
	// peeweeAmericanEnglishPhrase represents peewee american english phrase.
	peeweeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l dude"}

	// peeweeCanadianFrenchPhrase represents peewee canadian french phrase.
	peeweeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gadget"}

	// peeweeDutchPhrase represents peewee dutch phrase.
	peeweeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "dwerg"}

	// peeweeFrenchPhrase represents peewee french phrase.
	peeweeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gadget"}

	// peeweeGermanPhrase represents peewee german phrase.
	peeweeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwerg"}

	// peeweeItalianPhrase represents peewee italian phrase.
	peeweeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baaanane"}

	// peeweeJapanesePhrase represents peewee japanese phrase.
	peeweeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガオ"}

	// peeweeKoreanPhrase represents peewee korean phrase.
	peeweeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "겉멋"}

	// peeweeLatinAmericanSpanishPhrase represents peewee latin american spanish phrase.
	peeweeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cani"}

	// peeweeRussianPhrase represents peewee russian phrase.
	peeweeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мелкота"}

	// peeweeSimplifiedChinesePhrase represents peewee simplified chinese phrase.
	peeweeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "高高"}

	// peeweeSpanishPhrase represents peewee spanish phrase.
	peeweeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cani"}

	// peeweeTraditionalChinesePhrase represents peewee traditional chinese phrase.
	peeweeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "高高"}
)

var (
	// peeweePhrase represents peewee phrase.
	peeweePhrase = nook.Languages{
		language.AmericanEnglish:      peeweeAmericanEnglishPhrase,
		language.CanadianFrench:       peeweeCanadianFrenchPhrase,
		language.Dutch:                peeweeDutchPhrase,
		language.French:               peeweeFrenchPhrase,
		language.German:               peeweeGermanPhrase,
		language.Italian:              peeweeItalianPhrase,
		language.Japanese:             peeweeJapanesePhrase,
		language.Korean:               peeweeKoreanPhrase,
		language.LatinAmericanSpanish: peeweeLatinAmericanSpanishPhrase,
		language.Russian:              peeweeRussianPhrase,
		language.SimplifiedChinese:    peeweeSimplifiedChinesePhrase,
		language.Spanish:              peeweeSpanishPhrase,
		language.TraditionalChinese:   peeweeTraditionalChinesePhrase}
)

var (
	// Peewee represents peewee.
	Peewee = nook.Villager{
		Character:   peeweeCharacter,
		Personality: personality.Cranky,
		Phrase:      peeweePhrase}
)
