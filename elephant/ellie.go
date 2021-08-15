package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ellieBirthday = nook.Birthday{
		Day:   12,
		Month: time.May}
)

var (
	ellieCode = nook.Code{
		Value: "elp07"}
)

var (
	ellieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ellie"}

	ellieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ella"}

	ellieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ellie"}

	ellieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ella"}

	ellieGermanName = nook.Name{
		Language: language.German,
		Value:    "Elfi"}

	ellieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elly"}

	ellieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エクレア"}

	ellieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에끌레르"}

	ellieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eli"}

	ellieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элла"}

	ellieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "泡芙"}

	ellieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eli"}

	ellieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "泡芙"}
)

var (
	ellieName = nook.Languages{
		language.AmericanEnglish:      ellieAmericanEnglishName,
		language.CanadianFrench:       ellieCanadianFrenchName,
		language.Dutch:                ellieDutchName,
		language.French:               ellieFrenchName,
		language.German:               ellieGermanName,
		language.Italian:              ellieItalianName,
		language.Japanese:             ellieJapaneseName,
		language.Korean:               ellieKoreanName,
		language.LatinAmericanSpanish: ellieLatinAmericanSpanishName,
		language.Russian:              ellieRussianName,
		language.SimplifiedChinese:    ellieSimplifiedChineseName,
		language.Spanish:              ellieSpanishName,
		language.TraditionalChinese:   ellieTraditionalChineseName}
)

var (
	ellieCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: ellieBirthday,
		Code:     ellieCode,
		Gender:   nook.Female,
		Name:     ellieName}
)

var (
	ellieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee one"}

	ellieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sacrelotte"}

	ellieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "dreumes"}

	ellieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sacrelotte"}

	ellieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "polter"}

	ellieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caromé"}

	ellieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ララン"}

	ellieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "트랄라"}

	ellieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fi-fiú"}

	ellieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кроха"}

	ellieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啦啷"}

	ellieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trompi"}

	ellieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啦啷"}
)

var (
	elliePhrase = nook.Languages{
		language.AmericanEnglish:      ellieAmericanEnglishName,
		language.CanadianFrench:       ellieCanadianFrenchName,
		language.Dutch:                ellieDutchName,
		language.French:               ellieFrenchName,
		language.German:               ellieGermanName,
		language.Italian:              ellieItalianName,
		language.Japanese:             ellieJapaneseName,
		language.Korean:               ellieKoreanName,
		language.LatinAmericanSpanish: ellieLatinAmericanSpanishName,
		language.Russian:              ellieRussianName,
		language.SimplifiedChinese:    ellieSimplifiedChineseName,
		language.Spanish:              ellieSpanishName,
		language.TraditionalChinese:   ellieTraditionalChineseName}
)

var (
	Ellie = nook.Villager{
		Character:   ellieCharacter,
		Personality: nook.Normal,
		Phrase:      elliePhrase}
)
