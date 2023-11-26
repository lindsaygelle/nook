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
	// wolfgangBirthday represents Wolfgang's birthday.
	wolfgangBirthday = nook.Birthday{
		Day:   25,
		Month: time.November}
)

var (
	// wolfgangCode represents Wolfgang's unique code.
	wolfgangCode = nook.Code{
		Value: "wol02"}
)

var (
	// wolfgangAmericanEnglishName represents Wolfgang's name in American English.
	wolfgangAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wolfgang"}

	// wolfgangCanadianFrenchName represents Wolfgang's name in Canadian French.
	wolfgangCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wolfgang"}

	// wolfgangDutchName represents Wolfgang's name in Dutch.
	wolfgangDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wolfgang"}

	// wolfgangFrenchName represents Wolfgang's name in French.
	wolfgangFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wolfgang"}

	// wolfgangGermanName represents Wolfgang's name in German.
	wolfgangGermanName = nook.Name{
		Language: language.German,
		Value:    "Weber"}

	// wolfgangItalianName represents Wolfgang's name in Italian.
	wolfgangItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Wolfgang"}

	// wolfgangJapaneseName represents Wolfgang's name in Japanese.
	wolfgangJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロボ"}

	// wolfgangKoreanName represents Wolfgang's name in Korean.
	wolfgangKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로보"}

	// wolfgangLatinAmericanSpanishName represents Wolfgang's name in Latin American Spanish.
	wolfgangLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wolfi"}

	// wolfgangRussianName represents Wolfgang's name in Russian.
	wolfgangRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вольфи"}

	// wolfgangSimplifiedChineseName represents Wolfgang's name in Simplified Chinese.
	wolfgangSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗博"}

	// wolfgangSpanishName represents Wolfgang's name in Spanish.
	wolfgangSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Wolfi"}

	// wolfgangTraditionalChineseName represents Wolfgang's name in Traditional Chinese.
	wolfgangTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅博"}
)

var (
	// wolfgangName represents Wolfgang's name in different languages.
	wolfgangName = nook.Languages{
		language.AmericanEnglish:      wolfgangAmericanEnglishName,
		language.CanadianFrench:       wolfgangCanadianFrenchName,
		language.Dutch:                wolfgangDutchName,
		language.French:               wolfgangFrenchName,
		language.German:               wolfgangGermanName,
		language.Italian:              wolfgangItalianName,
		language.Japanese:             wolfgangJapaneseName,
		language.Korean:               wolfgangKoreanName,
		language.LatinAmericanSpanish: wolfgangLatinAmericanSpanishName,
		language.Russian:              wolfgangRussianName,
		language.SimplifiedChinese:    wolfgangSimplifiedChineseName,
		language.Spanish:              wolfgangSpanishName,
		language.TraditionalChinese:   wolfgangTraditionalChineseName}
)

var (
	// wolfgangCharacter represents Wolfgang's character information.
	wolfgangCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: wolfgangBirthday,
		Code:     wolfgangCode,
		Key:      character.Wolfgang,
		Gender:   gender.Male,
		Name:     wolfgangName,
		Special:  false}
)

var (
	// wolfgangAmericanEnglishPhrase represents Wolfgang's phrase in American English.
	wolfgangAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snarrrl"}

	// wolfgangCanadianFrenchPhrase represents Wolfgang's phrase in Canadian French.
	wolfgangCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "snurffff"}

	// wolfgangDutchPhrase represents Wolfgang's phrase in Dutch.
	wolfgangDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snauw"}

	// wolfgangFrenchPhrase represents Wolfgang's phrase in French.
	wolfgangFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "snurffff"}

	// wolfgangGermanPhrase represents Wolfgang's phrase in German.
	wolfgangGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knurrr"}

	// wolfgangItalianPhrase represents Wolfgang's phrase in Italian.
	wolfgangItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zanna"}

	// wolfgangJapanesePhrase represents Wolfgang's phrase in Japanese.
	wolfgangJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のな"}

	// wolfgangKoreanPhrase represents Wolfgang's phrase in Korean.
	wolfgangKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거봐"}

	// wolfgangLatinAmericanSpanishPhrase represents Wolfgang's phrase in Latin American Spanish.
	wolfgangLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grauuuh"}

	// wolfgangRussianPhrase represents Wolfgang's phrase in Russian.
	wolfgangRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-рр"}

	// wolfgangSimplifiedChinesePhrase represents Wolfgang's phrase in Simplified Chinese.
	wolfgangSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗罗"}

	// wolfgangSpanishPhrase represents Wolfgang's phrase in Spanish.
	wolfgangSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grauuuh"}

	// wolfgangTraditionalChinesePhrase represents Wolfgang's phrase in Traditional Chinese.
	wolfgangTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅羅"}
)

var (
	// wolfgangPhrase represents Wolfgang's phrases in different languages.
	wolfgangPhrase = nook.Languages{
		language.AmericanEnglish:      wolfgangAmericanEnglishPhrase,
		language.CanadianFrench:       wolfgangCanadianFrenchPhrase,
		language.Dutch:                wolfgangDutchPhrase,
		language.French:               wolfgangFrenchPhrase,
		language.German:               wolfgangGermanPhrase,
		language.Italian:              wolfgangItalianPhrase,
		language.Japanese:             wolfgangJapanesePhrase,
		language.Korean:               wolfgangKoreanPhrase,
		language.LatinAmericanSpanish: wolfgangLatinAmericanSpanishPhrase,
		language.Russian:              wolfgangRussianPhrase,
		language.SimplifiedChinese:    wolfgangSimplifiedChinesePhrase,
		language.Spanish:              wolfgangSpanishPhrase,
		language.TraditionalChinese:   wolfgangTraditionalChinesePhrase}
)

var (
	// Wolfgang represents the character Wolfgang with his complete information.
	Wolfgang = nook.Villager{
		Character:   wolfgangCharacter,
		Personality: personality.Cranky,
		Phrase:      wolfgangPhrase}
)
