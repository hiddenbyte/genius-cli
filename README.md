# genius-cli
genius-cli is a CLI tool that allows you to search and read your favorite songs lyrics without leaving the terminal window.

The songs lyrics are provided by [Genius](https://genius.com/) through their [API](https://docs.genius.com/).

## Build

### Build requirements

- Go 1.12 or newer

### How
Use the Go build command, in order to generate a binary of genius-cli. Binary output path: `./genius-cli`.

`go build`

## Usage

```
Usage of ./genius-cli:
  -s string
        Prints the lyrics from a Genius song document (short)
  -setup
        Setup genius-cli
  -song string
        Prints the lyrics from a Genius song document
```

### Setup
Before start using this tool you will need to authorize genius-cli to access Genius trough your account

1. Type `./genius-cli --setup`

2. Open auth URL 

3. Authorize

4. Copy the access_token into the genius-cli

### Examples

**Print 'A$AP Rocky ft Skept - Praise the Lord' chorus part**
```
➜  ~ ./genius-cli --song "A$AP Rocky Praise the Lord (Da Shine)" | grep Chorus -A 4
[Chorus: Skepta]
I came, I saw, I came, I saw
I praise the Lord, then break the law ( /A-ap-rocky-praise-the-lord-da-shine-lyrics#note-14632511 )
I take what's mine, then take some more
It rains, it pours, it rains, it pours ( /A-ap-rocky-praise-the-lord-da-shine-lyrics#note-14640053 )

...
```

**Print 'Profjam Mortalhas' lyrics**
```
➜  ~ ./genius-cli --song "Profjam Mortalhas"
[Letra de "Mortalhas"]

[Intro]
Yeah
Think music, yeah
ProfJam

Pacotes de mortalha no bolso ( /Profjam-mortalhas-lyrics#note-16752037 )
Não podes ter a medalha sem esforço
Vou corrigir a falha no posto

...
```
