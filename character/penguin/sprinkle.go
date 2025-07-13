package penguin

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
	// sprinkleBirthday represents sprinkle birthday.
	sprinkleBirthday = nook.Birthday{
		Day:   20,
		Month: time.February}
)

var (
	// sprinkleCode represents sprinkle code.
	sprinkleCode = nook.Code{
		Value: "pgn14"}
)

var (
	// sprinkleAmericanEnglishName represents sprinkle american english name.
	sprinkleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sprinkle"}

	// sprinkleCanadianFrenchName represents sprinkle canadian french name.
	sprinkleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Laurie"}

	// sprinkleDutchName represents sprinkle dutch name.
	sprinkleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sprinkle"}

	// sprinkleFrenchName represents sprinkle french name.
	sprinkleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laurie"}

	// sprinkleGermanName represents sprinkle german name.
	sprinkleGermanName = nook.Name{
		Language: language.German,
		Value:    "Svenja"}

	// sprinkleItalianName represents sprinkle italian name.
	sprinkleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brina"}

	// sprinkleJapaneseName represents sprinkle japanese name.
	sprinkleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フラッペ"}

	// sprinkleKoreanName represents sprinkle korean name.
	sprinkleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리미"}

	// sprinkleLatinAmericanSpanishName represents sprinkle latin american spanish name.
	sprinkleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rizolda"}

	// sprinkleRussianName represents sprinkle russian name.
	sprinkleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спринкл"}

	// sprinkleSimplifiedChineseName represents sprinkle simplified chinese name.
	sprinkleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "冰莎"}

	// sprinkleSpanishName represents sprinkle spanish name.
	sprinkleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rizolda"}

	// sprinkleTraditionalChineseName represents sprinkle traditional chinese name.
	sprinkleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "冰莎"}
)

var (
	// sprinkleName represents sprinkle name.
	sprinkleName = nook.Languages{
		language.AmericanEnglish:      sprinkleAmericanEnglishName,
		language.CanadianFrench:       sprinkleCanadianFrenchName,
		language.Dutch:                sprinkleDutchName,
		language.French:               sprinkleFrenchName,
		language.German:               sprinkleGermanName,
		language.Italian:              sprinkleItalianName,
		language.Japanese:             sprinkleJapaneseName,
		language.Korean:               sprinkleKoreanName,
		language.LatinAmericanSpanish: sprinkleLatinAmericanSpanishName,
		language.Russian:              sprinkleRussianName,
		language.SimplifiedChinese:    sprinkleSimplifiedChineseName,
		language.Spanish:              sprinkleSpanishName,
		language.TraditionalChinese:   sprinkleTraditionalChineseName}
)

var (
	// sprinkleCharacter represents sprinkle character.
	sprinkleCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: sprinkleBirthday,
		Code:     sprinkleCode,
		Key:      character.Sprinkle,
		Gender:   gender.Female,
		Name:     sprinkleName,
		Special:  false}
)

var (
	// sprinkleAmericanEnglishPhrase represents sprinkle american english phrase.
	sprinkleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "frappe"}

	// sprinkleCanadianFrenchPhrase represents sprinkle canadian french phrase.
	sprinkleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon glaçon"}

	// sprinkleDutchPhrase represents sprinkle dutch phrase.
	sprinkleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "frappé"}

	// sprinkleFrenchPhrase represents sprinkle french phrase.
	sprinkleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon glaçon"}

	// sprinkleGermanPhrase represents sprinkle german phrase.
	sprinkleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bibib"}

	// sprinkleItalianPhrase represents sprinkle italian phrase.
	sprinkleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "unz unz"}

	// sprinkleJapanesePhrase represents sprinkle japanese phrase.
	sprinkleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヒヤリ"}

	// sprinkleKoreanPhrase represents sprinkle korean phrase.
	sprinkleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꽁꽁"}

	// sprinkleLatinAmericanSpanishPhrase represents sprinkle latin american spanish phrase.
	sprinkleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "frapé"}

	// sprinkleRussianPhrase represents sprinkle russian phrase.
	sprinkleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пенка"}

	// sprinkleSimplifiedChinesePhrase represents sprinkle simplified chinese phrase.
	sprinkleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凉爽"}

	// sprinkleSpanishPhrase represents sprinkle spanish phrase.
	sprinkleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esmoquin"}

	// sprinkleTraditionalChinesePhrase represents sprinkle traditional chinese phrase.
	sprinkleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "涼爽"}
)

var (
	// sprinklePhrase represents sprinkle phrase.
	sprinklePhrase = nook.Languages{
		language.AmericanEnglish:      sprinkleAmericanEnglishPhrase,
		language.CanadianFrench:       sprinkleCanadianFrenchPhrase,
		language.Dutch:                sprinkleDutchPhrase,
		language.French:               sprinkleFrenchPhrase,
		language.German:               sprinkleGermanPhrase,
		language.Italian:              sprinkleItalianPhrase,
		language.Japanese:             sprinkleJapanesePhrase,
		language.Korean:               sprinkleKoreanPhrase,
		language.LatinAmericanSpanish: sprinkleLatinAmericanSpanishPhrase,
		language.Russian:              sprinkleRussianPhrase,
		language.SimplifiedChinese:    sprinkleSimplifiedChinesePhrase,
		language.Spanish:              sprinkleSpanishPhrase,
		language.TraditionalChinese:   sprinkleTraditionalChinesePhrase}
)

var (
	// Sprinkle represents sprinkle.
	Sprinkle = nook.Villager{
		Character:   sprinkleCharacter,
		Personality: personality.Peppy,
		Phrase:      sprinklePhrase}
)
