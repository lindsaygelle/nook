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
	// rolfBirthday represents Rolf's birthday.
	rolfBirthday = nook.Birthday{
		Day:   22,
		Month: time.August}
)

var (
	// rolfCode represents Rolf's unique code.
	rolfCode = nook.Code{
		Value: "tig00"}
)

var (
	// rolfAmericanEnglishName represents Rolf's name in American English.
	rolfAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rolf"}

	// rolfCanadianFrenchName represents Rolf's name in Canadian French.
	rolfCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ralf"}

	// rolfDutchName represents Rolf's name in Dutch.
	rolfDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rolf"}

	// rolfFrenchName represents Rolf's name in French.
	rolfFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ralf"}

	// rolfGermanName represents Rolf's name in German.
	rolfGermanName = nook.Name{
		Language: language.German,
		Value:    "Boris"}

	// rolfItalianName represents Rolf's name in Italian.
	rolfItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rolf"}

	// rolfJapaneseName represents Rolf's name in Japanese.
	rolfJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョモラン"}

	// rolfKoreanName represents Rolf's name in Korean.
	rolfKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호랭이"}

	// rolfLatinAmericanSpanishName represents Rolf's name in Latin American Spanish.
	rolfLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Albino"}

	// rolfRussianName represents Rolf's name in Russian.
	rolfRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рольф"}

	// rolfSimplifiedChineseName represents Rolf's name in Simplified Chinese.
	rolfSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱穆朗"}

	// rolfSpanishName represents Rolf's name in Spanish.
	rolfSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Albino"}

	// rolfTraditionalChineseName represents Rolf's name in Traditional Chinese.
	rolfTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱穆朗"}
)

var (
	// rolfName represents Rolf's name in different languages.
	rolfName = nook.Languages{
		language.AmericanEnglish:      rolfAmericanEnglishName,
		language.CanadianFrench:       rolfCanadianFrenchName,
		language.Dutch:                rolfDutchName,
		language.French:               rolfFrenchName,
		language.German:               rolfGermanName,
		language.Italian:              rolfItalianName,
		language.Japanese:             rolfJapaneseName,
		language.Korean:               rolfKoreanName,
		language.LatinAmericanSpanish: rolfLatinAmericanSpanishName,
		language.Russian:              rolfRussianName,
		language.SimplifiedChinese:    rolfSimplifiedChineseName,
		language.Spanish:              rolfSpanishName,
		language.TraditionalChinese:   rolfTraditionalChineseName}
)

var (
	// rolfCharacter represents Rolf's character information.
	rolfCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: rolfBirthday,
		Code:     rolfCode,
		Key:      character.Rolf,
		Gender:   gender.Male,
		Name:     rolfName,
		Special:  false}
)

var (
	// rolfAmericanEnglishPhrase represents Rolf's phrase in American English.
	rolfAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrrolf"}

	// rolfCanadianFrenchPhrase represents Rolf's phrase in Canadian French.
	rolfCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grrrufff"}

	// rolfDutchPhrase represents Rolf's phrase in Dutch.
	rolfDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grrrolf"}

	// rolfFrenchPhrase represents Rolf's phrase in French.
	rolfFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grrrufff"}

	// rolfGermanPhrase represents Rolf's phrase in German.
	rolfGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brüll"}

	// rolfItalianPhrase represents Rolf's phrase in Italian.
	rolfItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrrolf"}

	// rolfJapanesePhrase represents Rolf's phrase in Japanese.
	rolfJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だガー"}

	// rolfKoreanPhrase represents Rolf's phrase in Korean.
	rolfKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "근데"}

	// rolfLatinAmericanSpanishPhrase represents Rolf's phrase in Latin American Spanish.
	rolfLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grooorf"}

	// rolfRussianPhrase represents Rolf's phrase in Russian.
	rolfRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гр-р-рольф"}

	// rolfSimplifiedChinesePhrase represents Rolf's phrase in Simplified Chinese.
	rolfSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "王"}

	// rolfSpanishPhrase represents Rolf's phrase in Spanish.
	rolfSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grooorf"}

	// rolfTraditionalChinesePhrase represents Rolf's phrase in Traditional Chinese.
	rolfTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "王"}
)

var (
	// rolfPhrase represents Rolf's phrase in different languages.
	rolfPhrase = nook.Languages{
		language.AmericanEnglish:      rolfAmericanEnglishPhrase,
		language.CanadianFrench:       rolfCanadianFrenchPhrase,
		language.Dutch:                rolfDutchPhrase,
		language.French:               rolfFrenchPhrase,
		language.German:               rolfGermanPhrase,
		language.Italian:              rolfItalianPhrase,
		language.Japanese:             rolfJapanesePhrase,
		language.Korean:               rolfKoreanPhrase,
		language.LatinAmericanSpanish: rolfLatinAmericanSpanishPhrase,
		language.Russian:              rolfRussianPhrase,
		language.SimplifiedChinese:    rolfSimplifiedChinesePhrase,
		language.Spanish:              rolfSpanishPhrase,
		language.TraditionalChinese:   rolfTraditionalChinesePhrase}
)

var (
	// Rolf represents the character Rolf with his complete information.
	Rolf = nook.Villager{
		Character:   rolfCharacter,
		Personality: personality.Cranky,
		Phrase:      rolfPhrase}
)
