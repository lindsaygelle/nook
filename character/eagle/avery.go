package eagle

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
	// averyBirthday represents avery birthday.
	averyBirthday = nook.Birthday{
		Day:   22,
		Month: time.February}
)

var (
	// averyCode represents avery code.
	averyCode = nook.Code{
		Value: "pbr05"}
)

var (
	// averyAmericanEnglishName represents avery american english name.
	averyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Avery"}

	// averyCanadianFrenchName represents avery canadian french name.
	averyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Faust"}

	// averyDutchName represents avery dutch name.
	averyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Avery"}

	// averyFrenchName represents avery french name.
	averyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Faust"}

	// averyGermanName represents avery german name.
	averyGermanName = nook.Name{
		Language: language.German,
		Value:    "Quetzal"}

	// averyItalianName represents avery italian name.
	averyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Falco"}

	// averyJapaneseName represents avery japanese name.
	averyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クスケチャ"}

	// averyKoreanName represents avery korean name.
	averyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쿠스케처"}

	// averyLatinAmericanSpanishName represents avery latin american spanish name.
	averyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuzco"}

	// averyRussianName represents avery russian name.
	averyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Авери"}

	// averySimplifiedChineseName represents avery simplified chinese name.
	averySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安谷斯"}

	// averySpanishName represents avery spanish name.
	averySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuzco"}

	// averyTraditionalChineseName represents avery traditional chinese name.
	averyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安谷斯"}
)

var (
	// averyName represents avery name.
	averyName = nook.Languages{
		language.AmericanEnglish:      averyAmericanEnglishName,
		language.CanadianFrench:       averyCanadianFrenchName,
		language.Dutch:                averyDutchName,
		language.French:               averyFrenchName,
		language.German:               averyGermanName,
		language.Italian:              averyItalianName,
		language.Japanese:             averyJapaneseName,
		language.Korean:               averyKoreanName,
		language.LatinAmericanSpanish: averyLatinAmericanSpanishName,
		language.Russian:              averyRussianName,
		language.SimplifiedChinese:    averySimplifiedChineseName,
		language.Spanish:              averySpanishName,
		language.TraditionalChinese:   averyTraditionalChineseName}
)

var (
	// averyCharacter represents avery character.
	averyCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: averyBirthday,
		Code:     averyCode,
		Key:      character.Avery,
		Gender:   gender.Male,
		Name:     averyName,
		Special:  false}
)

var (
	// averyAmericanEnglishPhrase represents avery american english phrase.
	averyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "skree-haw"}

	// averyCanadianFrenchPhrase represents avery canadian french phrase.
	averyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grrraak"}

	// averyDutchPhrase represents avery dutch phrase.
	averyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kra-hoe"}

	// averyFrenchPhrase represents avery french phrase.
	averyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "phélès"}

	// averyGermanPhrase represents avery german phrase.
	averyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flapp"}

	// averyItalianPhrase represents avery italian phrase.
	averyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gawk"}

	// averyJapanesePhrase represents avery japanese phrase.
	averyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アリョイ"}

	// averyKoreanPhrase represents avery korean phrase.
	averyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "내놔"}

	// averyLatinAmericanSpanishPhrase represents avery latin american spanish phrase.
	averyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrraak"}

	// averyRussianPhrase represents avery russian phrase.
	averyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "курлы"}

	// averySimplifiedChinesePhrase represents avery simplified chinese phrase.
	averySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "溜溜"}

	// averySpanishPhrase represents avery spanish phrase.
	averySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grrraak"}

	// averyTraditionalChinesePhrase represents avery traditional chinese phrase.
	averyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "溜溜"}
)

var (
	// averyPhrase represents avery phrase.
	averyPhrase = nook.Languages{
		language.AmericanEnglish:      averyAmericanEnglishPhrase,
		language.CanadianFrench:       averyCanadianFrenchPhrase,
		language.Dutch:                averyDutchPhrase,
		language.French:               averyFrenchPhrase,
		language.German:               averyGermanPhrase,
		language.Italian:              averyItalianPhrase,
		language.Japanese:             averyJapanesePhrase,
		language.Korean:               averyKoreanPhrase,
		language.LatinAmericanSpanish: averyLatinAmericanSpanishPhrase,
		language.Russian:              averyRussianPhrase,
		language.SimplifiedChinese:    averySimplifiedChinesePhrase,
		language.Spanish:              averySpanishPhrase,
		language.TraditionalChinese:   averyTraditionalChinesePhrase}
)

var (
	// Avery represents avery.
	Avery = nook.Villager{
		Character:   averyCharacter,
		Personality: personality.Cranky,
		Phrase:      averyPhrase}
)
