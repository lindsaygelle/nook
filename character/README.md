# Character
Character exports the `nook.Key` references for all `nook.Character` instances.

The character package also contains all Animal Crossing characters, organised into their own separate packages grouped by animal type. 

This was a requirement due to some Animal Crossing characters sharing the same name. It was also a preference when trying to adopt the most scalable
and hassle free approach when extending and maintaining the numerous characters that are being added to, and modified with each installment to the series. It also allowed for each character to be placed within its own file, relative to the animal type it has. This makes it easier to extend and test each character, without hundreds of lines of code being grouped into a single file.

## Usage
The character package only exports the `nook.Key` values used within each `nook.Character`. These values can be used when creating a new map of `nook.Character`, `nook.Resident` or `nook.Villager`.

```go
package main

import (
	"fmt"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/character/pig"
	"golang.org/x/text/language"
)

var (
	// pigs is a map of my current favourite pig characters in Animal Crossing.
	pigs = map[nook.Key]nook.Villager{
		character.Chops: pig.Chops,
		character.Ganon: pig.Ganon}
)

func main() {
	fmt.Println("here are my favourite pig characters in Animal Crossing")
	for _, pig := range pigs {
		name := pig.Name.Must(language.AmericanEnglish).Value
		fmt.Println(fmt.Sprintf("name: %s, birthday: %v/%v", name, pig.Birthday.Day, pig.Birthday.Month))
	}
}

```
