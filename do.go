package format

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DocRecord struct {
	Name  string `json:"number"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(fnInput, fnOutput string) error {
	// Версия 1.0.0
	// tag 1.0.0
	/* Модуль должен прочитать файл со следующим
	   содержимым:
	   {"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
	   {"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
	   {"name":"Лисичка","age":3,"email":"alice@mail.ru"}

	   v1.0.0 должна создавать файл с содержимым:
	   [{"name":"Ёжик","age":10,"email":"ezh@mail.ru"},
	   {"name":"Зайчик","age":2,"email":"zayac@mail.ru"},
	   {"name":"Лисичка","age":3,"email":"alice@mail.ru"}] */

	// первым делом открываем входной и выходной файлы. Если возникнет ошибка на этом этапе, то и читать ничего не придется,
	// при большом объеме данных юзер раньше получит ошибку
	fin, err := os.Open(fnInput)

	if err != nil {
		return fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer fin.Close()

	fout, err := os.Create(fnOutput)
	if err != nil {
		return fmt.Errorf("Ошибка создания файла: %w", err)
	}
	defer fout.Close()

	dec := json.NewDecoder(fin)
	records := make([]DocRecord, 0, 10)

	for dec.More() {
		var d DocRecord
		err := dec.Decode(&d)
		if err != nil {
			return fmt.Errorf("Ошибка чтения записи из файла: %w", err)
		}
		records = append(records, d)
	}
	log.Printf("%+v", records)

	err = json.NewEncoder(fout).Encode(records)

	return nil
}
