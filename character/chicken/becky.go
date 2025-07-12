package chicken

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
	// beckyBirthday represents becky birthday.
	beckyBirthday = nook.Birthday{
		Day:   9,
		Month: time.December}
)

var (
	// beckyCode represents becky code.
	beckyCode = nook.Code{
		Value: "chn09"}
)

var (
	// beckyAmericanEnglishName represents becky american english name.
	beckyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Becky"}

	// beckyCanadianFrenchName represents becky canadian french name.
	beckyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sonya"}

	// beckyDutchName represents becky dutch name.
	beckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Becky"}

	// beckyFrenchName represents becky french name.
	beckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sonya"}

	// beckyGermanName represents becky german name.
	beckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Inga"}

	// beckyItalianName represents becky italian name.
	beckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annetta"}

	// beckyJapaneseName represents becky japanese name.
	beckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリア"}

	// beckyKoreanName represents becky korean name.
	beckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아리아"}

	// beckyLatinAmericanSpanishName represents becky latin american spanish name.
	beckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ramina"}

	// beckyRussianName represents becky russian name.
	beckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бекки"}

	// beckySimplifiedChineseName represents becky simplified chinese name.
	beckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咏旋"}

	// beckySpanishName represents becky spanish name.
	beckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ramina"}

	// beckyTraditionalChineseName represents becky traditional chinese name.
	beckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "詠旋"}
)

var (
	// beckyName represents becky name.
	beckyName = nook.Languages{
		language.AmericanEnglish:      beckyAmericanEnglishName,
		language.CanadianFrench:       beckyCanadianFrenchName,
		language.Dutch:                beckyDutchName,
		language.French:               beckyFrenchName,
		language.German:               beckyGermanName,
		language.Italian:              beckyItalianName,
		language.Japanese:             beckyJapaneseName,
		language.Korean:               beckyKoreanName,
		language.LatinAmericanSpanish: beckyLatinAmericanSpanishName,
		language.Russian:              beckyRussianName,
		language.SimplifiedChinese:    beckySimplifiedChineseName,
		language.Spanish:              beckySpanishName,
		language.TraditionalChinese:   beckyTraditionalChineseName}
)

var (
	// beckyCharacter represents becky character.
	beckyCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: beckyBirthday,
		Code:     beckyCode,
		Key:      character.Becky,
		Gender:   gender.Female,
		Name:     beckyName,
		Special:  false}
)

var (
	// beckyAmericanEnglishPhrase represents becky american english phrase.
	beckyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chicklet"}

	// beckyCanadianFrenchPhrase represents becky canadian french phrase.
	beckyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "côôôôt"}

	// beckyDutchPhrase represents becky dutch phrase.
	beckyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kuiken"}

	// beckyFrenchPhrase represents becky french phrase.
	beckyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "côôôôt"}

	// beckyGermanPhrase represents becky german phrase.
	beckyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bokbok"}

	// beckyItalianPhrase represents becky italian phrase.
	beckyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coccoricò"}

	// beckyJapanesePhrase represents becky japanese phrase.
	beckyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハレルヤ"}

	// beckyKoreanPhrase represents becky korean phrase.
	beckyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "도레미"}

	// beckyLatinAmericanSpanishPhrase represents becky latin american spanish phrase.
	beckyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cucurrús"}

	// beckyRussianPhrase represents becky russian phrase.
	beckyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "модненько"}

	// beckySimplifiedChinesePhrase represents becky simplified chinese phrase.
	beckySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈雷路亚"}

	// beckySpanishPhrase represents becky spanish phrase.
	beckySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cucurrús"}

	// beckyTraditionalChinesePhrase represents becky traditional chinese phrase.
	beckyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈雷路亞"}
)

var (
	// beckyPhrase represents becky phrase.
	beckyPhrase = nook.Languages{
		language.AmericanEnglish:      beckyAmericanEnglishPhrase,
		language.CanadianFrench:       beckyCanadianFrenchPhrase,
		language.Dutch:                beckyDutchPhrase,
		language.French:               beckyFrenchPhrase,
		language.German:               beckyGermanPhrase,
		language.Italian:              beckyItalianPhrase,
		language.Japanese:             beckyJapanesePhrase,
		language.Korean:               beckyKoreanPhrase,
		language.LatinAmericanSpanish: beckyLatinAmericanSpanishPhrase,
		language.Russian:              beckyRussianPhrase,
		language.SimplifiedChinese:    beckySimplifiedChinesePhrase,
		language.Spanish:              beckySpanishPhrase,
		language.TraditionalChinese:   beckyTraditionalChinesePhrase}
)

var (
	// Becky represents becky.
	Becky = nook.Villager{
		Character:   beckyCharacter,
		Personality: personality.Snooty,
		Phrase:      beckyPhrase}
)
