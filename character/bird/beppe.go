package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	beppeBirthday = nook.Birthday{
		Day:   18,
		Month: time.July}
)

var (
	beppeCode = nook.Code{
		Value: "cwc"}
)

var (
	beppeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beppe"}

	beppeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Choucac"}

	beppeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	beppeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Choucac"}

	beppeGermanName = nook.Name{
		Language: language.German,
		Value:    "Schorsch"}

	beppeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amper"}

	beppeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピンちゃん"}

	beppeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑군"}

	beppeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Grajardo"}

	beppeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	beppeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	beppeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Grajardo"}

	beppeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "車仔"}
)

var (
	beppeName = nook.Languages{
		language.AmericanEnglish:      beppeAmericanEnglishName,
		language.CanadianFrench:       beppeCanadianFrenchName,
		language.Dutch:                beppeDutchName,
		language.French:               beppeFrenchName,
		language.German:               beppeGermanName,
		language.Italian:              beppeItalianName,
		language.Japanese:             beppeJapaneseName,
		language.Korean:               beppeKoreanName,
		language.LatinAmericanSpanish: beppeLatinAmericanSpanishName,
		language.Russian:              beppeRussianName,
		language.SimplifiedChinese:    beppeSimplifiedChineseName,
		language.Spanish:              beppeSpanishName,
		language.TraditionalChinese:   beppeTraditionalChineseName}
)

var (
	beppeCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: beppeBirthday,
		Code:     beppeCode,
		Key:      character.Beppe,
		Gender:   gender.Male,
		Name:     beppeName,
		Special:  true}
)

var (
	Beppe = nook.Resident{
		Character: beppeCharacter}
)
