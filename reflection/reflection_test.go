package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Clebson"},
			[]string{"Clebson"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Clebson", "SP"},
			[]string{"Clebson", "SP"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Clebson", 35},
			[]string{"Clebson"},
		},
		{
			"nested fields",
			Person{
				"Clebson",
				Profile{"SP", 35},
			},
			[]string{"Clebson", "SP"},
		},
		{
			"pointers to things",
			&Person{
				"Clebson",
				Profile{"SP", 35},
			},
			[]string{"Clebson", "SP"},
		},
		{
			"slices",
			[]Profile{
				{"SP", 23},
				{"BSB", 19},
			},
			[]string{"SP", "BSB"},
		},
		{
			"arrays",
			[2]Profile{
				{"SP", 23},
				{"BSB", 19},
			},
			[]string{"SP", "BSB"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"SP", 23}
			aChannel <- Profile{"BSB", 19}
			close(aChannel)
		}()

		var got []string
		want := []string{"SP", "BSB"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"SP", 23}, Profile{"BSB", 19}
		}

		var got []string
		want := []string{"SP", "BSB"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
