package gomposer

import (
	"testing"
	"os"
)

func Test_Deletes_Folder(t *testing.T) {
	vendorDir := os.TempDir()  + "/vendors"
	dirName := vendorDir+"/symfony/symfony"
	os.MkdirAll(dirName, 0744)
	v := Version{
		Name: "symfony/symfony",
	}
	Remove(vendorDir, v)

	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		t.Errorf("Failed to delete %v", dirName)
	}
}
