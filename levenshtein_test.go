package levenshtein

import (
	"testing"
)

var (
	r int
)

/*
 * UTILITY
 */

func CallOneAlloc(t *testing.T, a, b string, expected int) {
	i := Dist([]byte(a), []byte(b))
	if i != expected {
		t.Logf("Failed on input: '%s', '%s' != %d but %d", a, b, expected, i)
		t.Fail()
	}
}

func CallReentrant(t *testing.T, l *Lev, a, b string, expected int) {
	i := l.Dist([]byte(a), []byte(b))
	if i != expected {
		t.Logf("Failed on input: '%s', '%s' != %d but %d", a, b, expected, i)
		t.Fail()
	}
}

/*
 * TESTS
 */

func TestLevOneAlloc(t *testing.T) {
	CallOneAlloc(t, "abc", "abc", 0)
	CallOneAlloc(t, "abc", "edc", 2)
	CallOneAlloc(t, "abc", "def", 3)
	CallOneAlloc(t, "abc", "deabc", 2)
	CallOneAlloc(t, "deabc", "abc", 2)
	CallOneAlloc(t, "abc", "abcde", 2)
	CallOneAlloc(t, "abcde", "abc", 2)
	CallOneAlloc(t, "Kasper", "Morten", 5)
	CallOneAlloc(t, "Kasper E Eenberg", "Kasper Eenber", 3)
	CallOneAlloc(t, "Kasper Eenberg", "Mor Eenberg", 5)
	CallOneAlloc(t, "Mor Eenberg", "Kasper Eenberg", 5)
}

func TestLevOneAlloc64(t *testing.T) {
	l := New(1, 1, 1)

	CallReentrant(t, l, "abc", "abc", 0)
	CallReentrant(t, l, "abc", "edc", 2)
	CallReentrant(t, l, "abc", "def", 3)
	CallReentrant(t, l, "abc", "deabc", 2)
	CallReentrant(t, l, "deabc", "abc", 2)
	CallReentrant(t, l, "abc", "abcde", 2)
	CallReentrant(t, l, "abcde", "abc", 2)
	CallReentrant(t, l, "Kasper", "Morten", 5)
	CallReentrant(t, l, "Kasper E Eenberg", "Kasper Eenber", 3)
	CallReentrant(t, l, "Kasper Eenberg", "Mor Eenberg", 5)
	CallReentrant(t, l, "Mor Eenberg", "Kasper Eenberg", 5)
}

/*
 * BENCHMARKS
 */

func BenchmarkLevOneAllocA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dist([]byte("abc"), []byte("def"))
	}
}

func BenchmarkLevOneAllocB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dist([]byte("Kasper Sacharias Roos Eenberg"), []byte("Morten Houmøller Nygaard"))
	}
}

func BenchmarkReentrantA(b *testing.B) {
	l := New(1, 1, 1)
	for i := 0; i < b.N; i++ {
		r = l.Dist([]byte("abc"), []byte("def"))
	}
}

func BenchmarkReentrantB(b *testing.B) {
	l := New(1, 1, 1)
	for i := 0; i < b.N; i++ {
		r = l.Dist([]byte("Kasper Sacharias Roos Eenberg"), []byte("Morten Houmøller Nygaard"))
	}
}

func BenchmarkReentrantC(b *testing.B) {
	l := New(1, 1, 1)
	for i := 0; i < b.N; i++ {
		r = l.Dist([]byte("asper Sas s Eerg"), []byte("Mort Houer Nygrd"))
	}
}

func BenchmarkReentrantD(b *testing.B) {
	l := New(1, 1, 1)
	for i := 0; i < b.N; i++ {
		r = l.Dist([]byte("Kasper Sas s Eenrg"), []byte("Morte Houer Nygrad"))
	}
}
