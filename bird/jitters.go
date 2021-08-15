package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	jittersBirthday = nook.Birthday{
		Day:   2,
		Month: time.February}
)

var (
	jittersCode = nook.Code{
		Value: "brd04"}
)

var (
	jittersAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jitters"}

	jittersCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gilbertbouerk"}

	jittersDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jitterstjulp"}

	jittersFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gilbertbouerk"}

	jittersGermanName = nook.Name{
		Language: language.German,
		Value:    "Alexzischhh"}

	jittersItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brividobzzert"}

	jittersJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジーニョだニョ"}

	jittersKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "딩요삼바"}

	jittersLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cameliopío-pío"}

	jittersRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джиттерсбзе-е-ерт"}

	jittersSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "纪诺诺诺"}

	jittersSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cameliopío-pío"}

	jittersTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "紀諾諾諾"}
)

var (
	jittersName = nook.Languages{
		language.AmericanEnglish:      jittersAmericanEnglishName,
		language.CanadianFrench:       jittersCanadianFrenchName,
		language.Dutch:                jittersDutchName,
		language.French:               jittersFrenchName,
		language.German:               jittersGermanName,
		language.Italian:              jittersItalianName,
		language.Japanese:             jittersJapaneseName,
		language.Korean:               jittersKoreanName,
		language.LatinAmericanSpanish: jittersLatinAmericanSpanishName,
		language.Russian:              jittersRussianName,
		language.SimplifiedChinese:    jittersSimplifiedChineseName,
		language.Spanish:              jittersSpanishName,
		language.TraditionalChinese:   jittersTraditionalChineseName}
)

var (
	jittersCharacter = nook.Character{
		Animal:   Bird,
		Birthday: jittersBirthday,
		Code:     jittersCode,
		Gender:   nook.Male,
		Name:     jittersName}
)

var (
	jittersAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bzzert"}

	jittersCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouerk"}

	jittersDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tjulp"}

	jittersFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouerk"}

	jittersGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zischhh"}

	jittersItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bzzert"}

	jittersJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だニョ"}

	jittersKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "삼바"}

	jittersLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pío-pío"}

	jittersRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бзе-е-ерт"}

	jittersSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "诺诺"}

	jittersSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pío-pío"}

	jittersTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "諾諾"}
)

var (
	jittersPhrase = nook.Languages{
		language.AmericanEnglish:      jittersAmericanEnglishName,
		language.CanadianFrench:       jittersCanadianFrenchName,
		language.Dutch:                jittersDutchName,
		language.French:               jittersFrenchName,
		language.German:               jittersGermanName,
		language.Italian:              jittersItalianName,
		language.Japanese:             jittersJapaneseName,
		language.Korean:               jittersKoreanName,
		language.LatinAmericanSpanish: jittersLatinAmericanSpanishName,
		language.Russian:              jittersRussianName,
		language.SimplifiedChinese:    jittersSimplifiedChineseName,
		language.Spanish:              jittersSpanishName,
		language.TraditionalChinese:   jittersTraditionalChineseName}
)

var (
	Jitters = nook.Villager{
		Character:   jittersCharacter,
		Personality: nook.Jock,
		Phrase:      jittersPhrase}
)
