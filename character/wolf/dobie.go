package wolf

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
	// dobieBirthday represents Dobie's birthday.
	dobieBirthday = nook.Birthday{
		Day:   17,
		Month: time.February}
)

var (
	// dobieCode represents Dobie's unique code.
	dobieCode = nook.Code{
		Value: "wol04"}
)

var (
	// dobieAmericanEnglishName represents Dobie's name in American English.
	dobieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dobie"}

	// dobieCanadianFrenchName represents Dobie's name in Canadian French.
	dobieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Loupiot"}

	// dobieDutchName represents Dobie's name in Dutch.
	dobieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dobie"}

	// dobieFrenchName represents Dobie's name in French.
	dobieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Loupiot"}

	// dobieGermanName represents Dobie's name in German.
	dobieGermanName = nook.Name{
		Language: language.German,
		Value:    "Sigmund"}

	// dobieItalianName represents Dobie's name in Italian.
	dobieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sigmund"}

	// dobieJapaneseName represents Dobie's name in Japanese.
	dobieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "けん"}

	// dobieKoreanName represents Dobie's name in Korean.
	dobieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "켄"}

	// dobieLatinAmericanSpanishName represents Dobie's name in Latin American Spanish.
	dobieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Germán"}

	// dobieRussianName represents Dobie's name in Russian.
	dobieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Доби"}

	// dobieSimplifiedChineseName represents Dobie's name in Simplified Chinese.
	dobieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿研"}

	// dobieSpanishName represents Dobie's name in Spanish.
	dobieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Germán"}

	// dobieTraditionalChineseName represents Dobie's name in Traditional Chinese.
	dobieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿研"}
)

var (
	// dobieName represents Dobie's name in different languages.
	dobieName = nook.Languages{
		language.AmericanEnglish:      dobieAmericanEnglishName,
		language.CanadianFrench:       dobieCanadianFrenchName,
		language.Dutch:                dobieDutchName,
		language.French:               dobieFrenchName,
		language.German:               dobieGermanName,
		language.Italian:              dobieItalianName,
		language.Japanese:             dobieJapaneseName,
		language.Korean:               dobieKoreanName,
		language.LatinAmericanSpanish: dobieLatinAmericanSpanishName,
		language.Russian:              dobieRussianName,
		language.SimplifiedChinese:    dobieSimplifiedChineseName,
		language.Spanish:              dobieSpanishName,
		language.TraditionalChinese:   dobieTraditionalChineseName}
)

var (
	// dobieCharacter represents Dobie's character information.
	dobieCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: dobieBirthday,
		Code:     dobieCode,
		Key:      character.Dobie,
		Gender:   gender.Male,
		Name:     dobieName,
		Special:  false}
)

var (
	// dobieAmericanEnglishPhrase represents Dobie's phrase in American English.
	dobieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ohmmm"}

	// dobieCanadianFrenchPhrase represents Dobie's phrase in Canadian French.
	dobieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "loudonc"}

	// dobieDutchPhrase represents Dobie's phrase in Dutch.
	dobieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eh"}

	// dobieFrenchPhrase represents Dobie's phrase in French.
	dobieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "peuchère"}

	// dobieGermanPhrase represents Dobie's phrase in German.
	dobieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heuuul"}

	// dobieItalianPhrase represents Dobie's phrase in Italian.
	dobieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnams"}

	// dobieJapanesePhrase represents Dobie's phrase in Japanese.
	dobieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゴホゴホ"}

	// dobieKoreanPhrase represents Dobie's phrase in Korean.
	dobieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음음"}

	// dobieLatinAmericanSpanishPhrase represents Dobie's phrase in Latin American Spanish.
	dobieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ooomm"}

	// dobieRussianPhrase represents Dobie's phrase in Russian.
	dobieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ом-м-м..."}

	// dobieSimplifiedChinesePhrase represents Dobie's phrase in Simplified Chinese.
	dobieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咳呵"}

	// dobieSpanishPhrase represents Dobie's phrase in Spanish.
	dobieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ooomm"}

	// dobieTraditionalChinesePhrase represents Dobie's phrase in Traditional Chinese.
	dobieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咳呵"}
)

var (
	// dobiePhrase represents Dobie's phrase in different languages.
	dobiePhrase = nook.Languages{
		language.AmericanEnglish:      dobieAmericanEnglishPhrase,
		language.CanadianFrench:       dobieCanadianFrenchPhrase,
		language.Dutch:                dobieDutchPhrase,
		language.French:               dobieFrenchPhrase,
		language.German:               dobieGermanPhrase,
		language.Italian:              dobieItalianPhrase,
		language.Japanese:             dobieJapanesePhrase,
		language.Korean:               dobieKoreanPhrase,
		language.LatinAmericanSpanish: dobieLatinAmericanSpanishPhrase,
		language.Russian:              dobieRussianPhrase,
		language.SimplifiedChinese:    dobieSimplifiedChinesePhrase,
		language.Spanish:              dobieSpanishPhrase,
		language.TraditionalChinese:   dobieTraditionalChinesePhrase}
)

var (
	// Dobie represents the character Dobie with his complete information.
	Dobie = nook.Villager{
		Character:   dobieCharacter,
		Personality: personality.Cranky,
		Phrase:      dobiePhrase}
)
