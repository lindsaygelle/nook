package pig

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
	// hughBirthday represents hugh birthday.
	hughBirthday = nook.Birthday{
		Day:   30,
		Month: time.December}
)

var (
	// hughCode represents hugh code.
	hughCode = nook.Code{
		Value: "pig03"}
)

var (
	// hughAmericanEnglishName represents hugh american english name.
	hughAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hugh"}

	// hughCanadianFrenchName represents hugh canadian french name.
	hughCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bonno"}

	// hughDutchName represents hugh dutch name.
	hughDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hugh"}

	// hughFrenchName represents hugh french name.
	hughFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bonno"}

	// hughGermanName represents hugh german name.
	hughGermanName = nook.Name{
		Language: language.German,
		Value:    "Hugo"}

	// hughItalianName represents hugh italian name.
	hughItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jambon"}

	// hughJapaneseName represents hugh japanese name.
	hughJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クッチャネ"}

	// hughKoreanName represents hugh korean name.
	hughKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "먹고파"}

	// hughLatinAmericanSpanishName represents hugh latin american spanish name.
	hughLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jacobo"}

	// hughRussianName represents hugh russian name.
	hughRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хью"}

	// hughSimplifiedChineseName represents hugh simplified chinese name.
	hughSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿猪"}

	// hughSpanishName represents hugh spanish name.
	hughSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jacobo"}

	// hughTraditionalChineseName represents hugh traditional chinese name.
	hughTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿豬"}
)

var (
	// hughName represents hugh name.
	hughName = nook.Languages{
		language.AmericanEnglish:      hughAmericanEnglishName,
		language.CanadianFrench:       hughCanadianFrenchName,
		language.Dutch:                hughDutchName,
		language.French:               hughFrenchName,
		language.German:               hughGermanName,
		language.Italian:              hughItalianName,
		language.Japanese:             hughJapaneseName,
		language.Korean:               hughKoreanName,
		language.LatinAmericanSpanish: hughLatinAmericanSpanishName,
		language.Russian:              hughRussianName,
		language.SimplifiedChinese:    hughSimplifiedChineseName,
		language.Spanish:              hughSpanishName,
		language.TraditionalChinese:   hughTraditionalChineseName}
)

var (
	// hughCharacter represents hugh character.
	hughCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: hughBirthday,
		Code:     hughCode,
		Key:      character.Hugh,
		Gender:   gender.Male,
		Name:     hughName,
		Special:  false}
)

var (
	// hughAmericanEnglishPhrase represents hugh american english phrase.
	hughAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snortle"}

	// hughCanadianFrenchPhrase represents hugh canadian french phrase.
	hughCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cuick"}

	// hughDutchPhrase represents hugh dutch phrase.
	hughDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuif"}

	// hughFrenchPhrase represents hugh french phrase.
	hughFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuick"}

	// hughGermanPhrase represents hugh german phrase.
	hughGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schniff"}

	// hughItalianPhrase represents hugh italian phrase.
	hughItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snort"}

	// hughJapanesePhrase represents hugh japanese phrase.
	hughJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とかね"}

	// hughKoreanPhrase represents hugh korean phrase.
	hughKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아님말구"}

	// hughLatinAmericanSpanishPhrase represents hugh latin american spanish phrase.
	hughLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "groink"}

	// hughRussianPhrase represents hugh russian phrase.
	hughRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюк-хрюк"}

	// hughSimplifiedChinesePhrase represents hugh simplified chinese phrase.
	hughSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "懒懒"}

	// hughSpanishPhrase represents hugh spanish phrase.
	hughSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "groink"}

	// hughTraditionalChinesePhrase represents hugh traditional chinese phrase.
	hughTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "懶懶"}
)

var (
	// hughPhrase represents hugh phrase.
	hughPhrase = nook.Languages{
		language.AmericanEnglish:      hughAmericanEnglishPhrase,
		language.CanadianFrench:       hughCanadianFrenchPhrase,
		language.Dutch:                hughDutchPhrase,
		language.French:               hughFrenchPhrase,
		language.German:               hughGermanPhrase,
		language.Italian:              hughItalianPhrase,
		language.Japanese:             hughJapanesePhrase,
		language.Korean:               hughKoreanPhrase,
		language.LatinAmericanSpanish: hughLatinAmericanSpanishPhrase,
		language.Russian:              hughRussianPhrase,
		language.SimplifiedChinese:    hughSimplifiedChinesePhrase,
		language.Spanish:              hughSpanishPhrase,
		language.TraditionalChinese:   hughTraditionalChinesePhrase}
)

var (
	// Hugh represents hugh.
	Hugh = nook.Villager{
		Character:   hughCharacter,
		Personality: personality.Lazy,
		Phrase:      hughPhrase}
)
