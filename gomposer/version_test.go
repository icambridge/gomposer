package gomposer

import (
	"testing"
)

func Test_VersionComparer_GreaterThan_False(t *testing.T) {

	vc := VersionComparer{}

	output := vc.GreaterThan("1.2.1", "1.2.2")

	if output == true {
		t.Error("Expected a false response, but got a true response")
	}
}

func Test_VersionComparer_GreaterThan_True(t *testing.T) {

	vc := VersionComparer{}

	output := vc.GreaterThan("1.2.3", "1.2.2")

	if output == false {
		t.Error("Expected a true response, but got a false response")
	}
}

func Test_VersionComparer_LessThan_False(t *testing.T) {

	vc := VersionComparer{}

	output := vc.LessThan("1.2.4", "1.2.2")

	if output == true {
		t.Error("Expected a false response, but got a true response")
	}
}

func Test_VersionComparer_LessThan_True(t *testing.T) {

	vc := VersionComparer{}

	output := vc.LessThan("1.2.1", "1.2.2")

	if output == false {
		t.Error("Expected a true response, but got a false response")
	}
}

func Test_VersionComparer_Equal_False(t *testing.T) {

	vc := VersionComparer{}

	output := vc.Equal("1.2.4", "1.2.2")

	if output == true {
		t.Error("Expected a false response, but got a true response")
	}
}

func Test_VersionComparer_Equal_True(t *testing.T) {

	vc := VersionComparer{}

	output := vc.Equal("1.2.2", "1.2.2")

	if output == false {
		t.Error("Expected a true response, but got a false response")
	}
}

func Test_VersionComparer_GreaterThanOrEqual_False(t *testing.T) {

	vc := VersionComparer{}

	output := vc.GreaterThanOrEqual("1.2.1", "1.2.2")

	if output == true {
		t.Error("Expected a false response, but got a true response")
	}
}

func Test_VersionComparer_GreaterThanOrEqual_True(t *testing.T) {

	vc := VersionComparer{}

	output := vc.GreaterThanOrEqual("1.2.2", "1.2.2")

	if output == false {
		t.Error("Expected a true response, but got a false response")
	}
}

func Test_VersionComparer_LessThanOrEqual_False(t *testing.T) {

	vc := VersionComparer{}

	output := vc.LessThanOrEqual("1.2.4", "1.2.2")

	if output == true {
		t.Error("Expected a false response, but got a true response")
	}
}

func Test_VersionComparer_LessThanOrEqual_True(t *testing.T) {

	vc := VersionComparer{}

	output := vc.LessThanOrEqual("1.2.1", "1.2.2")

	if output == false {
		t.Error("Expected a true response, but got a false response")
	}
}
