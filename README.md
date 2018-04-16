# XKCD Password Generator

Creates [XKCD-style password](http://xkcd.com/936/) based on your parameters.

**Fair warning: I'm not a cryptographer or any sort of expert in security. Use at your own risk.**

## Installation
`go get github.com/jlpellicer/xkcdpass`

### Usage
```Go
import "github.com/jlpellicer/xkcdpass"

config := xkcdpass.Config{
	MinLength: 12,
	MaxLength: 18,
	Language:  "es",
	Separator: ".",
	Numbers:   3,
}

password, err := xkcdpass.Generate(config)

if err != nil {
	log.Printf("Error: %s", err)
}

log.Printf("Password: %s", password)
```

## Notes

`xkcdpass` will try it's best to give you a password based on your parameters.

* `MinLength`: a minimum possible length will be calculated, based on the language file of your choice and an error will be returned if `MinLength` is less than the shortest word in the file
* `MaxLength`: also, if `MaxLength` is less than the shortest word, an error will be returned (ok, fringe case, but, hey, one tries to be thorough)
* `Language`: currently English (`en`) and Spanish (`es`) are supported. You can add your own `%s_words.txt` to `static` and if you would like to add your language to this repo then send me a PR. These lists have been curated to contain mostly words that are reasonably less common to be used in each language. If you want to compile a list of words for your language, try to find a source that can help you apply the same kind of filter and keep a high number of words
* `Separator`: can be any string you want, probably best to keep it to one character, such as a period, dash, space or none (use an empty string or leave it out)
* `Numbers`: add an amount of random numbers to the end of word list

Use along with [Dropbox's zxcvbn](https://github.com/dropbox/zxcvbn) password strength estimator for better, stronger passwords.
