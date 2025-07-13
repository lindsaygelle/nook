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
	// jubeiBirthday represents jubei birthday.
	jubeiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// jubeiCode represents jubei code.
	jubeiCode = nook.Code{
		Value: ""}
)

var (
	// jubeiAmericanEnglishName represents jubei american english name.
	jubeiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jūbei"}

	// jubeiCanadianFrenchName represents jubei canadian french name.
	jubeiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// jubeiDutchName represents jubei dutch name.
	jubeiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// jubeiFrenchName represents jubei french name.
	jubeiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// jubeiGermanName represents jubei german name.
	jubeiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// jubeiItalianName represents jubei italian name.
	jubeiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// jubeiJapaneseName represents jubei japanese name.
	jubeiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュウベエ"}

	// jubeiKoreanName represents jubei korean name.
	jubeiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// jubeiLatinAmericanSpanishName represents jubei latin american spanish name.
	jubeiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// jubeiRussianName represents jubei russian name.
	jubeiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// jubeiSimplifiedChineseName represents jubei simplified chinese name.
	jubeiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// jubeiSpanishName represents jubei spanish name.
	jubeiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// jubeiTraditionalChineseName represents jubei traditional chinese name.
	jubeiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// jubeiName represents jubei name.
	jubeiName = nook.Languages{
		language.AmericanEnglish:      jubeiAmericanEnglishName,
		language.CanadianFrench:       jubeiCanadianFrenchName,
		language.Dutch:                jubeiDutchName,
		language.French:               jubeiFrenchName,
		language.German:               jubeiGermanName,
		language.Italian:              jubeiItalianName,
		language.Japanese:             jubeiJapaneseName,
		language.Korean:               jubeiKoreanName,
		language.LatinAmericanSpanish: jubeiLatinAmericanSpanishName,
		language.Russian:              jubeiRussianName,
		language.SimplifiedChinese:    jubeiSimplifiedChineseName,
		language.Spanish:              jubeiSpanishName,
		language.TraditionalChinese:   jubeiTraditionalChineseName}
)

var (
	// jubeiCharacter represents jubei character.
	jubeiCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: jubeiBirthday,
		Code:     jubeiCode,
		Key:      character.Jubei,
		Gender:   gender.Male,
		Name:     jubeiName,
		Special:  false}
)

var (
	// jubeiAmericanEnglishPhrase represents jubei american english phrase.
	jubeiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "いかにも"}

	// jubeiCanadianFrenchPhrase represents jubei canadian french phrase.
	jubeiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// jubeiDutchPhrase represents jubei dutch phrase.
	jubeiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// jubeiFrenchPhrase represents jubei french phrase.
	jubeiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// jubeiGermanPhrase represents jubei german phrase.
	jubeiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// jubeiItalianPhrase represents jubei italian phrase.
	jubeiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// jubeiJapanesePhrase represents jubei japanese phrase.
	jubeiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いかにも"}

	// jubeiKoreanPhrase represents jubei korean phrase.
	jubeiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// jubeiLatinAmericanSpanishPhrase represents jubei latin american spanish phrase.
	jubeiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// jubeiRussianPhrase represents jubei russian phrase.
	jubeiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// jubeiSimplifiedChinesePhrase represents jubei simplified chinese phrase.
	jubeiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// jubeiSpanishPhrase represents jubei spanish phrase.
	jubeiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// jubeiTraditionalChinesePhrase represents jubei traditional chinese phrase.
	jubeiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// jubeiPhrase represents jubei phrase.
	jubeiPhrase = nook.Languages{
		language.AmericanEnglish:      jubeiAmericanEnglishPhrase,
		language.CanadianFrench:       jubeiCanadianFrenchPhrase,
		language.Dutch:                jubeiDutchPhrase,
		language.French:               jubeiFrenchPhrase,
		language.German:               jubeiGermanPhrase,
		language.Italian:              jubeiItalianPhrase,
		language.Japanese:             jubeiJapanesePhrase,
		language.Korean:               jubeiKoreanPhrase,
		language.LatinAmericanSpanish: jubeiLatinAmericanSpanishPhrase,
		language.Russian:              jubeiRussianPhrase,
		language.SimplifiedChinese:    jubeiSimplifiedChinesePhrase,
		language.Spanish:              jubeiSpanishPhrase,
		language.TraditionalChinese:   jubeiTraditionalChinesePhrase}
)

var (
	// Jubei represents jubei.
	Jubei = nook.Villager{
		Character:   jubeiCharacter,
		Personality: personality.Cranky,
		Phrase:      jubeiPhrase}
)
