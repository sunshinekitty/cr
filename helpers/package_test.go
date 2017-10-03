package helpers

import (
	"strings"
	"testing"
)

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
	if ValidPackageName("i") {
		t.Error("Package name \"i\" should be invalid")
	}
	if ValidPackageName("") {
		t.Error("Package name \"\" should be invalid")
	}
	// 51 chars
	if ValidPackageName("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy") {
		t.Error("Package name \"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy\" should be invalid")
	}
	if ValidPackageName("!@#$%^&*()invalid/\\") {
		t.Error("Package name \"!@#$%^&*()invalid\\/\" should be invalid")
	}
	if ValidPackageName("\"") {
		t.Error("Package name \"\"\" should be invalid")
	}
	if ValidPackageName("ABC") {
		t.Error("Package name \"ABC\" should be invalid")
	}
}

func TestValidRepositoryName(t *testing.T) {
	if ValidRepositoryName("\"") {
		t.Error("Repository name \"\"\" should be invalid")
	}
	// 142 chars
	a := strings.Repeat("a", 142)
	if ValidRepositoryName(a) {
		t.Error("Repository name \"", a, "\" should be invalid")
	}
	if ValidRepositoryName("space ") {
		t.Error("Repository name \"space \" should be invalid")
	}
	if ValidRepositoryName("a/") {
		t.Error("Repository name \"a/\" should be invalid")
	}
	if !ValidRepositoryName("123123123.dkr.ecr.us-west-2.amazonaws.com/amazonlinux:latest") {
		t.Error("Repository name \"123123123.dkr.ecr.us-west-2.amazonaws.com/amazonlinux:latest\" should be valid")
	}
	if !ValidRepositoryName("us.gcr.io/my-project/my-image:test") {
		t.Error("Repository name \"\" should be valid")
	}
	if !ValidRepositoryName("sunshinekitty/testing") {
		t.Error("Repository name \"sunshinekitty/testing\" should be valid")
	}
}

// https://tools.ietf.org/html/rfc793
func TestValidPort(t *testing.T) {
	if ValidPort(-1) {
		t.Error("Port \"-1\" should be invalid")
	}
	if ValidPort(0) {
		t.Error("Port \"0\" should be invalid")
	}
	if ValidPort(65536) {
		t.Error("Port \"65536\" should be invalid")
	}
	if !ValidPort(1) {
		t.Error("Port \"1\" should be valid")
	}
	if !ValidPort(65535) {
		t.Error("Port \"65535\" should be valid")
	}
}
