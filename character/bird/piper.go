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
	// piperBirthday represents piper birthday.
	piperBirthday = nook.Birthday{
		Day:   18,
		Month: time.April}
)

var (
	// piperCode represents piper code.
	piperCode = nook.Code{
		Value: "brd05"}
)

var (
	// piperAmericanEnglishName represents piper american english name.
	piperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Piper"}

	// piperCanadianFrenchName represents piper canadian french name.
	piperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Neige"}

	// piperDutchName represents piper dutch name.
	piperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Piper"}

	// piperFrenchName represents piper french name.
	piperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neige"}

	// piperGermanName represents piper german name.
	piperGermanName = nook.Name{
		Language: language.German,
		Value:    "Iris"}

	// piperItalianName represents piper italian name.
	piperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giuliva"}

	// piperJapaneseName represents piper japanese name.
	piperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイコ"}

	// piperKoreanName represents piper korean name.
	piperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파이프"}

	// piperLatinAmericanSpanishName represents piper latin american spanish name.
	piperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bárbara"}

	// piperRussianName represents piper russian name.
	piperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пайпер"}

	// piperSimplifiedChineseName represents piper simplified chinese name.
	piperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丽婷"}

	// piperSpanishName represents piper spanish name.
	piperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bárbara"}

	// piperTraditionalChineseName represents piper traditional chinese name.
	piperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麗婷"}
)

var (
	// piperName represents piper name.
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
	// piperCharacter represents piper character.
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
	// piperAmericanEnglishPhrase represents piper american english phrase.
	piperAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chickadee"}

	// piperCanadianFrenchPhrase represents piper canadian french phrase.
	piperCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "puic puic"}

	// piperDutchPhrase represents piper dutch phrase.
	piperDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tjiep tjiep"}

	// piperFrenchPhrase represents piper french phrase.
	piperFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "puic puic"}

	// piperGermanPhrase represents piper german phrase.
	piperGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschikadii"}

	// piperItalianPhrase represents piper italian phrase.
	piperItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cocoricò"}

	// piperJapanesePhrase represents piper japanese phrase.
	piperJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "けどー"}

	// piperKoreanPhrase represents piper korean phrase.
	piperKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "허나"}

	// piperLatinAmericanSpanishPhrase represents piper latin american spanish phrase.
	piperLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chupiyá"}

	// piperRussianPhrase represents piper russian phrase.
	piperRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чики-чик"}

	// piperSimplifiedChinesePhrase represents piper simplified chinese phrase.
	piperSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "山雀"}

	// piperSpanishPhrase represents piper spanish phrase.
	piperSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chupiyá"}

	// piperTraditionalChinesePhrase represents piper traditional chinese phrase.
	piperTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "山雀"}
)

var (
	// piperPhrase represents piper phrase.
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
	// Piper represents piper.
	Piper = nook.Villager{
		Character:   piperCharacter,
		Personality: personality.Peppy,
		Phrase:      piperPhrase}
)
