package bird

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
	// jittersBirthday represents jitters birthday.
	jittersBirthday = nook.Birthday{
		Day:   2,
		Month: time.February}
)

var (
	// jittersCode represents jitters code.
	jittersCode = nook.Code{
		Value: "brd04"}
)

var (
	// jittersAmericanEnglishName represents jitters american english name.
	jittersAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jitters"}

	// jittersCanadianFrenchName represents jitters canadian french name.
	jittersCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gilbert"}

	// jittersDutchName represents jitters dutch name.
	jittersDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jitters"}

	// jittersFrenchName represents jitters french name.
	jittersFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gilbert"}

	// jittersGermanName represents jitters german name.
	jittersGermanName = nook.Name{
		Language: language.German,
		Value:    "Alex"}

	// jittersItalianName represents jitters italian name.
	jittersItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brivido"}

	// jittersJapaneseName represents jitters japanese name.
	jittersJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジーニョ"}

	// jittersKoreanName represents jitters korean name.
	jittersKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "딩요"}

	// jittersLatinAmericanSpanishName represents jitters latin american spanish name.
	jittersLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camelio"}

	// jittersRussianName represents jitters russian name.
	jittersRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джиттерс"}

	// jittersSimplifiedChineseName represents jitters simplified chinese name.
	jittersSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "纪诺"}

	// jittersSpanishName represents jitters spanish name.
	jittersSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camelio"}

	// jittersTraditionalChineseName represents jitters traditional chinese name.
	jittersTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "紀諾"}
)

var (
	// jittersName represents jitters name.
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
	// jittersCharacter represents jitters character.
	jittersCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: jittersBirthday,
		Code:     jittersCode,
		Key:      character.Jitters,
		Gender:   gender.Male,
		Name:     jittersName,
		Special:  false}
)

var (
	// jittersAmericanEnglishPhrase represents jitters american english phrase.
	jittersAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bzzert"}

	// jittersCanadianFrenchPhrase represents jitters canadian french phrase.
	jittersCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouerk"}

	// jittersDutchPhrase represents jitters dutch phrase.
	jittersDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tjulp"}

	// jittersFrenchPhrase represents jitters french phrase.
	jittersFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouerk"}

	// jittersGermanPhrase represents jitters german phrase.
	jittersGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zischhh"}

	// jittersItalianPhrase represents jitters italian phrase.
	jittersItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bzzert"}

	// jittersJapanesePhrase represents jitters japanese phrase.
	jittersJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だニョ"}

	// jittersKoreanPhrase represents jitters korean phrase.
	jittersKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "삼바"}

	// jittersLatinAmericanSpanishPhrase represents jitters latin american spanish phrase.
	jittersLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pío-pío"}

	// jittersRussianPhrase represents jitters russian phrase.
	jittersRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бзе-е-ерт"}

	// jittersSimplifiedChinesePhrase represents jitters simplified chinese phrase.
	jittersSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "诺诺"}

	// jittersSpanishPhrase represents jitters spanish phrase.
	jittersSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pío-pío"}

	// jittersTraditionalChinesePhrase represents jitters traditional chinese phrase.
	jittersTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "諾諾"}
)

var (
	// jittersPhrase represents jitters phrase.
	jittersPhrase = nook.Languages{
		language.AmericanEnglish:      jittersAmericanEnglishPhrase,
		language.CanadianFrench:       jittersCanadianFrenchPhrase,
		language.Dutch:                jittersDutchPhrase,
		language.French:               jittersFrenchPhrase,
		language.German:               jittersGermanPhrase,
		language.Italian:              jittersItalianPhrase,
		language.Japanese:             jittersJapanesePhrase,
		language.Korean:               jittersKoreanPhrase,
		language.LatinAmericanSpanish: jittersLatinAmericanSpanishPhrase,
		language.Russian:              jittersRussianPhrase,
		language.SimplifiedChinese:    jittersSimplifiedChinesePhrase,
		language.Spanish:              jittersSpanishPhrase,
		language.TraditionalChinese:   jittersTraditionalChinesePhrase}
)

var (
	// Jitters represents jitters.
	Jitters = nook.Villager{
		Character:   jittersCharacter,
		Personality: personality.Jock,
		Phrase:      jittersPhrase}
)
