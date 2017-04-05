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
	expectKeyExists(t, keys, "key")
	expectKeyExists(t, keys, "secondKey")
}

func expectKeyExists(t *testing.T, keys []string, key string) {
	for _, k := range keys {
		if key == k {
			return
		}
	}
	t.Fatalf("Couldn't find key %s in keys: %s", key, keys)
}
