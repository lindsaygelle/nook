package pig

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
	// pancettiBirthday represents pancetti birthday.
	pancettiBirthday = nook.Birthday{
		Day:   14,
		Month: time.November}
)

var (
	// pancettiCode represents pancetti code.
	pancettiCode = nook.Code{
		Value: "pig16"}
)

var (
	// pancettiAmericanEnglishName represents pancetti american english name.
	pancettiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pancetti"}

	// pancettiCanadianFrenchName represents pancetti canadian french name.
	pancettiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lydie"}

	// pancettiDutchName represents pancetti dutch name.
	pancettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pancetti"}

	// pancettiFrenchName represents pancetti french name.
	pancettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lydie"}

	// pancettiGermanName represents pancetti german name.
	pancettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Brigitte"}

	// pancettiItalianName represents pancetti italian name.
	pancettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicciola"}

	// pancettiJapaneseName represents pancetti japanese name.
	pancettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブリトニー"}

	// pancettiKoreanName represents pancetti korean name.
	pancettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브리트니"}

	// pancettiLatinAmericanSpanishName represents pancetti latin american spanish name.
	pancettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Talía"}

	// pancettiRussianName represents pancetti russian name.
	pancettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панчетти"}

	// pancettiSimplifiedChineseName represents pancetti simplified chinese name.
	pancettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布兰妮"}

	// pancettiSpanishName represents pancetti spanish name.
	pancettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Talía"}

	// pancettiTraditionalChineseName represents pancetti traditional chinese name.
	pancettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布蘭妮"}
)

var (
	// pancettiName represents pancetti name.
	pancettiName = nook.Languages{
		language.AmericanEnglish:      pancettiAmericanEnglishName,
		language.CanadianFrench:       pancettiCanadianFrenchName,
		language.Dutch:                pancettiDutchName,
		language.French:               pancettiFrenchName,
		language.German:               pancettiGermanName,
		language.Italian:              pancettiItalianName,
		language.Japanese:             pancettiJapaneseName,
		language.Korean:               pancettiKoreanName,
		language.LatinAmericanSpanish: pancettiLatinAmericanSpanishName,
		language.Russian:              pancettiRussianName,
		language.SimplifiedChinese:    pancettiSimplifiedChineseName,
		language.Spanish:              pancettiSpanishName,
		language.TraditionalChinese:   pancettiTraditionalChineseName}
)

var (
	// pancettiCharacter represents pancetti character.
	pancettiCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: pancettiBirthday,
		Code:     pancettiCode,
		Key:      character.Pancetti,
		Gender:   gender.Female,
		Name:     pancettiName,
		Special:  false}
)

var (
	// pancettiAmericanEnglishPhrase represents pancetti american english phrase.
	pancettiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sooey"}

	// pancettiCanadianFrenchPhrase represents pancetti canadian french phrase.
	pancettiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sabots"}

	// pancettiDutchPhrase represents pancetti dutch phrase.
	pancettiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kom kom"}

	// pancettiFrenchPhrase represents pancetti french phrase.
	pancettiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sabots"}

	// pancettiGermanPhrase represents pancetti german phrase.
	pancettiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grubbel"}

	// pancettiItalianPhrase represents pancetti italian phrase.
	pancettiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grufidù"}

	// pancettiJapanesePhrase represents pancetti japanese phrase.
	pancettiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やだも～"}

	// pancettiKoreanPhrase represents pancetti korean phrase.
	pancettiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어우야"}

	// pancettiLatinAmericanSpanishPhrase represents pancetti latin american spanish phrase.
	pancettiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuinoink"}

	// pancettiRussianPhrase represents pancetti russian phrase.
	pancettiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюшка"}

	// pancettiSimplifiedChinesePhrase represents pancetti simplified chinese phrase.
	pancettiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "讨厌啦"}

	// pancettiSpanishPhrase represents pancetti spanish phrase.
	pancettiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "morrito"}

	// pancettiTraditionalChinesePhrase represents pancetti traditional chinese phrase.
	pancettiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "討厭啦"}
)

var (
	// pancettiPhrase represents pancetti phrase.
	pancettiPhrase = nook.Languages{
		language.AmericanEnglish:      pancettiAmericanEnglishPhrase,
		language.CanadianFrench:       pancettiCanadianFrenchPhrase,
		language.Dutch:                pancettiDutchPhrase,
		language.French:               pancettiFrenchPhrase,
		language.German:               pancettiGermanPhrase,
		language.Italian:              pancettiItalianPhrase,
		language.Japanese:             pancettiJapanesePhrase,
		language.Korean:               pancettiKoreanPhrase,
		language.LatinAmericanSpanish: pancettiLatinAmericanSpanishPhrase,
		language.Russian:              pancettiRussianPhrase,
		language.SimplifiedChinese:    pancettiSimplifiedChinesePhrase,
		language.Spanish:              pancettiSpanishPhrase,
		language.TraditionalChinese:   pancettiTraditionalChinesePhrase}
)

var (
	// Pancetti represents pancetti.
	Pancetti = nook.Villager{
		Character:   pancettiCharacter,
		Personality: personality.Snooty,
		Phrase:      pancettiPhrase}
)
