package object

import "testing"

func TestSetAndGet(t *testing.T) {
	env := NewEnvironment()
	str := &String{Value: "hello"}

	obj := env.Set("key", str)
	if obj != str {
		t.Fatal("expected", obj, "to be", str)
	}
	result, ok := env.Get("key")
	if !ok {
		t.Fatal("was unable to get string back")
	}

	if result != str {
		t.Fatal("expected", result, "to be", str)
	}
}

func TestGetKeys(t *testing.T) {
	env := NewEnvironment()
	str := &String{Value: "hello"}
	env.Set("key", str)
	env.Set("secondKey", str)

	keys := env.GetKeys()
	if !contains(keys, "key") || !contains(keys, "secondKey") {
		t.Fatal("Couldn't find key in environment")
	}
}

func contains(keys []string, key string) bool {
	for _, k := range keys {
		if key == k {
			return true
		}
	}
	return false
}
