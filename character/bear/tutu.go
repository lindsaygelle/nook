package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// tutuBirthday represents Tutu's birthday.
var (
	// tutuBirthday represents tutu birthday.
	tutuBirthday = nook.Birthday{
		Day:   15,
		Month: time.September}
)

// tutuCode represents Tutu's unique code.
var (
	// tutuCode represents tutu code.
	tutuCode = nook.Code{
		Value: "bea07"}
)

// Different names for Tutu in various languages.
var (
	// tutuAmericanEnglishName represents tutu american english name.
	tutuAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tutu"}

	// tutuCanadianFrenchName represents tutu canadian french name.
	tutuCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tutu"}

	// tutuDutchName represents tutu dutch name.
	tutuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tutu"}

	// tutuFrenchName represents tutu french name.
	tutuFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tutu"}

	// tutuGermanName represents tutu german name.
	tutuGermanName = nook.Name{
		Language: language.German,
		Value:    "Mandy"}

	// tutuItalianName represents tutu italian name.
	tutuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lola"}

	// tutuJapaneseName represents tutu japanese name.
	tutuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "れんにゅう"}

	// tutuKoreanName represents tutu korean name.
	tutuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "연유"}

	// tutuLatinAmericanSpanishName represents tutu latin american spanish name.
	tutuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tutú"}

	// tutuRussianName represents tutu russian name.
	tutuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Туту"}

	// tutuSimplifiedChineseName represents tutu simplified chinese name.
	tutuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "恋恋"}

	// tutuSpanishName represents tutu spanish name.
	tutuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tutú"}

	// tutuTraditionalChineseName represents tutu traditional chinese name.
	tutuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戀戀"}
)

// tutuName represents Tutu's name in different languages.
var (
	// tutuName represents tutu name.
	tutuName = nook.Languages{
		language.AmericanEnglish:      tutuAmericanEnglishName,
		language.CanadianFrench:       tutuCanadianFrenchName,
		language.Dutch:                tutuDutchName,
		language.French:               tutuFrenchName,
		language.German:               tutuGermanName,
		language.Italian:              tutuItalianName,
		language.Japanese:             tutuJapaneseName,
		language.Korean:               tutuKoreanName,
		language.LatinAmericanSpanish: tutuLatinAmericanSpanishName,
		language.Russian:              tutuRussianName,
		language.SimplifiedChinese:    tutuSimplifiedChineseName,
		language.Spanish:              tutuSpanishName,
		language.TraditionalChinese:   tutuTraditionalChineseName}
)

// tutuCharacter represents Tutu's character information.
var (
	// tutuCharacter represents tutu character.
	tutuCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: tutuBirthday,
		Code:     tutuCode,
		Key:      character.Tutu,
		Gender:   gender.Female,
		Name:     tutuName,
		Special:  false}
)

// Different phrases spoken by Tutu in various languages.
var (
	// tutuAmericanEnglishPhrase represents tutu american english phrase.
	tutuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "twinkles"}

	// tutuCanadianFrenchPhrase represents tutu canadian french phrase.
	tutuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pic pic"}

	// tutuDutchPhrase represents tutu dutch phrase.
	tutuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "honingbij"}

	// tutuFrenchPhrase represents tutu french phrase.
	tutuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pic pic"}

	// tutuGermanPhrase represents tutu german phrase.
	tutuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hooonig"}

	// tutuItalianPhrase represents tutu italian phrase.
	tutuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnifa"}

	// tutuJapanesePhrase represents tutu japanese phrase.
	tutuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ファイト"}

	// tutuKoreanPhrase represents tutu korean phrase.
	tutuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "사르르"}

	// tutuLatinAmericanSpanishPhrase represents tutu latin american spanish phrase.
	tutuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fisflash"}

	// tutuRussianPhrase represents tutu russian phrase.
	tutuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ме-е-ед"}

	// tutuSimplifiedChinesePhrase represents tutu simplified chinese phrase.
	tutuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "加油"}

	// tutuSpanishPhrase represents tutu spanish phrase.
	tutuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fisflash"}

	// tutuTraditionalChinesePhrase represents tutu traditional chinese phrase.
	tutuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "加油"}
)

// tutuPhrase represents Tutu's phrases in different languages.
var (
	// tutuPhrase represents tutu phrase.
	tutuPhrase = nook.Languages{
		language.AmericanEnglish:      tutuAmericanEnglishPhrase,
		language.CanadianFrench:       tutuCanadianFrenchPhrase,
		language.Dutch:                tutuDutchPhrase,
		language.French:               tutuFrenchPhrase,
		language.German:               tutuGermanPhrase,
		language.Italian:              tutuItalianPhrase,
		language.Japanese:             tutuJapanesePhrase,
		language.Korean:               tutuKoreanPhrase,
		language.LatinAmericanSpanish: tutuLatinAmericanSpanishPhrase,
		language.Russian:              tutuRussianPhrase,
		language.SimplifiedChinese:    tutuSimplifiedChinesePhrase,
		language.Spanish:              tutuSpanishPhrase,
		language.TraditionalChinese:   tutuTraditionalChinesePhrase}
)

// Tutu represents the character Tutu with her complete information.
var (
	// Tutu represents tutu.
	Tutu = nook.Villager{
		Character:   tutuCharacter,
		Personality: personality.Peppy,
		Phrase:      tutuPhrase}
)
