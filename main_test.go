// Cover testing -v -cover
package main

import (
	"testing"
	"time"
)

// Unit-testing
func TestDecode(t *testing.T) {
	// Load file
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	// Check result
	// Test case 1: Test Id
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got ", post.Id)
	}

	// Test case 2: Test Content
	if post.Content != "Hello world" {
		t.Error("Wrong content, was expecting 'Hello world' but got ", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}

// Skip long running -short(skip long running)
func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}

// Parallel testing -parallel
func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

// Benchmark testing -bench .
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
