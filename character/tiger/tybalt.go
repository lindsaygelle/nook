package tiger

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
	// tybaltBirthday represents Tybalt's birthday.
	tybaltBirthday = nook.Birthday{
		Day:   19,
		Month: time.August}
)

var (
	// tybaltCode represents Tybalt's unique code.
	tybaltCode = nook.Code{
		Value: "tig02"}
)

var (
	// tybaltAmericanEnglishName represents Tybalt's name in American English.
	tybaltAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tybalt"}

	// tybaltCanadianFrenchName represents Tybalt's name in Canadian French.
	tybaltCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jeff"}

	// tybaltDutchName represents Tybalt's name in Dutch.
	tybaltDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tybalt"}

	// tybaltFrenchName represents Tybalt's name in French.
	tybaltFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jeff"}

	// tybaltGermanName represents Tybalt's name in German.
	tybaltGermanName = nook.Name{
		Language: language.German,
		Value:    "Arne"}

	// tybaltItalianName represents Tybalt's name in Italian.
	tybaltItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ubaldo"}

	// tybaltJapaneseName represents Tybalt's name in Japanese.
	tybaltJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハリマオ"}

	// tybaltKoreanName represents Tybalt's name in Korean.
	tybaltKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티볼트"}

	// tybaltLatinAmericanSpanishName represents Tybalt's name in Latin American Spanish.
	tybaltLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Teobaldo"}

	// tybaltRussianName represents Tybalt's name in Russian.
	tybaltRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тибальт"}

	// tybaltSimplifiedChineseName represents Tybalt's name in Simplified Chinese.
	tybaltSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谷丰"}

	// tybaltSpanishName represents Tybalt's name in Spanish.
	tybaltSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Teobaldo"}

	// tybaltTraditionalChineseName represents Tybalt's name in Traditional Chinese.
	tybaltTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "谷豐"}
)

var (
	// tybaltName represents Tybalt's name in different languages.
	tybaltName = nook.Languages{
		language.AmericanEnglish:      tybaltAmericanEnglishName,
		language.CanadianFrench:       tybaltCanadianFrenchName,
		language.Dutch:                tybaltDutchName,
		language.French:               tybaltFrenchName,
		language.German:               tybaltGermanName,
		language.Italian:              tybaltItalianName,
		language.Japanese:             tybaltJapaneseName,
		language.Korean:               tybaltKoreanName,
		language.LatinAmericanSpanish: tybaltLatinAmericanSpanishName,
		language.Russian:              tybaltRussianName,
		language.SimplifiedChinese:    tybaltSimplifiedChineseName,
		language.Spanish:              tybaltSpanishName,
		language.TraditionalChinese:   tybaltTraditionalChineseName}
)

var (
	// tybaltCharacter represents Tybalt's character information.
	tybaltCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: tybaltBirthday,
		Code:     tybaltCode,
		Key:      character.Tybalt,
		Gender:   gender.Male,
		Name:     tybaltName,
		Special:  false}
)

var (
	// tybaltAmericanEnglishPhrase represents Tybalt's phrase in American English.
	tybaltAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrrRAH"}

	// tybaltCanadianFrenchPhrase represents Tybalt's phrase in Canadian French.
	tybaltCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grouaaah"}

	// tybaltDutchPhrase represents Tybalt's phrase in Dutch.
	tybaltDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrrulll"}

	// tybaltFrenchPhrase represents Tybalt's phrase in French.
	tybaltFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grouaaah"}

	// tybaltGermanPhrase represents Tybalt's phrase in German.
	tybaltGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrrummel"}

	// tybaltItalianPhrase represents Tybalt's phrase in Italian.
	tybaltItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrrRAH"}

	// tybaltJapanesePhrase represents Tybalt's phrase in Japanese.
	tybaltJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だトラ"}

	// tybaltKoreanPhrase represents Tybalt's phrase in Korean.
	tybaltKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으르렁"}

	// tybaltLatinAmericanSpanishPhrase represents Tybalt's phrase in Latin American Spanish.
	tybaltLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "raaarruf"}

	// tybaltRussianPhrase represents Tybalt's phrase in Russian.
	tybaltRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "р-рык"}

	// tybaltSimplifiedChinesePhrase represents Tybalt's phrase in Simplified Chinese.
	tybaltSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "虎虎"}

	// tybaltSpanishPhrase represents Tybalt's phrase in Spanish.
	tybaltSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bengala"}

	// tybaltTraditionalChinesePhrase represents Tybalt's phrase in Traditional Chinese.
	tybaltTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "虎虎"}
)

var (
	// tybaltPhrase represents Tybalt's phrase in different languages.
	tybaltPhrase = nook.Languages{
		language.AmericanEnglish:      tybaltAmericanEnglishPhrase,
		language.CanadianFrench:       tybaltCanadianFrenchPhrase,
		language.Dutch:                tybaltDutchPhrase,
		language.French:               tybaltFrenchPhrase,
		language.German:               tybaltGermanPhrase,
		language.Italian:              tybaltItalianPhrase,
		language.Japanese:             tybaltJapanesePhrase,
		language.Korean:               tybaltKoreanPhrase,
		language.LatinAmericanSpanish: tybaltLatinAmericanSpanishPhrase,
		language.Russian:              tybaltRussianPhrase,
		language.SimplifiedChinese:    tybaltSimplifiedChinesePhrase,
		language.Spanish:              tybaltSpanishPhrase,
		language.TraditionalChinese:   tybaltTraditionalChinesePhrase}
)

var (
	// Tybalt represents the character Tybalt with his complete information.
	Tybalt = nook.Villager{
		Character:   tybaltCharacter,
		Personality: personality.Jock,
		Phrase:      tybaltPhrase}
)
