# Nook
Nook exports basic information for characters that appear in the video game series Animal Crossing.

Specifically this package centers on the organisation and classification of residents versus villagers, as well as
defining the common attributes between them.

Residents in the context of this package are Animal Crossing characters that cannot
be invited as home owners within the players town, city, or island. Residents perform a special role within the world
and often are quest givers. An example of a Resident would be Isabelle, Tom Nook or Guliver.

Villagers are Animal Crossing characters that can be invited to stay in the players town, city or island,
and make up the larger world population. Characters such as Alfonso, Tabby and Eggbert are all examples of Villagers. Villagers
can be given gifts, have personalities and homes.

This package separates and organises Animal Crossing characters by their animal type. This is a requirement due to the large number
of characters there are in the series, and that some characters overlap with other components in name. 
An example of this is the name Petunia, who is both a Cow and a Rhinoceros. Another example is Snooty, which is the name of one of the wolves and a personality type.

The information for this package has largely been sourced from various Wiki pages across the internet.

All contributions and corrections are welcome and encouraged.

## Go
This package can be added as a dependency using Go's package manager.

### Install
Adding the package.

```sh
go get -u github.com/lindsaygelle/nook
```

## Docker
This code can be run and executed from within a Docker Container. 
Below are the instructions to use the provided Dockerfile as the development and testing environment.

### Building 
Building the Container.

```sh

docker build . -t nook
```

### Running
Developing and running Go from within the Container.

```sh

docker run -it --rm --name nook nook
```

## Contents
Nook at the package level only exports the concrete types consumed by this package. The philosophy was to keep each package extensible, without the concern of
searching through various subpackages to find all types that make up the package. Through this, ideally, the pattern will keep the imports relatively straightforward and simple. It should also allow each child package to focus on using only the known set of building blocks to create its exports, and limit
the exposure to coding towards circular dependencies. Most of the decisions were shaped by the content of the Animal Crossing series, versus developer preference. For example, the initial concept was to have all characters live at the root package, but name conflicts, data types, and maintance became an issue. Having each character exist in its own directory was also attempted, but imports became confusing. All recommendations welcome, but for now this appears to be the best trade off.

## Extending
Nook was built with extensibility in mind. The simplest way to add to the functionality of the package is to import the concrete type and apply the required changes.

### Character
Below is an example of adding functionality to the `nook.Character` type.

```go
package main

import (
    "fmt"

    "github.com/lindsaygelle/nook"
    "github.com/lindsaygelle/nook/character/alligator"
    "golang.org/x/text/language"
)

var (
    alligators = []nook.Villager{
        alligator.Alfonso,
        alligator.Alli}
)

type Villager struct{
    nook.Villager
}

func (v Villager) Greet() string {
    return fmt.Sprintf("%s! My name is %s. Nice to meet you.", v.Phrase.Must(language.AmericanEnglish).Vale, v.Name.Must(language.AmericanEnglish).Value)
}

func makeVillager(villager nook.Villager) Villager {
    return Villager{villager}
}

func makeVillagers(villagers ...nook.Villager) []Villager {
    v := make([]Villager, len(villagers))
    for i, villager := range villagers {
        v[i] = makeVillager(villager)
    }
    return villagers
}

func main() {
    villagers := makeVillagers(alligators...) 
    for _, villager := range villagers {
        fmt.Println(villager.Greet())
    }
}
```