package hamster

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
	// appleBirthday represents apple birthday.
	appleBirthday = nook.Birthday{
		Day:   24,
		Month: time.September}
)

var (
	// appleCode represents apple code.
	appleCode = nook.Code{
		Value: "ham01"}
)

var (
	// appleAmericanEnglishName represents apple american english name.
	appleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Apple"}

	// appleCanadianFrenchName represents apple canadian french name.
	appleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Esther"}

	// appleDutchName represents apple dutch name.
	appleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Apple"}

	// appleFrenchName represents apple french name.
	appleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Esther"}

	// appleGermanName represents apple german name.
	appleGermanName = nook.Name{
		Language: language.German,
		Value:    "Jessi"}

	// appleItalianName represents apple italian name.
	appleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicci"}

	// appleJapaneseName represents apple japanese name.
	appleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アップル"}

	// appleKoreanName represents apple korean name.
	appleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "애플"}

	// appleLatinAmericanSpanishName represents apple latin american spanish name.
	appleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosi"}

	// appleRussianName represents apple russian name.
	appleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эпл"}

	// appleSimplifiedChineseName represents apple simplified chinese name.
	appleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "苹果"}

	// appleSpanishName represents apple spanish name.
	appleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosi"}

	// appleTraditionalChineseName represents apple traditional chinese name.
	appleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蘋果"}
)

var (
	// appleName represents apple name.
	appleName = nook.Languages{
		language.AmericanEnglish:      appleAmericanEnglishName,
		language.CanadianFrench:       appleCanadianFrenchName,
		language.Dutch:                appleDutchName,
		language.French:               appleFrenchName,
		language.German:               appleGermanName,
		language.Italian:              appleItalianName,
		language.Japanese:             appleJapaneseName,
		language.Korean:               appleKoreanName,
		language.LatinAmericanSpanish: appleLatinAmericanSpanishName,
		language.Russian:              appleRussianName,
		language.SimplifiedChinese:    appleSimplifiedChineseName,
		language.Spanish:              appleSpanishName,
		language.TraditionalChinese:   appleTraditionalChineseName}
)

var (
	// appleCharacter represents apple character.
	appleCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: appleBirthday,
		Code:     appleCode,
		Key:      character.Apple,
		Gender:   gender.Female,
		Name:     appleName,
		Special:  false}
)

var (
	// appleAmericanEnglishPhrase represents apple american english phrase.
	appleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheekers"}

	// appleCanadianFrenchPhrase represents apple canadian french phrase.
	appleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bibille"}

	// appleDutchPhrase represents apple dutch phrase.
	appleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wangetjes"}

	// appleFrenchPhrase represents apple french phrase.
	appleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bibille"}

	// appleGermanPhrase represents apple german phrase.
	appleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fiep"}

	// appleItalianPhrase represents apple italian phrase.
	appleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "triplo uau"}

	// appleJapanesePhrase represents apple japanese phrase.
	appleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キュルン"}

	// appleKoreanPhrase represents apple korean phrase.
	appleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "큐룽"}

	// appleLatinAmericanSpanishPhrase represents apple latin american spanish phrase.
	appleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "do-re-mi"}

	// appleRussianPhrase represents apple russian phrase.
	appleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щечки"}

	// appleSimplifiedChinesePhrase represents apple simplified chinese phrase.
	appleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "转转"}

	// appleSpanishPhrase represents apple spanish phrase.
	appleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "do-re-mi"}

	// appleTraditionalChinesePhrase represents apple traditional chinese phrase.
	appleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "轉轉"}
)

var (
	// applePhrase represents apple phrase.
	applePhrase = nook.Languages{
		language.AmericanEnglish:      appleAmericanEnglishPhrase,
		language.CanadianFrench:       appleCanadianFrenchPhrase,
		language.Dutch:                appleDutchPhrase,
		language.French:               appleFrenchPhrase,
		language.German:               appleGermanPhrase,
		language.Italian:              appleItalianPhrase,
		language.Japanese:             appleJapanesePhrase,
		language.Korean:               appleKoreanPhrase,
		language.LatinAmericanSpanish: appleLatinAmericanSpanishPhrase,
		language.Russian:              appleRussianPhrase,
		language.SimplifiedChinese:    appleSimplifiedChinesePhrase,
		language.Spanish:              appleSpanishPhrase,
		language.TraditionalChinese:   appleTraditionalChinesePhrase}
)

var (
	// Apple represents apple.
	Apple = nook.Villager{
		Character:   appleCharacter,
		Personality: personality.Peppy,
		Phrase:      applePhrase}
)
