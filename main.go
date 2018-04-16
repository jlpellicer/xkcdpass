package xkcdpass

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	mathRand "math/rand"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	MinLength int
	MaxLength int
	Language  string
	Separator string
	Numbers   int
}

const filepath = "../github.com/jlpellicer/xkcdpass/static"

func Generate(config Config) (string, error) {
	if config.MaxLength < config.MinLength {
		return "", errors.New("xkcdpass: MaxLength should be greater than or equal to MinLength")
	}
	password := ""
	minWordLen := 100
	maxWordLen := 0
	filename := getFilename(config)
	data, err := readFile(filename)
	if err != nil {
		return "", errors.New("xkcdpass: failed to read file " + filename)
	}
	dict := string(data)
	words := strings.Split(dict, "\n")
	passwordWords := []string{}
	for index, w := range words {
		cleanWord := strings.TrimSpace(w)
		if cleanWord != "" {
			words[index] = cleanWord
			if len(words[index]) < minWordLen {
				minWordLen = len(words[index])
			}
			if len(words[index]) > maxWordLen {
				maxWordLen = len(words[index])
			}
		}
	}
	minPossibleLength := minWordLen
	if config.MinLength < minPossibleLength {
		return "", errors.New("xkcdpass: min possible length " + strconv.Itoa(minPossibleLength))
	}
	if config.MaxLength < minPossibleLength {
		return "", errors.New("xkcdpass: max length cannot be less than " + strconv.Itoa(minPossibleLength))
	}

	max := big.NewInt(int64(len(words)) - 1)

	for {
		passwordWords, err = generatePassword(config, words, max)
		password, err = joinPasswordWords(passwordWords, config)
		if err != nil {
			return "", err
		}
		if len(password) >= config.MinLength && len(password) <= config.MaxLength {
			break
		}
	}

	return password, nil
}

func joinPasswordWords(passwordWords []string, config Config) (password string, err error) {
	if len(passwordWords) == 0 {
		return password, errors.New("xkcdpass: no words to join")
	}
	password = strings.Join(passwordWords, config.Separator)
	return
}

func generatePassword(config Config, words []string, max *big.Int) (passwordWords []string, err error) {
	passwordWords = []string{}
	seen := make(map[string]struct{}, len(words))
	errorMsg := ""
	i := 0
	for len(strings.Join(passwordWords, "")) < config.MinLength {
		n, err := rand.Int(rand.Reader, max)
		index := int(n.Int64())
		if err != nil {
			errorMsg = ": failed to generate random number"
			break
		}
		word := words[index]
		if word == "" {
			errorMsg = ": failed to read word"
			err = errors.New("empty word")
			break
		}
		_, ok := seen[word]
		if ok {
			i--
			continue
		}
		seen[word] = struct{}{}
		passwordWords = append(passwordWords, word)
		i++
	}
	if err != nil {
		return []string{}, errors.New("xkcdpass " + errorMsg + " : " + err.Error())
	}
	if config.Numbers > 0 {
		var float float64
		for float < 0.1 {
			float = mathRand.Float64()
		}
		power := math.Pow(10, float64(config.Numbers))
		mult := float64(power)
		res := mult * float
		num := int16(res)
		number := strconv.Itoa(int(num))
		passwordWords = append(passwordWords, number)
	}
	return
}

func getFilename(config Config) (filename string) {
	if config.Language == "" {
		config.Language = "en"
	}
	filename = fmt.Sprintf("%s/%s_words.txt", filepath, config.Language)
	return
}

func readFile(filepath string) (data []byte, error error) {
	data, error = ioutil.ReadFile(filepath)
	if error != nil {
		return
	}
	return
}

func readDir(root string) {
	file, err := os.Open(root)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fmt.Println(name)
	}
}
