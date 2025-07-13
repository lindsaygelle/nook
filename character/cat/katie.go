package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// katieBirthday represents katie birthday.
	katieBirthday = nook.Birthday{
		Day:   22,
		Month: time.October}
)

var (
	// katieCode represents katie code.
	katieCode = nook.Code{
		Value: "los/lom"}
)

var (
	// katieAmericanEnglishName represents katie american english name.
	katieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Katie"}

	// katieCanadianFrenchName represents katie canadian french name.
	katieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cathie"}

	// katieDutchName represents katie dutch name.
	katieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Katie"}

	// katieFrenchName represents katie french name.
	katieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cathie"}

	// katieGermanName represents katie german name.
	katieGermanName = nook.Name{
		Language: language.German,
		Value:    "Katja"}

	// katieItalianName represents katie italian name.
	katieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Micetta"}

	// katieJapaneseName represents katie japanese name.
	katieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まいこちゃん"}

	// katieKoreanName represents katie korean name.
	katieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미야"}

	// katieLatinAmericanSpanishName represents katie latin american spanish name.
	katieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cati"}

	// katieRussianName represents katie russian name.
	katieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кейти"}

	// katieSimplifiedChineseName represents katie simplified chinese name.
	katieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咪露"}

	// katieSpanishName represents katie spanish name.
	katieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cati"}

	// katieTraditionalChineseName represents katie traditional chinese name.
	katieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咪露"}
)

var (
	// katieName represents katie name.
	katieName = nook.Languages{
		language.AmericanEnglish:      katieAmericanEnglishName,
		language.CanadianFrench:       katieCanadianFrenchName,
		language.Dutch:                katieDutchName,
		language.French:               katieFrenchName,
		language.German:               katieGermanName,
		language.Italian:              katieItalianName,
		language.Japanese:             katieJapaneseName,
		language.Korean:               katieKoreanName,
		language.LatinAmericanSpanish: katieLatinAmericanSpanishName,
		language.Russian:              katieRussianName,
		language.SimplifiedChinese:    katieSimplifiedChineseName,
		language.Spanish:              katieSpanishName,
		language.TraditionalChinese:   katieTraditionalChineseName}
)

var (
	// katieCharacter represents katie character.
	katieCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: katieBirthday,
		Code:     katieCode,
		Key:      character.Katie,
		Gender:   gender.Female,
		Name:     katieName,
		Special:  true}
)

var (
	// Katie represents katie.
	Katie = nook.Resident{
		Character: katieCharacter}
)
