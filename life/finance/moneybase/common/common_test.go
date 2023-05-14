package common

import "testing"

func TestAddDate(t *testing.T) {
	demo := "20060102"

	start, end, err := AddDate(demo, 1)
	if err != nil {
		t.Fatal(err)
	}

	if start != 20060103000001 {
		t.Fatalf("wrong wrong [%d]", start)
	}

	if end != 20060103235959 {
		t.Fatalf("wrong wrong wrong [%d]", end)
	}

	start, end, err = AddDate(demo, -2)
	if err != nil {
		t.Fatal(err)
	}

	if start != 20051231000001 {
		t.Fatalf("wrong wrong 2 [%d]", start)
	}

	if end != 20051231235959 {
		t.Fatalf("wrong wrong wrong 2 [%d]", end)
	}
}
