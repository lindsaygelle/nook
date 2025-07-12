package dog

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
	// butchBirthday represents butch birthday.
	butchBirthday = nook.Birthday{
		Day:   1,
		Month: time.November}
)

var (
	// butchCode represents butch code.
	butchCode = nook.Code{
		Value: "dog01"}
)

var (
	// butchAmericanEnglishName represents butch american english name.
	butchAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Butch"}

	// butchCanadianFrenchName represents butch canadian french name.
	butchCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Avril"}

	// butchDutchName represents butch dutch name.
	butchDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Butch"}

	// butchFrenchName represents butch french name.
	butchFrenchName = nook.Name{
		Language: language.French,
		Value:    "Avril"}

	// butchGermanName represents butch german name.
	butchGermanName = nook.Name{
		Language: language.German,
		Value:    "Hasso"}

	// butchItalianName represents butch italian name.
	butchItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Attila"}

	// butchJapaneseName represents butch japanese name.
	butchJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョン"}

	// butchKoreanName represents butch korean name.
	butchKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "존"}

	// butchLatinAmericanSpanishName represents butch latin american spanish name.
	butchLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bruno"}

	// butchRussianName represents butch russian name.
	butchRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бутч"}

	// butchSimplifiedChineseName represents butch simplified chinese name.
	butchSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "约翰"}

	// butchSpanishName represents butch spanish name.
	butchSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bruno"}

	// butchTraditionalChineseName represents butch traditional chinese name.
	butchTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "約翰"}
)

var (
	// butchName represents butch name.
	butchName = nook.Languages{
		language.AmericanEnglish:      butchAmericanEnglishName,
		language.CanadianFrench:       butchCanadianFrenchName,
		language.Dutch:                butchDutchName,
		language.French:               butchFrenchName,
		language.German:               butchGermanName,
		language.Italian:              butchItalianName,
		language.Japanese:             butchJapaneseName,
		language.Korean:               butchKoreanName,
		language.LatinAmericanSpanish: butchLatinAmericanSpanishName,
		language.Russian:              butchRussianName,
		language.SimplifiedChinese:    butchSimplifiedChineseName,
		language.Spanish:              butchSpanishName,
		language.TraditionalChinese:   butchTraditionalChineseName}
)

var (
	// butchCharacter represents butch character.
	butchCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: butchBirthday,
		Code:     butchCode,
		Key:      character.Butch,
		Gender:   gender.Male,
		Name:     butchName,
		Special:  false}
)

var (
	// butchAmericanEnglishPhrase represents butch american english phrase.
	butchAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ROOOOOWF"}

	// butchCanadianFrenchPhrase represents butch canadian french phrase.
	butchCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "WROOOOUF"}

	// butchDutchPhrase represents butch dutch phrase.
	butchDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "BROEEEF"}

	// butchFrenchPhrase represents butch french phrase.
	butchFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "WROOOOUF"}

	// butchGermanPhrase represents butch german phrase.
	butchGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "GRRUFF"}

	// butchItalianPhrase represents butch italian phrase.
	butchItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ROOOOOF"}

	// butchJapanesePhrase represents butch japanese phrase.
	butchJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ノン"}

	// butchKoreanPhrase represents butch korean phrase.
	butchKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아니"}

	// butchLatinAmericanSpanishPhrase represents butch latin american spanish phrase.
	butchLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruf-gruf"}

	// butchRussianPhrase represents butch russian phrase.
	butchRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ррр-гав"}

	// butchSimplifiedChinesePhrase represents butch simplified chinese phrase.
	butchSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "侬"}

	// butchSpanishPhrase represents butch spanish phrase.
	butchSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruf-gruf"}

	// butchTraditionalChinesePhrase represents butch traditional chinese phrase.
	butchTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "儂"}
)

var (
	// butchPhrase represents butch phrase.
	butchPhrase = nook.Languages{
		language.AmericanEnglish:      butchAmericanEnglishPhrase,
		language.CanadianFrench:       butchCanadianFrenchPhrase,
		language.Dutch:                butchDutchPhrase,
		language.French:               butchFrenchPhrase,
		language.German:               butchGermanPhrase,
		language.Italian:              butchItalianPhrase,
		language.Japanese:             butchJapanesePhrase,
		language.Korean:               butchKoreanPhrase,
		language.LatinAmericanSpanish: butchLatinAmericanSpanishPhrase,
		language.Russian:              butchRussianPhrase,
		language.SimplifiedChinese:    butchSimplifiedChinesePhrase,
		language.Spanish:              butchSpanishPhrase,
		language.TraditionalChinese:   butchTraditionalChinesePhrase}
)

var (
	// Butch represents butch.
	Butch = nook.Villager{
		Character:   butchCharacter,
		Personality: personality.Cranky,
		Phrase:      butchPhrase}
)
