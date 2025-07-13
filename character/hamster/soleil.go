package hamster

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
	// soleilBirthday represents soleil birthday.
	soleilBirthday = nook.Birthday{
		Day:   9,
		Month: time.August}
)

var (
	// soleilCode represents soleil code.
	soleilCode = nook.Code{
		Value: "ham04"}
)

var (
	// soleilAmericanEnglishName represents soleil american english name.
	soleilAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Soleil"}

	// soleilCanadianFrenchName represents soleil canadian french name.
	soleilCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Stella"}

	// soleilDutchName represents soleil dutch name.
	soleilDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Soleil"}

	// soleilFrenchName represents soleil french name.
	soleilFrenchName = nook.Name{
		Language: language.French,
		Value:    "Stella"}

	// soleilGermanName represents soleil german name.
	soleilGermanName = nook.Name{
		Language: language.German,
		Value:    "Samira"}

	// soleilItalianName represents soleil italian name.
	soleilItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cettina"}

	// soleilJapaneseName represents soleil japanese name.
	soleilJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャンティ"}

	// soleilKoreanName represents soleil korean name.
	soleilKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샨티"}

	// soleilLatinAmericanSpanishName represents soleil latin american spanish name.
	soleilLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Soraya"}

	// soleilRussianName represents soleil russian name.
	soleilRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Солей"}

	// soleilSimplifiedChineseName represents soleil simplified chinese name.
	soleilSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安安"}

	// soleilSpanishName represents soleil spanish name.
	soleilSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Soraya"}

	// soleilTraditionalChineseName represents soleil traditional chinese name.
	soleilTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安安"}
)

var (
	// soleilName represents soleil name.
	soleilName = nook.Languages{
		language.AmericanEnglish:      soleilAmericanEnglishName,
		language.CanadianFrench:       soleilCanadianFrenchName,
		language.Dutch:                soleilDutchName,
		language.French:               soleilFrenchName,
		language.German:               soleilGermanName,
		language.Italian:              soleilItalianName,
		language.Japanese:             soleilJapaneseName,
		language.Korean:               soleilKoreanName,
		language.LatinAmericanSpanish: soleilLatinAmericanSpanishName,
		language.Russian:              soleilRussianName,
		language.SimplifiedChinese:    soleilSimplifiedChineseName,
		language.Spanish:              soleilSpanishName,
		language.TraditionalChinese:   soleilTraditionalChineseName}
)

var (
	// soleilCharacter represents soleil character.
	soleilCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: soleilBirthday,
		Code:     soleilCode,
		Key:      character.Soleil,
		Gender:   gender.Female,
		Name:     soleilName,
		Special:  false}
)

var (
	// soleilAmericanEnglishPhrase represents soleil american english phrase.
	soleilAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tarnation"}

	// soleilCanadianFrenchPhrase represents soleil canadian french phrase.
	soleilCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ciel"}

	// soleilDutchPhrase represents soleil dutch phrase.
	soleilDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "korenwolf"}

	// soleilFrenchPhrase represents soleil french phrase.
	soleilFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gnognotte"}

	// soleilGermanPhrase represents soleil german phrase.
	soleilGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hümpf"}

	// soleilItalianPhrase represents soleil italian phrase.
	soleilItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vriiin"}

	// soleilJapanesePhrase represents soleil japanese phrase.
	soleilJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だってば"}

	// soleilKoreanPhrase represents soleil korean phrase.
	soleilKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파샤샤"}

	// soleilLatinAmericanSpanishPhrase represents soleil latin american spanish phrase.
	soleilLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiqui-ñú"}

	// soleilRussianPhrase represents soleil russian phrase.
	soleilRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "о-хо-хо"}

	// soleilSimplifiedChinesePhrase represents soleil simplified chinese phrase.
	soleilSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "相信我"}

	// soleilSpanishPhrase represents soleil spanish phrase.
	soleilSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "moflete"}

	// soleilTraditionalChinesePhrase represents soleil traditional chinese phrase.
	soleilTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "相信我"}
)

var (
	// soleilPhrase represents soleil phrase.
	soleilPhrase = nook.Languages{
		language.AmericanEnglish:      soleilAmericanEnglishPhrase,
		language.CanadianFrench:       soleilCanadianFrenchPhrase,
		language.Dutch:                soleilDutchPhrase,
		language.French:               soleilFrenchPhrase,
		language.German:               soleilGermanPhrase,
		language.Italian:              soleilItalianPhrase,
		language.Japanese:             soleilJapanesePhrase,
		language.Korean:               soleilKoreanPhrase,
		language.LatinAmericanSpanish: soleilLatinAmericanSpanishPhrase,
		language.Russian:              soleilRussianPhrase,
		language.SimplifiedChinese:    soleilSimplifiedChinesePhrase,
		language.Spanish:              soleilSpanishPhrase,
		language.TraditionalChinese:   soleilTraditionalChinesePhrase}
)

var (
	// Soleil represents soleil.
	Soleil = nook.Villager{
		Character:   soleilCharacter,
		Personality: personality.Snooty,
		Phrase:      soleilPhrase}
)
