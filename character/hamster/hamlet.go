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
	// hamletBirthday represents hamlet birthday.
	hamletBirthday = nook.Birthday{
		Day:   30,
		Month: time.May}
)

var (
	// hamletCode represents hamlet code.
	hamletCode = nook.Code{
		Value: "ham00"}
)

var (
	// hamletAmericanEnglishName represents hamlet american english name.
	hamletAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hamlet"}

	// hamletCanadianFrenchName represents hamlet canadian french name.
	hamletCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jojo"}

	// hamletDutchName represents hamlet dutch name.
	hamletDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hamlet"}

	// hamletFrenchName represents hamlet french name.
	hamletFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jojo"}

	// hamletGermanName represents hamlet german name.
	hamletGermanName = nook.Name{
		Language: language.German,
		Value:    "Hamid"}

	// hamletItalianName represents hamlet italian name.
	hamletItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amleto"}

	// hamletJapaneseName represents hamlet japanese name.
	hamletJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムスケ"}

	// hamletKoreanName represents hamlet korean name.
	hamletKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄스틴"}

	// hamletLatinAmericanSpanishName represents hamlet latin american spanish name.
	hamletLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bombo"}

	// hamletRussianName represents hamlet russian name.
	hamletRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гамлет"}

	// hamletSimplifiedChineseName represents hamlet simplified chinese name.
	hamletSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈姆"}

	// hamletSpanishName represents hamlet spanish name.
	hamletSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bombo"}

	// hamletTraditionalChineseName represents hamlet traditional chinese name.
	hamletTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈姆"}
)

var (
	// hamletName represents hamlet name.
	hamletName = nook.Languages{
		language.AmericanEnglish:      hamletAmericanEnglishName,
		language.CanadianFrench:       hamletCanadianFrenchName,
		language.Dutch:                hamletDutchName,
		language.French:               hamletFrenchName,
		language.German:               hamletGermanName,
		language.Italian:              hamletItalianName,
		language.Japanese:             hamletJapaneseName,
		language.Korean:               hamletKoreanName,
		language.LatinAmericanSpanish: hamletLatinAmericanSpanishName,
		language.Russian:              hamletRussianName,
		language.SimplifiedChinese:    hamletSimplifiedChineseName,
		language.Spanish:              hamletSpanishName,
		language.TraditionalChinese:   hamletTraditionalChineseName}
)

var (
	// hamletCharacter represents hamlet character.
	hamletCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: hamletBirthday,
		Code:     hamletCode,
		Key:      character.Hamlet,
		Gender:   gender.Male,
		Name:     hamletName,
		Special:  false}
)

var (
	// hamletAmericanEnglishPhrase represents hamlet american english phrase.
	hamletAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hammie"}

	// hamletCanadianFrenchPhrase represents hamlet canadian french phrase.
	hamletCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chicots"}

	// hamletDutchPhrase represents hamlet dutch phrase.
	hamletDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hammie"}

	// hamletFrenchPhrase represents hamlet french phrase.
	hamletFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chicots"}

	// hamletGermanPhrase represents hamlet german phrase.
	hamletGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "twirps"}

	// hamletItalianPhrase represents hamlet italian phrase.
	hamletItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "puff pant"}

	// hamletJapanesePhrase represents hamlet japanese phrase.
	hamletJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハム"}

	// hamletKoreanPhrase represents hamlet korean phrase.
	hamletKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "햄햄"}

	// hamletLatinAmericanSpanishPhrase represents hamlet latin american spanish phrase.
	hamletLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chucufá"}

	// hamletRussianPhrase represents hamlet russian phrase.
	hamletRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хомячок"}

	// hamletSimplifiedChinesePhrase represents hamlet simplified chinese phrase.
	hamletSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "火腿"}

	// hamletSpanishPhrase represents hamlet spanish phrase.
	hamletSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chucufá"}

	// hamletTraditionalChinesePhrase represents hamlet traditional chinese phrase.
	hamletTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "火腿"}
)

var (
	// hamletPhrase represents hamlet phrase.
	hamletPhrase = nook.Languages{
		language.AmericanEnglish:      hamletAmericanEnglishPhrase,
		language.CanadianFrench:       hamletCanadianFrenchPhrase,
		language.Dutch:                hamletDutchPhrase,
		language.French:               hamletFrenchPhrase,
		language.German:               hamletGermanPhrase,
		language.Italian:              hamletItalianPhrase,
		language.Japanese:             hamletJapanesePhrase,
		language.Korean:               hamletKoreanPhrase,
		language.LatinAmericanSpanish: hamletLatinAmericanSpanishPhrase,
		language.Russian:              hamletRussianPhrase,
		language.SimplifiedChinese:    hamletSimplifiedChinesePhrase,
		language.Spanish:              hamletSpanishPhrase,
		language.TraditionalChinese:   hamletTraditionalChinesePhrase}
)

var (
	// Hamlet represents hamlet.
	Hamlet = nook.Villager{
		Character:   hamletCharacter,
		Personality: personality.Jock,
		Phrase:      hamletPhrase}
)
