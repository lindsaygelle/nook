package bearcub

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
		Value:    "Eddie"}

	murphyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Murphy"}

	murphyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eddie"}

	murphyGermanName = nook.Name{
		Language: language.German,
		Value:    "Michael"}

	murphyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marcello"}

	murphyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりお"}

	murphyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "머피"}

	murphyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Orsón"}

	murphyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мерфи"}

	murphySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宪雄"}

	murphySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Orsón"}

	murphyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "憲雄"}
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
		Animal:   animal.Bearcub,
		Birthday: murphyBirthday,
		Code:     murphyCode,
		Key:      character.Murphy,
		Gender:   gender.Male,
		Name:     murphyName}
)

var (
	murphyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "laddie"}

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
		language.AmericanEnglish:      murphyAmericanEnglishPhrase,
		language.CanadianFrench:       murphyCanadianFrenchPhrase,
		language.Dutch:                murphyDutchPhrase,
		language.French:               murphyFrenchPhrase,
		language.German:               murphyGermanPhrase,
		language.Italian:              murphyItalianPhrase,
		language.Japanese:             murphyJapanesePhrase,
		language.Korean:               murphyKoreanPhrase,
		language.LatinAmericanSpanish: murphyLatinAmericanSpanishPhrase,
		language.Russian:              murphyRussianPhrase,
		language.SimplifiedChinese:    murphySimplifiedChinesePhrase,
		language.Spanish:              murphySpanishPhrase,
		language.TraditionalChinese:   murphyTraditionalChinesePhrase}
)

var (
	Murphy = nook.Villager{
		Character:   murphyCharacter,
		Personality: personality.Cranky,
		Phrase:      murphyPhrase}
)
