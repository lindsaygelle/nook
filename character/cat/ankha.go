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
	// ankhaBirthday represents ankha birthday.
	ankhaBirthday = nook.Birthday{
		Day:   22,
		Month: time.September}
)

var (
	// ankhaCode represents ankha code.
	ankhaCode = nook.Code{
		Value: "cat19"}
)

var (
	// ankhaAmericanEnglishName represents ankha american english name.
	ankhaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ankha"}

	// ankhaCanadianFrenchName represents ankha canadian french name.
	ankhaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Neferti"}

	// ankhaDutchName represents ankha dutch name.
	ankhaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ankha"}

	// ankhaFrenchName represents ankha french name.
	ankhaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neferti"}

	// ankhaGermanName represents ankha german name.
	ankhaGermanName = nook.Name{
		Language: language.German,
		Value:    "Kleo"}

	// ankhaItalianName represents ankha italian name.
	ankhaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cleo"}

	// ankhaJapaneseName represents ankha japanese name.
	ankhaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナイル"}

	// ankhaKoreanName represents ankha korean name.
	ankhaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "클레오"}

	// ankhaLatinAmericanSpanishName represents ankha latin american spanish name.
	ankhaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Patri"}

	// ankhaRussianName represents ankha russian name.
	ankhaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анка"}

	// ankhaSimplifiedChineseName represents ankha simplified chinese name.
	ankhaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艳后"}

	// ankhaSpanishName represents ankha spanish name.
	ankhaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Patri"}

	// ankhaTraditionalChineseName represents ankha traditional chinese name.
	ankhaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艷后"}
)

var (
	// ankhaName represents ankha name.
	ankhaName = nook.Languages{
		language.AmericanEnglish:      ankhaAmericanEnglishName,
		language.CanadianFrench:       ankhaCanadianFrenchName,
		language.Dutch:                ankhaDutchName,
		language.French:               ankhaFrenchName,
		language.German:               ankhaGermanName,
		language.Italian:              ankhaItalianName,
		language.Japanese:             ankhaJapaneseName,
		language.Korean:               ankhaKoreanName,
		language.LatinAmericanSpanish: ankhaLatinAmericanSpanishName,
		language.Russian:              ankhaRussianName,
		language.SimplifiedChinese:    ankhaSimplifiedChineseName,
		language.Spanish:              ankhaSpanishName,
		language.TraditionalChinese:   ankhaTraditionalChineseName}
)

var (
	// ankhaCharacter represents ankha character.
	ankhaCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: ankhaBirthday,
		Code:     ankhaCode,
		Key:      character.Ankha,
		Gender:   gender.Female,
		Name:     ankhaName,
		Special:  false}
)

var (
	// ankhaAmericanEnglishPhrase represents ankha american english phrase.
	ankhaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "me meow"}

	// ankhaCanadianFrenchPhrase represents ankha canadian french phrase.
	ankhaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mi-i-i-ouh"}

	// ankhaDutchPhrase represents ankha dutch phrase.
	ankhaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sfinx"}

	// ankhaFrenchPhrase represents ankha french phrase.
	ankhaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mi-i-i-ouh"}

	// ankhaGermanPhrase represents ankha german phrase.
	ankhaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mi-miau"}

	// ankhaItalianPhrase represents ankha italian phrase.
	ankhaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miao miao"}

	// ankhaJapanesePhrase represents ankha japanese phrase.
	ankhaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クフフ"}

	// ankhaKoreanPhrase represents ankha korean phrase.
	ankhaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파트라"}

	// ankhaLatinAmericanSpanishPhrase represents ankha latin american spanish phrase.
	ankhaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "marramius"}

	// ankhaRussianPhrase represents ankha russian phrase.
	ankhaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ми-мяу"}

	// ankhaSimplifiedChinesePhrase represents ankha simplified chinese phrase.
	ankhaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尼罗"}

	// ankhaSpanishPhrase represents ankha spanish phrase.
	ankhaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "marramius"}

	// ankhaTraditionalChinesePhrase represents ankha traditional chinese phrase.
	ankhaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "尼羅"}
)

var (
	// ankhaPhrase represents ankha phrase.
	ankhaPhrase = nook.Languages{
		language.AmericanEnglish:      ankhaAmericanEnglishPhrase,
		language.CanadianFrench:       ankhaCanadianFrenchPhrase,
		language.Dutch:                ankhaDutchPhrase,
		language.French:               ankhaFrenchPhrase,
		language.German:               ankhaGermanPhrase,
		language.Italian:              ankhaItalianPhrase,
		language.Japanese:             ankhaJapanesePhrase,
		language.Korean:               ankhaKoreanPhrase,
		language.LatinAmericanSpanish: ankhaLatinAmericanSpanishPhrase,
		language.Russian:              ankhaRussianPhrase,
		language.SimplifiedChinese:    ankhaSimplifiedChinesePhrase,
		language.Spanish:              ankhaSpanishPhrase,
		language.TraditionalChinese:   ankhaTraditionalChinesePhrase}
)

var (
	// Ankha represents ankha.
	Ankha = nook.Villager{
		Character:   ankhaCharacter,
		Personality: personality.Snooty,
		Phrase:      ankhaPhrase}
)
