package cow

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
	// tipperBirthday represents tipper birthday.
	tipperBirthday = nook.Birthday{
		Day:   25,
		Month: time.August}
)

var (
	// tipperCode represents tipper code.
	tipperCode = nook.Code{
		Value: "cow01"}
)

var (
	// tipperAmericanEnglishName represents tipper american english name.
	tipperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tipper"}

	// tipperCanadianFrenchName represents tipper canadian french name.
	tipperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Valé"}

	// tipperDutchName represents tipper dutch name.
	tipperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tipper"}

	// tipperFrenchName represents tipper french name.
	tipperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Valé"}

	// tipperGermanName represents tipper german name.
	tipperGermanName = nook.Name{
		Language: language.German,
		Value:    "Angela"}

	// tipperItalianName represents tipper italian name.
	tipperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tipper"}

	// tipperJapaneseName represents tipper japanese name.
	tipperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まきば"}

	// tipperKoreanName represents tipper korean name.
	tipperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마틸다"}

	// tipperLatinAmericanSpanishName represents tipper latin american spanish name.
	tipperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pinta"}

	// tipperRussianName represents tipper russian name.
	tipperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Типпер"}

	// tipperSimplifiedChineseName represents tipper simplified chinese name.
	tipperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牧牧"}

	// tipperSpanishName represents tipper spanish name.
	tipperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pinta"}

	// tipperTraditionalChineseName represents tipper traditional chinese name.
	tipperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牧牧"}
)

var (
	// tipperName represents tipper name.
	tipperName = nook.Languages{
		language.AmericanEnglish:      tipperAmericanEnglishName,
		language.CanadianFrench:       tipperCanadianFrenchName,
		language.Dutch:                tipperDutchName,
		language.French:               tipperFrenchName,
		language.German:               tipperGermanName,
		language.Italian:              tipperItalianName,
		language.Japanese:             tipperJapaneseName,
		language.Korean:               tipperKoreanName,
		language.LatinAmericanSpanish: tipperLatinAmericanSpanishName,
		language.Russian:              tipperRussianName,
		language.SimplifiedChinese:    tipperSimplifiedChineseName,
		language.Spanish:              tipperSpanishName,
		language.TraditionalChinese:   tipperTraditionalChineseName}
)

var (
	// tipperCharacter represents tipper character.
	tipperCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: tipperBirthday,
		Code:     tipperCode,
		Key:      character.Tipper,
		Gender:   gender.Female,
		Name:     tipperName,
		Special:  false}
)

var (
	// tipperAmericanEnglishPhrase represents tipper american english phrase.
	tipperAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pushy"}

	// tipperCanadianFrenchPhrase represents tipper canadian french phrase.
	tipperCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "choupette"}

	// tipperDutchPhrase represents tipper dutch phrase.
	tipperDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boelala"}

	// tipperFrenchPhrase represents tipper french phrase.
	tipperFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "choupette"}

	// tipperGermanPhrase represents tipper german phrase.
	tipperGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "dingdong"}

	// tipperItalianPhrase represents tipper italian phrase.
	tipperItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sì sì sì"}

	// tipperJapanesePhrase represents tipper japanese phrase.
	tipperJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ミルミル"}

	// tipperKoreanPhrase represents tipper korean phrase.
	tipperKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우유우유"}

	// tipperLatinAmericanSpanishPhrase represents tipper latin american spanish phrase.
	tipperLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muamuuu"}

	// tipperRussianPhrase represents tipper russian phrase.
	tipperRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "теленок"}

	// tipperSimplifiedChinesePhrase represents tipper simplified chinese phrase.
	tipperSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牛奶"}

	// tipperSpanishPhrase represents tipper spanish phrase.
	tipperSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuernitos"}

	// tipperTraditionalChinesePhrase represents tipper traditional chinese phrase.
	tipperTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牛奶"}
)

var (
	// tipperPhrase represents tipper phrase.
	tipperPhrase = nook.Languages{
		language.AmericanEnglish:      tipperAmericanEnglishPhrase,
		language.CanadianFrench:       tipperCanadianFrenchPhrase,
		language.Dutch:                tipperDutchPhrase,
		language.French:               tipperFrenchPhrase,
		language.German:               tipperGermanPhrase,
		language.Italian:              tipperItalianPhrase,
		language.Japanese:             tipperJapanesePhrase,
		language.Korean:               tipperKoreanPhrase,
		language.LatinAmericanSpanish: tipperLatinAmericanSpanishPhrase,
		language.Russian:              tipperRussianPhrase,
		language.SimplifiedChinese:    tipperSimplifiedChinesePhrase,
		language.Spanish:              tipperSpanishPhrase,
		language.TraditionalChinese:   tipperTraditionalChinesePhrase}
)

var (
	// Tipper represents tipper.
	Tipper = nook.Villager{
		Character:   tipperCharacter,
		Personality: personality.Snooty,
		Phrase:      tipperPhrase}
)
