package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	appleBirthday = nook.Birthday{
		Day:   24,
		Month: time.September}
)

var (
	appleCode = nook.Code{
		Value: "ham01"}
)

var (
	appleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Apple"}

	appleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Esther"}

	appleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Apple"}

	appleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Esther"}

	appleGermanName = nook.Name{
		Language: language.German,
		Value:    "Jessi"}

	appleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicci"}

	appleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アップル"}

	appleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "애플"}

	appleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosi"}

	appleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эпл"}

	appleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "苹果"}

	appleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosi"}

	appleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蘋果"}
)

var (
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
	appleCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: appleBirthday,
		Code:     appleCode,
		Gender:   gender.Female,
		Name:     appleName}
)

var (
	appleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheekers"}

	appleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bibille"}

	appleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wangetjes"}

	appleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bibille"}

	appleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fiep"}

	appleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "triplo uau"}

	appleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キュルン"}

	appleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "큐룽"}

	appleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "do-re-mi"}

	appleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щечки"}

	appleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "转转"}

	appleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "do-re-mi"}

	appleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "轉轉"}
)

var (
	applePhrase = nook.Languages{
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
	Apple = nook.Villager{
		Character:   appleCharacter,
		Personality: personality.Peppy,
		Phrase:      applePhrase}
)
