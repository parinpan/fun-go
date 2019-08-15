package fungo

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	maxChildren = 1000000
)

type child struct {
	Name    string
	Age     int64
	Hobbies []string
}

func generateChildrenList(maxData int) []child {
	var listOfChildren = make([]child, 0)

	for i := 0; i < maxData; i++ {
		listOfChildren = append(listOfChildren, child{
			Name:    "Fachrin",
			Age:     rand.Int63n(100),
			Hobbies: []string{"Guitar", "Software Engineering", "Reading"},
		})
	}

	return listOfChildren
}

func getFilteredChildrenNormalAgeMoreThan50(childrenList []child) []child {
	childrenSlice := make([]child, 0)

	for _, child := range childrenList {
		if child.Age > 50 {
			childrenSlice = append(childrenSlice, child)
		}
	}

	return childrenSlice
}

func normalForLoopChildren(childrenList []child) {
	for _, child := range childrenList {
		if child.Age > 0 {

		}
	}
}

func getMappedChildrenNameAndHobby(childrenList []child) []string {
	var nameWithFirstHobby = make([]string, 0)

	for _, child := range childrenList {
		nameWithFirstHobby = append(nameWithFirstHobby, fmt.Sprint(
			child.Name,
			child.Hobbies[0],
		))
	}

	return nameWithFirstHobby
}

func getReducedChildrenAge(childrenList []child) int64 {
	var accumulator = int64(0)

	for _, child := range childrenList {
		accumulator += child.Age
	}

	return accumulator
}

func BenchmarkFunctionalIteratorMap(b *testing.B) {
	childrenList := generateChildrenList(maxChildren)
	childrenSlice := NewSlice(childrenList)

	b.Run("functionalIterator", func(b *testing.B) {
		_ = childrenSlice.Map([]string{}, func(value interface{}, index int) interface{} {
			child := value.(child)
			return fmt.Sprint(child.Name, child.Hobbies[0])
		}).([]string)
	})

	b.Run("normalIterator", func(b *testing.B) {
		_ = getMappedChildrenNameAndHobby(childrenList)
	})
}

func BenchmarkFunctionalIteratorFilter(b *testing.B) {
	childrenList := generateChildrenList(maxChildren)
	childrenSlice := NewSlice(childrenList)

	b.Run("functionalIterator", func(b *testing.B) {
		_ = childrenSlice.Filter(func(value interface{}, index int) bool {
			child := value.(child)
			return child.Age > 50
		}).([]child)
	})

	b.Run("normalIterator", func(b *testing.B) {
		_ = getFilteredChildrenNormalAgeMoreThan50(childrenList)
	})
}

func BenchmarkFunctionalIteratorForEach(b *testing.B) {
	childrenList := generateChildrenList(maxChildren)
	childrenSlice := NewSlice(childrenList)

	b.Run("functionalIterator", func(b *testing.B) {
		childrenSlice.ForEach(func(value interface{}, index int) {
			if value.(child).Age > 0 {

			}
		})
	})

	b.Run("normalIterator", func(b *testing.B) {
		normalForLoopChildren(childrenList)
	})
}

func BenchmarkFunctionalIteratorReduce(b *testing.B) {
	childrenList := generateChildrenList(maxChildren)
	childrenSlice := NewSlice(childrenList)

	b.Run("functionalIterator", func(b *testing.B) {
		_ = childrenSlice.Reduce(func(accumulator int64, value interface{}) int64 {
			return accumulator + value.(child).Age
		})
	})

	b.Run("normalIterator", func(b *testing.B) {
		_ = getReducedChildrenAge(childrenList)
	})
}
