package concurrency

import (
	"reflect"
	"testing"
)

// Test (should be in test file tho)
func mockWebsiteChecker(string) bool {
	return true
}
func TestCheckWebsites(t *testing.T) {
	t.Run("check for calling the dependency", func(t *testing.T) {

		// mock websitechecker func
		url1 := "https://google.com"
		url2 := "https://notion.so"
		websites := []string{url1, url2}
		expectedData := map[string]bool{
			url1: true,
			url2: true,
		}
		resultData := CheckWebsites(mockWebsiteChecker, websites)
		if !reflect.DeepEqual(expectedData, resultData) {
			t.Errorf("Expected %#v but got %#v", expectedData, resultData)
		}
	})
}
