package helpers

import "testing"

func TestValidPackageName(t *testing.T) {
	// Valid package names
	if !ValidPackageName("va") {
		t.Error("Package name \"va\" should be valid")
	}
	if !ValidPackageName("va-lid") {
		t.Error("Package name \"va-lid\" should be valid")
	}
	if !ValidPackageName("va-l-id") {
		t.Error("Package name \"va-l-id\" should be valid")
	}
	// 50 chars
	if !ValidPackageName("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwx") {
		t.Error("Package name \"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwx\" should be valid")
	}
	// Invalid package names
	if ValidPackageName("invalid-") {
		t.Error("Package name \"invalid-\" should be invalid")
	}
	if ValidPackageName("-invalid") {
		t.Error("Package name \"-invalid\" should be invalid")
	}
	if ValidPackageName("invalid_") {
		t.Error("Package name \"invalid-\" should be invalid")
	}
	if ValidPackageName("_invalid") {
		t.Error("Package name \"-invalid\" should be invalid")
	}
	// 51 chars
	if ValidPackageName("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy") {
		t.Error("Package name \"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy\" should be invalid")
	}
	if ValidPackageName("!@#$%^&*()invalid/\\") {
		t.Error("Package name \"!@#$%^&*()invalid\\/\" should be invalid")
	}
	if ValidPackageName("ABC") {
		t.Error("Package name \"ABC\" should be invalid")
	}
}
