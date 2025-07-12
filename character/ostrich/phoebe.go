package ostrich

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
	// phoebeBirthday represents phoebe birthday.
	phoebeBirthday = nook.Birthday{
		Day:   22,
		Month: time.April}
)

var (
	// phoebeCode represents phoebe code.
	phoebeCode = nook.Code{
		Value: "ost10"}
)

var (
	// phoebeAmericanEnglishName represents phoebe american english name.
	phoebeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phoebe"}

	// phoebeCanadianFrenchName represents phoebe canadian french name.
	phoebeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Faustine"}

	// phoebeDutchName represents phoebe dutch name.
	phoebeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phoebe"}

	// phoebeFrenchName represents phoebe french name.
	phoebeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Faustine"}

	// phoebeGermanName represents phoebe german name.
	phoebeGermanName = nook.Name{
		Language: language.German,
		Value:    "Philippa"}

	// phoebeItalianName represents phoebe italian name.
	phoebeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fenicia"}

	// phoebeJapaneseName represents phoebe japanese name.
	phoebeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒノコ"}

	// phoebeKoreanName represents phoebe korean name.
	phoebeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "휘니"}

	// phoebeLatinAmericanSpanishName represents phoebe latin american spanish name.
	phoebeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Avelina"}

	// phoebeRussianName represents phoebe russian name.
	phoebeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фиби"}

	// phoebeSimplifiedChineseName represents phoebe simplified chinese name.
	phoebeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "火鸟"}

	// phoebeSpanishName represents phoebe spanish name.
	phoebeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Avelina"}

	// phoebeTraditionalChineseName represents phoebe traditional chinese name.
	phoebeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "火鳥"}
)

var (
	// phoebeName represents phoebe name.
	phoebeName = nook.Languages{
		language.AmericanEnglish:      phoebeAmericanEnglishName,
		language.CanadianFrench:       phoebeCanadianFrenchName,
		language.Dutch:                phoebeDutchName,
		language.French:               phoebeFrenchName,
		language.German:               phoebeGermanName,
		language.Italian:              phoebeItalianName,
		language.Japanese:             phoebeJapaneseName,
		language.Korean:               phoebeKoreanName,
		language.LatinAmericanSpanish: phoebeLatinAmericanSpanishName,
		language.Russian:              phoebeRussianName,
		language.SimplifiedChinese:    phoebeSimplifiedChineseName,
		language.Spanish:              phoebeSpanishName,
		language.TraditionalChinese:   phoebeTraditionalChineseName}
)

var (
	// phoebeCharacter represents phoebe character.
	phoebeCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: phoebeBirthday,
		Code:     phoebeCode,
		Key:      character.Phoebe,
		Gender:   gender.Female,
		Name:     phoebeName,
		Special:  false}
)

var (
	// phoebeAmericanEnglishPhrase represents phoebe american english phrase.
	phoebeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sparky"}

	// phoebeCanadianFrenchPhrase represents phoebe canadian french phrase.
	phoebeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chanmé"}

	// phoebeDutchPhrase represents phoebe dutch phrase.
	phoebeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vonkje"}

	// phoebeFrenchPhrase represents phoebe french phrase.
	phoebeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chanmé"}

	// phoebeGermanPhrase represents phoebe german phrase.
	phoebeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pieks"}

	// phoebeItalianPhrase represents phoebe italian phrase.
	phoebeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "behmù"}

	// phoebeJapanesePhrase represents phoebe japanese phrase.
	phoebeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アツイわ"}

	// phoebeKoreanPhrase represents phoebe korean phrase.
	phoebeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아뜨뜨"}

	// phoebeLatinAmericanSpanishPhrase represents phoebe latin american spanish phrase.
	phoebeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "equilicuá"}

	// phoebeRussianPhrase represents phoebe russian phrase.
	phoebeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "искорка"}

	// phoebeSimplifiedChinesePhrase represents phoebe simplified chinese phrase.
	phoebeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好热"}

	// phoebeSpanishPhrase represents phoebe spanish phrase.
	phoebeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "equilicuá"}

	// phoebeTraditionalChinesePhrase represents phoebe traditional chinese phrase.
	phoebeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好熱"}
)

var (
	// phoebePhrase represents phoebe phrase.
	phoebePhrase = nook.Languages{
		language.AmericanEnglish:      phoebeAmericanEnglishPhrase,
		language.CanadianFrench:       phoebeCanadianFrenchPhrase,
		language.Dutch:                phoebeDutchPhrase,
		language.French:               phoebeFrenchPhrase,
		language.German:               phoebeGermanPhrase,
		language.Italian:              phoebeItalianPhrase,
		language.Japanese:             phoebeJapanesePhrase,
		language.Korean:               phoebeKoreanPhrase,
		language.LatinAmericanSpanish: phoebeLatinAmericanSpanishPhrase,
		language.Russian:              phoebeRussianPhrase,
		language.SimplifiedChinese:    phoebeSimplifiedChinesePhrase,
		language.Spanish:              phoebeSpanishPhrase,
		language.TraditionalChinese:   phoebeTraditionalChinesePhrase}
)

var (
	// Phoebe represents phoebe.
	Phoebe = nook.Villager{
		Character:   phoebeCharacter,
		Personality: personality.BigSister,
		Phrase:      phoebePhrase}
)
