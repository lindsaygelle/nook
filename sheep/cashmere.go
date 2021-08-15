package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Cashmirbêllissime"}

	cashmereDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cashmerebè-⁠ste"}

	cashmereFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cashmirbêllissime"}

	cashmereGermanName = nook.Name{
		Language: language.German,
		Value:    "Lanabäääby"}

	cashmereItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cashmerebeensì"}

	cashmereJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラムールありんす"}

	cashmereKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐시미어포근포근"}

	cashmereLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cachemirbaaa-buá"}

	cashmereRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кашмирбэ-эби"}

	cashmereSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱睦爱"}

	cashmereSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cachemirpeeeque"}

	cashmereTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛睦愛"}
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
		Animal:   Sheep,
		Birthday: cashmereBirthday,
		Code:     cashmereCode,
		Gender:   nook.Female,
		Name:     cashmereName}
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
	Cashmere = nook.Villager{
		Character:   cashmereCharacter,
		Personality: nook.Snooty,
		Phrase:      cashmerePhrase}
)
