package sheep

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
	// cashmereBirthday represents Cashmere's birthday (April 2nd).
	cashmereBirthday = nook.Birthday{
		Day:   2,
		Month: time.April}
)

var (
	// cashmereCode represents Cashmere's unique code ("shp04").
	cashmereCode = nook.Code{
		Value: "shp04"}
)

var (
	// cashmereAmericanEnglishName represents Cashmere's name in American English.
	cashmereAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cashmere"}

	// cashmereCanadianFrenchName represents Cashmere's name in Canadian French.
	cashmereCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cashmir"}

	// cashmereDutchName represents Cashmere's name in Dutch.
	cashmereDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cashmere"}

	// cashmereFrenchName represents Cashmere's name in French.
	cashmereFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cashmir"}

	// cashmereGermanName represents Cashmere's name in German.
	cashmereGermanName = nook.Name{
		Language: language.German,
		Value:    "Lana"}

	// cashmereItalianName represents Cashmere's name in Italian.
	cashmereItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cashmere"}

	// cashmereJapaneseName represents Cashmere's name in Japanese.
	cashmereJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラムール"}

	// cashmereKoreanName represents Cashmere's name in Korean.
	cashmereKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐시미어"}

	// cashmereLatinAmericanSpanishName represents Cashmere's name in Latin American Spanish.
	cashmereLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cachemir"}

	// cashmereRussianName represents Cashmere's name in Russian.
	cashmereRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кашмир"}

	// cashmereSimplifiedChineseName represents Cashmere's name in Simplified Chinese.
	cashmereSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱睦"}

	// cashmereSpanishName represents Cashmere's name in Spanish.
	cashmereSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cachemir"}

	// cashmereTraditionalChineseName represents Cashmere's name in Traditional Chinese.
	cashmereTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛睦"}
)

var (
	// cashmereName represents Cashmere's name in different languages.
	cashmereName = nook.Languages{
		language.AmericanEnglish:      cashmereAmericanEnglishName,
		language.CanadianFrench:       cashmereCanadianFrenchName,
		language.Dutch:                cashmereDutchName,
		language.French:               cashmereFrenchName,
		language.German:               cashmereGermanName,
		language.Italian:              cashmereItalianName,
		language.Japanese:             cashmereJapaneseName,
		language.Korean:               cashmereKoreanName,
		language.LatinAmericanSpanish: cashmereLatinAmericanSpanishName,
		language.Russian:              cashmereRussianName,
		language.SimplifiedChinese:    cashmereSimplifiedChineseName,
		language.Spanish:              cashmereSpanishName,
		language.TraditionalChinese:   cashmereTraditionalChineseName}
)

var (
	// cashmereCharacter represents Cashmere's character information.
	cashmereCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: cashmereBirthday,
		Code:     cashmereCode,
		Key:      character.Cashmere,
		Gender:   gender.Female,
		Name:     cashmereName,
		Special:  false}
)

var (
	// cashmereAmericanEnglishPhrase represents Cashmere's phrase in American English.
	cashmereAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaaby"}

	// cashmereCanadianFrenchPhrase represents Cashmere's phrase in Canadian French.
	cashmereCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bêllissime"}

	// cashmereDutchPhrase represents Cashmere's phrase in Dutch.
	cashmereDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bè-⁠ste"}

	// cashmereFrenchPhrase represents Cashmere's phrase in French.
	cashmereFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bêllissime"}

	// cashmereGermanPhrase represents Cashmere's phrase in German.
	cashmereGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bäääby"}

	// cashmereItalianPhrase represents Cashmere's phrase in Italian.
	cashmereItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beensì"}

	// cashmereJapanesePhrase represents Cashmere's phrase in Japanese.
	cashmereJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ありんす"}

	// cashmereKoreanPhrase represents Cashmere's phrase in Korean.
	cashmereKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "포근포근"}

	// cashmereLatinAmericanSpanishPhrase represents Cashmere's phrase in Latin American Spanish.
	cashmereLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "baaa-buá"}

	// cashmereRussianPhrase represents Cashmere's phrase in Russian.
	cashmereRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бэ-эби"}

	// cashmereSimplifiedChinesePhrase represents Cashmere's phrase in Simplified Chinese.
	cashmereSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱"}

	// cashmereSpanishPhrase represents Cashmere's phrase in Spanish.
	cashmereSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peeeque"}

	// cashmereTraditionalChinesePhrase represents Cashmere's phrase in Traditional Chinese.
	cashmereTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛"}
)

var (
	// cashmerePhrase represents Cashmere's phrases in different languages.
	cashmerePhrase = nook.Languages{
		language.AmericanEnglish:      cashmereAmericanEnglishPhrase,
		language.CanadianFrench:       cashmereCanadianFrenchPhrase,
		language.Dutch:                cashmereDutchPhrase,
		language.French:               cashmereFrenchPhrase,
		language.German:               cashmereGermanPhrase,
		language.Italian:              cashmereItalianPhrase,
		language.Japanese:             cashmereJapanesePhrase,
		language.Korean:               cashmereKoreanPhrase,
		language.LatinAmericanSpanish: cashmereLatinAmericanSpanishPhrase,
		language.Russian:              cashmereRussianPhrase,
		language.SimplifiedChinese:    cashmereSimplifiedChinesePhrase,
		language.Spanish:              cashmereSpanishPhrase,
		language.TraditionalChinese:   cashmereTraditionalChinesePhrase}
)

var (
	// Cashmere represents the character Cashmere with her complete information.
	Cashmere = nook.Villager{
		Character:   cashmereCharacter,
		Personality: personality.Snooty,
		Phrase:      cashmerePhrase}
)
