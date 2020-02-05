package dvcardlib

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LoadAllCardLibsFromFolder загружает все библиотеки карточек из папки folder
func LoadAllCardLibsFromFolder(folder string) ([]CardLib, error) {
	cardLibs := []CardLib{}

	log.Printf("Получение файлов из папки %s \n", folder)

	err := filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {

			// Только файлы и только XML
			if info.IsDir() || filepath.Ext(path) != ".xml" {
				return nil
			}

			reader, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("Ошибка чтения файла %s\n%s", path, err.Error())
			}
			defer reader.Close()

			log.Printf("Поиск признаков библиотеки карточек в файле %s", path)
			isCardLib, err := IsCardLibFile(bufio.NewReader(reader))

			// Была ошибка или не библиотека карточек - выходим
			if err != nil || isCardLib == false {
				return err
			}

			// Возвращаем читателя к началу
			reader.Seek(0, 0)

			// Загружаем файл библиотеки карточек
			cardLib, err := LoadCardLibFromFile(bufio.NewReader(reader))
			if err != nil {
				return err
			}

			cardLibs = append(cardLibs, cardLib)
			fmt.Printf("Добавлена библиотека карточек %s\n", cardLib.Alias)
			return nil
		})

	if err != nil {
		return cardLibs, err
	}

	return cardLibs, nil
}

// IsCardLibFile возвращает true, если файл библиотеки карточек
func IsCardLibFile(reader *bufio.Reader) (bool, error) {

	// Читаем первые 80 символов из файла - достаточно, чтобы понять, что это библиотека карточек
	var b = make([]byte, 80)
	_, err := reader.Read(b)

	if err != nil {
		return false, fmt.Errorf("Ошибка получении данных из буфера в IsCardLibFile.\n%s", err.Error())
	}

	// Если в полученной строке есть тег CardLibrary, значит - библиотека карточек
	if strings.Contains(string(b), "<CardLibrary ") {
		return true, nil
	}
	return false, nil
}

// LoadCardLibFromFile загружает библиотеку карточек из файла file
func LoadCardLibFromFile(reader *bufio.Reader) (CardLib, error) {
	var cardlib CardLib
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return cardlib, fmt.Errorf("Ошибка чтения потока из файла библиотеки карточек\n%s", err.Error())
	}

	err = xml.Unmarshal(b, &cardlib)
	if err != nil {
		return cardlib, fmt.Errorf("Ошибка десериализации файла библиотеки карточек\n%s", err.Error())
	}

	return cardlib, nil
}

// CardLib описывает схему библиотеки карточек в XML
type CardLib struct {
	ID    string            `xml:"ID,attr"`
	Alias string            `xml:"Alias,attr"`
	Names []LocalizedString `xml:"Name>LocalizedString"`
}

// LocalizedString описывает узел локализованного значения
type LocalizedString struct {
	Lang  string `xml:"Language,attr"`
	Value string `xml:",chardata"`
}
