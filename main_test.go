package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsyncMaxChunks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := splitTochanks(data, 3)
	max := maxChunks(chunks)
	require.Equal(t, 10, max)
}

func TestAsyncMaxChunksInEmptyArr(t *testing.T) {
	data := []int{}
	chunks := splitTochanks(data, 3)
	max := maxChunks(chunks)
	require.Equal(t, 0, max)
}

func TestMaxInEmtyArr(t *testing.T) {
	data := []int{}
	max := maximum(data)
	require.Equal(t, 0, max)
}

func TestMaxInOneElementArr(t *testing.T) {
	data := []int{1}
	max := maximum(data)
	require.Equal(t, 1, max)
}

func TestMaxInTwoElementArr(t *testing.T) {
	data := []int{1, 2}
	max := maximum(data)
	require.Equal(t, 2, max)
}

func TestMaxInManyElementArr(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max := maximum(data)
	require.Equal(t, 10, max)
}

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(10)
	require.Equal(t, 10, len(data))
}

func TestSplitTochanks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := splitTochanks(data, 2)
	require.Equal(t, 2, len(chunks))
	require.Equal(t, []int{1, 2, 3, 4, 5}, chunks[0])
	require.Equal(t, []int{6, 7, 8, 9, 10}, chunks[1])
}

func TestSplitTochanks1(t *testing.T) {
	data := []int{}
	chunks := splitTochanks(data, 2)
	require.Equal(t, 2, len(chunks))
	require.Equal(t, []int{}, chunks[0])
	require.Equal(t, []int{}, chunks[1])
}
