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
	cashmereBirthday = nook.Birthday{
		Day:   2,
		Month: time.April}
)

var (
	cashmereCode = nook.Code{
		Value: "shp04"}
)

var (
	cashmereAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cashmere"}

	cashmereCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cashmir"}

	cashmereDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cashmere"}

	cashmereFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cashmir"}

	cashmereGermanName = nook.Name{
		Language: language.German,
		Value:    "Lana"}

	cashmereItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cashmere"}

	cashmereJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラムール"}

	cashmereKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐시미어"}

	cashmereLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cachemir"}

	cashmereRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кашмир"}

	cashmereSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱睦"}

	cashmereSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cachemir"}

	cashmereTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛睦"}
)

var (
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
	cashmereAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaaby"}

	cashmereCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bêllissime"}

	cashmereDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bè-⁠ste"}

	cashmereFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bêllissime"}

	cashmereGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bäääby"}

	cashmereItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beensì"}

	cashmereJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ありんす"}

	cashmereKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "포근포근"}

	cashmereLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "baaa-buá"}

	cashmereRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бэ-эби"}

	cashmereSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱"}

	cashmereSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peeeque"}

	cashmereTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛"}
)

var (
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
	Cashmere = nook.Villager{
		Character:   cashmereCharacter,
		Personality: personality.Snooty,
		Phrase:      cashmerePhrase}
)
