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
	// medliBirthday represents medli birthday.
	medliBirthday = nook.Birthday{
		Day:   13,
		Month: time.December}
)

var (
	// medliCode represents medli code.
	medliCode = nook.Code{
		Value: "brd19"}
)

var (
	// medliAmericanEnglishName represents medli american english name.
	medliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Medli"}

	// medliCanadianFrenchName represents medli canadian french name.
	medliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Médolie"}

	// medliDutchName represents medli dutch name.
	medliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// medliFrenchName represents medli french name.
	medliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Médolie"}

	// medliGermanName represents medli german name.
	medliGermanName = nook.Name{
		Language: language.German,
		Value:    "Medolie"}

	// medliItalianName represents medli italian name.
	medliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Famirè"}

	// medliJapaneseName represents medli japanese name.
	medliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メドリ"}

	// medliKoreanName represents medli korean name.
	medliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메들리"}

	// medliLatinAmericanSpanishName represents medli latin american spanish name.
	medliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Medli"}

	// medliRussianName represents medli russian name.
	medliRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// medliSimplifiedChineseName represents medli simplified chinese name.
	medliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// medliSpanishName represents medli spanish name.
	medliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Medli"}

	// medliTraditionalChineseName represents medli traditional chinese name.
	medliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// medliName represents medli name.
	medliName = nook.Languages{
		language.AmericanEnglish:      medliAmericanEnglishName,
		language.CanadianFrench:       medliCanadianFrenchName,
		language.Dutch:                medliDutchName,
		language.French:               medliFrenchName,
		language.German:               medliGermanName,
		language.Italian:              medliItalianName,
		language.Japanese:             medliJapaneseName,
		language.Korean:               medliKoreanName,
		language.LatinAmericanSpanish: medliLatinAmericanSpanishName,
		language.Russian:              medliRussianName,
		language.SimplifiedChinese:    medliSimplifiedChineseName,
		language.Spanish:              medliSpanishName,
		language.TraditionalChinese:   medliTraditionalChineseName}
)

var (
	// medliCharacter represents medli character.
	medliCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: medliBirthday,
		Code:     medliCode,
		Key:      character.Medli,
		Gender:   gender.Female,
		Name:     medliName,
		Special:  false}
)

var (
	// medliAmericanEnglishPhrase represents medli american english phrase.
	medliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gimme"}

	// medliCanadianFrenchPhrase represents medli canadian french phrase.
	medliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lalalyre"}

	// medliDutchPhrase represents medli dutch phrase.
	medliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// medliFrenchPhrase represents medli french phrase.
	medliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lalalyre"}

	// medliGermanPhrase represents medli german phrase.
	medliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hilfhilf"}

	// medliItalianPhrase represents medli italian phrase.
	medliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "plinplon"}

	// medliJapanesePhrase represents medli japanese phrase.
	medliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "くらさい"}

	// medliKoreanPhrase represents medli korean phrase.
	medliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "핑그르르"}

	// medliLatinAmericanSpanishPhrase represents medli latin american spanish phrase.
	medliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plimplín"}

	// medliRussianPhrase represents medli russian phrase.
	medliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// medliSimplifiedChinesePhrase represents medli simplified chinese phrase.
	medliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// medliSpanishPhrase represents medli spanish phrase.
	medliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plimplín"}

	// medliTraditionalChinesePhrase represents medli traditional chinese phrase.
	medliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// medliPhrase represents medli phrase.
	medliPhrase = nook.Languages{
		language.AmericanEnglish:      medliAmericanEnglishPhrase,
		language.CanadianFrench:       medliCanadianFrenchPhrase,
		language.Dutch:                medliDutchPhrase,
		language.French:               medliFrenchPhrase,
		language.German:               medliGermanPhrase,
		language.Italian:              medliItalianPhrase,
		language.Japanese:             medliJapanesePhrase,
		language.Korean:               medliKoreanPhrase,
		language.LatinAmericanSpanish: medliLatinAmericanSpanishPhrase,
		language.Russian:              medliRussianPhrase,
		language.SimplifiedChinese:    medliSimplifiedChinesePhrase,
		language.Spanish:              medliSpanishPhrase,
		language.TraditionalChinese:   medliTraditionalChinesePhrase}
)

var (
	// Medli represents medli.
	Medli = nook.Villager{
		Character:   medliCharacter,
		Personality: personality.Normal,
		Phrase:      medliPhrase}
)
