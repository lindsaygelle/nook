# Nook
Nook exports basic information for characters that appear in the video game series Animal Crossing.

Specifically this package focuses on the organisation and classification of residents and villagers.
Residents in the context of this package are Animal Crossing characters that cannot
be invited to live within the players town, city, village or island.

Villagers are Animal Crossing characters that live in homes within the players world but do not have
a special function, such as Tom Nook or Isabelle.

This package separates Animal Crossing characters by their animal type. This is a requirement due to the number
of characters there are in the series, and that some characters overlap with other villagers in name
or other elements that compose a character. An example of this is the name Petunia, who is both a Cow and a Rhinoceros villager. 
Another example is Snooty, which is the name of one of the wolves and a personality type.

The information for this package has largely been sourced from various Wiki pages across the internet.

All contributions and corrections are welcome and encouraged.

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
