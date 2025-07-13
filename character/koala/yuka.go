package koala

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
	// yukaBirthday represents yuka birthday.
	yukaBirthday = nook.Birthday{
		Day:   20,
		Month: time.July}
)

var (
	// yukaCode represents yuka code.
	yukaCode = nook.Code{
		Value: "kal00"}
)

var (
	// yukaAmericanEnglishName represents yuka american english name.
	yukaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Yuka"}

	// yukaCanadianFrenchName represents yuka canadian french name.
	yukaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Calypso"}

	// yukaDutchName represents yuka dutch name.
	yukaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Yuka"}

	// yukaFrenchName represents yuka french name.
	yukaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Calypso"}

	// yukaGermanName represents yuka german name.
	yukaGermanName = nook.Name{
		Language: language.German,
		Value:    "Ute"}

	// yukaItalianName represents yuka italian name.
	yukaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Yuka"}

	// yukaJapaneseName represents yuka japanese name.
	yukaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユーカリ"}

	// yukaKoreanName represents yuka korean name.
	yukaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유카리"}

	// yukaLatinAmericanSpanishName represents yuka latin american spanish name.
	yukaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Yuka"}

	// yukaRussianName represents yuka russian name.
	yukaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юка"}

	// yukaSimplifiedChineseName represents yuka simplified chinese name.
	yukaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尤嘉莉"}

	// yukaSpanishName represents yuka spanish name.
	yukaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Yuka"}

	// yukaTraditionalChineseName represents yuka traditional chinese name.
	yukaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "尤嘉莉"}
)

var (
	// yukaName represents yuka name.
	yukaName = nook.Languages{
		language.AmericanEnglish:      yukaAmericanEnglishName,
		language.CanadianFrench:       yukaCanadianFrenchName,
		language.Dutch:                yukaDutchName,
		language.French:               yukaFrenchName,
		language.German:               yukaGermanName,
		language.Italian:              yukaItalianName,
		language.Japanese:             yukaJapaneseName,
		language.Korean:               yukaKoreanName,
		language.LatinAmericanSpanish: yukaLatinAmericanSpanishName,
		language.Russian:              yukaRussianName,
		language.SimplifiedChinese:    yukaSimplifiedChineseName,
		language.Spanish:              yukaSpanishName,
		language.TraditionalChinese:   yukaTraditionalChineseName}
)

var (
	// yukaCharacter represents yuka character.
	yukaCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: yukaBirthday,
		Code:     yukaCode,
		Key:      character.Yuka,
		Gender:   gender.Female,
		Name:     yukaName,
		Special:  false}
)

var (
	// yukaAmericanEnglishPhrase represents yuka american english phrase.
	yukaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tsk tsk"}

	// yukaCanadianFrenchPhrase represents yuka canadian french phrase.
	yukaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tss tss"}

	// yukaDutchPhrase represents yuka dutch phrase.
	yukaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tss tss"}

	// yukaFrenchPhrase represents yuka french phrase.
	yukaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tss tss"}

	// yukaGermanPhrase represents yuka german phrase.
	yukaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grins"}

	// yukaItalianPhrase represents yuka italian phrase.
	yukaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tsé tsé"}

	// yukaJapanesePhrase represents yuka japanese phrase.
	yukaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アラ"}

	// yukaKoreanPhrase represents yuka korean phrase.
	yukaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어맛"}

	// yukaLatinAmericanSpanishPhrase represents yuka latin american spanish phrase.
	yukaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tsé tsé"}

	// yukaRussianPhrase represents yuka russian phrase.
	yukaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ц-ц"}

	// yukaSimplifiedChinesePhrase represents yuka simplified chinese phrase.
	yukaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啊啦"}

	// yukaSpanishPhrase represents yuka spanish phrase.
	yukaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pues"}

	// yukaTraditionalChinesePhrase represents yuka traditional chinese phrase.
	yukaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啊啦"}
)

var (
	// yukaPhrase represents yuka phrase.
	yukaPhrase = nook.Languages{
		language.AmericanEnglish:      yukaAmericanEnglishPhrase,
		language.CanadianFrench:       yukaCanadianFrenchPhrase,
		language.Dutch:                yukaDutchPhrase,
		language.French:               yukaFrenchPhrase,
		language.German:               yukaGermanPhrase,
		language.Italian:              yukaItalianPhrase,
		language.Japanese:             yukaJapanesePhrase,
		language.Korean:               yukaKoreanPhrase,
		language.LatinAmericanSpanish: yukaLatinAmericanSpanishPhrase,
		language.Russian:              yukaRussianPhrase,
		language.SimplifiedChinese:    yukaSimplifiedChinesePhrase,
		language.Spanish:              yukaSpanishPhrase,
		language.TraditionalChinese:   yukaTraditionalChinesePhrase}
)

var (
	// Yuka represents yuka.
	Yuka = nook.Villager{
		Character:   yukaCharacter,
		Personality: personality.Snooty,
		Phrase:      yukaPhrase}
)
