package chicken

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
	// avaBirthday represents ava birthday.
	avaBirthday = nook.Birthday{
		Day:   28,
		Month: time.April}
)

var (
	// avaCode represents ava code.
	avaCode = nook.Code{
		Value: "chn05"}
)

var (
	// avaAmericanEnglishName represents ava american english name.
	avaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ava"}

	// avaCanadianFrenchName represents ava canadian french name.
	avaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eva"}

	// avaDutchName represents ava dutch name.
	avaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ava"}

	// avaFrenchName represents ava french name.
	avaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eva"}

	// avaGermanName represents ava german name.
	avaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gisela"}

	// avaItalianName represents ava italian name.
	avaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ava"}

	// avaJapaneseName represents ava japanese name.
	avaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドミグラ"}

	// avaKoreanName represents ava korean name.
	avaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에바"}

	// avaLatinAmericanSpanishName represents ava latin american spanish name.
	avaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ava"}

	// avaRussianName represents ava russian name.
	avaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ава"}

	// avaSimplifiedChineseName represents ava simplified chinese name.
	avaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "陶米咕"}

	// avaSpanishName represents ava spanish name.
	avaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ava"}

	// avaTraditionalChineseName represents ava traditional chinese name.
	avaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陶米咕"}
)

var (
	// avaName represents ava name.
	avaName = nook.Languages{
		language.AmericanEnglish:      avaAmericanEnglishName,
		language.CanadianFrench:       avaCanadianFrenchName,
		language.Dutch:                avaDutchName,
		language.French:               avaFrenchName,
		language.German:               avaGermanName,
		language.Italian:              avaItalianName,
		language.Japanese:             avaJapaneseName,
		language.Korean:               avaKoreanName,
		language.LatinAmericanSpanish: avaLatinAmericanSpanishName,
		language.Russian:              avaRussianName,
		language.SimplifiedChinese:    avaSimplifiedChineseName,
		language.Spanish:              avaSpanishName,
		language.TraditionalChinese:   avaTraditionalChineseName}
)

var (
	// avaCharacter represents ava character.
	avaCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: avaBirthday,
		Code:     avaCode,
		Key:      character.Ava,
		Gender:   gender.Female,
		Name:     avaName,
		Special:  false}
)

var (
	// avaAmericanEnglishPhrase represents ava american english phrase.
	avaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "beaker"}

	// avaCanadianFrenchPhrase represents ava canadian french phrase.
	avaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cô-cômment"}

	// avaDutchPhrase represents ava dutch phrase.
	avaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snavel"}

	// avaFrenchPhrase represents ava french phrase.
	avaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cô-cômment"}

	// avaGermanPhrase represents ava german phrase.
	avaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "biiika"}

	// avaItalianPhrase represents ava italian phrase.
	avaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbecsbec"}

	// avaJapanesePhrase represents ava japanese phrase.
	avaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のよぉ"}

	// avaKoreanPhrase represents ava korean phrase.
	avaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬대요"}

	// avaLatinAmericanSpanishPhrase represents ava latin american spanish phrase.
	avaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kikiú"}

	// avaRussianPhrase represents ava russian phrase.
	avaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "клювик"}

	// avaSimplifiedChinesePhrase represents ava simplified chinese phrase.
	avaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是唷"}

	// avaSpanishPhrase represents ava spanish phrase.
	avaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piquito"}

	// avaTraditionalChinesePhrase represents ava traditional chinese phrase.
	avaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是唷"}
)

var (
	// avaPhrase represents ava phrase.
	avaPhrase = nook.Languages{
		language.AmericanEnglish:      avaAmericanEnglishPhrase,
		language.CanadianFrench:       avaCanadianFrenchPhrase,
		language.Dutch:                avaDutchPhrase,
		language.French:               avaFrenchPhrase,
		language.German:               avaGermanPhrase,
		language.Italian:              avaItalianPhrase,
		language.Japanese:             avaJapanesePhrase,
		language.Korean:               avaKoreanPhrase,
		language.LatinAmericanSpanish: avaLatinAmericanSpanishPhrase,
		language.Russian:              avaRussianPhrase,
		language.SimplifiedChinese:    avaSimplifiedChinesePhrase,
		language.Spanish:              avaSpanishPhrase,
		language.TraditionalChinese:   avaTraditionalChinesePhrase}
)

var (
	// Ava represents ava.
	Ava = nook.Villager{
		Character:   avaCharacter,
		Personality: personality.Normal,
		Phrase:      avaPhrase}
)
