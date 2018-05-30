package url

import (
	"testing"
	"fmt"
)

func TestAbsolute(t *testing.T) {
	fqdn := "localhost"
	port := "5000"
	path := "asd"

	absolute := Absolute(fqdn, port, path)
	excepted := fmt.Sprintf("http://%s:%s/%s", fqdn, port, path)
	if excepted != absolute {
		t.Fatalf("Failed to assert, that %s == %s", absolute, excepted)
	}

	port = "80"
	excepted = fmt.Sprintf("http://%s/%s", fqdn, path)
	absolute = Absolute(fqdn, port, path)
	if excepted != absolute {
		t.Fatalf("Failed to assert, that %s == %s", absolute, excepted)
	}

	path = "/asd"
	absolute = Absolute(fqdn, port, path)
	excepted = fmt.Sprintf("http://%s%s", fqdn, path)
	if excepted != absolute {
		t.Fatalf("Failed to assert, that %s == %s", absolute, excepted)
	}

}
