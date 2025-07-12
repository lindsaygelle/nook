package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// delBirthday represents Del's birthday.
var (
	// delBirthday represents del birthday.
	delBirthday = nook.Birthday{
		Day:   27,
		Month: time.May}
)

// delCode represents Del's unique code.
var (
	// delCode represents del code.
	delCode = nook.Code{
		Value: "crd04"}
)

// Different names for Del in various languages.
var (
	// delAmericanEnglishName represents del american english name.
	delAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Del"}

	// delCanadianFrenchName represents del canadian french name.
	delCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hector"}

	// delDutchName represents del dutch name.
	delDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Del"}

	// delFrenchName represents del french name.
	delFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hector"}

	// delGermanName represents del german name.
	delGermanName = nook.Name{
		Language: language.German,
		Value:    "Krokki"}

	// delItalianName represents del italian name.
	delItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Krokki"}

	// delJapaneseName represents del japanese name.
	delJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヤマト"}

	// delKoreanName represents del korean name.
	delKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파도맨"}

	// delLatinAmericanSpanishName represents del latin american spanish name.
	delLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Croco"}

	// delRussianName represents del russian name.
	delRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дел"}

	// delSimplifiedChineseName represents del simplified chinese name.
	delSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大和"}

	// delSpanishName represents del spanish name.
	delSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Croco"}

	// delTraditionalChineseName represents del traditional chinese name.
	delTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大和"}
)

// delName represents Del's name in different languages.
var (
	// delName represents del name.
	delName = nook.Languages{
		language.AmericanEnglish:      delAmericanEnglishName,
		language.CanadianFrench:       delCanadianFrenchName,
		language.Dutch:                delDutchName,
		language.French:               delFrenchName,
		language.German:               delGermanName,
		language.Italian:              delItalianName,
		language.Japanese:             delJapaneseName,
		language.Korean:               delKoreanName,
		language.LatinAmericanSpanish: delLatinAmericanSpanishName,
		language.Russian:              delRussianName,
		language.SimplifiedChinese:    delSimplifiedChineseName,
		language.Spanish:              delSpanishName,
		language.TraditionalChinese:   delTraditionalChineseName}
)

// delCharacter represents Del's character information.
var (
	// delCharacter represents del character.
	delCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: delBirthday,
		Code:     delCode,
		Key:      character.Del,
		Gender:   gender.Male,
		Name:     delName,
		Special:  false}
)

// Different phrases spoken by Del in various languages.
var (
	// delAmericanEnglishPhrase represents del american english phrase.
	delAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gronk"}

	// delCanadianFrenchPhrase represents del canadian french phrase.
	delCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "clap"}

	// delDutchPhrase represents del dutch phrase.
	delDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gronk"}

	// delFrenchPhrase represents del french phrase.
	delFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "clap"}

	// delGermanPhrase represents del german phrase.
	delGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "happ"}

	// delItalianPhrase represents del italian phrase.
	delItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chomp"}

	// delJapanesePhrase represents del japanese phrase.
	delJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "プシュー"}

	// delKoreanPhrase represents del korean phrase.
	delKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "처얼썩"}

	// delLatinAmericanSpanishPhrase represents del latin american spanish phrase.
	delLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chomp"}

	// delRussianPhrase represents del russian phrase.
	delRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щелк"}

	// delSimplifiedChinesePhrase represents del simplified chinese phrase.
	delSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗咻"}

	// delSpanishPhrase represents del spanish phrase.
	delSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chomp"}

	// delTraditionalChinesePhrase represents del traditional chinese phrase.
	delTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗咻"}
)

// delPhrase represents Del's phrases in different languages.
var (
	// delPhrase represents del phrase.
	delPhrase = nook.Languages{
		language.AmericanEnglish:      delAmericanEnglishPhrase,
		language.CanadianFrench:       delCanadianFrenchPhrase,
		language.Dutch:                delDutchPhrase,
		language.French:               delFrenchPhrase,
		language.German:               delGermanPhrase,
		language.Italian:              delItalianPhrase,
		language.Japanese:             delJapanesePhrase,
		language.Korean:               delKoreanPhrase,
		language.LatinAmericanSpanish: delLatinAmericanSpanishPhrase,
		language.Russian:              delRussianPhrase,
		language.SimplifiedChinese:    delSimplifiedChinesePhrase,
		language.Spanish:              delSpanishPhrase,
		language.TraditionalChinese:   delTraditionalChinesePhrase}
)

// Del represents the character Del with his complete information.
var (
	// Del represents del.
	Del = nook.Villager{
		Character:   delCharacter,
		Personality: personality.Cranky,
		Phrase:      delPhrase}
)
