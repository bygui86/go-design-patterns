package singleton

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInstance(t *testing.T) {
	// given 1
	sing1 := GetInstance()
	require.NotNil(t, sing1, "First singleton is nil")
	title1 := "Title"
	// when 1
	sing1.SetTitle(title1)
	// then 1
	assert.Equal(t, title1, sing1.GetTitle(), "Title is not set")

	// given 2
	// -
	// when 2
	sing2 := GetInstance()
	// then 2
	assert.Equal(t, sing1, sing2, "New instance different")

	// given 3
	title2 := "New title"
	// when 3
	sing2.SetTitle(title2)
	// then 3
	assert.Equal(t, sing2.GetTitle(), sing1.GetTitle(), "Title different after change")
}

func TestGetInstance_Concurrency(t *testing.T) {
	// given
	sing1 := GetInstance()
	sing2 := GetInstance()

	// when
	var waitGroup sync.WaitGroup
	for routineNum := 0; routineNum < 3000; routineNum++ {
		titleId := routineNum

		waitGroup.Add(1)
		go func() {
			sing1.SetTitle(fmt.Sprintf("title_%d", titleId))
			waitGroup.Done()
		}()

		waitGroup.Add(1)
		go func() {
			sing2.SetTitle(fmt.Sprintf("title_2_%d", titleId))
			waitGroup.Done()
		}()
	}

	// then
	assert.Equal(t, sing1.GetTitle(), sing2.GetTitle())
}
