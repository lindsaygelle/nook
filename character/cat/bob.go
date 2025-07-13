package cat

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
	// bobBirthday represents bob birthday.
	bobBirthday = nook.Birthday{
		Day:   1,
		Month: time.January}
)

var (
	// bobCode represents bob code.
	bobCode = nook.Code{
		Value: "cat00"}
)

var (
	// bobAmericanEnglishName represents bob american english name.
	bobAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bob"}

	// bobCanadianFrenchName represents bob canadian french name.
	bobCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Robert"}

	// bobDutchName represents bob dutch name.
	bobDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bob"}

	// bobFrenchName represents bob french name.
	bobFrenchName = nook.Name{
		Language: language.French,
		Value:    "Robert"}

	// bobGermanName represents bob german name.
	bobGermanName = nook.Name{
		Language: language.German,
		Value:    "Jens"}

	// bobItalianName represents bob italian name.
	bobItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bob"}

	// bobJapaneseName represents bob japanese name.
	bobJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニコバン"}

	// bobKoreanName represents bob korean name.
	bobKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "히죽"}

	// bobLatinAmericanSpanishName represents bob latin american spanish name.
	bobLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arándano"}

	// bobRussianName represents bob russian name.
	bobRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Боб"}

	// bobSimplifiedChineseName represents bob simplified chinese name.
	bobSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿判"}

	// bobSpanishName represents bob spanish name.
	bobSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arándano"}

	// bobTraditionalChineseName represents bob traditional chinese name.
	bobTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿判"}
)

var (
	// bobName represents bob name.
	bobName = nook.Languages{
		language.AmericanEnglish:      bobAmericanEnglishName,
		language.CanadianFrench:       bobCanadianFrenchName,
		language.Dutch:                bobDutchName,
		language.French:               bobFrenchName,
		language.German:               bobGermanName,
		language.Italian:              bobItalianName,
		language.Japanese:             bobJapaneseName,
		language.Korean:               bobKoreanName,
		language.LatinAmericanSpanish: bobLatinAmericanSpanishName,
		language.Russian:              bobRussianName,
		language.SimplifiedChinese:    bobSimplifiedChineseName,
		language.Spanish:              bobSpanishName,
		language.TraditionalChinese:   bobTraditionalChineseName}
)

var (
	// bobCharacter represents bob character.
	bobCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: bobBirthday,
		Code:     bobCode,
		Key:      character.Bob,
		Gender:   gender.Male,
		Name:     bobName,
		Special:  false}
)

var (
	// bobAmericanEnglishPhrase represents bob american english phrase.
	bobAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pthhpth"}

	// bobCanadianFrenchPhrase represents bob canadian french phrase.
	bobCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pfioupf"}

	// bobDutchPhrase represents bob dutch phrase.
	bobDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poekie"}

	// bobFrenchPhrase represents bob french phrase.
	bobFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pfioupf"}

	// bobGermanPhrase represents bob german phrase.
	bobGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "psssps"}

	// bobItalianPhrase represents bob italian phrase.
	bobItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baffo"}

	// bobJapanesePhrase represents bob japanese phrase.
	bobJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ネコ"}

	// bobKoreanPhrase represents bob korean phrase.
	bobKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "고양이"}

	// bobLatinAmericanSpanishPhrase represents bob latin american spanish phrase.
	bobLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fishfish"}

	// bobRussianPhrase represents bob russian phrase.
	bobRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кх-кх-кх"}

	// bobSimplifiedChinesePhrase represents bob simplified chinese phrase.
	bobSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "猫猫"}

	// bobSpanishPhrase represents bob spanish phrase.
	bobSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zarpas"}

	// bobTraditionalChinesePhrase represents bob traditional chinese phrase.
	bobTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貓貓"}
)

var (
	// bobPhrase represents bob phrase.
	bobPhrase = nook.Languages{
		language.AmericanEnglish:      bobAmericanEnglishPhrase,
		language.CanadianFrench:       bobCanadianFrenchPhrase,
		language.Dutch:                bobDutchPhrase,
		language.French:               bobFrenchPhrase,
		language.German:               bobGermanPhrase,
		language.Italian:              bobItalianPhrase,
		language.Japanese:             bobJapanesePhrase,
		language.Korean:               bobKoreanPhrase,
		language.LatinAmericanSpanish: bobLatinAmericanSpanishPhrase,
		language.Russian:              bobRussianPhrase,
		language.SimplifiedChinese:    bobSimplifiedChinesePhrase,
		language.Spanish:              bobSpanishPhrase,
		language.TraditionalChinese:   bobTraditionalChinesePhrase}
)

var (
	// Bob represents bob.
	Bob = nook.Villager{
		Character:   bobCharacter,
		Personality: personality.Lazy,
		Phrase:      bobPhrase}
)
