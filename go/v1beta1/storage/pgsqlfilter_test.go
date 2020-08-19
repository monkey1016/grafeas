package storage_test

import (
	"fmt"
	"testing"

	"github.com/grafeas/grafeas/go/v1beta1/storage"
)

var myFilter storage.PgSQLFilterSql

func TestSimpleParseFilter(t *testing.T) {
	filter := `name="test_note_1"`
	actual := myFilter.ParseFilter(filter)
	expected := `data @@ '$.name == "test_note_1"'`
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("Expecting: " + expected + "\nGet: " + actual)
	}
}

func TestComplexParseFilter(t *testing.T) {
	filter := `vulnerability.details[*].package="dbus"`
	actual := myFilter.ParseFilter(filter)
	expected := `data @@ '$.vulnerability.details[*].package == "dbus"'`
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("Expecting: " + expected + "\nGet: " + actual)
	}
}

func TestMoreComplexParseFilter(t *testing.T) {
	filter := `resource.uri="registry.access.redhat.com/ubi8/ubi:8.0-199" AND vulnerability.severity="HIGH"`
	actual := myFilter.ParseFilter(filter)
	expected := `data @@ '$.resource.uri == "registry.access.redhat.com/ubi8/ubi:8.0-199"' AND data @@ '$.vulnerability.severity == "HIGH"'`
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("Expecting: " + expected + "\nGet: " + actual)
	}
}

func TestMostComplexParseFilter(t *testing.T) {
	filter := `resource.uri="registry.access.redhat.com/ubi8/ubi:8.0-199" AND vulnerability.severity="HIGH" AND name="tester"`
	actual := myFilter.ParseFilter(filter)
	expected := `data @@ '$.resource.uri == "registry.access.redhat.com/ubi8/ubi:8.0-199"' AND data @@ '$.vulnerability.severity == "HIGH"' AND data @@ '$.name == "tester"'`
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("Expecting: " + expected + "\nGet: " + actual)
	}
}
