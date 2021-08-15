package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	moeBirthday = nook.Birthday{
		Day:   12,
		Month: time.January}
)

var (
	moeCode = nook.Code{
		Value: "cat08"}
)

var (
	moeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Moe"}

	moeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Momomiou"}

	moeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Moegeeuweldig"}

	moeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Momomiou"}

	moeGermanName = nook.Name{
		Language: language.German,
		Value:    "Tristanschluchz"}

	moeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moemyawn"}

	moeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジンペイな"}

	moeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "진상그치"}

	moeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tristángrrrroou"}

	moeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Моэмяу-а-а-а"}

	moeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仁平呐"}

	moeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tristángrrrroou"}

	moeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仁平吶"}
)

var (
	moeName = nook.Languages{
		language.AmericanEnglish:      moeAmericanEnglishName,
		language.CanadianFrench:       moeCanadianFrenchName,
		language.Dutch:                moeDutchName,
		language.French:               moeFrenchName,
		language.German:               moeGermanName,
		language.Italian:              moeItalianName,
		language.Japanese:             moeJapaneseName,
		language.Korean:               moeKoreanName,
		language.LatinAmericanSpanish: moeLatinAmericanSpanishName,
		language.Russian:              moeRussianName,
		language.SimplifiedChinese:    moeSimplifiedChineseName,
		language.Spanish:              moeSpanishName,
		language.TraditionalChinese:   moeTraditionalChineseName}
)

var (
	moeCharacter = nook.Character{
		Animal:   Cat,
		Birthday: moeBirthday,
		Code:     moeCode,
		Gender:   nook.Male,
		Name:     moeName}
)

var (
	moeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "myawn"}

	moeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miou"}

	moeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "geeuweldig"}

	moeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miou"}

	moeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schluchz"}

	moeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "myawn"}

	moeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "な"}

	moeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그치"}

	moeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrrroou"}

	moeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мяу-а-а-а"}

	moeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呐"}

	moeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grrrroou"}

	moeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吶"}
)

var (
	moePhrase = nook.Languages{
		language.AmericanEnglish:      moeAmericanEnglishName,
		language.CanadianFrench:       moeCanadianFrenchName,
		language.Dutch:                moeDutchName,
		language.French:               moeFrenchName,
		language.German:               moeGermanName,
		language.Italian:              moeItalianName,
		language.Japanese:             moeJapaneseName,
		language.Korean:               moeKoreanName,
		language.LatinAmericanSpanish: moeLatinAmericanSpanishName,
		language.Russian:              moeRussianName,
		language.SimplifiedChinese:    moeSimplifiedChineseName,
		language.Spanish:              moeSpanishName,
		language.TraditionalChinese:   moeTraditionalChineseName}
)

var (
	Moe = nook.Villager{
		Character:   moeCharacter,
		Personality: nook.Lazy,
		Phrase:      moePhrase}
)
