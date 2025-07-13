package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// olafBirthday represents Olaf's birthday.
var (
	// olafBirthday represents olaf birthday.
	olafBirthday = nook.Birthday{
		Day:   19,
		Month: time.May}
)

// olafCode represents Olaf's unique code.
var (
	// olafCode represents olaf code.
	olafCode = nook.Code{
		Value: "ant09"}
)

// Different names for Olaf in various languages.
var (
	// olafAmericanEnglishName represents olaf american english name.
	olafAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Olaf"}

	// olafCanadianFrenchName represents olaf canadian french name.
	olafCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blair"}

	// olafDutchName represents olaf dutch name.
	olafDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olaf"}

	// olafFrenchName represents olaf french name.
	olafFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blair"}

	// olafGermanName represents olaf german name.
	olafGermanName = nook.Name{
		Language: language.German,
		Value:    "Olaf"}

	// olafItalianName represents olaf italian name.
	olafItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ettore"}

	// olafJapaneseName represents olaf japanese name.
	olafJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アントニオ"}

	// olafKoreanName represents olaf korean name.
	olafKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안토니오"}

	// olafLatinAmericanSpanishName represents olaf latin american spanish name.
	olafLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Olaf"}

	// olafRussianName represents olaf russian name.
	olafRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Олаф"}

	// olafSimplifiedChineseName represents olaf simplified chinese name.
	olafSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安东尼"}

	// olafSpanishName represents olaf spanish name.
	olafSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Olaf"}

	// olafTraditionalChineseName represents olaf traditional chinese name.
	olafTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安東尼"}
)

// olafName represents Olaf's name in different languages.
var (
	// olafName represents olaf name.
	olafName = nook.Languages{
		language.AmericanEnglish:      olafAmericanEnglishName,
		language.CanadianFrench:       olafCanadianFrenchName,
		language.Dutch:                olafDutchName,
		language.French:               olafFrenchName,
		language.German:               olafGermanName,
		language.Italian:              olafItalianName,
		language.Japanese:             olafJapaneseName,
		language.Korean:               olafKoreanName,
		language.LatinAmericanSpanish: olafLatinAmericanSpanishName,
		language.Russian:              olafRussianName,
		language.SimplifiedChinese:    olafSimplifiedChineseName,
		language.Spanish:              olafSpanishName,
		language.TraditionalChinese:   olafTraditionalChineseName}
)

// olafCharacter represents Olaf's character information.
var (
	// olafCharacter represents olaf character.
	olafCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: olafBirthday,
		Code:     olafCode,
		Key:      character.Olaf,
		Gender:   gender.Male,
		Name:     olafName,
		Special:  false}
)

// Different phrases spoken by Olaf in various languages.
var (
	// olafAmericanEnglishPhrase represents olaf american english phrase.
	olafAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whiffa"}

	// olafCanadianFrenchPhrase represents olaf canadian french phrase.
	olafCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yesss"}

	// olafDutchPhrase represents olaf dutch phrase.
	olafDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toet"}

	// olafFrenchPhrase represents olaf french phrase.
	olafFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tartarin"}

	// olafGermanPhrase represents olaf german phrase.
	olafGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mwahaaa"}

	// olafItalianPhrase represents olaf italian phrase.
	olafItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "swish"}

	// olafJapanesePhrase represents olaf japanese phrase.
	olafJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムーチョ"}

	// olafKoreanPhrase represents olaf korean phrase.
	olafKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "무쵸"}

	// olafLatinAmericanSpanishPhrase represents olaf latin american spanish phrase.
	olafLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fruf"}

	// olafRussianPhrase represents olaf russian phrase.
	olafRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уф-уф"}

	// olafSimplifiedChinesePhrase represents olaf simplified chinese phrase.
	olafSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧啦"}

	// olafSpanishPhrase represents olaf spanish phrase.
	olafSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fruf"}

	// olafTraditionalChinesePhrase represents olaf traditional chinese phrase.
	olafTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐啦"}
)

// olafPhrase represents Olaf's phrases in different languages.
var (
	// olafPhrase represents olaf phrase.
	olafPhrase = nook.Languages{
		language.AmericanEnglish:      olafAmericanEnglishPhrase,
		language.CanadianFrench:       olafCanadianFrenchPhrase,
		language.Dutch:                olafDutchPhrase,
		language.French:               olafFrenchPhrase,
		language.German:               olafGermanPhrase,
		language.Italian:              olafItalianPhrase,
		language.Japanese:             olafJapanesePhrase,
		language.Korean:               olafKoreanPhrase,
		language.LatinAmericanSpanish: olafLatinAmericanSpanishPhrase,
		language.Russian:              olafRussianPhrase,
		language.SimplifiedChinese:    olafSimplifiedChinesePhrase,
		language.Spanish:              olafSpanishPhrase,
		language.TraditionalChinese:   olafTraditionalChinesePhrase}
)

// Olaf represents the character Olaf with his complete information.
var (
	// Olaf represents olaf.
	Olaf = nook.Villager{
		Character:   olafCharacter,
		Personality: personality.Smug,
		Phrase:      olafPhrase}
)
