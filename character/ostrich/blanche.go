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
	// blancheBirthday represents blanche birthday.
	blancheBirthday = nook.Birthday{
		Day:   21,
		Month: time.December}
)

var (
	// blancheCode represents blanche code.
	blancheCode = nook.Code{
		Value: "ost08"}
)

var (
	// blancheAmericanEnglishName represents blanche american english name.
	blancheAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blanche"}

	// blancheCanadianFrenchName represents blanche canadian french name.
	blancheCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sophie"}

	// blancheDutchName represents blanche dutch name.
	blancheDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blanche"}

	// blancheFrenchName represents blanche french name.
	blancheFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sophie"}

	// blancheGermanName represents blanche german name.
	blancheGermanName = nook.Name{
		Language: language.German,
		Value:    "Christa"}

	// blancheItalianName represents blanche italian name.
	blancheItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Emilia"}

	// blancheJapaneseName represents blanche japanese name.
	blancheJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しのぶ"}

	// blancheKoreanName represents blanche korean name.
	blancheKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "신옥"}

	// blancheLatinAmericanSpanishName represents blanche latin american spanish name.
	blancheLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rocío"}

	// blancheRussianName represents blanche russian name.
	blancheRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бланш"}

	// blancheSimplifiedChineseName represents blanche simplified chinese name.
	blancheSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小偲"}

	// blancheSpanishName represents blanche spanish name.
	blancheSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rocío"}

	// blancheTraditionalChineseName represents blanche traditional chinese name.
	blancheTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小偲"}
)

var (
	// blancheName represents blanche name.
	blancheName = nook.Languages{
		language.AmericanEnglish:      blancheAmericanEnglishName,
		language.CanadianFrench:       blancheCanadianFrenchName,
		language.Dutch:                blancheDutchName,
		language.French:               blancheFrenchName,
		language.German:               blancheGermanName,
		language.Italian:              blancheItalianName,
		language.Japanese:             blancheJapaneseName,
		language.Korean:               blancheKoreanName,
		language.LatinAmericanSpanish: blancheLatinAmericanSpanishName,
		language.Russian:              blancheRussianName,
		language.SimplifiedChinese:    blancheSimplifiedChineseName,
		language.Spanish:              blancheSpanishName,
		language.TraditionalChinese:   blancheTraditionalChineseName}
)

var (
	// blancheCharacter represents blanche character.
	blancheCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: blancheBirthday,
		Code:     blancheCode,
		Key:      character.Blanche,
		Gender:   gender.Female,
		Name:     blancheName,
		Special:  false}
)

var (
	// blancheAmericanEnglishPhrase represents blanche american english phrase.
	blancheAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quite so"}

	// blancheCanadianFrenchPhrase represents blanche canadian french phrase.
	blancheCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "viiiite"}

	// blancheDutchPhrase represents blanche dutch phrase.
	blancheDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "juist ja"}

	// blancheFrenchPhrase represents blanche french phrase.
	blancheFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "viiiite"}

	// blancheGermanPhrase represents blanche german phrase.
	blancheGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flaum"}

	// blancheItalianPhrase represents blanche italian phrase.
	blancheItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bien sur"}

	// blancheJapanesePhrase represents blanche japanese phrase.
	blancheJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですのね"}

	// blancheKoreanPhrase represents blanche korean phrase.
	blancheKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬구나"}

	// blancheLatinAmericanSpanishPhrase represents blanche latin american spanish phrase.
	blancheLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "igh-igh"}

	// blancheRussianPhrase represents blanche russian phrase.
	blancheRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вот так"}

	// blancheSimplifiedChinesePhrase represents blanche simplified chinese phrase.
	blancheSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是嘛"}

	// blancheSpanishPhrase represents blanche spanish phrase.
	blancheSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "goticas"}

	// blancheTraditionalChinesePhrase represents blanche traditional chinese phrase.
	blancheTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是嘛"}
)

var (
	// blanchePhrase represents blanche phrase.
	blanchePhrase = nook.Languages{
		language.AmericanEnglish:      blancheAmericanEnglishPhrase,
		language.CanadianFrench:       blancheCanadianFrenchPhrase,
		language.Dutch:                blancheDutchPhrase,
		language.French:               blancheFrenchPhrase,
		language.German:               blancheGermanPhrase,
		language.Italian:              blancheItalianPhrase,
		language.Japanese:             blancheJapanesePhrase,
		language.Korean:               blancheKoreanPhrase,
		language.LatinAmericanSpanish: blancheLatinAmericanSpanishPhrase,
		language.Russian:              blancheRussianPhrase,
		language.SimplifiedChinese:    blancheSimplifiedChinesePhrase,
		language.Spanish:              blancheSpanishPhrase,
		language.TraditionalChinese:   blancheTraditionalChinesePhrase}
)

var (
	// Blanche represents blanche.
	Blanche = nook.Villager{
		Character:   blancheCharacter,
		Personality: personality.Snooty,
		Phrase:      blanchePhrase}
)
