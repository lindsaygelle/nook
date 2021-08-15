package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pierceBirthday = nook.Birthday{
		Day:   8,
		Month: time.January}
)

var (
	pierceCode = nook.Code{
		Value: "pbr02"}
)

var (
	pierceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pierce"}

	pierceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Napoléontout vu"}

	pierceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Piercearendsoog"}

	pierceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Napoléonœil agile"}

	pierceGermanName = nook.Name{
		Language: language.German,
		Value:    "Adrianschwinge"}

	pierceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Attiliopronti via"}

	pierceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セバスチャンバサバサ"}

	pierceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "세바스찬퍼덕퍼덕"}

	pierceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hugoaguilííí"}

	pierceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пирсглаз-алмаз"}

	pierceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谢博强啪沙啪沙"}

	pierceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hugoaguililla"}

	pierceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "謝博強啪沙啪沙"}
)

var (
	pierceName = nook.Languages{
		language.AmericanEnglish:      pierceAmericanEnglishName,
		language.CanadianFrench:       pierceCanadianFrenchName,
		language.Dutch:                pierceDutchName,
		language.French:               pierceFrenchName,
		language.German:               pierceGermanName,
		language.Italian:              pierceItalianName,
		language.Japanese:             pierceJapaneseName,
		language.Korean:               pierceKoreanName,
		language.LatinAmericanSpanish: pierceLatinAmericanSpanishName,
		language.Russian:              pierceRussianName,
		language.SimplifiedChinese:    pierceSimplifiedChineseName,
		language.Spanish:              pierceSpanishName,
		language.TraditionalChinese:   pierceTraditionalChineseName}
)

var (
	pierceCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: pierceBirthday,
		Code:     pierceCode,
		Gender:   nook.Male,
		Name:     pierceName}
)

var (
	pierceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wingerhawkeye"}

	pierceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tout vu"}

	pierceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "arendsoog"}

	pierceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "œil agile"}

	pierceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schwinge"}

	pierceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pronti via"}

	pierceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バサバサ"}

	pierceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "퍼덕퍼덕"}

	pierceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aguilííí"}

	pierceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "глаз-алмаз"}

	pierceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啪沙啪沙"}

	pierceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aguililla"}

	pierceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啪沙啪沙"}
)

var (
	piercePhrase = nook.Languages{
		language.AmericanEnglish:      pierceAmericanEnglishName,
		language.CanadianFrench:       pierceCanadianFrenchName,
		language.Dutch:                pierceDutchName,
		language.French:               pierceFrenchName,
		language.German:               pierceGermanName,
		language.Italian:              pierceItalianName,
		language.Japanese:             pierceJapaneseName,
		language.Korean:               pierceKoreanName,
		language.LatinAmericanSpanish: pierceLatinAmericanSpanishName,
		language.Russian:              pierceRussianName,
		language.SimplifiedChinese:    pierceSimplifiedChineseName,
		language.Spanish:              pierceSpanishName,
		language.TraditionalChinese:   pierceTraditionalChineseName}
)

var (
	Pierce = nook.Villager{
		Character:   pierceCharacter,
		Personality: nook.Jock,
		Phrase:      piercePhrase}
)
