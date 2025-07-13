package cat

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
	// rosieBirthday represents rosie birthday.
	rosieBirthday = nook.Birthday{
		Day:   27,
		Month: time.February}
)

var (
	// rosieCode represents rosie code.
	rosieCode = nook.Code{
		Value: "cat02"}
)

var (
	// rosieAmericanEnglishName represents rosie american english name.
	rosieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rosie"}

	// rosieCanadianFrenchName represents rosie canadian french name.
	rosieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosie"}

	// rosieDutchName represents rosie dutch name.
	rosieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rosie"}

	// rosieFrenchName represents rosie french name.
	rosieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosie"}

	// rosieGermanName represents rosie german name.
	rosieGermanName = nook.Name{
		Language: language.German,
		Value:    "Sophie"}

	// rosieItalianName represents rosie italian name.
	rosieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grinfia"}

	// rosieJapaneseName represents rosie japanese name.
	rosieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブーケ"}

	// rosieKoreanName represents rosie korean name.
	rosieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "부케"}

	// rosieLatinAmericanSpanishName represents rosie latin american spanish name.
	rosieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Minina"}

	// rosieRussianName represents rosie russian name.
	rosieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рози"}

	// rosieSimplifiedChineseName represents rosie simplified chinese name.
	rosieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彭花"}

	// rosieSpanishName represents rosie spanish name.
	rosieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Minina"}

	// rosieTraditionalChineseName represents rosie traditional chinese name.
	rosieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彭花"}
)

var (
	// rosieName represents rosie name.
	rosieName = nook.Languages{
		language.AmericanEnglish:      rosieAmericanEnglishName,
		language.CanadianFrench:       rosieCanadianFrenchName,
		language.Dutch:                rosieDutchName,
		language.French:               rosieFrenchName,
		language.German:               rosieGermanName,
		language.Italian:              rosieItalianName,
		language.Japanese:             rosieJapaneseName,
		language.Korean:               rosieKoreanName,
		language.LatinAmericanSpanish: rosieLatinAmericanSpanishName,
		language.Russian:              rosieRussianName,
		language.SimplifiedChinese:    rosieSimplifiedChineseName,
		language.Spanish:              rosieSpanishName,
		language.TraditionalChinese:   rosieTraditionalChineseName}
)

var (
	// rosieCharacter represents rosie character.
	rosieCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: rosieBirthday,
		Code:     rosieCode,
		Key:      character.Rosie,
		Gender:   gender.Female,
		Name:     rosieName,
		Special:  false}
)

var (
	// rosieAmericanEnglishPhrase represents rosie american english phrase.
	rosieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "silly"}

	// rosieCanadianFrenchPhrase represents rosie canadian french phrase.
	rosieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flûte"}

	// rosieDutchPhrase represents rosie dutch phrase.
	rosieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gekkie"}

	// rosieFrenchPhrase represents rosie french phrase.
	rosieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flûte"}

	// rosieGermanPhrase represents rosie german phrase.
	rosieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flöt"}

	// rosieItalianPhrase represents rosie italian phrase.
	rosieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tontolon"}

	// rosieJapanesePhrase represents rosie japanese phrase.
	rosieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チェキ"}

	// rosieKoreanPhrase represents rosie korean phrase.
	rosieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "헤이"}

	// rosieLatinAmericanSpanishPhrase represents rosie latin american spanish phrase.
	rosieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miaaau"}

	// rosieRussianPhrase represents rosie russian phrase.
	rosieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "рыбка"}

	// rosieSimplifiedChinesePhrase represents rosie simplified chinese phrase.
	rosieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "看看"}

	// rosieSpanishPhrase represents rosie spanish phrase.
	rosieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miaaau"}

	// rosieTraditionalChinesePhrase represents rosie traditional chinese phrase.
	rosieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "看看"}
)

var (
	// rosiePhrase represents rosie phrase.
	rosiePhrase = nook.Languages{
		language.AmericanEnglish:      rosieAmericanEnglishPhrase,
		language.CanadianFrench:       rosieCanadianFrenchPhrase,
		language.Dutch:                rosieDutchPhrase,
		language.French:               rosieFrenchPhrase,
		language.German:               rosieGermanPhrase,
		language.Italian:              rosieItalianPhrase,
		language.Japanese:             rosieJapanesePhrase,
		language.Korean:               rosieKoreanPhrase,
		language.LatinAmericanSpanish: rosieLatinAmericanSpanishPhrase,
		language.Russian:              rosieRussianPhrase,
		language.SimplifiedChinese:    rosieSimplifiedChinesePhrase,
		language.Spanish:              rosieSpanishPhrase,
		language.TraditionalChinese:   rosieTraditionalChinesePhrase}
)

var (
	// Rosie represents rosie.
	Rosie = nook.Villager{
		Character:   rosieCharacter,
		Personality: personality.Peppy,
		Phrase:      rosiePhrase}
)
