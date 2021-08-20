package duck

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
	weberBirthday = nook.Birthday{
		Day:   30,
		Month: time.June}
)

var (
	weberCode = nook.Code{
		Value: "duk11"}
)

var (
	weberAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Weber"}

	weberCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bébert"}

	weberDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Weber"}

	weberFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bébert"}

	weberGermanName = nook.Name{
		Language: language.German,
		Value:    "Volker"}

	weberItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pasquale"}

	weberJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アチョット"}

	weberKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아잠만"}

	weberLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paticio"}

	weberRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вебер"}

	weberSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卿德"}

	weberSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paticio"}

	weberTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "卿德"}
)

var (
	weberName = nook.Languages{
		language.AmericanEnglish:      weberAmericanEnglishName,
		language.CanadianFrench:       weberCanadianFrenchName,
		language.Dutch:                weberDutchName,
		language.French:               weberFrenchName,
		language.German:               weberGermanName,
		language.Italian:              weberItalianName,
		language.Japanese:             weberJapaneseName,
		language.Korean:               weberKoreanName,
		language.LatinAmericanSpanish: weberLatinAmericanSpanishName,
		language.Russian:              weberRussianName,
		language.SimplifiedChinese:    weberSimplifiedChineseName,
		language.Spanish:              weberSpanishName,
		language.TraditionalChinese:   weberTraditionalChineseName}
)

var (
	weberCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: weberBirthday,
		Code:     weberCode,
		Key:      character.Weber,
		Gender:   gender.Male,
		Name:     weberName,
		Special:  false}
)

var (
	weberAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quaa"}

	weberCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "plumeau"}

	weberDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaaah"}

	weberFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "plumeau"}

	weberGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaa"}

	weberItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaaa"}

	weberJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピヨ"}

	weberKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꽥꽥"}

	weberLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuaa-cuaa"}

	weberRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кря-а"}

	weberSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唧唧"}

	weberSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuaa-cuaa"}

	weberTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唧唧"}
)

var (
	weberPhrase = nook.Languages{
		language.AmericanEnglish:      weberAmericanEnglishPhrase,
		language.CanadianFrench:       weberCanadianFrenchPhrase,
		language.Dutch:                weberDutchPhrase,
		language.French:               weberFrenchPhrase,
		language.German:               weberGermanPhrase,
		language.Italian:              weberItalianPhrase,
		language.Japanese:             weberJapanesePhrase,
		language.Korean:               weberKoreanPhrase,
		language.LatinAmericanSpanish: weberLatinAmericanSpanishPhrase,
		language.Russian:              weberRussianPhrase,
		language.SimplifiedChinese:    weberSimplifiedChinesePhrase,
		language.Spanish:              weberSpanishPhrase,
		language.TraditionalChinese:   weberTraditionalChinesePhrase}
)

var (
	Weber = nook.Villager{
		Character:   weberCharacter,
		Personality: personality.Lazy,
		Phrase:      weberPhrase}
)
