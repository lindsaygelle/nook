package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	bobBirthday = nook.Birthday{
		Day:   1,
		Month: time.January}
)

var (
	bobCode = nook.Code{
		Value: "cat00"}
)

var (
	bobAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bob"}

	bobCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Robert"}

	bobDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bob"}

	bobFrenchName = nook.Name{
		Language: language.French,
		Value:    "Robert"}

	bobGermanName = nook.Name{
		Language: language.German,
		Value:    "Jens"}

	bobItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bob"}

	bobJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニコバン"}

	bobKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "히죽"}

	bobLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arándano"}

	bobRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Боб"}

	bobSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿判"}

	bobSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arándano"}

	bobTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿判"}
)

var (
	bobName = nook.Languages{
		language.AmericanEnglish:      bobAmericanEnglishName,
		language.CanadianFrench:       bobCanadianFrenchName,
		language.Dutch:                bobDutchName,
		language.French:               bobFrenchName,
		language.German:               bobGermanName,
		language.Italian:              bobItalianName,
		language.Japanese:             bobJapaneseName,
		language.Korean:               bobKoreanName,
		language.LatinAmericanSpanish: bobLatinAmericanSpanishName,
		language.Russian:              bobRussianName,
		language.SimplifiedChinese:    bobSimplifiedChineseName,
		language.Spanish:              bobSpanishName,
		language.TraditionalChinese:   bobTraditionalChineseName}
)

var (
	bobCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: bobBirthday,
		Code:     bobCode,
		Gender:   gender.Male,
		Name:     bobName}
)

var (
	bobAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pthhpth"}

	bobCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pfioupf"}

	bobDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poekie"}

	bobFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pfioupf"}

	bobGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "psssps"}

	bobItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baffo"}

	bobJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ネコ"}

	bobKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "고양이"}

	bobLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fishfish"}

	bobRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кх-кх-кх"}

	bobSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "猫猫"}

	bobSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zarpas"}

	bobTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貓貓"}
)

var (
	bobPhrase = nook.Languages{
		language.AmericanEnglish:      bobAmericanEnglishName,
		language.CanadianFrench:       bobCanadianFrenchName,
		language.Dutch:                bobDutchName,
		language.French:               bobFrenchName,
		language.German:               bobGermanName,
		language.Italian:              bobItalianName,
		language.Japanese:             bobJapaneseName,
		language.Korean:               bobKoreanName,
		language.LatinAmericanSpanish: bobLatinAmericanSpanishName,
		language.Russian:              bobRussianName,
		language.SimplifiedChinese:    bobSimplifiedChineseName,
		language.Spanish:              bobSpanishName,
		language.TraditionalChinese:   bobTraditionalChineseName}
)

var (
	Bob = nook.Villager{
		Character:   bobCharacter,
		Personality: personality.Lazy,
		Phrase:      bobPhrase}
)
