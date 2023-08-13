package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// atua sobre slice
func Test_append(t *testing.T) {
	// Arrange
	a := []int{1, 5, 6}
	b := []int{2, 3, 4}
	expected := []int{1, 5, 6, 2, 3, 4}

	// Act
	a = append(a, b...)

	// Assert
	assert.Equal(t, expected, a)
}

// atua sobre slice
func Test_copy(t *testing.T) {
	// Arrange
	a := []int{1, 5, 6}
	aCopy := make([]int, len(a))

	// Act
	copy(aCopy, a)

	// Assert
	assert.Equal(t, a, aCopy)
}

// atua sobre map
func Test_delete(t *testing.T) {
	// Arrange
	a := map[int]bool{1: true, 2: true, 3: true}
	expected := map[int]bool{1: true, 3: true}

	// Act
	delete(a, 2)

	// Assert
	assert.Equal(t, expected, a)
}

// atua sobre qualquer tipo
func Test_len(t *testing.T) {
	// Arrange
	nums := []int{1, 5, 6}
	text := "abc"

	// Act
	lenNums := len(nums)
	lenText := len(text)

	// Assert
	assert.Equal(t, 3, lenNums)
	assert.Equal(t, 3, lenText)
}

// atua sobre array e slice
func Test_cap(t *testing.T) {
	// Arrange
	var nums [3]int
	numsSlice := []int{1, 5, 6}

	// Act
	capNums := cap(nums)
	capNumsSlice := cap(numsSlice)

	// Assert
	assert.Equal(t, 3, capNums)
	assert.Equal(t, 3, capNumsSlice)
}

// atua sobre slice, map e channel
func Test_make(t *testing.T) {
	// Arrange
	// Act
	a := make([]int, 0)
	b := make(map[int]bool)
	c := make(chan int)

	// Assert
	assert.Equal(t, 0, len(a))
	assert.Equal(t, 0, len(b))
	assert.Equal(t, 0, len(c))
}

// atua sobre slice e map
func Test_clear(t *testing.T) {
	// Arrange
	a := []int{1, 2}
	b := map[int]bool{1: true, 2: true}

	// Act
	clear(a)
	clear(b)

	// Assert
	assert.Equal(t, 2, len(a))
	assert.Equal(t, 0, a[0])
	assert.Equal(t, 0, a[1])
	assert.Equal(t, 0, len(b))
}

// atua sobre int, fload e string
func Test_max(t *testing.T) {
	// Arrange
	// Act
	max1 := max(1, 2, 3)
	max2 := max(1.0, 2.0, 3.0)
	max3 := max("as", "c", "bola")

	// Assert
	assert.Equal(t, 3, max1)
	assert.Equal(t, 3.0, max2)
	assert.Equal(t, "c", max3)
}

func Test_min(t *testing.T) {
	// Arrange
	// Act
	min1 := min(1, 2, 3)
	min2 := min(1.0, 2.0, 3.0)
	min3 := min("as", "c", "bola")

	// Assert
	assert.Equal(t, 1, min1)
	assert.Equal(t, 1.0, min2)
	assert.Equal(t, "as", min3)
}

// atua sobre channel
func Test_close(t *testing.T) {
	// Arrange
	c1 := make(chan int)
	c2 := make(chan int)

	// Act
	go func() {
		c1 <- 1
		close(c2)
	}()
	_, c1IsOpen := <-c1
	_, c2IsOpen := <-c2

	// Assert
	assert.Equal(t, true, c1IsOpen)
	assert.Equal(t, false, c2IsOpen)
}

// se recupera de um panic
func Test_recover(t *testing.T) {
	// Arrange
	// Act
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		} else {
			t.Error("The code did not panic")
		}
	}()
	panic("panic")
}

func Test_init_var(t *testing.T) {
	// Arrange
	// Act
	var a *int
	var b int
	c := new(int)
	d := make([]int, 0)

	// Assert
	assert.Equal(t, (*int)(nil), a)
	assert.Equal(t, 0, b)
	assert.Equal(t, 0, *c)
	assert.Equal(t, 0, len(d))
}
