package cat

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
	// tabbyBirthday represents tabby birthday.
	tabbyBirthday = nook.Birthday{
		Day:   13,
		Month: time.August}
)

var (
	// tabbyCode represents tabby code.
	tabbyCode = nook.Code{
		Value: "cat12"}
)

var (
	// tabbyAmericanEnglishName represents tabby american english name.
	tabbyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tabby"}

	// tabbyCanadianFrenchName represents tabby canadian french name.
	tabbyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tigri"}

	// tabbyDutchName represents tabby dutch name.
	tabbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tabby"}

	// tabbyFrenchName represents tabby french name.
	tabbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tigri"}

	// tabbyGermanName represents tabby german name.
	tabbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Zita"}

	// tabbyItalianName represents tabby italian name.
	tabbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lisca"}

	// tabbyJapaneseName represents tabby japanese name.
	tabbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トラこ"}

	// tabbyKoreanName represents tabby korean name.
	tabbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호냥이"}

	// tabbyLatinAmericanSpanishName represents tabby latin american spanish name.
	tabbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Liana"}

	// tabbyRussianName represents tabby russian name.
	tabbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэбби"}

	// tabbySimplifiedChineseName represents tabby simplified chinese name.
	tabbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小虎"}

	// tabbySpanishName represents tabby spanish name.
	tabbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Liana"}

	// tabbyTraditionalChineseName represents tabby traditional chinese name.
	tabbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小虎"}
)

var (
	// tabbyName represents tabby name.
	tabbyName = nook.Languages{
		language.AmericanEnglish:      tabbyAmericanEnglishName,
		language.CanadianFrench:       tabbyCanadianFrenchName,
		language.Dutch:                tabbyDutchName,
		language.French:               tabbyFrenchName,
		language.German:               tabbyGermanName,
		language.Italian:              tabbyItalianName,
		language.Japanese:             tabbyJapaneseName,
		language.Korean:               tabbyKoreanName,
		language.LatinAmericanSpanish: tabbyLatinAmericanSpanishName,
		language.Russian:              tabbyRussianName,
		language.SimplifiedChinese:    tabbySimplifiedChineseName,
		language.Spanish:              tabbySpanishName,
		language.TraditionalChinese:   tabbyTraditionalChineseName}
)

var (
	// tabbyCharacter represents tabby character.
	tabbyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: tabbyBirthday,
		Code:     tabbyCode,
		Key:      character.Tabby,
		Gender:   gender.Female,
		Name:     tabbyName,
		Special:  false}
)

var (
	// tabbyAmericanEnglishPhrase represents tabby american english phrase.
	tabbyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "me-WOW"}

	// tabbyCanadianFrenchPhrase represents tabby canadian french phrase.
	tabbyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mi-AOUUH"}

	// tabbyDutchPhrase represents tabby dutch phrase.
	tabbyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mi-WAUW"}

	// tabbyFrenchPhrase represents tabby french phrase.
	tabbyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mi-AOUUH"}

	// tabbyGermanPhrase represents tabby german phrase.
	tabbyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mrriau"}

	// tabbyItalianPhrase represents tabby italian phrase.
	tabbyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "meOW"}

	// tabbyJapanesePhrase represents tabby japanese phrase.
	tabbyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃは"}

	// tabbyKoreanPhrase represents tabby korean phrase.
	tabbyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "삐뽀"}

	// tabbyLatinAmericanSpanishPhrase represents tabby latin american spanish phrase.
	tabbyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "animiaul"}

	// tabbyRussianPhrase represents tabby russian phrase.
	tabbyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мя-ого"}

	// tabbySimplifiedChinesePhrase represents tabby simplified chinese phrase.
	tabbySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喵吼"}

	// tabbySpanishPhrase represents tabby spanish phrase.
	tabbySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "animiaul"}

	// tabbyTraditionalChinesePhrase represents tabby traditional chinese phrase.
	tabbyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喵吼"}
)

var (
	// tabbyPhrase represents tabby phrase.
	tabbyPhrase = nook.Languages{
		language.AmericanEnglish:      tabbyAmericanEnglishPhrase,
		language.CanadianFrench:       tabbyCanadianFrenchPhrase,
		language.Dutch:                tabbyDutchPhrase,
		language.French:               tabbyFrenchPhrase,
		language.German:               tabbyGermanPhrase,
		language.Italian:              tabbyItalianPhrase,
		language.Japanese:             tabbyJapanesePhrase,
		language.Korean:               tabbyKoreanPhrase,
		language.LatinAmericanSpanish: tabbyLatinAmericanSpanishPhrase,
		language.Russian:              tabbyRussianPhrase,
		language.SimplifiedChinese:    tabbySimplifiedChinesePhrase,
		language.Spanish:              tabbySpanishPhrase,
		language.TraditionalChinese:   tabbyTraditionalChinesePhrase}
)

var (
	// Tabby represents tabby.
	Tabby = nook.Villager{
		Character:   tabbyCharacter,
		Personality: personality.Peppy,
		Phrase:      tabbyPhrase}
)
