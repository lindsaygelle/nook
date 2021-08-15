package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	murphyBirthday = nook.Birthday{
		Day:   29,
		Month: time.December}
)

var (
	murphyCode = nook.Code{
		Value: "cbr07"}
)

var (
	murphyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Murphy"}

	murphyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eddiefait que"}

	murphyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Murphyspruit"}

	murphyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eddiecanaille"}

	murphyGermanName = nook.Name{
		Language: language.German,
		Value:    "Michaelliebelein"}

	murphyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marcellogrizzel"}

	murphyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりおですばい"}

	murphyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "머피샐리"}

	murphyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Orsónheee"}

	murphyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мерфитоварищ"}

	murphySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宪雄罢了"}

	murphySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Orsónpelusita"}

	murphyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "憲雄罷了"}
)

var (
	murphyName = nook.Languages{
		language.AmericanEnglish:      murphyAmericanEnglishName,
		language.CanadianFrench:       murphyCanadianFrenchName,
		language.Dutch:                murphyDutchName,
		language.French:               murphyFrenchName,
		language.German:               murphyGermanName,
		language.Italian:              murphyItalianName,
		language.Japanese:             murphyJapaneseName,
		language.Korean:               murphyKoreanName,
		language.LatinAmericanSpanish: murphyLatinAmericanSpanishName,
		language.Russian:              murphyRussianName,
		language.SimplifiedChinese:    murphySimplifiedChineseName,
		language.Spanish:              murphySpanishName,
		language.TraditionalChinese:   murphyTraditionalChineseName}
)

var (
	murphyCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: murphyBirthday,
		Code:     murphyCode,
		Gender:   nook.Male,
		Name:     murphyName}
)

var (
	murphyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "laddiemalarkey"}

	murphyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fait que"}

	murphyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "spruit"}

	murphyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canaille"}

	murphyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "liebelein"}

	murphyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grizzel"}

	murphyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですばい"}

	murphyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "샐리"}

	murphyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "heee"}

	murphyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "товарищ"}

	murphySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罢了"}

	murphySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pelusita"}

	murphyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "罷了"}
)

var (
	murphyPhrase = nook.Languages{
		language.AmericanEnglish:      murphyAmericanEnglishName,
		language.CanadianFrench:       murphyCanadianFrenchName,
		language.Dutch:                murphyDutchName,
		language.French:               murphyFrenchName,
		language.German:               murphyGermanName,
		language.Italian:              murphyItalianName,
		language.Japanese:             murphyJapaneseName,
		language.Korean:               murphyKoreanName,
		language.LatinAmericanSpanish: murphyLatinAmericanSpanishName,
		language.Russian:              murphyRussianName,
		language.SimplifiedChinese:    murphySimplifiedChineseName,
		language.Spanish:              murphySpanishName,
		language.TraditionalChinese:   murphyTraditionalChineseName}
)

var (
	Murphy = nook.Villager{
		Character:   murphyCharacter,
		Personality: nook.Cranky,
		Phrase:      murphyPhrase}
)
