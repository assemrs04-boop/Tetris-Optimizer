package tetromino

import (
	"os"
	"testing"
)

func TestReadAndValidate(t *testing.T) {
	goodData := "####\n....\n....\n....\n\n.#..\n.#..\n.#..\n.#..\n"
	badData := "###.\n....\n....\n....\n"

	_ = os.WriteFile("test_good.txt", []byte(goodData), 0644)
	defer os.Remove("test_good.txt")

	_ = os.WriteFile("test_bad.txt", []byte(badData), 0644)
	defer os.Remove("test_bad.txt")

	res, err := ReadAndValidate("test_good.txt")
	if err != nil || len(res) != 2 {
		t.Errorf("Expected 2 valid tetrominoes, got %v", len(res))
	}

	res, _ = ReadAndValidate("test_bad.txt")
	if res != nil {
		t.Errorf("Expected nil for invalid tetromino")
	}
}
