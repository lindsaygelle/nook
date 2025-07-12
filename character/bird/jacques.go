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
	// jacquesBirthday represents jacques birthday.
	jacquesBirthday = nook.Birthday{
		Day:   22,
		Month: time.June}
)

var (
	// jacquesCode represents jacques code.
	jacquesCode = nook.Code{
		Value: "brd16"}
)

var (
	// jacquesAmericanEnglishName represents jacques american english name.
	jacquesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jacques"}

	// jacquesCanadianFrenchName represents jacques canadian french name.
	jacquesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jacky"}

	// jacquesDutchName represents jacques dutch name.
	jacquesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jacques"}

	// jacquesFrenchName represents jacques french name.
	jacquesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jacky"}

	// jacquesGermanName represents jacques german name.
	jacquesGermanName = nook.Name{
		Language: language.German,
		Value:    "Pierre"}

	// jacquesItalianName represents jacques italian name.
	jacquesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampiero"}

	// jacquesJapaneseName represents jacques japanese name.
	jacquesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョッキー"}

	// jacquesKoreanName represents jacques korean name.
	jacquesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쪼끼"}

	// jacquesLatinAmericanSpanishName represents jacques latin american spanish name.
	jacquesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gorrelmo"}

	// jacquesRussianName represents jacques russian name.
	jacquesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Жак"}

	// jacquesSimplifiedChineseName represents jacques simplified chinese name.
	jacquesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "赵奇"}

	// jacquesSpanishName represents jacques spanish name.
	jacquesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gorrelmo"}

	// jacquesTraditionalChineseName represents jacques traditional chinese name.
	jacquesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "趙奇"}
)

var (
	// jacquesName represents jacques name.
	jacquesName = nook.Languages{
		language.AmericanEnglish:      jacquesAmericanEnglishName,
		language.CanadianFrench:       jacquesCanadianFrenchName,
		language.Dutch:                jacquesDutchName,
		language.French:               jacquesFrenchName,
		language.German:               jacquesGermanName,
		language.Italian:              jacquesItalianName,
		language.Japanese:             jacquesJapaneseName,
		language.Korean:               jacquesKoreanName,
		language.LatinAmericanSpanish: jacquesLatinAmericanSpanishName,
		language.Russian:              jacquesRussianName,
		language.SimplifiedChinese:    jacquesSimplifiedChineseName,
		language.Spanish:              jacquesSpanishName,
		language.TraditionalChinese:   jacquesTraditionalChineseName}
)

var (
	// jacquesCharacter represents jacques character.
	jacquesCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: jacquesBirthday,
		Code:     jacquesCode,
		Key:      character.Jacques,
		Gender:   gender.Male,
		Name:     jacquesName,
		Special:  false}
)

var (
	// jacquesAmericanEnglishPhrase represents jacques american english phrase.
	jacquesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zut alors"}

	// jacquesCanadianFrenchPhrase represents jacques canadian french phrase.
	jacquesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "why not"}

	// jacquesDutchPhrase represents jacques dutch phrase.
	jacquesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zut alors"}

	// jacquesFrenchPhrase represents jacques french phrase.
	jacquesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ploc ploc"}

	// jacquesGermanPhrase represents jacques german phrase.
	jacquesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piepiep"}

	// jacquesItalianPhrase represents jacques italian phrase.
	jacquesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pio pio"}

	// jacquesJapanesePhrase represents jacques japanese phrase.
	jacquesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チェケラ"}

	// jacquesKoreanPhrase represents jacques korean phrase.
	jacquesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "체키라웃"}

	// jacquesLatinAmericanSpanishPhrase represents jacques latin american spanish phrase.
	jacquesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chapó"}

	// jacquesRussianPhrase represents jacques russian phrase.
	jacquesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фу ты"}

	// jacquesSimplifiedChinesePhrase represents jacques simplified chinese phrase.
	jacquesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷唷唷"}

	// jacquesSpanishPhrase represents jacques spanish phrase.
	jacquesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chapó"}

	// jacquesTraditionalChinesePhrase represents jacques traditional chinese phrase.
	jacquesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷唷唷"}
)

var (
	// jacquesPhrase represents jacques phrase.
	jacquesPhrase = nook.Languages{
		language.AmericanEnglish:      jacquesAmericanEnglishPhrase,
		language.CanadianFrench:       jacquesCanadianFrenchPhrase,
		language.Dutch:                jacquesDutchPhrase,
		language.French:               jacquesFrenchPhrase,
		language.German:               jacquesGermanPhrase,
		language.Italian:              jacquesItalianPhrase,
		language.Japanese:             jacquesJapanesePhrase,
		language.Korean:               jacquesKoreanPhrase,
		language.LatinAmericanSpanish: jacquesLatinAmericanSpanishPhrase,
		language.Russian:              jacquesRussianPhrase,
		language.SimplifiedChinese:    jacquesSimplifiedChinesePhrase,
		language.Spanish:              jacquesSpanishPhrase,
		language.TraditionalChinese:   jacquesTraditionalChinesePhrase}
)

var (
	// Jacques represents jacques.
	Jacques = nook.Villager{
		Character:   jacquesCharacter,
		Personality: personality.Smug,
		Phrase:      jacquesPhrase}
)
