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
	// pierceBirthday represents pierce birthday.
	pierceBirthday = nook.Birthday{
		Day:   8,
		Month: time.January}
)

var (
	// pierceCode represents pierce code.
	pierceCode = nook.Code{
		Value: "pbr02"}
)

var (
	// pierceAmericanEnglishName represents pierce american english name.
	pierceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pierce"}

	// pierceCanadianFrenchName represents pierce canadian french name.
	pierceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Napoléon"}

	// pierceDutchName represents pierce dutch name.
	pierceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pierce"}

	// pierceFrenchName represents pierce french name.
	pierceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Napoléon"}

	// pierceGermanName represents pierce german name.
	pierceGermanName = nook.Name{
		Language: language.German,
		Value:    "Adrian"}

	// pierceItalianName represents pierce italian name.
	pierceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Attilio"}

	// pierceJapaneseName represents pierce japanese name.
	pierceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セバスチャン"}

	// pierceKoreanName represents pierce korean name.
	pierceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "세바스찬"}

	// pierceLatinAmericanSpanishName represents pierce latin american spanish name.
	pierceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hugo"}

	// pierceRussianName represents pierce russian name.
	pierceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пирс"}

	// pierceSimplifiedChineseName represents pierce simplified chinese name.
	pierceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谢博强"}

	// pierceSpanishName represents pierce spanish name.
	pierceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hugo"}

	// pierceTraditionalChineseName represents pierce traditional chinese name.
	pierceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "謝博強"}
)

var (
	// pierceName represents pierce name.
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
	// pierceCharacter represents pierce character.
	pierceCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: pierceBirthday,
		Code:     pierceCode,
		Key:      character.Pierce,
		Gender:   gender.Male,
		Name:     pierceName,
		Special:  false}
)

var (
	// pierceAmericanEnglishPhrase represents pierce american english phrase.
	pierceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "winger"}

	// pierceCanadianFrenchPhrase represents pierce canadian french phrase.
	pierceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tout vu"}

	// pierceDutchPhrase represents pierce dutch phrase.
	pierceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "arendsoog"}

	// pierceFrenchPhrase represents pierce french phrase.
	pierceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "œil agile"}

	// pierceGermanPhrase represents pierce german phrase.
	pierceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schwinge"}

	// pierceItalianPhrase represents pierce italian phrase.
	pierceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pronti via"}

	// pierceJapanesePhrase represents pierce japanese phrase.
	pierceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バサバサ"}

	// pierceKoreanPhrase represents pierce korean phrase.
	pierceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "퍼덕퍼덕"}

	// pierceLatinAmericanSpanishPhrase represents pierce latin american spanish phrase.
	pierceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aguilííí"}

	// pierceRussianPhrase represents pierce russian phrase.
	pierceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "глаз-алмаз"}

	// pierceSimplifiedChinesePhrase represents pierce simplified chinese phrase.
	pierceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啪沙啪沙"}

	// pierceSpanishPhrase represents pierce spanish phrase.
	pierceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aguililla"}

	// pierceTraditionalChinesePhrase represents pierce traditional chinese phrase.
	pierceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啪沙啪沙"}
)

var (
	// piercePhrase represents pierce phrase.
	piercePhrase = nook.Languages{
		language.AmericanEnglish:      pierceAmericanEnglishPhrase,
		language.CanadianFrench:       pierceCanadianFrenchPhrase,
		language.Dutch:                pierceDutchPhrase,
		language.French:               pierceFrenchPhrase,
		language.German:               pierceGermanPhrase,
		language.Italian:              pierceItalianPhrase,
		language.Japanese:             pierceJapanesePhrase,
		language.Korean:               pierceKoreanPhrase,
		language.LatinAmericanSpanish: pierceLatinAmericanSpanishPhrase,
		language.Russian:              pierceRussianPhrase,
		language.SimplifiedChinese:    pierceSimplifiedChinesePhrase,
		language.Spanish:              pierceSpanishPhrase,
		language.TraditionalChinese:   pierceTraditionalChinesePhrase}
)

var (
	// Pierce represents pierce.
	Pierce = nook.Villager{
		Character:   pierceCharacter,
		Personality: personality.Jock,
		Phrase:      piercePhrase}
)
