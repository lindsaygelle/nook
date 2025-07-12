package duck

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
	// weberBirthday represents weber birthday.
	weberBirthday = nook.Birthday{
		Day:   30,
		Month: time.June}
)

var (
	// weberCode represents weber code.
	weberCode = nook.Code{
		Value: "duk11"}
)

var (
	// weberAmericanEnglishName represents weber american english name.
	weberAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Weber"}

	// weberCanadianFrenchName represents weber canadian french name.
	weberCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bébert"}

	// weberDutchName represents weber dutch name.
	weberDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Weber"}

	// weberFrenchName represents weber french name.
	weberFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bébert"}

	// weberGermanName represents weber german name.
	weberGermanName = nook.Name{
		Language: language.German,
		Value:    "Volker"}

	// weberItalianName represents weber italian name.
	weberItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pasquale"}

	// weberJapaneseName represents weber japanese name.
	weberJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アチョット"}

	// weberKoreanName represents weber korean name.
	weberKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아잠만"}

	// weberLatinAmericanSpanishName represents weber latin american spanish name.
	weberLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paticio"}

	// weberRussianName represents weber russian name.
	weberRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вебер"}

	// weberSimplifiedChineseName represents weber simplified chinese name.
	weberSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卿德"}

	// weberSpanishName represents weber spanish name.
	weberSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paticio"}

	// weberTraditionalChineseName represents weber traditional chinese name.
	weberTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "卿德"}
)

var (
	// weberName represents weber name.
	weberName = nook.Languages{
		language.AmericanEnglish:      weberAmericanEnglishName,
		language.CanadianFrench:       weberCanadianFrenchName,
		language.Dutch:                weberDutchName,
		language.French:               weberFrenchName,
		language.German:               weberGermanName,
		language.Italian:              weberItalianName,
		language.Japanese:             weberJapaneseName,
		language.Korean:               weberKoreanName,
		language.LatinAmericanSpanish: weberLatinAmericanSpanishName,
		language.Russian:              weberRussianName,
		language.SimplifiedChinese:    weberSimplifiedChineseName,
		language.Spanish:              weberSpanishName,
		language.TraditionalChinese:   weberTraditionalChineseName}
)

var (
	// weberCharacter represents weber character.
	weberCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: weberBirthday,
		Code:     weberCode,
		Key:      character.Weber,
		Gender:   gender.Male,
		Name:     weberName,
		Special:  false}
)

var (
	// weberAmericanEnglishPhrase represents weber american english phrase.
	weberAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quaa"}

	// weberCanadianFrenchPhrase represents weber canadian french phrase.
	weberCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "plumeau"}

	// weberDutchPhrase represents weber dutch phrase.
	weberDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaaah"}

	// weberFrenchPhrase represents weber french phrase.
	weberFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "plumeau"}

	// weberGermanPhrase represents weber german phrase.
	weberGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaa"}

	// weberItalianPhrase represents weber italian phrase.
	weberItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaaa"}

	// weberJapanesePhrase represents weber japanese phrase.
	weberJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピヨ"}

	// weberKoreanPhrase represents weber korean phrase.
	weberKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꽥꽥"}

	// weberLatinAmericanSpanishPhrase represents weber latin american spanish phrase.
	weberLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuaa-cuaa"}

	// weberRussianPhrase represents weber russian phrase.
	weberRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кря-а"}

	// weberSimplifiedChinesePhrase represents weber simplified chinese phrase.
	weberSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唧唧"}

	// weberSpanishPhrase represents weber spanish phrase.
	weberSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuaa-cuaa"}

	// weberTraditionalChinesePhrase represents weber traditional chinese phrase.
	weberTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唧唧"}
)

var (
	// weberPhrase represents weber phrase.
	weberPhrase = nook.Languages{
		language.AmericanEnglish:      weberAmericanEnglishPhrase,
		language.CanadianFrench:       weberCanadianFrenchPhrase,
		language.Dutch:                weberDutchPhrase,
		language.French:               weberFrenchPhrase,
		language.German:               weberGermanPhrase,
		language.Italian:              weberItalianPhrase,
		language.Japanese:             weberJapanesePhrase,
		language.Korean:               weberKoreanPhrase,
		language.LatinAmericanSpanish: weberLatinAmericanSpanishPhrase,
		language.Russian:              weberRussianPhrase,
		language.SimplifiedChinese:    weberSimplifiedChinesePhrase,
		language.Spanish:              weberSpanishPhrase,
		language.TraditionalChinese:   weberTraditionalChinesePhrase}
)

var (
	// Weber represents weber.
	Weber = nook.Villager{
		Character:   weberCharacter,
		Personality: personality.Lazy,
		Phrase:      weberPhrase}
)
