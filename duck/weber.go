package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Bébertplumeau"}

	weberDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Weberkwaaah"}

	weberFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bébertplumeau"}

	weberGermanName = nook.Name{
		Language: language.German,
		Value:    "Volkerquaa"}

	weberItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pasqualequaaa"}

	weberJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アチョットピヨ"}

	weberKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아잠만꽥꽥"}

	weberLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paticiocuaa-cuaa"}

	weberRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Веберкря-а"}

	weberSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卿德唧唧"}

	weberSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paticiocuaa-cuaa"}

	weberTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "卿德唧唧"}
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
		Animal:   Duck,
		Birthday: weberBirthday,
		Code:     weberCode,
		Gender:   nook.Male,
		Name:     weberName}
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
	Weber = nook.Villager{
		Character:   weberCharacter,
		Personality: nook.Lazy,
		Phrase:      weberPhrase}
)
