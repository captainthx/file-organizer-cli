package organizer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const historyFile = "history/move.json"

type MoveRecord struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func SaveHistory(records []MoveRecord) error {
	_ = os.MkdirAll(filepath.Dir(historyFile), 0755)
	f, err := os.Create(historyFile)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(records)
}

func loadHistory() ([]MoveRecord, error) {
	f, err := os.Open(historyFile)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var records []MoveRecord
	if err := json.NewDecoder(f).Decode(&records); err != nil {
		return nil, err
	}
	return records, nil
}

func RevertLastMove() error {
	records, err := loadHistory()
	if err != nil {
		return errors.New("no history to revert")
	}

	for _, r := range records {
		fmt.Printf("Reverting: %s -> %s\n", r.To, r.From)
		err := os.Rename(r.To, r.From)
		if err != nil {
			return err
		}
	}

	return os.Remove(historyFile)
}
