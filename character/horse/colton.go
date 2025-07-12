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
	// coltonBirthday represents colton birthday.
	coltonBirthday = nook.Birthday{
		Day:   22,
		Month: time.May}
)

var (
	// coltonCode represents colton code.
	coltonCode = nook.Code{
		Value: "hrs11"}
)

var (
	// coltonAmericanEnglishName represents colton american english name.
	coltonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Colton"}

	// coltonCanadianFrenchName represents colton canadian french name.
	coltonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tony"}

	// coltonDutchName represents colton dutch name.
	coltonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Colton"}

	// coltonFrenchName represents colton french name.
	coltonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tony"}

	// coltonGermanName represents colton german name.
	coltonGermanName = nook.Name{
		Language: language.German,
		Value:    "Carsten"}

	// coltonItalianName represents colton italian name.
	coltonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Furio"}

	// coltonJapaneseName represents colton japanese name.
	coltonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンソニー"}

	// coltonKoreanName represents colton korean name.
	coltonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안소니"}

	// coltonLatinAmericanSpanishName represents colton latin american spanish name.
	coltonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Furio"}

	// coltonRussianName represents colton russian name.
	coltonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Колтон"}

	// coltonSimplifiedChineseName represents colton simplified chinese name.
	coltonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安索尼"}

	// coltonSpanishName represents colton spanish name.
	coltonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Furio"}

	// coltonTraditionalChineseName represents colton traditional chinese name.
	coltonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安索尼"}
)

var (
	// coltonName represents colton name.
	coltonName = nook.Languages{
		language.AmericanEnglish:      coltonAmericanEnglishName,
		language.CanadianFrench:       coltonCanadianFrenchName,
		language.Dutch:                coltonDutchName,
		language.French:               coltonFrenchName,
		language.German:               coltonGermanName,
		language.Italian:              coltonItalianName,
		language.Japanese:             coltonJapaneseName,
		language.Korean:               coltonKoreanName,
		language.LatinAmericanSpanish: coltonLatinAmericanSpanishName,
		language.Russian:              coltonRussianName,
		language.SimplifiedChinese:    coltonSimplifiedChineseName,
		language.Spanish:              coltonSpanishName,
		language.TraditionalChinese:   coltonTraditionalChineseName}
)

var (
	// coltonCharacter represents colton character.
	coltonCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: coltonBirthday,
		Code:     coltonCode,
		Key:      character.Colton,
		Gender:   gender.Male,
		Name:     coltonName,
		Special:  false}
)

var (
	// coltonAmericanEnglishPhrase represents colton american english phrase.
	coltonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "check it"}

	// coltonCanadianFrenchPhrase represents colton canadian french phrase.
	coltonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hue hue"}

	// coltonDutchPhrase represents colton dutch phrase.
	coltonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hortsik"}

	// coltonFrenchPhrase represents colton french phrase.
	coltonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hue hue"}

	// coltonGermanPhrase represents colton german phrase.
	coltonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hamham"}

	// coltonItalianPhrase represents colton italian phrase.
	coltonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hiii"}

	// coltonJapanesePhrase represents colton japanese phrase.
	coltonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ごらんよ"}

	// coltonKoreanPhrase represents colton korean phrase.
	coltonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이거봐"}

	// coltonLatinAmericanSpanishPhrase represents colton latin american spanish phrase.
	coltonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jia"}

	// coltonRussianPhrase represents colton russian phrase.
	coltonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "глянь"}

	// coltonSimplifiedChinesePhrase represents colton simplified chinese phrase.
	coltonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "请看"}

	// coltonSpanishPhrase represents colton spanish phrase.
	coltonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jia"}

	// coltonTraditionalChinesePhrase represents colton traditional chinese phrase.
	coltonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "請看"}
)

var (
	// coltonPhrase represents colton phrase.
	coltonPhrase = nook.Languages{
		language.AmericanEnglish:      coltonAmericanEnglishPhrase,
		language.CanadianFrench:       coltonCanadianFrenchPhrase,
		language.Dutch:                coltonDutchPhrase,
		language.French:               coltonFrenchPhrase,
		language.German:               coltonGermanPhrase,
		language.Italian:              coltonItalianPhrase,
		language.Japanese:             coltonJapanesePhrase,
		language.Korean:               coltonKoreanPhrase,
		language.LatinAmericanSpanish: coltonLatinAmericanSpanishPhrase,
		language.Russian:              coltonRussianPhrase,
		language.SimplifiedChinese:    coltonSimplifiedChinesePhrase,
		language.Spanish:              coltonSpanishPhrase,
		language.TraditionalChinese:   coltonTraditionalChinesePhrase}
)

var (
	// Colton represents colton.
	Colton = nook.Villager{
		Character:   coltonCharacter,
		Personality: personality.Smug,
		Phrase:      coltonPhrase}
)
