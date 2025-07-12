package monkey

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
	// flipBirthday represents flip birthday.
	flipBirthday = nook.Birthday{
		Day:   21,
		Month: time.November}
)

var (
	// flipCode represents flip code.
	flipCode = nook.Code{
		Value: "mnk06"}
)

var (
	// flipAmericanEnglishName represents flip american english name.
	flipAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flip"}

	// flipCanadianFrenchName represents flip canadian french name.
	flipCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rudy"}

	// flipDutchName represents flip dutch name.
	flipDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flip"}

	// flipFrenchName represents flip french name.
	flipFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rudy"}

	// flipGermanName represents flip german name.
	flipGermanName = nook.Name{
		Language: language.German,
		Value:    "Pippo"}

	// flipItalianName represents flip italian name.
	flipItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Benny"}

	// flipJapaneseName represents flip japanese name.
	flipJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さすけ"}

	// flipKoreanName represents flip korean name.
	flipKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "원승"}

	// flipLatinAmericanSpanishName represents flip latin american spanish name.
	flipLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Monet"}

	// flipRussianName represents flip russian name.
	flipRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флип"}

	// flipSimplifiedChineseName represents flip simplified chinese name.
	flipSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "天佑"}

	// flipSpanishName represents flip spanish name.
	flipSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Monet"}

	// flipTraditionalChineseName represents flip traditional chinese name.
	flipTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "天佑"}
)

var (
	// flipName represents flip name.
	flipName = nook.Languages{
		language.AmericanEnglish:      flipAmericanEnglishName,
		language.CanadianFrench:       flipCanadianFrenchName,
		language.Dutch:                flipDutchName,
		language.French:               flipFrenchName,
		language.German:               flipGermanName,
		language.Italian:              flipItalianName,
		language.Japanese:             flipJapaneseName,
		language.Korean:               flipKoreanName,
		language.LatinAmericanSpanish: flipLatinAmericanSpanishName,
		language.Russian:              flipRussianName,
		language.SimplifiedChinese:    flipSimplifiedChineseName,
		language.Spanish:              flipSpanishName,
		language.TraditionalChinese:   flipTraditionalChineseName}
)

var (
	// flipCharacter represents flip character.
	flipCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: flipBirthday,
		Code:     flipCode,
		Key:      character.Flip,
		Gender:   gender.Male,
		Name:     flipName,
		Special:  false}
)

var (
	// flipAmericanEnglishPhrase represents flip american english phrase.
	flipAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rerack"}

	// flipCanadianFrenchPhrase represents flip canadian french phrase.
	flipCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh hisse"}

	// flipDutchPhrase represents flip dutch phrase.
	flipDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oe-⁠a-⁠a"}

	// flipFrenchPhrase represents flip french phrase.
	flipFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh hisse"}

	// flipGermanPhrase represents flip german phrase.
	flipGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kikiki"}

	// flipItalianPhrase represents flip italian phrase.
	flipItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clap clap"}

	// flipJapanesePhrase represents flip japanese phrase.
	flipJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どっこい"}

	// flipKoreanPhrase represents flip korean phrase.
	flipKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "빠샤"}

	// flipLatinAmericanSpanishPhrase represents flip latin american spanish phrase.
	flipLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "afuá"}

	// flipRussianPhrase represents flip russian phrase.
	flipRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "у-у-у"}

	// flipSimplifiedChinesePhrase represents flip simplified chinese phrase.
	flipSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "差不多"}

	// flipSpanishPhrase represents flip spanish phrase.
	flipSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "afuá"}

	// flipTraditionalChinesePhrase represents flip traditional chinese phrase.
	flipTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "差不多"}
)

var (
	// flipPhrase represents flip phrase.
	flipPhrase = nook.Languages{
		language.AmericanEnglish:      flipAmericanEnglishPhrase,
		language.CanadianFrench:       flipCanadianFrenchPhrase,
		language.Dutch:                flipDutchPhrase,
		language.French:               flipFrenchPhrase,
		language.German:               flipGermanPhrase,
		language.Italian:              flipItalianPhrase,
		language.Japanese:             flipJapanesePhrase,
		language.Korean:               flipKoreanPhrase,
		language.LatinAmericanSpanish: flipLatinAmericanSpanishPhrase,
		language.Russian:              flipRussianPhrase,
		language.SimplifiedChinese:    flipSimplifiedChinesePhrase,
		language.Spanish:              flipSpanishPhrase,
		language.TraditionalChinese:   flipTraditionalChinesePhrase}
)

var (
	// Flip represents flip.
	Flip = nook.Villager{
		Character:   flipCharacter,
		Personality: personality.Jock,
		Phrase:      flipPhrase}
)
