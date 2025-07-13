package horse

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
	// roscoeBirthday represents roscoe birthday.
	roscoeBirthday = nook.Birthday{
		Day:   16,
		Month: time.June}
)

var (
	// roscoeCode represents roscoe code.
	roscoeCode = nook.Code{
		Value: "hrs04"}
)

var (
	// roscoeAmericanEnglishName represents roscoe american english name.
	roscoeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Roscoe"}

	// roscoeCanadianFrenchName represents roscoe canadian french name.
	roscoeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosco"}

	// roscoeDutchName represents roscoe dutch name.
	roscoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Roscoe"}

	// roscoeFrenchName represents roscoe french name.
	roscoeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosco"}

	// roscoeGermanName represents roscoe german name.
	roscoeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jolly"}

	// roscoeItalianName represents roscoe italian name.
	roscoeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Roscoe"}

	// roscoeJapaneseName represents roscoe japanese name.
	roscoeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シュバルツ"}

	// roscoeKoreanName represents roscoe korean name.
	roscoeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "슈베르트"}

	// roscoeLatinAmericanSpanishName represents roscoe latin american spanish name.
	roscoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jereza"}

	// roscoeRussianName represents roscoe russian name.
	roscoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роско"}

	// roscoeSimplifiedChineseName represents roscoe simplified chinese name.
	roscoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "黑马"}

	// roscoeSpanishName represents roscoe spanish name.
	roscoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jereza"}

	// roscoeTraditionalChineseName represents roscoe traditional chinese name.
	roscoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "黑馬"}
)

var (
	// roscoeName represents roscoe name.
	roscoeName = nook.Languages{
		language.AmericanEnglish:      roscoeAmericanEnglishName,
		language.CanadianFrench:       roscoeCanadianFrenchName,
		language.Dutch:                roscoeDutchName,
		language.French:               roscoeFrenchName,
		language.German:               roscoeGermanName,
		language.Italian:              roscoeItalianName,
		language.Japanese:             roscoeJapaneseName,
		language.Korean:               roscoeKoreanName,
		language.LatinAmericanSpanish: roscoeLatinAmericanSpanishName,
		language.Russian:              roscoeRussianName,
		language.SimplifiedChinese:    roscoeSimplifiedChineseName,
		language.Spanish:              roscoeSpanishName,
		language.TraditionalChinese:   roscoeTraditionalChineseName}
)

var (
	// roscoeCharacter represents roscoe character.
	roscoeCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: roscoeBirthday,
		Code:     roscoeCode,
		Key:      character.Roscoe,
		Gender:   gender.Male,
		Name:     roscoeName,
		Special:  false}
)

var (
	// roscoeAmericanEnglishPhrase represents roscoe american english phrase.
	roscoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nay"}

	// roscoeCanadianFrenchPhrase represents roscoe canadian french phrase.
	roscoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "shérif"}

	// roscoeDutchPhrase represents roscoe dutch phrase.
	roscoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nihihiet"}

	// roscoeFrenchPhrase represents roscoe french phrase.
	roscoeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "shérif"}

	// roscoeGermanPhrase represents roscoe german phrase.
	roscoeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "galoppp"}

	// roscoeItalianPhrase represents roscoe italian phrase.
	roscoeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nay"}

	// roscoeJapanesePhrase represents roscoe japanese phrase.
	roscoeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブルル"}

	// roscoeKoreanPhrase represents roscoe korean phrase.
	roscoeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부르르"}

	// roscoeLatinAmericanSpanishPhrase represents roscoe latin american spanish phrase.
	roscoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiii"}

	// roscoeRussianPhrase represents roscoe russian phrase.
	roscoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "не-е-ет"}

	// roscoeSimplifiedChinesePhrase represents roscoe simplified chinese phrase.
	roscoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布鲁鲁"}

	// roscoeSpanishPhrase represents roscoe spanish phrase.
	roscoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bípedo"}

	// roscoeTraditionalChinesePhrase represents roscoe traditional chinese phrase.
	roscoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布魯魯"}
)

var (
	// roscoePhrase represents roscoe phrase.
	roscoePhrase = nook.Languages{
		language.AmericanEnglish:      roscoeAmericanEnglishPhrase,
		language.CanadianFrench:       roscoeCanadianFrenchPhrase,
		language.Dutch:                roscoeDutchPhrase,
		language.French:               roscoeFrenchPhrase,
		language.German:               roscoeGermanPhrase,
		language.Italian:              roscoeItalianPhrase,
		language.Japanese:             roscoeJapanesePhrase,
		language.Korean:               roscoeKoreanPhrase,
		language.LatinAmericanSpanish: roscoeLatinAmericanSpanishPhrase,
		language.Russian:              roscoeRussianPhrase,
		language.SimplifiedChinese:    roscoeSimplifiedChinesePhrase,
		language.Spanish:              roscoeSpanishPhrase,
		language.TraditionalChinese:   roscoeTraditionalChinesePhrase}
)

var (
	// Roscoe represents roscoe.
	Roscoe = nook.Villager{
		Character:   roscoeCharacter,
		Personality: personality.Cranky,
		Phrase:      roscoePhrase}
)
