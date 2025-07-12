package cat

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
	// moeBirthday represents moe birthday.
	moeBirthday = nook.Birthday{
		Day:   12,
		Month: time.January}
)

var (
	// moeCode represents moe code.
	moeCode = nook.Code{
		Value: "cat08"}
)

var (
	// moeAmericanEnglishName represents moe american english name.
	moeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Moe"}

	// moeCanadianFrenchName represents moe canadian french name.
	moeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Momo"}

	// moeDutchName represents moe dutch name.
	moeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Moe"}

	// moeFrenchName represents moe french name.
	moeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Momo"}

	// moeGermanName represents moe german name.
	moeGermanName = nook.Name{
		Language: language.German,
		Value:    "Tristan"}

	// moeItalianName represents moe italian name.
	moeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moe"}

	// moeJapaneseName represents moe japanese name.
	moeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジンペイ"}

	// moeKoreanName represents moe korean name.
	moeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "진상"}

	// moeLatinAmericanSpanishName represents moe latin american spanish name.
	moeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tristán"}

	// moeRussianName represents moe russian name.
	moeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Моэ"}

	// moeSimplifiedChineseName represents moe simplified chinese name.
	moeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仁平"}

	// moeSpanishName represents moe spanish name.
	moeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tristán"}

	// moeTraditionalChineseName represents moe traditional chinese name.
	moeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仁平"}
)

var (
	// moeName represents moe name.
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
	// moeCharacter represents moe character.
	moeCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: moeBirthday,
		Code:     moeCode,
		Key:      character.Moe,
		Gender:   gender.Male,
		Name:     moeName,
		Special:  false}
)

var (
	// moeAmericanEnglishPhrase represents moe american english phrase.
	moeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "myawn"}

	// moeCanadianFrenchPhrase represents moe canadian french phrase.
	moeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miou"}

	// moeDutchPhrase represents moe dutch phrase.
	moeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "geeuweldig"}

	// moeFrenchPhrase represents moe french phrase.
	moeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miou"}

	// moeGermanPhrase represents moe german phrase.
	moeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schluchz"}

	// moeItalianPhrase represents moe italian phrase.
	moeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "myawn"}

	// moeJapanesePhrase represents moe japanese phrase.
	moeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "な"}

	// moeKoreanPhrase represents moe korean phrase.
	moeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그치"}

	// moeLatinAmericanSpanishPhrase represents moe latin american spanish phrase.
	moeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrrroou"}

	// moeRussianPhrase represents moe russian phrase.
	moeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мяу-а-а-а"}

	// moeSimplifiedChinesePhrase represents moe simplified chinese phrase.
	moeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呐"}

	// moeSpanishPhrase represents moe spanish phrase.
	moeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grrrroou"}

	// moeTraditionalChinesePhrase represents moe traditional chinese phrase.
	moeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吶"}
)

var (
	// moePhrase represents moe phrase.
	moePhrase = nook.Languages{
		language.AmericanEnglish:      moeAmericanEnglishPhrase,
		language.CanadianFrench:       moeCanadianFrenchPhrase,
		language.Dutch:                moeDutchPhrase,
		language.French:               moeFrenchPhrase,
		language.German:               moeGermanPhrase,
		language.Italian:              moeItalianPhrase,
		language.Japanese:             moeJapanesePhrase,
		language.Korean:               moeKoreanPhrase,
		language.LatinAmericanSpanish: moeLatinAmericanSpanishPhrase,
		language.Russian:              moeRussianPhrase,
		language.SimplifiedChinese:    moeSimplifiedChinesePhrase,
		language.Spanish:              moeSpanishPhrase,
		language.TraditionalChinese:   moeTraditionalChinesePhrase}
)

var (
	// Moe represents moe.
	Moe = nook.Villager{
		Character:   moeCharacter,
		Personality: personality.Lazy,
		Phrase:      moePhrase}
)
