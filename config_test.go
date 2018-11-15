package JSONConfig

import (
	"os"
	"testing"
)

type testData struct {
	TestString string `json:"foo"`
	TestInt    int
	TestArray  []string
}

func TestConfig(t *testing.T) {
	conf := Config{}

	data := testData{
		"foo",
		123,
		[]string{
			"foo",
			"bar",
		},
	}

	t.Run("should create", func(t *testing.T) {
		err := conf.CreateConfig("testData")
		if err != nil {
			t.Error("Received error: ", err)
		}

		if conf.fileName != "testData.json" {
			t.Errorf("Expected: testData.json, got: %v", conf.fileName)
		}
	})

	t.Run("should write", func(t *testing.T) {
		if err := conf.Write(data); err != nil {
			t.Error("Received error: ", err)
		}
	})

	t.Run("should get", func(t *testing.T) {
		retrieveTest := &testData{}

		if err := conf.Get(&retrieveTest); err != nil {
			t.Error("Received error: ", err)
		}

		if retrieveTest.TestInt != data.TestInt {
			t.Errorf("Expected: %d, got: %d", data.TestInt, retrieveTest.TestInt)
		}
	})

	t.Run("should override old config", func(t *testing.T) {
		data.TestInt = 321
		data.TestArray[1] = "testOverride"

		if err := conf.Write(data); err != nil {
			t.Error("Received error: ", err)
		}

		override := &testData{}

		if err := conf.Get(&override); err != nil {
			t.Error("Received error: ", err)
		}

		if override.TestInt != 321 {
			t.Errorf("Expected: %d, got: %d", 321, override.TestInt)
		}

		if override.TestArray != nil {
			if override.TestArray[1] != "testOverride" {
				t.Errorf("Expected: testOverride, got: %v", override.TestArray[1])
			}
		} else {
			t.Error("Test array []")
		}
	})

	t.Run("Open should get existing fileName", func(t *testing.T) {
		conf = Config{}

		if err := conf.Open("testData.json"); err != nil {
			t.Error("Could not open existing config")
		}
	})

	t.Run("Should reject files which are not json", func(t *testing.T) {
		conf2 := Config{}

		t.Run("Wrong extension", func(t *testing.T) {
			err := conf2.Open("testData.png")
			if err == nil {
				t.Error("Error expected: ", err)
			}
			if err.Error() != "File extension is not of type json" {
				t.Error(err)
			}
		})

		t.Run("Empty extension", func(t *testing.T) {
			err := conf2.Open("testData")
			if err == nil {
				t.Error("Error expected: ", err)
			}
			if err.Error() != "File extension is not of type json" {
				t.Error(err)
			}
		})
	})

	if err := os.Remove(conf.fileName); err != nil {
		t.Error("Could not delete testData.json, manual actions required")
	}
}