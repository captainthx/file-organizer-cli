package organizer

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Rules map[string]string

func LoadRules(path string) (Rules, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rules Rules
	err = json.NewDecoder(file).Decode(&rules)
	return rules, err
}

func OrganizeFiles(target string, rules Rules, interactive bool) error {
	var moves []MoveRecord

	// ถ้า interactive: ยืนยันก่อนเริ่มย้ายทั้งหมด
	if interactive {
		fmt.Print("Do you want to organize all matched files? [y/N]: ")
		var resp string
		fmt.Scanln(&resp)
		if strings.ToLower(resp) != "y" {
			fmt.Println("Aborted.")
			return nil
		}
	}

	err := filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		folder, ok := rules[ext]
		if !ok {
			return nil
		}
		destDir := filepath.Join(target, folder)
		_ = os.MkdirAll(destDir, 0755)
		destPath := filepath.Join(destDir, filepath.Base(path))

		if path == destPath {
			return nil
		}
		if _, err := os.Stat(destPath); err == nil {
			fmt.Println("[Skipped]", destPath, "already exists")
			return nil
		}

		fmt.Printf("Moving %s -> %s\n", path, destPath)
		if err := os.Rename(path, destPath); err == nil {
			moves = append(moves, MoveRecord{From: path, To: destPath})
		}
		return nil
	})

	if err != nil {
		return err
	}

	if len(moves) > 0 {
		fmt.Printf("Moved %d files.\n", len(moves))
	}
	return SaveHistory(moves)
}

func PreviewFiles(target string, rules Rules) error {
	fmt.Println("Previewing file organization:")
	return filepath.Walk(target, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		ext := filepath.Ext(path)
		folder, ok := rules[ext]
		if !ok {
			return nil
		}

		destDir := filepath.Join(target, folder)
		destPath := filepath.Join(destDir, filepath.Base(path))
		fmt.Printf("[Preview] %s -> %s\n", path, destPath)
		return nil
	})
}
