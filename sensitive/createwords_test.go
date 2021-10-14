package main

import "testing"

func TestCreateKeywords_portionToString(t *testing.T) {
	items := portionToString("piao")
	t.Log(items)
}

func TestCreateKeywords(t *testing.T) {
	items := CreateKeywords()
	t.Log(items)
}
