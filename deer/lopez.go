package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lopezBirthday = nook.Birthday{
		Day:   20,
		Month: time.August}
)

var (
	lopezCode = nook.Code{
		Value: "der05"}
)

var (
	lopezAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lopez"}

	lopezCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jonyou bet"}

	lopezDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lopezbokje"}

	lopezFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jonjaunisse"}

	lopezGermanName = nook.Name{
		Language: language.German,
		Value:    "Eckartbocko"}

	lopezItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lopezsdleng"}

	lopezJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トムソンがぜん"}

	lopezKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "톰슨엣헴"}

	lopezLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Agrestetrisca"}

	lopezRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лопесбубух"}

	lopezSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "汤姆变了"}

	lopezSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Agrestetrisca"}

	lopezTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "湯姆變了"}
)

var (
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
	lopezCharacter = nook.Character{
		Animal:   Deer,
		Birthday: lopezBirthday,
		Code:     lopezCode,
		Gender:   nook.Male,
		Name:     lopezName}
)

var (
	lopezAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buckaroobadoom"}

	lopezCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "you bet"}

	lopezDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bokje"}

	lopezFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "jaunisse"}

	lopezGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bocko"}

	lopezItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sdleng"}

	lopezJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "がぜん"}

	lopezKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엣헴"}

	lopezLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trisca"}

	lopezRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бубух"}

	lopezSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "变了"}

	lopezSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trisca"}

	lopezTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "變了"}
)

var (
	lopezPhrase = nook.Languages{
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
	Lopez = nook.Villager{
		Character:   lopezCharacter,
		Personality: nook.Smug,
		Phrase:      lopezPhrase}
)
