package koala

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
	// eugeneBirthday represents eugene birthday.
	eugeneBirthday = nook.Birthday{
		Day:   26,
		Month: time.October}
)

var (
	// eugeneCode represents eugene code.
	eugeneCode = nook.Code{
		Value: "kal10"}
)

var (
	// eugeneAmericanEnglishName represents eugene american english name.
	eugeneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eugene"}

	// eugeneCanadianFrenchName represents eugene canadian french name.
	eugeneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jamy"}

	// eugeneDutchName represents eugene dutch name.
	eugeneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eugene"}

	// eugeneFrenchName represents eugene french name.
	eugeneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jamy"}

	// eugeneGermanName represents eugene german name.
	eugeneGermanName = nook.Name{
		Language: language.German,
		Value:    "Sunny"}

	// eugeneItalianName represents eugene italian name.
	eugeneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Corrado"}

	// eugeneJapaneseName represents eugene japanese name.
	eugeneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロッキー"}

	// eugeneKoreanName represents eugene korean name.
	eugeneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "코알"}

	// eugeneLatinAmericanSpanishName represents eugene latin american spanish name.
	eugeneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eucalín"}

	// eugeneRussianName represents eugene russian name.
	eugeneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юджин"}

	// eugeneSimplifiedChineseName represents eugene simplified chinese name.
	eugeneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗奇"}

	// eugeneSpanishName represents eugene spanish name.
	eugeneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eucalín"}

	// eugeneTraditionalChineseName represents eugene traditional chinese name.
	eugeneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅奇"}
)

var (
	// eugeneName represents eugene name.
	eugeneName = nook.Languages{
		language.AmericanEnglish:      eugeneAmericanEnglishName,
		language.CanadianFrench:       eugeneCanadianFrenchName,
		language.Dutch:                eugeneDutchName,
		language.French:               eugeneFrenchName,
		language.German:               eugeneGermanName,
		language.Italian:              eugeneItalianName,
		language.Japanese:             eugeneJapaneseName,
		language.Korean:               eugeneKoreanName,
		language.LatinAmericanSpanish: eugeneLatinAmericanSpanishName,
		language.Russian:              eugeneRussianName,
		language.SimplifiedChinese:    eugeneSimplifiedChineseName,
		language.Spanish:              eugeneSpanishName,
		language.TraditionalChinese:   eugeneTraditionalChineseName}
)

var (
	// eugeneCharacter represents eugene character.
	eugeneCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: eugeneBirthday,
		Code:     eugeneCode,
		Key:      character.Eugene,
		Gender:   gender.Male,
		Name:     eugeneName,
		Special:  false}
)

var (
	// eugeneAmericanEnglishPhrase represents eugene american english phrase.
	eugeneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yeah buddy"}

	// eugeneCanadianFrenchPhrase represents eugene canadian french phrase.
	eugeneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh yeah"}

	// eugeneDutchPhrase represents eugene dutch phrase.
	eugeneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "yeah"}

	// eugeneFrenchPhrase represents eugene french phrase.
	eugeneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yapaphoto"}

	// eugeneGermanPhrase represents eugene german phrase.
	eugeneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "yeah"}

	// eugeneItalianPhrase represents eugene italian phrase.
	eugeneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "koahola"}

	// eugeneJapanesePhrase represents eugene japanese phrase.
	eugeneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おいおい"}

	// eugeneKoreanPhrase represents eugene korean phrase.
	eugeneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야야야"}

	// eugeneLatinAmericanSpanishPhrase represents eugene latin american spanish phrase.
	eugeneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ouyeah"}

	// eugeneRussianPhrase represents eugene russian phrase.
	eugeneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "приятель"}

	// eugeneSimplifiedChinesePhrase represents eugene simplified chinese phrase.
	eugeneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂喂"}

	// eugeneSpanishPhrase represents eugene spanish phrase.
	eugeneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ouyeah"}

	// eugeneTraditionalChinesePhrase represents eugene traditional chinese phrase.
	eugeneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂喂"}
)

var (
	// eugenePhrase represents eugene phrase.
	eugenePhrase = nook.Languages{
		language.AmericanEnglish:      eugeneAmericanEnglishPhrase,
		language.CanadianFrench:       eugeneCanadianFrenchPhrase,
		language.Dutch:                eugeneDutchPhrase,
		language.French:               eugeneFrenchPhrase,
		language.German:               eugeneGermanPhrase,
		language.Italian:              eugeneItalianPhrase,
		language.Japanese:             eugeneJapanesePhrase,
		language.Korean:               eugeneKoreanPhrase,
		language.LatinAmericanSpanish: eugeneLatinAmericanSpanishPhrase,
		language.Russian:              eugeneRussianPhrase,
		language.SimplifiedChinese:    eugeneSimplifiedChinesePhrase,
		language.Spanish:              eugeneSpanishPhrase,
		language.TraditionalChinese:   eugeneTraditionalChinesePhrase}
)

var (
	// Eugene represents eugene.
	Eugene = nook.Villager{
		Character:   eugeneCharacter,
		Personality: personality.Smug,
		Phrase:      eugenePhrase}
)
