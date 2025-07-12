package horse

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
	// victoriaBirthday represents victoria birthday.
	victoriaBirthday = nook.Birthday{
		Day:   11,
		Month: time.July}
)

var (
	// victoriaCode represents victoria code.
	victoriaCode = nook.Code{
		Value: "hrs01"}
)

var (
	// victoriaAmericanEnglishName represents victoria american english name.
	victoriaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Victoria"}

	// victoriaCanadianFrenchName represents victoria canadian french name.
	victoriaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Victoria"}

	// victoriaDutchName represents victoria dutch name.
	victoriaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Victoria"}

	// victoriaFrenchName represents victoria french name.
	victoriaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Victoria"}

	// victoriaGermanName represents victoria german name.
	victoriaGermanName = nook.Name{
		Language: language.German,
		Value:    "Emma"}

	// victoriaItalianName represents victoria italian name.
	victoriaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vittoria"}

	// victoriaJapaneseName represents victoria japanese name.
	victoriaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セントアロー"}

	// victoriaKoreanName represents victoria korean name.
	victoriaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "센트엘로"}

	// victoriaLatinAmericanSpanishName represents victoria latin american spanish name.
	victoriaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Victoria"}

	// victoriaRussianName represents victoria russian name.
	victoriaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Виктория"}

	// victoriaSimplifiedChineseName represents victoria simplified chinese name.
	victoriaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "圣萝"}

	// victoriaSpanishName represents victoria spanish name.
	victoriaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Victoria"}

	// victoriaTraditionalChineseName represents victoria traditional chinese name.
	victoriaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聖蘿"}
)

var (
	// victoriaName represents victoria name.
	victoriaName = nook.Languages{
		language.AmericanEnglish:      victoriaAmericanEnglishName,
		language.CanadianFrench:       victoriaCanadianFrenchName,
		language.Dutch:                victoriaDutchName,
		language.French:               victoriaFrenchName,
		language.German:               victoriaGermanName,
		language.Italian:              victoriaItalianName,
		language.Japanese:             victoriaJapaneseName,
		language.Korean:               victoriaKoreanName,
		language.LatinAmericanSpanish: victoriaLatinAmericanSpanishName,
		language.Russian:              victoriaRussianName,
		language.SimplifiedChinese:    victoriaSimplifiedChineseName,
		language.Spanish:              victoriaSpanishName,
		language.TraditionalChinese:   victoriaTraditionalChineseName}
)

var (
	// victoriaCharacter represents victoria character.
	victoriaCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: victoriaBirthday,
		Code:     victoriaCode,
		Key:      character.Victoria,
		Gender:   gender.Female,
		Name:     victoriaName,
		Special:  false}
)

var (
	// victoriaAmericanEnglishPhrase represents victoria american english phrase.
	victoriaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sugar cube"}

	// victoriaCanadianFrenchPhrase represents victoria canadian french phrase.
	victoriaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "susucre"}

	// victoriaDutchPhrase represents victoria dutch phrase.
	victoriaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klontje"}

	// victoriaFrenchPhrase represents victoria french phrase.
	victoriaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "susucre"}

	// victoriaGermanPhrase represents victoria german phrase.
	victoriaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gimazuka"}

	// victoriaItalianPhrase represents victoria italian phrase.
	victoriaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zolletta"}

	// victoriaJapanesePhrase represents victoria japanese phrase.
	victoriaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いくわよ"}

	// victoriaKoreanPhrase represents victoria korean phrase.
	victoriaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달리자"}

	// victoriaLatinAmericanSpanishPhrase represents victoria latin american spanish phrase.
	victoriaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "clic-cloc"}

	// victoriaRussianPhrase represents victoria russian phrase.
	victoriaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сахарок"}

	// victoriaSimplifiedChinesePhrase represents victoria simplified chinese phrase.
	victoriaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "走哦"}

	// victoriaSpanishPhrase represents victoria spanish phrase.
	victoriaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "terroncito"}

	// victoriaTraditionalChinesePhrase represents victoria traditional chinese phrase.
	victoriaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "走囉"}
)

var (
	// victoriaPhrase represents victoria phrase.
	victoriaPhrase = nook.Languages{
		language.AmericanEnglish:      victoriaAmericanEnglishPhrase,
		language.CanadianFrench:       victoriaCanadianFrenchPhrase,
		language.Dutch:                victoriaDutchPhrase,
		language.French:               victoriaFrenchPhrase,
		language.German:               victoriaGermanPhrase,
		language.Italian:              victoriaItalianPhrase,
		language.Japanese:             victoriaJapanesePhrase,
		language.Korean:               victoriaKoreanPhrase,
		language.LatinAmericanSpanish: victoriaLatinAmericanSpanishPhrase,
		language.Russian:              victoriaRussianPhrase,
		language.SimplifiedChinese:    victoriaSimplifiedChinesePhrase,
		language.Spanish:              victoriaSpanishPhrase,
		language.TraditionalChinese:   victoriaTraditionalChinesePhrase}
)

var (
	// Victoria represents victoria.
	Victoria = nook.Villager{
		Character:   victoriaCharacter,
		Personality: personality.Peppy,
		Phrase:      victoriaPhrase}
)
