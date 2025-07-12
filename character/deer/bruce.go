package deer

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
	// bruceBirthday represents bruce birthday.
	bruceBirthday = nook.Birthday{
		Day:   26,
		Month: time.May}
)

var (
	// bruceCode represents bruce code.
	bruceCode = nook.Code{
		Value: "der03"}
)

var (
	// bruceAmericanEnglishName represents bruce american english name.
	bruceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bruce"}

	// bruceCanadianFrenchName represents bruce canadian french name.
	bruceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Boubou"}

	// bruceDutchName represents bruce dutch name.
	bruceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bruce"}

	// bruceFrenchName represents bruce french name.
	bruceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Boubou"}

	// bruceGermanName represents bruce german name.
	bruceGermanName = nook.Name{
		Language: language.German,
		Value:    "Oswald"}

	// bruceItalianName represents bruce italian name.
	bruceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Arturo"}

	// bruceJapaneseName represents bruce japanese name.
	bruceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブルース"}

	// bruceKoreanName represents bruce korean name.
	bruceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브루스"}

	// bruceLatinAmericanSpanishName represents bruce latin american spanish name.
	bruceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aristo"}

	// bruceRussianName represents bruce russian name.
	bruceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брюс"}

	// bruceSimplifiedChineseName represents bruce simplified chinese name.
	bruceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布鹿斯"}

	// bruceSpanishName represents bruce spanish name.
	bruceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aristo"}

	// bruceTraditionalChineseName represents bruce traditional chinese name.
	bruceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布鹿斯"}
)

var (
	// bruceName represents bruce name.
	bruceName = nook.Languages{
		language.AmericanEnglish:      bruceAmericanEnglishName,
		language.CanadianFrench:       bruceCanadianFrenchName,
		language.Dutch:                bruceDutchName,
		language.French:               bruceFrenchName,
		language.German:               bruceGermanName,
		language.Italian:              bruceItalianName,
		language.Japanese:             bruceJapaneseName,
		language.Korean:               bruceKoreanName,
		language.LatinAmericanSpanish: bruceLatinAmericanSpanishName,
		language.Russian:              bruceRussianName,
		language.SimplifiedChinese:    bruceSimplifiedChineseName,
		language.Spanish:              bruceSpanishName,
		language.TraditionalChinese:   bruceTraditionalChineseName}
)

var (
	// bruceCharacter represents bruce character.
	bruceCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: bruceBirthday,
		Code:     bruceCode,
		Key:      character.Bruce,
		Gender:   gender.Male,
		Name:     bruceName,
		Special:  false}
)

var (
	// bruceAmericanEnglishPhrase represents bruce american english phrase.
	bruceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gruff"}

	// bruceCanadianFrenchPhrase represents bruce canadian french phrase.
	bruceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouain"}

	// bruceDutchPhrase represents bruce dutch phrase.
	bruceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "burl"}

	// bruceFrenchPhrase represents bruce french phrase.
	bruceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nunul"}

	// bruceGermanPhrase represents bruce german phrase.
	bruceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "röööhr"}

	// bruceItalianPhrase represents bruce italian phrase.
	bruceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uauwaii"}

	// bruceJapanesePhrase represents bruce japanese phrase.
	bruceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "しかしな"}

	// bruceKoreanPhrase represents bruce korean phrase.
	bruceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아비옹"}

	// bruceLatinAmericanSpanishPhrase represents bruce latin american spanish phrase.
	bruceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aquesí"}

	// bruceRussianPhrase represents bruce russian phrase.
	bruceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "забияка"}

	// bruceSimplifiedChinesePhrase represents bruce simplified chinese phrase.
	bruceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "福鹿寿"}

	// bruceSpanishPhrase represents bruce spanish phrase.
	bruceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aquesí"}

	// bruceTraditionalChinesePhrase represents bruce traditional chinese phrase.
	bruceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "福鹿壽"}
)

var (
	// brucePhrase represents bruce phrase.
	brucePhrase = nook.Languages{
		language.AmericanEnglish:      bruceAmericanEnglishPhrase,
		language.CanadianFrench:       bruceCanadianFrenchPhrase,
		language.Dutch:                bruceDutchPhrase,
		language.French:               bruceFrenchPhrase,
		language.German:               bruceGermanPhrase,
		language.Italian:              bruceItalianPhrase,
		language.Japanese:             bruceJapanesePhrase,
		language.Korean:               bruceKoreanPhrase,
		language.LatinAmericanSpanish: bruceLatinAmericanSpanishPhrase,
		language.Russian:              bruceRussianPhrase,
		language.SimplifiedChinese:    bruceSimplifiedChinesePhrase,
		language.Spanish:              bruceSpanishPhrase,
		language.TraditionalChinese:   bruceTraditionalChinesePhrase}
)

var (
	// Bruce represents bruce.
	Bruce = nook.Villager{
		Character:   bruceCharacter,
		Personality: personality.Cranky,
		Phrase:      brucePhrase}
)
