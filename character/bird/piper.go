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
	piperBirthday = nook.Birthday{
		Day:   18,
		Month: time.April}
)

var (
	piperCode = nook.Code{
		Value: "brd05"}
)

var (
	piperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Piper"}

	piperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Neige"}

	piperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Piper"}

	piperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neige"}

	piperGermanName = nook.Name{
		Language: language.German,
		Value:    "Iris"}

	piperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giuliva"}

	piperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイコ"}

	piperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파이프"}

	piperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bárbara"}

	piperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пайпер"}

	piperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丽婷"}

	piperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bárbara"}

	piperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麗婷"}
)

var (
	piperName = nook.Languages{
		language.AmericanEnglish:      piperAmericanEnglishName,
		language.CanadianFrench:       piperCanadianFrenchName,
		language.Dutch:                piperDutchName,
		language.French:               piperFrenchName,
		language.German:               piperGermanName,
		language.Italian:              piperItalianName,
		language.Japanese:             piperJapaneseName,
		language.Korean:               piperKoreanName,
		language.LatinAmericanSpanish: piperLatinAmericanSpanishName,
		language.Russian:              piperRussianName,
		language.SimplifiedChinese:    piperSimplifiedChineseName,
		language.Spanish:              piperSpanishName,
		language.TraditionalChinese:   piperTraditionalChineseName}
)

var (
	piperCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: piperBirthday,
		Code:     piperCode,
		Key:      character.Piper,
		Gender:   gender.Female,
		Name:     piperName,
		Special:  false}
)

var (
	piperAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chickadee"}

	piperCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "puic puic"}

	piperDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tjiep tjiep"}

	piperFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "puic puic"}

	piperGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschikadii"}

	piperItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cocoricò"}

	piperJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "けどー"}

	piperKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "허나"}

	piperLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chupiyá"}

	piperRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чики-чик"}

	piperSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "山雀"}

	piperSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chupiyá"}

	piperTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "山雀"}
)

var (
	piperPhrase = nook.Languages{
		language.AmericanEnglish:      piperAmericanEnglishPhrase,
		language.CanadianFrench:       piperCanadianFrenchPhrase,
		language.Dutch:                piperDutchPhrase,
		language.French:               piperFrenchPhrase,
		language.German:               piperGermanPhrase,
		language.Italian:              piperItalianPhrase,
		language.Japanese:             piperJapanesePhrase,
		language.Korean:               piperKoreanPhrase,
		language.LatinAmericanSpanish: piperLatinAmericanSpanishPhrase,
		language.Russian:              piperRussianPhrase,
		language.SimplifiedChinese:    piperSimplifiedChinesePhrase,
		language.Spanish:              piperSpanishPhrase,
		language.TraditionalChinese:   piperTraditionalChinesePhrase}
)

var (
	Piper = nook.Villager{
		Character:   piperCharacter,
		Personality: personality.Peppy,
		Phrase:      piperPhrase}
)
