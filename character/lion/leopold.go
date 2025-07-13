package lion

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
	// leopoldBirthday represents leopold birthday.
	leopoldBirthday = nook.Birthday{
		Day:   14,
		Month: time.August}
)

var (
	// leopoldCode represents leopold code.
	leopoldCode = nook.Code{
		Value: "lon04"}
)

var (
	// leopoldAmericanEnglishName represents leopold american english name.
	leopoldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leopold"}

	// leopoldCanadianFrenchName represents leopold canadian french name.
	leopoldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Leandro"}

	// leopoldDutchName represents leopold dutch name.
	leopoldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leopold"}

	// leopoldFrenchName represents leopold french name.
	leopoldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Leandro"}

	// leopoldGermanName represents leopold german name.
	leopoldGermanName = nook.Name{
		Language: language.German,
		Value:    "Leandro"}

	// leopoldItalianName represents leopold italian name.
	leopoldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leandro"}

	// leopoldJapaneseName represents leopold japanese name.
	leopoldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティーチャー"}

	// leopoldKoreanName represents leopold korean name.
	leopoldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티처"}

	// leopoldLatinAmericanSpanishName represents leopold latin american spanish name.
	leopoldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leopoldo"}

	// leopoldRussianName represents leopold russian name.
	leopoldRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Леопольд"}

	// leopoldSimplifiedChineseName represents leopold simplified chinese name.
	leopoldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "老狮"}

	// leopoldSpanishName represents leopold spanish name.
	leopoldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leopoldo"}

	// leopoldTraditionalChineseName represents leopold traditional chinese name.
	leopoldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "老獅"}
)

var (
	// leopoldName represents leopold name.
	leopoldName = nook.Languages{
		language.AmericanEnglish:      leopoldAmericanEnglishName,
		language.CanadianFrench:       leopoldCanadianFrenchName,
		language.Dutch:                leopoldDutchName,
		language.French:               leopoldFrenchName,
		language.German:               leopoldGermanName,
		language.Italian:              leopoldItalianName,
		language.Japanese:             leopoldJapaneseName,
		language.Korean:               leopoldKoreanName,
		language.LatinAmericanSpanish: leopoldLatinAmericanSpanishName,
		language.Russian:              leopoldRussianName,
		language.SimplifiedChinese:    leopoldSimplifiedChineseName,
		language.Spanish:              leopoldSpanishName,
		language.TraditionalChinese:   leopoldTraditionalChineseName}
)

var (
	// leopoldCharacter represents leopold character.
	leopoldCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: leopoldBirthday,
		Code:     leopoldCode,
		Key:      character.Leopold,
		Gender:   gender.Male,
		Name:     leopoldName,
		Special:  false}
)

var (
	// leopoldAmericanEnglishPhrase represents leopold american english phrase.
	leopoldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lion cub"}

	// leopoldCanadianFrenchPhrase represents leopold canadian french phrase.
	leopoldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "minus"}

	// leopoldDutchPhrase represents leopold dutch phrase.
	leopoldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jong"}

	// leopoldFrenchPhrase represents leopold french phrase.
	leopoldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "minus"}

	// leopoldGermanPhrase represents leopold german phrase.
	leopoldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "löwieh"}

	// leopoldItalianPhrase represents leopold italian phrase.
	leopoldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sonré"}

	// leopoldJapanesePhrase represents leopold japanese phrase.
	leopoldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よな"}

	// leopoldKoreanPhrase represents leopold korean phrase.
	leopoldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "졸지마"}

	// leopoldLatinAmericanSpanishPhrase represents leopold latin american spanish phrase.
	leopoldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gro-grau"}

	// leopoldRussianPhrase represents leopold russian phrase.
	leopoldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "левушка"}

	// leopoldSimplifiedChinesePhrase represents leopold simplified chinese phrase.
	leopoldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "上课了"}

	// leopoldSpanishPhrase represents leopold spanish phrase.
	leopoldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fiera"}

	// leopoldTraditionalChinesePhrase represents leopold traditional chinese phrase.
	leopoldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "上課了"}
)

var (
	// leopoldPhrase represents leopold phrase.
	leopoldPhrase = nook.Languages{
		language.AmericanEnglish:      leopoldAmericanEnglishPhrase,
		language.CanadianFrench:       leopoldCanadianFrenchPhrase,
		language.Dutch:                leopoldDutchPhrase,
		language.French:               leopoldFrenchPhrase,
		language.German:               leopoldGermanPhrase,
		language.Italian:              leopoldItalianPhrase,
		language.Japanese:             leopoldJapanesePhrase,
		language.Korean:               leopoldKoreanPhrase,
		language.LatinAmericanSpanish: leopoldLatinAmericanSpanishPhrase,
		language.Russian:              leopoldRussianPhrase,
		language.SimplifiedChinese:    leopoldSimplifiedChinesePhrase,
		language.Spanish:              leopoldSpanishPhrase,
		language.TraditionalChinese:   leopoldTraditionalChinesePhrase}
)

var (
	// Leopold represents leopold.
	Leopold = nook.Villager{
		Character:   leopoldCharacter,
		Personality: personality.Smug,
		Phrase:      leopoldPhrase}
)
