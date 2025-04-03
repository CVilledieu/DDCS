package util

import (
	C "DCS/source/char"
	"encoding/json"
	"os"
)

// A list of characters is read from a separate json file with limited character info
// This list is used to display currently created characters
// User then either selects a character or creates a new one.
// Newly created characters are appended to the limited info list and a new json file is created for that character.
// Loading a character from the list will load the full character info from a separate json file.

func saveCharacter(c C.Character) {
	fileName := "CharList" + ".json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	j, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(j)
	if err != nil {
		panic(err)
	}
}

func LoadCharacter(name string) C.Character {
	fileName := name + ".json"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var c C.Character
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}

//Pulls all saved characters from a single json file

func GetListOfCharacters() []C.Character {
	fileName := "CharList" + ".json"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var c []C.Character
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}

func Save(clist []C.Character) {
	fileName := "CharList" + ".json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	j, err := json.Marshal(clist)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(j)
	if err != nil {
		panic(err)
	}
}
