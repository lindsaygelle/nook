package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	soleilBirthday = nook.Birthday{
		Day:   9,
		Month: time.August}
)

var (
	soleilCode = nook.Code{
		Value: "ham04"}
)

var (
	soleilAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Soleil"}

	soleilCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Stellaciel"}

	soleilDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Soleilkorenwolf"}

	soleilFrenchName = nook.Name{
		Language: language.French,
		Value:    "Stellagnognotte"}

	soleilGermanName = nook.Name{
		Language: language.German,
		Value:    "Samirahümpf"}

	soleilItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cettinavriiin"}

	soleilJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャンティだってば"}

	soleilKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샨티파샤샤"}

	soleilLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sorayañiqui-ñú"}

	soleilRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Солейо-хо-хо"}

	soleilSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安安相信我"}

	soleilSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sorayamoflete"}

	soleilTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安安相信我"}
)

var (
	soleilName = nook.Languages{
		language.AmericanEnglish:      soleilAmericanEnglishName,
		language.CanadianFrench:       soleilCanadianFrenchName,
		language.Dutch:                soleilDutchName,
		language.French:               soleilFrenchName,
		language.German:               soleilGermanName,
		language.Italian:              soleilItalianName,
		language.Japanese:             soleilJapaneseName,
		language.Korean:               soleilKoreanName,
		language.LatinAmericanSpanish: soleilLatinAmericanSpanishName,
		language.Russian:              soleilRussianName,
		language.SimplifiedChinese:    soleilSimplifiedChineseName,
		language.Spanish:              soleilSpanishName,
		language.TraditionalChinese:   soleilTraditionalChineseName}
)

var (
	soleilCharacter = nook.Character{
		Animal:   Hamster,
		Birthday: soleilBirthday,
		Code:     soleilCode,
		Gender:   nook.Female,
		Name:     soleilName}
)

var (
	soleilAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tarnation"}

	soleilCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ciel"}

	soleilDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "korenwolf"}

	soleilFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gnognotte"}

	soleilGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hümpf"}

	soleilItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vriiin"}

	soleilJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だってば"}

	soleilKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파샤샤"}

	soleilLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiqui-ñú"}

	soleilRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "о-хо-хо"}

	soleilSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "相信我"}

	soleilSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "moflete"}

	soleilTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "相信我"}
)

var (
	soleilPhrase = nook.Languages{
		language.AmericanEnglish:      soleilAmericanEnglishName,
		language.CanadianFrench:       soleilCanadianFrenchName,
		language.Dutch:                soleilDutchName,
		language.French:               soleilFrenchName,
		language.German:               soleilGermanName,
		language.Italian:              soleilItalianName,
		language.Japanese:             soleilJapaneseName,
		language.Korean:               soleilKoreanName,
		language.LatinAmericanSpanish: soleilLatinAmericanSpanishName,
		language.Russian:              soleilRussianName,
		language.SimplifiedChinese:    soleilSimplifiedChineseName,
		language.Spanish:              soleilSpanishName,
		language.TraditionalChinese:   soleilTraditionalChineseName}
)

var (
	Soleil = nook.Villager{
		Character:   soleilCharacter,
		Personality: nook.Snooty,
		Phrase:      soleilPhrase}
)
