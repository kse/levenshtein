// levenshtein implements the Levenshtein edit distance using the
// Wagner-Fisher algorithm
// (http://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm).
package levenshtein

/* TODO:
 * Lazy evalutation could be implemented.
 * Why is non-threadsafe Dist faster than threadsafe for long strings?
 * Is calling C functions faster?
 */

type Lev struct {
	cost                         []int
	Delete, Substitution, Insert int
}

func New(deletion, substitution, insert int) (l *Lev) {
	l = &Lev{
		Delete:       deletion,
		Substitution: substitution,
		Insert:       insert,
		cost:         make([]int, 96),
	}

	return
}

var (
	cost []int
)

// Dist is a non-goroutine safe version of the Levenshtein edit distance.
func Dist(a, b []byte) (d int) {
	var (
		m int = len(a)
		n int = len(b)

		i, j, p, pMod int = 0, 0, n + 1, 0
	)

	if cost == nil || len(cost) < 2*(n+1) {
		cost = make([]int, 2*(n+1))
	}

	for i = 0; i < n+1; i++ {
		cost[i] = i
	}

	for j = 1; j < m+1; j++ {
		cost[p] = j

		for i = 1; i < n+1; i++ {

			if a[j-1] == b[i-1] {
				cost[p+i] = cost[pMod+i-1]
			} else {
				d = cost[p+i-1] + 1

				if cost[pMod+i]+1 < d {
					d = cost[pMod+i] + 1
				}

				if cost[pMod+i-1]+1 < d {
					d = cost[pMod+i-1] + 1
				}

				cost[p+i] = d
			}
		}

		if p == 0 {
			p = n + 1
			pMod = 0
		} else {
			pMod = n + 1
			p = 0
		}
	}

	d = cost[pMod+n]

	return
}

// Dist is a goroutine safe version of the Levenshtein edit distance.
func (l *Lev) Dist(a, b []byte) (d int) {
	var (
		m int = len(a)
		n int = len(b)

		i, j, p, pMod int = 0, 0, n + 1, 0
	)

	if len(l.cost) < 2*(n+1) {
		l.cost = make([]int, 2*(n+1))
	}

	for i = 0; i < n+1; i++ {
		l.cost[i] = i
	}

	for j = 1; j < m+1; j++ {
		l.cost[p] = j

		for i = 1; i < n+1; i++ {

			if a[j-1] == b[i-1] {
				l.cost[p+i] = l.cost[pMod+i-1]
			} else {
				// DELETION
				d = l.cost[p+i-1] + l.Delete

				// INSERTION
				if l.cost[pMod+i]+1 < d {
					d = l.cost[pMod+i] + l.Insert
				}

				// SUBSTITUTION
				if l.cost[pMod+i-1]+1 < d {
					d = l.cost[pMod+i-1] + l.Substitution
				}

				l.cost[p+i] = d
			}
		}

		if p == 0 {
			p = n + 1
			pMod = 0
		} else {
			pMod = n + 1
			p = 0
		}
	}

	d = l.cost[pMod+n]

	return
}
