package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	samsonBirthday = nook.Birthday{
		Day:   5,
		Month: time.July}
)

var (
	samsonCode = nook.Code{
		Value: "mus04"}
)

var (
	samsonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Samson"}

	samsonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Samsonscouic"}

	samsonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Samsonpiepklein"}

	samsonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Samsonscouic"}

	samsonGermanName = nook.Name{
		Language: language.German,
		Value:    "Samsonpieeeps"}

	samsonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sansonepisquak"}

	samsonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピースチュー"}

	samsonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피스쪽"}

	samsonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sansónescuác"}

	samsonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Самсонревушка"}

	samsonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "和平吱"}

	samsonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sansónflojeras"}

	samsonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "和平吱"}
)

var (
	samsonName = nook.Languages{
		language.AmericanEnglish:      samsonAmericanEnglishName,
		language.CanadianFrench:       samsonCanadianFrenchName,
		language.Dutch:                samsonDutchName,
		language.French:               samsonFrenchName,
		language.German:               samsonGermanName,
		language.Italian:              samsonItalianName,
		language.Japanese:             samsonJapaneseName,
		language.Korean:               samsonKoreanName,
		language.LatinAmericanSpanish: samsonLatinAmericanSpanishName,
		language.Russian:              samsonRussianName,
		language.SimplifiedChinese:    samsonSimplifiedChineseName,
		language.Spanish:              samsonSpanishName,
		language.TraditionalChinese:   samsonTraditionalChineseName}
)

var (
	samsonCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: samsonBirthday,
		Code:     samsonCode,
		Gender:   nook.Male,
		Name:     samsonName}
)

var (
	samsonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pipsqueak"}

	samsonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "scouic"}

	samsonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepklein"}

	samsonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "scouic"}

	samsonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pieeeps"}

	samsonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pisquak"}

	samsonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チュー"}

	samsonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쪽"}

	samsonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escuác"}

	samsonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ревушка"}

	samsonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱"}

	samsonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flojeras"}

	samsonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吱"}
)

var (
	samsonPhrase = nook.Languages{
		language.AmericanEnglish:      samsonAmericanEnglishName,
		language.CanadianFrench:       samsonCanadianFrenchName,
		language.Dutch:                samsonDutchName,
		language.French:               samsonFrenchName,
		language.German:               samsonGermanName,
		language.Italian:              samsonItalianName,
		language.Japanese:             samsonJapaneseName,
		language.Korean:               samsonKoreanName,
		language.LatinAmericanSpanish: samsonLatinAmericanSpanishName,
		language.Russian:              samsonRussianName,
		language.SimplifiedChinese:    samsonSimplifiedChineseName,
		language.Spanish:              samsonSpanishName,
		language.TraditionalChinese:   samsonTraditionalChineseName}
)

var (
	Samson = nook.Villager{
		Character:   samsonCharacter,
		Personality: nook.Jock,
		Phrase:      samsonPhrase}
)
