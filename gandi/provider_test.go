package gandi

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"gandi": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GANDI_PERSONAL_ACCESS_TOKEN"); v == "" {
		t.Fatal("GANDI_PERSONAL_ACCESS_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("GANDI_URL"); v == "" {
		t.Fatal("GANDI_URL must be set for acceptance tests")
	}
}
