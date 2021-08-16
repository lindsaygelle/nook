package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	rodBirthday = nook.Birthday{
		Day:   14,
		Month: time.August}
)

var (
	rodCode = nook.Code{
		Value: "mus05"}
)

var (
	rodAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rod"}

	rodCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marcel"}

	rodDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rod"}

	rodFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marcel"}

	rodGermanName = nook.Name{
		Language: language.German,
		Value:    "Manni"}

	rodItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Riccardo"}

	rodJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャン"}

	rodKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쟝"}

	rodLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodi"}

	rodRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Род"}

	rodSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿将"}

	rodSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodi"}

	rodTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿將"}
)

var (
	rodName = nook.Languages{
		language.AmericanEnglish:      rodAmericanEnglishName,
		language.CanadianFrench:       rodCanadianFrenchName,
		language.Dutch:                rodDutchName,
		language.French:               rodFrenchName,
		language.German:               rodGermanName,
		language.Italian:              rodItalianName,
		language.Japanese:             rodJapaneseName,
		language.Korean:               rodKoreanName,
		language.LatinAmericanSpanish: rodLatinAmericanSpanishName,
		language.Russian:              rodRussianName,
		language.SimplifiedChinese:    rodSimplifiedChineseName,
		language.Spanish:              rodSpanishName,
		language.TraditionalChinese:   rodTraditionalChineseName}
)

var (
	rodCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: rodBirthday,
		Code:     rodCode,
		Gender:   gender.Male,
		Name:     rodName}
)

var (
	rodAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ace"}

	rodCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tip top"}

	rodDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toppie"}

	rodFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tip top"}

	rodGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "boa-ey"}

	rodItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flapflap"}

	rodJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "すっげぇ"}

	rodKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "무진장"}

	rodLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chopchop"}

	rodRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мастерски"}

	rodSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "厉害"}

	rodSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quesito"}

	rodTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "厲害"}
)

var (
	rodPhrase = nook.Languages{
		language.AmericanEnglish:      rodAmericanEnglishName,
		language.CanadianFrench:       rodCanadianFrenchName,
		language.Dutch:                rodDutchName,
		language.French:               rodFrenchName,
		language.German:               rodGermanName,
		language.Italian:              rodItalianName,
		language.Japanese:             rodJapaneseName,
		language.Korean:               rodKoreanName,
		language.LatinAmericanSpanish: rodLatinAmericanSpanishName,
		language.Russian:              rodRussianName,
		language.SimplifiedChinese:    rodSimplifiedChineseName,
		language.Spanish:              rodSpanishName,
		language.TraditionalChinese:   rodTraditionalChineseName}
)

var (
	Rod = nook.Villager{
		Character:   rodCharacter,
		Personality: personality.Jock,
		Phrase:      rodPhrase}
)
