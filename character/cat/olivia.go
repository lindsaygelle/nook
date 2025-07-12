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
	// oliviaBirthday represents olivia birthday.
	oliviaBirthday = nook.Birthday{
		Day:   3,
		Month: time.February}
)

var (
	// oliviaCode represents olivia code.
	oliviaCode = nook.Code{
		Value: "cat03"}
)

var (
	// oliviaAmericanEnglishName represents olivia american english name.
	oliviaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Olivia"}

	// oliviaCanadianFrenchName represents olivia canadian french name.
	oliviaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Olivia"}

	// oliviaDutchName represents olivia dutch name.
	oliviaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olivia"}

	// oliviaFrenchName represents olivia french name.
	oliviaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Olivia"}

	// oliviaGermanName represents olivia german name.
	oliviaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bianca"}

	// oliviaItalianName represents olivia italian name.
	oliviaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Micina"}

	// oliviaJapaneseName represents olivia japanese name.
	oliviaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オリビア"}

	// oliviaKoreanName represents olivia korean name.
	oliviaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리비아"}

	// oliviaLatinAmericanSpanishName represents olivia latin american spanish name.
	oliviaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Olivia"}

	// oliviaRussianName represents olivia russian name.
	oliviaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оливия"}

	// oliviaSimplifiedChineseName represents olivia simplified chinese name.
	oliviaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "奥莉"}

	// oliviaSpanishName represents olivia spanish name.
	oliviaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Olivia"}

	// oliviaTraditionalChineseName represents olivia traditional chinese name.
	oliviaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "奧莉"}
)

var (
	// oliviaName represents olivia name.
	oliviaName = nook.Languages{
		language.AmericanEnglish:      oliviaAmericanEnglishName,
		language.CanadianFrench:       oliviaCanadianFrenchName,
		language.Dutch:                oliviaDutchName,
		language.French:               oliviaFrenchName,
		language.German:               oliviaGermanName,
		language.Italian:              oliviaItalianName,
		language.Japanese:             oliviaJapaneseName,
		language.Korean:               oliviaKoreanName,
		language.LatinAmericanSpanish: oliviaLatinAmericanSpanishName,
		language.Russian:              oliviaRussianName,
		language.SimplifiedChinese:    oliviaSimplifiedChineseName,
		language.Spanish:              oliviaSpanishName,
		language.TraditionalChinese:   oliviaTraditionalChineseName}
)

var (
	// oliviaCharacter represents olivia character.
	oliviaCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: oliviaBirthday,
		Code:     oliviaCode,
		Key:      character.Olivia,
		Gender:   gender.Female,
		Name:     oliviaName,
		Special:  false}
)

var (
	// oliviaAmericanEnglishPhrase represents olivia american english phrase.
	oliviaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "purrr"}

	// oliviaCanadianFrenchPhrase represents olivia canadian french phrase.
	oliviaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rrrrrrrrr"}

	// oliviaDutchPhrase represents olivia dutch phrase.
	oliviaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "prrrrrr"}

	// oliviaFrenchPhrase represents olivia french phrase.
	oliviaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rrrrrrrrr"}

	// oliviaGermanPhrase represents olivia german phrase.
	oliviaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnurr"}

	// oliviaItalianPhrase represents olivia italian phrase.
	oliviaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "purrr"}

	// oliviaJapanesePhrase represents olivia japanese phrase.
	oliviaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんやん"}

	// oliviaKoreanPhrase represents olivia korean phrase.
	oliviaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "머냐"}

	// oliviaLatinAmericanSpanishPhrase represents olivia latin american spanish phrase.
	oliviaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michimiau"}

	// oliviaRussianPhrase represents olivia russian phrase.
	oliviaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мур-мяу"}

	// oliviaSimplifiedChinesePhrase represents olivia simplified chinese phrase.
	oliviaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "什么"}

	// oliviaSpanishPhrase represents olivia spanish phrase.
	oliviaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "raspas"}

	// oliviaTraditionalChinesePhrase represents olivia traditional chinese phrase.
	oliviaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "什麼"}
)

var (
	// oliviaPhrase represents olivia phrase.
	oliviaPhrase = nook.Languages{
		language.AmericanEnglish:      oliviaAmericanEnglishPhrase,
		language.CanadianFrench:       oliviaCanadianFrenchPhrase,
		language.Dutch:                oliviaDutchPhrase,
		language.French:               oliviaFrenchPhrase,
		language.German:               oliviaGermanPhrase,
		language.Italian:              oliviaItalianPhrase,
		language.Japanese:             oliviaJapanesePhrase,
		language.Korean:               oliviaKoreanPhrase,
		language.LatinAmericanSpanish: oliviaLatinAmericanSpanishPhrase,
		language.Russian:              oliviaRussianPhrase,
		language.SimplifiedChinese:    oliviaSimplifiedChinesePhrase,
		language.Spanish:              oliviaSpanishPhrase,
		language.TraditionalChinese:   oliviaTraditionalChinesePhrase}
)

var (
	// Olivia represents olivia.
	Olivia = nook.Villager{
		Character:   oliviaCharacter,
		Personality: personality.Snooty,
		Phrase:      oliviaPhrase}
)
