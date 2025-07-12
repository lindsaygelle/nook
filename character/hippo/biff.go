package hippo

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
	// biffBirthday represents biff birthday.
	biffBirthday = nook.Birthday{
		Day:   29,
		Month: time.March}
)

var (
	// biffCode represents biff code.
	biffCode = nook.Code{
		Value: "hip04"}
)

var (
	// biffAmericanEnglishName represents biff american english name.
	biffAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Biff"}

	// biffCanadianFrenchName represents biff canadian french name.
	biffCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Biff"}

	// biffDutchName represents biff dutch name.
	biffDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Biff"}

	// biffFrenchName represents biff french name.
	biffFrenchName = nook.Name{
		Language: language.French,
		Value:    "Biff"}

	// biffGermanName represents biff german name.
	biffGermanName = nook.Name{
		Language: language.German,
		Value:    "Norbert"}

	// biffItalianName represents biff italian name.
	biffItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pippo"}

	// biffJapaneseName represents biff japanese name.
	biffJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガブリエル"}

	// biffKoreanName represents biff korean name.
	biffKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가브리엘"}

	// biffLatinAmericanSpanishName represents biff latin american spanish name.
	biffLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pipo"}

	// biffRussianName represents biff russian name.
	biffRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бифф"}

	// biffSimplifiedChineseName represents biff simplified chinese name.
	biffSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾宝为"}

	// biffSpanishName represents biff spanish name.
	biffSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pipo"}

	// biffTraditionalChineseName represents biff traditional chinese name.
	biffTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈寶為"}
)

var (
	// biffName represents biff name.
	biffName = nook.Languages{
		language.AmericanEnglish:      biffAmericanEnglishName,
		language.CanadianFrench:       biffCanadianFrenchName,
		language.Dutch:                biffDutchName,
		language.French:               biffFrenchName,
		language.German:               biffGermanName,
		language.Italian:              biffItalianName,
		language.Japanese:             biffJapaneseName,
		language.Korean:               biffKoreanName,
		language.LatinAmericanSpanish: biffLatinAmericanSpanishName,
		language.Russian:              biffRussianName,
		language.SimplifiedChinese:    biffSimplifiedChineseName,
		language.Spanish:              biffSpanishName,
		language.TraditionalChinese:   biffTraditionalChineseName}
)

var (
	// biffCharacter represents biff character.
	biffCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: biffBirthday,
		Code:     biffCode,
		Key:      character.Biff,
		Gender:   gender.Male,
		Name:     biffName,
		Special:  false}
)

var (
	// biffAmericanEnglishPhrase represents biff american english phrase.
	biffAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squirt"}

	// biffCanadianFrenchPhrase represents biff canadian french phrase.
	biffCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hips hips"}

	// biffDutchPhrase represents biff dutch phrase.
	biffDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "spetter"}

	// biffFrenchPhrase represents biff french phrase.
	biffFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hips hips"}

	// biffGermanPhrase represents biff german phrase.
	biffGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "skwiirt"}

	// biffItalianPhrase represents biff italian phrase.
	biffItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squirt"}

	// biffJapanesePhrase represents biff japanese phrase.
	biffJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃけん"}

	// biffKoreanPhrase represents biff korean phrase.
	biffKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오도독"}

	// biffLatinAmericanSpanishPhrase represents biff latin american spanish phrase.
	biffLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hipo-pipo"}

	// biffRussianPhrase represents biff russian phrase.
	biffRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малек"}

	// biffSimplifiedChinesePhrase represents biff simplified chinese phrase.
	biffSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "因为"}

	// biffSpanishPhrase represents biff spanish phrase.
	biffSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hipo-pipo"}

	// biffTraditionalChinesePhrase represents biff traditional chinese phrase.
	biffTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "因為"}
)

var (
	// biffPhrase represents biff phrase.
	biffPhrase = nook.Languages{
		language.AmericanEnglish:      biffAmericanEnglishPhrase,
		language.CanadianFrench:       biffCanadianFrenchPhrase,
		language.Dutch:                biffDutchPhrase,
		language.French:               biffFrenchPhrase,
		language.German:               biffGermanPhrase,
		language.Italian:              biffItalianPhrase,
		language.Japanese:             biffJapanesePhrase,
		language.Korean:               biffKoreanPhrase,
		language.LatinAmericanSpanish: biffLatinAmericanSpanishPhrase,
		language.Russian:              biffRussianPhrase,
		language.SimplifiedChinese:    biffSimplifiedChinesePhrase,
		language.Spanish:              biffSpanishPhrase,
		language.TraditionalChinese:   biffTraditionalChinesePhrase}
)

var (
	// Biff represents biff.
	Biff = nook.Villager{
		Character:   biffCharacter,
		Personality: personality.Jock,
		Phrase:      biffPhrase}
)
