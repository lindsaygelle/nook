package bird

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
	// joeBirthday represents joe birthday.
	joeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// joeCode represents joe code.
	joeCode = nook.Code{
		Value: ""}
)

var (
	// joeAmericanEnglishName represents joe american english name.
	joeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Joe"}

	// joeCanadianFrenchName represents joe canadian french name.
	joeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// joeDutchName represents joe dutch name.
	joeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// joeFrenchName represents joe french name.
	joeFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// joeGermanName represents joe german name.
	joeGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// joeItalianName represents joe italian name.
	joeItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// joeJapaneseName represents joe japanese name.
	joeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョー"}

	// joeKoreanName represents joe korean name.
	joeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// joeLatinAmericanSpanishName represents joe latin american spanish name.
	joeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// joeRussianName represents joe russian name.
	joeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// joeSimplifiedChineseName represents joe simplified chinese name.
	joeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// joeSpanishName represents joe spanish name.
	joeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// joeTraditionalChineseName represents joe traditional chinese name.
	joeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// joeName represents joe name.
	joeName = nook.Languages{
		language.AmericanEnglish:      joeAmericanEnglishName,
		language.CanadianFrench:       joeCanadianFrenchName,
		language.Dutch:                joeDutchName,
		language.French:               joeFrenchName,
		language.German:               joeGermanName,
		language.Italian:              joeItalianName,
		language.Japanese:             joeJapaneseName,
		language.Korean:               joeKoreanName,
		language.LatinAmericanSpanish: joeLatinAmericanSpanishName,
		language.Russian:              joeRussianName,
		language.SimplifiedChinese:    joeSimplifiedChineseName,
		language.Spanish:              joeSpanishName,
		language.TraditionalChinese:   joeTraditionalChineseName}
)

var (
	// joeCharacter represents joe character.
	joeCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: joeBirthday,
		Code:     joeCode,
		Key:      character.Joe,
		Gender:   gender.Male,
		Name:     joeName,
		Special:  false}
)

var (
	// joeAmericanEnglishPhrase represents joe american english phrase.
	joeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "やれやれ"}

	// joeCanadianFrenchPhrase represents joe canadian french phrase.
	joeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// joeDutchPhrase represents joe dutch phrase.
	joeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// joeFrenchPhrase represents joe french phrase.
	joeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// joeGermanPhrase represents joe german phrase.
	joeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// joeItalianPhrase represents joe italian phrase.
	joeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// joeJapanesePhrase represents joe japanese phrase.
	joeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やれやれ"}

	// joeKoreanPhrase represents joe korean phrase.
	joeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// joeLatinAmericanSpanishPhrase represents joe latin american spanish phrase.
	joeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// joeRussianPhrase represents joe russian phrase.
	joeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// joeSimplifiedChinesePhrase represents joe simplified chinese phrase.
	joeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// joeSpanishPhrase represents joe spanish phrase.
	joeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// joeTraditionalChinesePhrase represents joe traditional chinese phrase.
	joeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// joePhrase represents joe phrase.
	joePhrase = nook.Languages{
		language.AmericanEnglish:      joeAmericanEnglishPhrase,
		language.CanadianFrench:       joeCanadianFrenchPhrase,
		language.Dutch:                joeDutchPhrase,
		language.French:               joeFrenchPhrase,
		language.German:               joeGermanPhrase,
		language.Italian:              joeItalianPhrase,
		language.Japanese:             joeJapanesePhrase,
		language.Korean:               joeKoreanPhrase,
		language.LatinAmericanSpanish: joeLatinAmericanSpanishPhrase,
		language.Russian:              joeRussianPhrase,
		language.SimplifiedChinese:    joeSimplifiedChinesePhrase,
		language.Spanish:              joeSpanishPhrase,
		language.TraditionalChinese:   joeTraditionalChinesePhrase}
)

var (
	// Joe represents joe.
	Joe = nook.Villager{
		Character:   joeCharacter,
		Personality: personality.Cranky,
		Phrase:      joePhrase}
)
