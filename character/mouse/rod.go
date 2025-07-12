package mouse

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
	// rodBirthday represents rod birthday.
	rodBirthday = nook.Birthday{
		Day:   14,
		Month: time.August}
)

var (
	// rodCode represents rod code.
	rodCode = nook.Code{
		Value: "mus05"}
)

var (
	// rodAmericanEnglishName represents rod american english name.
	rodAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rod"}

	// rodCanadianFrenchName represents rod canadian french name.
	rodCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marcel"}

	// rodDutchName represents rod dutch name.
	rodDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rod"}

	// rodFrenchName represents rod french name.
	rodFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marcel"}

	// rodGermanName represents rod german name.
	rodGermanName = nook.Name{
		Language: language.German,
		Value:    "Manni"}

	// rodItalianName represents rod italian name.
	rodItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Riccardo"}

	// rodJapaneseName represents rod japanese name.
	rodJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャン"}

	// rodKoreanName represents rod korean name.
	rodKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쟝"}

	// rodLatinAmericanSpanishName represents rod latin american spanish name.
	rodLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodi"}

	// rodRussianName represents rod russian name.
	rodRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Род"}

	// rodSimplifiedChineseName represents rod simplified chinese name.
	rodSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿将"}

	// rodSpanishName represents rod spanish name.
	rodSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodi"}

	// rodTraditionalChineseName represents rod traditional chinese name.
	rodTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿將"}
)

var (
	// rodName represents rod name.
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
	// rodCharacter represents rod character.
	rodCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: rodBirthday,
		Code:     rodCode,
		Key:      character.Rod,
		Gender:   gender.Male,
		Name:     rodName,
		Special:  false}
)

var (
	// rodAmericanEnglishPhrase represents rod american english phrase.
	rodAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ace"}

	// rodCanadianFrenchPhrase represents rod canadian french phrase.
	rodCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tip top"}

	// rodDutchPhrase represents rod dutch phrase.
	rodDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toppie"}

	// rodFrenchPhrase represents rod french phrase.
	rodFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tip top"}

	// rodGermanPhrase represents rod german phrase.
	rodGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "boa-ey"}

	// rodItalianPhrase represents rod italian phrase.
	rodItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flapflap"}

	// rodJapanesePhrase represents rod japanese phrase.
	rodJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "すっげぇ"}

	// rodKoreanPhrase represents rod korean phrase.
	rodKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "무진장"}

	// rodLatinAmericanSpanishPhrase represents rod latin american spanish phrase.
	rodLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chopchop"}

	// rodRussianPhrase represents rod russian phrase.
	rodRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мастерски"}

	// rodSimplifiedChinesePhrase represents rod simplified chinese phrase.
	rodSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "厉害"}

	// rodSpanishPhrase represents rod spanish phrase.
	rodSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quesito"}

	// rodTraditionalChinesePhrase represents rod traditional chinese phrase.
	rodTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "厲害"}
)

var (
	// rodPhrase represents rod phrase.
	rodPhrase = nook.Languages{
		language.AmericanEnglish:      rodAmericanEnglishPhrase,
		language.CanadianFrench:       rodCanadianFrenchPhrase,
		language.Dutch:                rodDutchPhrase,
		language.French:               rodFrenchPhrase,
		language.German:               rodGermanPhrase,
		language.Italian:              rodItalianPhrase,
		language.Japanese:             rodJapanesePhrase,
		language.Korean:               rodKoreanPhrase,
		language.LatinAmericanSpanish: rodLatinAmericanSpanishPhrase,
		language.Russian:              rodRussianPhrase,
		language.SimplifiedChinese:    rodSimplifiedChinesePhrase,
		language.Spanish:              rodSpanishPhrase,
		language.TraditionalChinese:   rodTraditionalChinesePhrase}
)

var (
	// Rod represents rod.
	Rod = nook.Villager{
		Character:   rodCharacter,
		Personality: personality.Jock,
		Phrase:      rodPhrase}
)
