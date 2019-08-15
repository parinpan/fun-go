package fungo

import (
	"strings"
	"testing"
)

var (
	children = []child{
		child{Name: "Fachrin", Age: 20, Hobbies: []string{"Software Engineering", "Sing"}},
		child{Name: "Nobby", Age: 30, Hobbies: []string{"Software Architecture", "Jump"}},
		child{Name: "Thing", Age: 10, Hobbies: []string{"Software Infra", "Sit"}},
		child{Name: "Edward", Age: 40, Hobbies: []string{"Software Dev", "Stand"}},
	}
)

func TestMap(t *testing.T) {
	childrenSlice := NewSlice(children)
	expectedFirstNames := []string{"fachrin", "nobby", "thing", "edward"}

	sliceFirstNames := childrenSlice.Map([]string{}, func(value interface{}, index int) interface{} {
		return strings.ToLower(value.(child).Name)
	}).([]string)

	if len(expectedFirstNames) != len(sliceFirstNames) {
		t.Errorf("Slice length is unexpected")
		return
	}

	for index := range expectedFirstNames {
		if expectedFirstNames[index] != sliceFirstNames[index] {
			t.Errorf("Data doesn't match")
			return
		}
	}
}

func TestFilter(t *testing.T) {
	childrenSlice := NewSlice(children)
	expectedChildren := []child{children[0], children[1]}

	filteredChildren := childrenSlice.Filter(func(value interface{}, index int) bool {
		age := value.(child).Age
		return age > 10 && age < 40
	}).([]child)

	if len(expectedChildren) != len(filteredChildren) {
		t.Errorf("Slice length is unexpected")
		return
	}

	for index := range expectedChildren {
		if expectedChildren[index].Age != filteredChildren[index].Age ||
			expectedChildren[index].Name != filteredChildren[index].Name {
			t.Errorf("Data doesn't match")
			return
		}

		if len(expectedChildren[index].Hobbies) != len(filteredChildren[index].Hobbies) {
			t.Errorf("Slice (hobbies) length is unexpected")
			return
		}

		for hIndex := range expectedChildren[index].Hobbies {
			if expectedChildren[index].Hobbies[hIndex] != filteredChildren[index].Hobbies[hIndex] {
				t.Errorf("Data (hobbies) doesn't match")
				return
			}
		}
	}
}

func TestForEach(t *testing.T) {
	var childrenSlice = NewSlice(children)

	var ages = []int64{20, 30, 10, 40}
	var collectedAges = []int64{}

	childrenSlice.ForEach(func(value interface{}, index int) {
		collectedAges = append(collectedAges, value.(child).Age)
	})

	if len(ages) != len(collectedAges) {
		t.Errorf("Slice length is unexpected")
		return
	}

	for index := range ages {
		if ages[index] != collectedAges[index] {
			t.Errorf("Data doesn't match")
			return
		}
	}
}

func TestReduce(t *testing.T) {
	var childrenSlice = NewSlice(children)
	var totalAge = int64(100)

	var reducedTotalAge = childrenSlice.Reduce(func(accumulator int64, value interface{}) int64 {
		return accumulator + value.(child).Age
	})

	if totalAge != reducedTotalAge {
		t.Errorf("Age is not match")
		return
	}
}
