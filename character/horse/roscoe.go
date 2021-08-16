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
	roscoeBirthday = nook.Birthday{
		Day:   16,
		Month: time.June}
)

var (
	roscoeCode = nook.Code{
		Value: "hrs04"}
)

var (
	roscoeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Roscoe"}

	roscoeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosco"}

	roscoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Roscoe"}

	roscoeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosco"}

	roscoeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jolly"}

	roscoeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Roscoe"}

	roscoeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シュバルツ"}

	roscoeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "슈베르트"}

	roscoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jereza"}

	roscoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роско"}

	roscoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "黑马"}

	roscoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jereza"}

	roscoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "黑馬"}
)

var (
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
	roscoeCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: roscoeBirthday,
		Code:     roscoeCode,
		Key:      character.Roscoe,
		Gender:   gender.Male,
		Name:     roscoeName}
)

var (
	roscoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nay"}

	roscoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "shérif"}

	roscoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nihihiet"}

	roscoeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "shérif"}

	roscoeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "galoppp"}

	roscoeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nay"}

	roscoeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブルル"}

	roscoeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부르르"}

	roscoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiii"}

	roscoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "не-е-ет"}

	roscoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布鲁鲁"}

	roscoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bípedo"}

	roscoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布魯魯"}
)

var (
	roscoePhrase = nook.Languages{
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
	Roscoe = nook.Villager{
		Character:   roscoeCharacter,
		Personality: personality.Cranky,
		Phrase:      roscoePhrase}
)
