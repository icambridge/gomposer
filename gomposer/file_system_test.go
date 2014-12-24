package gomposer

import (
	"os"
	"testing"
)

func Test_Remove_Deletes_Folder(t *testing.T) {
	vendorDir := os.TempDir() + "/vendors"
	dirName := vendorDir + "/symfony/symfony"
	os.MkdirAll(dirName, 0744)
	v := ComposerPackage{
		Name: "symfony/symfony",
	}
	Remove(vendorDir, v)

	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		t.Errorf("Failed to delete %v", dirName)
	}
}

func Test_Remove_Deletes_Parent_Folder(t *testing.T) {
	vendorDir := os.TempDir() + "/vendors"
	dirName := vendorDir + "/symfony"
	os.MkdirAll(dirName+"/symfony", 0744)
	v := ComposerPackage{
		Name: "symfony/symfony",
	}
	Remove(vendorDir, v)

	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		t.Errorf("Failed to delete %v", dirName)
	}
}

func Test_Remove_Keeps_Parent_Folder_When_Not_Empty(t *testing.T) {
	vendorDir := os.TempDir() + "/vendors"
	dirName := vendorDir + "/symfony"
	os.MkdirAll(dirName+"/symfony", 0744)
	fp, err := os.Create(dirName + "/remove_deletes_parent_folder")

	if err != nil {
		t.Errorf("%v", err)
	}

	fp.Write([]byte("hello world"))
	v := ComposerPackage{
		Name: "symfony/symfony",
	}
	Remove(vendorDir, v)

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		t.Errorf("Failed to keep %v when not empty", dirName)
	}
	os.Remove(dirName + "/remove_deletes_parent_folder")
}
