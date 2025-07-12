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
	// stitchesBirthday represents stitches birthday.
	stitchesBirthday = nook.Birthday{
		Day:   10,
		Month: time.February}
)

var (
	// stitchesCode represents stitches code.
	stitchesCode = nook.Code{
		Value: "cbr05"}
)

var (
	// stitchesAmericanEnglishName represents stitches american english name.
	stitchesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stitches"}

	// stitchesCanadianFrenchName represents stitches canadian french name.
	stitchesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miro"}

	// stitchesDutchName represents stitches dutch name.
	stitchesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stitches"}

	// stitchesFrenchName represents stitches french name.
	stitchesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miro"}

	// stitchesGermanName represents stitches german name.
	stitchesGermanName = nook.Name{
		Language: language.German,
		Value:    "Berry"}

	// stitchesItalianName represents stitches italian name.
	stitchesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Toppetta"}

	// stitchesJapaneseName represents stitches japanese name.
	stitchesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パッチ"}

	// stitchesKoreanName represents stitches korean name.
	stitchesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패치"}

	// stitchesLatinAmericanSpanishName represents stitches latin american spanish name.
	stitchesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Parches"}

	// stitchesRussianName represents stitches russian name.
	stitchesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стичес"}

	// stitchesSimplifiedChineseName represents stitches simplified chinese name.
	stitchesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玩具熊"}

	// stitchesSpanishName represents stitches spanish name.
	stitchesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Parches"}

	// stitchesTraditionalChineseName represents stitches traditional chinese name.
	stitchesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玩具熊"}
)

var (
	// stitchesName represents stitches name.
	stitchesName = nook.Languages{
		language.AmericanEnglish:      stitchesAmericanEnglishName,
		language.CanadianFrench:       stitchesCanadianFrenchName,
		language.Dutch:                stitchesDutchName,
		language.French:               stitchesFrenchName,
		language.German:               stitchesGermanName,
		language.Italian:              stitchesItalianName,
		language.Japanese:             stitchesJapaneseName,
		language.Korean:               stitchesKoreanName,
		language.LatinAmericanSpanish: stitchesLatinAmericanSpanishName,
		language.Russian:              stitchesRussianName,
		language.SimplifiedChinese:    stitchesSimplifiedChineseName,
		language.Spanish:              stitchesSpanishName,
		language.TraditionalChinese:   stitchesTraditionalChineseName}
)

var (
	// stitchesCharacter represents stitches character.
	stitchesCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: stitchesBirthday,
		Code:     stitchesCode,
		Key:      character.Stitches,
		Gender:   gender.Male,
		Name:     stitchesName,
		Special:  false}
)

var (
	// stitchesAmericanEnglishPhrase represents stitches american english phrase.
	stitchesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stuffin'"}

	// stitchesCanadianFrenchPhrase represents stitches canadian french phrase.
	stitchesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ou bien"}

	// stitchesDutchPhrase represents stitches dutch phrase.
	stitchesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pluche"}

	// stitchesFrenchPhrase represents stitches french phrase.
	stitchesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ou bien"}

	// stitchesGermanPhrase represents stitches german phrase.
	stitchesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brummm"}

	// stitchesItalianPhrase represents stitches italian phrase.
	stitchesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ohilà"}

	// stitchesJapanesePhrase represents stitches japanese phrase.
	stitchesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのれす"}

	// stitchesKoreanPhrase represents stitches korean phrase.
	stitchesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그런거죠"}

	// stitchesLatinAmericanSpanishPhrase represents stitches latin american spanish phrase.
	stitchesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "paguahhh"}

	// stitchesRussianPhrase represents stitches russian phrase.
	stitchesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "плюш-плюш"}

	// stitchesSimplifiedChinesePhrase represents stitches simplified chinese phrase.
	stitchesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玩玩"}

	// stitchesSpanishPhrase represents stitches spanish phrase.
	stitchesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peluche"}

	// stitchesTraditionalChinesePhrase represents stitches traditional chinese phrase.
	stitchesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玩玩"}
)

var (
	// stitchesPhrase represents stitches phrase.
	stitchesPhrase = nook.Languages{
		language.AmericanEnglish:      stitchesAmericanEnglishPhrase,
		language.CanadianFrench:       stitchesCanadianFrenchPhrase,
		language.Dutch:                stitchesDutchPhrase,
		language.French:               stitchesFrenchPhrase,
		language.German:               stitchesGermanPhrase,
		language.Italian:              stitchesItalianPhrase,
		language.Japanese:             stitchesJapanesePhrase,
		language.Korean:               stitchesKoreanPhrase,
		language.LatinAmericanSpanish: stitchesLatinAmericanSpanishPhrase,
		language.Russian:              stitchesRussianPhrase,
		language.SimplifiedChinese:    stitchesSimplifiedChinesePhrase,
		language.Spanish:              stitchesSpanishPhrase,
		language.TraditionalChinese:   stitchesTraditionalChinesePhrase}
)

var (
	// Stitches represents stitches.
	Stitches = nook.Villager{
		Character:   stitchesCharacter,
		Personality: personality.Lazy,
		Phrase:      stitchesPhrase}
)
