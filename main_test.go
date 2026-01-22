package randomuser

import (
	"bytes"
	"testing"
)

func TestRandomImage(t *testing.T) {
	portrait := RandomImage()
	if portrait == nil {
		t.Error("RandomImage() returned nil")
	}
	if len(portrait) == 0 {
		t.Error("RandomImage() returned empty byte slice")
	}
	// Check for JPEG magic number
	if !bytes.HasPrefix(portrait, []byte{0xFF, 0xD8}) {
		t.Error("RandomImage() did not return a valid JPEG")
	}
}

func TestRandomManImage(t *testing.T) {
	portrait := RandomManImage()
	if portrait == nil {
		t.Error("RandomManImage() returned nil")
	}
	if len(portrait) == 0 {
		t.Error("RandomManImage() returned empty byte slice")
	}
	// Check for JPEG magic number
	if !bytes.HasPrefix(portrait, []byte{0xFF, 0xD8}) {
		t.Error("RandomManImage() did not return a valid JPEG")
	}
}

func TestRandomWomanImage(t *testing.T) {
	portrait := RandomWomanImage()
	if portrait == nil {
		t.Error("RandomWomanImage() returned nil")
	}
	if len(portrait) == 0 {
		t.Error("RandomWomanImage() returned empty byte slice")
	}
	// Check for JPEG magic number
	if !bytes.HasPrefix(portrait, []byte{0xFF, 0xD8}) {
		t.Error("RandomWomanImage() did not return a valid JPEG")
	}
}

func TestRandomDistribution(t *testing.T) {
	// Test that RandomImage() returns both men and women over multiple calls
	const iterations = 100
	menCount := 0
	womenCount := 0

	for i := 0; i < iterations; i++ {
		portrait := RandomImage()
		// Check if it's a man (compare with men images)
		for j := 0; j < len(menImages); j++ {
			if bytes.Equal(portrait, menImages[j]) {
				menCount++
				break
			}
		}
		// Check if it's a woman (compare with women images)
		for j := 0; j < len(womenImages); j++ {
			if bytes.Equal(portrait, womenImages[j]) {
				womenCount++
				break
			}
		}
	}

	if menCount == 0 {
		t.Error("RandomImage() never returned a man's portrait in 100 iterations")
	}
	if womenCount == 0 {
		t.Error("RandomImage() never returned a woman's portrait in 100 iterations")
	}
	if menCount+womenCount != iterations {
		t.Errorf("Count mismatch: men=%d, women=%d, total=%d", menCount, womenCount, iterations)
	}
}

func BenchmarkRandomImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomImage()
	}
}

func BenchmarkRandomManImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomManImage()
	}
}

func BenchmarkRandomWomanImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomWomanImage()
	}
}