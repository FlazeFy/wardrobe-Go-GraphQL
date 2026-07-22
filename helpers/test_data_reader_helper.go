package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type TestDataReader struct{}

var testDataMu sync.Mutex

func testDataPath() string {
	root, err := filepath.Abs(".")
	if err != nil {
		root = "."
	}
	return filepath.Join(findModuleRoot(root), "tests", "test_data", "data.csv")
}

func findModuleRoot(dir string) string {
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return dir
		}
		dir = parent
	}
}

func GetValue(key string) (*string, error) {
	testDataMu.Lock()
	defer testDataMu.Unlock()

	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	f, err := os.Open(testDataPath())
	if err != nil {
		return nil, fmt.Errorf("failed to read test data: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	// Skip header
	if _, err := r.Read(); err != nil {
		return nil, nil
	}

	for {
		row, err := r.Read()
		if err != nil {
			break // EOF or malformed row, stop scanning
		}
		if len(row) >= 2 && row[0] == key {
			value := row[1]
			return &value, nil
		}
	}

	return nil, nil
}

func SetValue(key, value string) error {
	testDataMu.Lock()
	defer testDataMu.Unlock()

	if err := ensureFileExists(); err != nil {
		return err
	}

	rows, err := readAllRows()
	if err != nil {
		return err
	}

	found := false
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	for i := 1; i < len(rows); i++ {
		if len(rows[i]) > 0 && rows[i][0] == key {
			rows[i] = []string{key, value, timestamp}
			found = true
			break
		}
	}

	if !found {
		rows = append(rows, []string{key, value, timestamp})
	}

	f, err := os.Create(testDataPath())
	if err != nil {
		return fmt.Errorf("failed to save test data: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.WriteAll(rows); err != nil {
		return fmt.Errorf("failed to save test data: %w", err)
	}

	return nil
}

func readAllRows() ([][]string, error) {
	f, err := os.Open(testDataPath())
	if err != nil {
		return nil, fmt.Errorf("failed to read test data: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read test data: %w", err)
	}

	return rows, nil
}

func ensureFileExists() error {
	path := testDataPath()
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("failed to create test data file: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create test data file: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write([]string{"key", "value", "created_at"}); err != nil {
		return fmt.Errorf("failed to create test data file: %w", err)
	}
	w.Flush()

	return w.Error()
}
