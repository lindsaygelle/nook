package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Neigepuic puic"}

	piperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pipertjiep tjiep"}

	piperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neigepuic puic"}

	piperGermanName = nook.Name{
		Language: language.German,
		Value:    "Iristschikadii"}

	piperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giulivacocoricò"}

	piperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイコけどー"}

	piperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파이프허나"}

	piperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bárbarachupiyá"}

	piperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пайперчики-чик"}

	piperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丽婷山雀"}

	piperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bárbarachupiyá"}

	piperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麗婷山雀"}
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
		Animal:   Bird,
		Birthday: piperBirthday,
		Code:     piperCode,
		Gender:   nook.Female,
		Name:     piperName}
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
	Piper = nook.Villager{
		Character:   piperCharacter,
		Personality: nook.Peppy,
		Phrase:      piperPhrase}
)
