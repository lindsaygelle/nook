package duck

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
	// derwinBirthday represents derwin birthday.
	derwinBirthday = nook.Birthday{
		Day:   25,
		Month: time.May}
)

var (
	// derwinCode represents derwin code.
	derwinCode = nook.Code{
		Value: "duk08"}
)

var (
	// derwinAmericanEnglishName represents derwin american english name.
	derwinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Derwin"}

	// derwinCanadianFrenchName represents derwin canadian french name.
	derwinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Prof"}

	// derwinDutchName represents derwin dutch name.
	derwinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Derwin"}

	// derwinFrenchName represents derwin french name.
	derwinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Prof"}

	// derwinGermanName represents derwin german name.
	derwinGermanName = nook.Name{
		Language: language.German,
		Value:    "Erwin"}

	// derwinItalianName represents derwin italian name.
	derwinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mike"}

	// derwinJapaneseName represents derwin japanese name.
	derwinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボン"}

	// derwinKoreanName represents derwin korean name.
	derwinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "봉"}

	// derwinLatinAmericanSpanishName represents derwin latin american spanish name.
	derwinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Torcuato"}

	// derwinRussianName represents derwin russian name.
	derwinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дервин"}

	// derwinSimplifiedChineseName represents derwin simplified chinese name.
	derwinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿坊"}

	// derwinSpanishName represents derwin spanish name.
	derwinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Torcuato"}

	// derwinTraditionalChineseName represents derwin traditional chinese name.
	derwinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿坊"}
)

var (
	// derwinName represents derwin name.
	derwinName = nook.Languages{
		language.AmericanEnglish:      derwinAmericanEnglishName,
		language.CanadianFrench:       derwinCanadianFrenchName,
		language.Dutch:                derwinDutchName,
		language.French:               derwinFrenchName,
		language.German:               derwinGermanName,
		language.Italian:              derwinItalianName,
		language.Japanese:             derwinJapaneseName,
		language.Korean:               derwinKoreanName,
		language.LatinAmericanSpanish: derwinLatinAmericanSpanishName,
		language.Russian:              derwinRussianName,
		language.SimplifiedChinese:    derwinSimplifiedChineseName,
		language.Spanish:              derwinSpanishName,
		language.TraditionalChinese:   derwinTraditionalChineseName}
)

var (
	// derwinCharacter represents derwin character.
	derwinCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: derwinBirthday,
		Code:     derwinCode,
		Key:      character.Derwin,
		Gender:   gender.Male,
		Name:     derwinName,
		Special:  false}
)

var (
	// derwinAmericanEnglishPhrase represents derwin american english phrase.
	derwinAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "derrrr"}

	// derwinCanadianFrenchPhrase represents derwin canadian french phrase.
	derwinCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "derrrr"}

	// derwinDutchPhrase represents derwin dutch phrase.
	derwinDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ergo"}

	// derwinFrenchPhrase represents derwin french phrase.
	derwinFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "derrrr"}

	// derwinGermanPhrase represents derwin german phrase.
	derwinGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krook"}

	// derwinItalianPhrase represents derwin italian phrase.
	derwinItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaccolo"}

	// derwinJapanesePhrase represents derwin japanese phrase.
	derwinJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ね！ママ"}

	// derwinKoreanPhrase represents derwin korean phrase.
	derwinKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엄마"}

	// derwinLatinAmericanSpanishPhrase represents derwin latin american spanish phrase.
	derwinLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuato"}

	// derwinRussianPhrase represents derwin russian phrase.
	derwinRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кррря"}

	// derwinSimplifiedChinesePhrase represents derwin simplified chinese phrase.
	derwinSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂！妈妈"}

	// derwinSpanishPhrase represents derwin spanish phrase.
	derwinSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuato"}

	// derwinTraditionalChinesePhrase represents derwin traditional chinese phrase.
	derwinTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂！媽媽"}
)

var (
	// derwinPhrase represents derwin phrase.
	derwinPhrase = nook.Languages{
		language.AmericanEnglish:      derwinAmericanEnglishPhrase,
		language.CanadianFrench:       derwinCanadianFrenchPhrase,
		language.Dutch:                derwinDutchPhrase,
		language.French:               derwinFrenchPhrase,
		language.German:               derwinGermanPhrase,
		language.Italian:              derwinItalianPhrase,
		language.Japanese:             derwinJapanesePhrase,
		language.Korean:               derwinKoreanPhrase,
		language.LatinAmericanSpanish: derwinLatinAmericanSpanishPhrase,
		language.Russian:              derwinRussianPhrase,
		language.SimplifiedChinese:    derwinSimplifiedChinesePhrase,
		language.Spanish:              derwinSpanishPhrase,
		language.TraditionalChinese:   derwinTraditionalChinesePhrase}
)

var (
	// Derwin represents derwin.
	Derwin = nook.Villager{
		Character:   derwinCharacter,
		Personality: personality.Lazy,
		Phrase:      derwinPhrase}
)
