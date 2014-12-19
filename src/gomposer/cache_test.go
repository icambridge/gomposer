package gomposer

import (
	"fmt"
	"os/user"
	"testing"
)

func Test_GetCacheFileName(t *testing.T) {

	usr, err := user.Current()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	packageName := "behat/behat"
	safePackageName := "behat$behat"
	expected := fmt.Sprintf("%s/.composer/cache/repo/https---packagist.org/provider-%s.json", usr.HomeDir, safePackageName)

	actual := GetCacheFilename(packageName)

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
