package clients

import (
	"encoding/json"
	"errors"
	"net/http"
	"task0/models"
)

// GetKittyQuote fetches a quote from the catfact API and decodes it into
// models.NekoQuote. It returns the quote or an error.
func GetKittyQuote() (models.NekoQuote, error) {
	url := "https://catfact.ninja/fact"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.NekoQuote{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.NekoQuote{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.NekoQuote{}, errors.New("unexpected status from catfact API: " + resp.Status)
	}

	var nekoQuote models.NekoQuote
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&nekoQuote); err != nil {
		return models.NekoQuote{}, err
	}

	return nekoQuote, nil
}
