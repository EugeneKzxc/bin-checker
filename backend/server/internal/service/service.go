package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type binData struct {
	Bank struct {
		Name string `json:"name"`
	} `json:"bank"`
}

func SearchBinInfo(bin int) (string, error) {
	url := fmt.Sprintf("https://binlist.io/lookup/%d/", bin) //почти всегда 429 код ответа

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(strconv.Itoa(resp.StatusCode))
	}

	var data binData
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.Bank.Name, nil
}
