package mouse

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
	// bettinaBirthday represents bettina birthday.
	bettinaBirthday = nook.Birthday{
		Day:   12,
		Month: time.June}
)

var (
	// bettinaCode represents bettina code.
	bettinaCode = nook.Code{
		Value: "mus15"}
)

var (
	// bettinaAmericanEnglishName represents bettina american english name.
	bettinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bettina"}

	// bettinaCanadianFrenchName represents bettina canadian french name.
	bettinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sabrina"}

	// bettinaDutchName represents bettina dutch name.
	bettinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bettina"}

	// bettinaFrenchName represents bettina french name.
	bettinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sabrina"}

	// bettinaGermanName represents bettina german name.
	bettinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Schoki"}

	// bettinaItalianName represents bettina italian name.
	bettinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rattina"}

	// bettinaJapaneseName represents bettina japanese name.
	bettinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マルコ"}

	// bettinaKoreanName represents bettina korean name.
	bettinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마르카"}

	// bettinaLatinAmericanSpanishName represents bettina latin american spanish name.
	bettinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ottina"}

	// bettinaRussianName represents bettina russian name.
	bettinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беттина"}

	// bettinaSimplifiedChineseName represents bettina simplified chinese name.
	bettinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丸子"}

	// bettinaSpanishName represents bettina spanish name.
	bettinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ottina"}

	// bettinaTraditionalChineseName represents bettina traditional chinese name.
	bettinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丸子"}
)

var (
	// bettinaName represents bettina name.
	bettinaName = nook.Languages{
		language.AmericanEnglish:      bettinaAmericanEnglishName,
		language.CanadianFrench:       bettinaCanadianFrenchName,
		language.Dutch:                bettinaDutchName,
		language.French:               bettinaFrenchName,
		language.German:               bettinaGermanName,
		language.Italian:              bettinaItalianName,
		language.Japanese:             bettinaJapaneseName,
		language.Korean:               bettinaKoreanName,
		language.LatinAmericanSpanish: bettinaLatinAmericanSpanishName,
		language.Russian:              bettinaRussianName,
		language.SimplifiedChinese:    bettinaSimplifiedChineseName,
		language.Spanish:              bettinaSpanishName,
		language.TraditionalChinese:   bettinaTraditionalChineseName}
)

var (
	// bettinaCharacter represents bettina character.
	bettinaCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: bettinaBirthday,
		Code:     bettinaCode,
		Key:      character.Bettina,
		Gender:   gender.Female,
		Name:     bettinaName,
		Special:  false}
)

var (
	// bettinaAmericanEnglishPhrase represents bettina american english phrase.
	bettinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eekers"}

	// bettinaCanadianFrenchPhrase represents bettina canadian french phrase.
	bettinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "skouik-ik"}

	// bettinaDutchPhrase represents bettina dutch phrase.
	bettinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepers"}

	// bettinaFrenchPhrase represents bettina french phrase.
	bettinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "skouik-ik"}

	// bettinaGermanPhrase represents bettina german phrase.
	bettinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "achja"}

	// bettinaItalianPhrase represents bettina italian phrase.
	bettinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zink zink"}

	// bettinaJapanesePhrase represents bettina japanese phrase.
	bettinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですよね"}

	// bettinaKoreanPhrase represents bettina korean phrase.
	bettinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "내말이요"}

	// bettinaLatinAmericanSpanishPhrase represents bettina latin american spanish phrase.
	bettinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escui-cui"}

	// bettinaRussianPhrase represents bettina russian phrase.
	bettinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пи-пи-пи"}

	// bettinaSimplifiedChinesePhrase represents bettina simplified chinese phrase.
	bettinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就说吧"}

	// bettinaSpanishPhrase represents bettina spanish phrase.
	bettinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piruleta"}

	// bettinaTraditionalChinesePhrase represents bettina traditional chinese phrase.
	bettinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就說吧"}
)

var (
	// bettinaPhrase represents bettina phrase.
	bettinaPhrase = nook.Languages{
		language.AmericanEnglish:      bettinaAmericanEnglishPhrase,
		language.CanadianFrench:       bettinaCanadianFrenchPhrase,
		language.Dutch:                bettinaDutchPhrase,
		language.French:               bettinaFrenchPhrase,
		language.German:               bettinaGermanPhrase,
		language.Italian:              bettinaItalianPhrase,
		language.Japanese:             bettinaJapanesePhrase,
		language.Korean:               bettinaKoreanPhrase,
		language.LatinAmericanSpanish: bettinaLatinAmericanSpanishPhrase,
		language.Russian:              bettinaRussianPhrase,
		language.SimplifiedChinese:    bettinaSimplifiedChinesePhrase,
		language.Spanish:              bettinaSpanishPhrase,
		language.TraditionalChinese:   bettinaTraditionalChinesePhrase}
)

var (
	// Bettina represents bettina.
	Bettina = nook.Villager{
		Character:   bettinaCharacter,
		Personality: personality.Normal,
		Phrase:      bettinaPhrase}
)
