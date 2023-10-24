package towersofhanoi_test

import (
	towersofhanoi "godsa/recursion-backtracking/towersOfHanoi"
	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	t.Run("Tower of Hanoi with 1 disk", func(t *testing.T) {
		n := 1
		fromPeg, toPeg, auxPeg := "A", "C", "B"
		expectedMoves := []string{"A->C"}

		moves := towersofhanoi.TowerOfHanoi(n, fromPeg, toPeg, auxPeg)

		for i := 0; i < len(moves); i++ {
			if moves[i] != expectedMoves[i] {
				t.Errorf("Move %d is incorrect. Expected: %s, Got: %s", i, expectedMoves[i], moves[i])
			}
		}
	})

	t.Run("Tower of Hanoi with 2 disks", func(t *testing.T) {
		n := 2
		fromPeg, toPeg, auxPeg := "A", "C", "B"
		expectedMoves := []string{
			"A->B",
			"A->C",
			"B->C",
		}

		moves := towersofhanoi.TowerOfHanoi(n, fromPeg, toPeg, auxPeg)

		for i := 0; i < len(moves); i++ {
			if moves[i] != expectedMoves[i] {
				t.Errorf("Move %d is incorrect. Expected: %s, Got: %s", i, expectedMoves[i], moves[i])
			}
		}
	})

	t.Run("Tower of Hanoi with 3 disks", func(t *testing.T) {
		n := 3
		fromPeg, toPeg, auxPeg := "A", "C", "B"
		expectedMoves := []string{
			"A->C",
			"A->B",
			"C->B",
			"A->C",
			"B->A",
			"B->C",
			"A->C",
		}

		moves := towersofhanoi.TowerOfHanoi(n, fromPeg, toPeg, auxPeg)

		for i := 0; i < len(moves); i++ {
			if moves[i] != expectedMoves[i] {
				t.Errorf("Move %d is incorrect. Expected: %s, Got: %s", i, expectedMoves[i], moves[i])
			}
		}
	})

	t.Run("Tower of Hanoi with 4 disks", func(t *testing.T) {
		n := 4
		fromPeg, toPeg, auxPeg := "A", "C", "B"
		expectedMoves := []string{
			"A->B",
			"A->C",
			"B->C",
			"A->B",
			"C->A",
			"C->B",
			"A->B",
			"A->C",
			"B->C",
			"B->A",
			"C->A",
			"B->C",
			"A->B",
			"A->C",
			"B->C",
		}

		moves := towersofhanoi.TowerOfHanoi(n, fromPeg, toPeg, auxPeg)

		for i := 0; i < len(moves); i++ {
			if moves[i] != expectedMoves[i] {
				t.Errorf("Move %d is incorrect. Expected: %s, Got: %s", i, expectedMoves[i], moves[i])
			}
		}
	})

	t.Run("Tower of Hanoi with 5 disks", func(t *testing.T) {
		n := 5
		fromPeg, toPeg, auxPeg := "A", "C", "B"
		expectedMoves := []string{
			"A->C",
			"A->B",
			"C->B",
			"A->C",
			"B->A",
			"B->C",
			"A->C",
			"A->B",
			"C->B",
			"C->A",
			"B->A",
			"C->B",
			"A->C",
			"A->B",
			"C->B",
			"A->C",
			"B->A",
			"B->C",
			"A->C",
			"B->A",
			"C->B",
			"C->A",
			"B->A",
			"B->C",
			"A->C",
			"A->B",
			"C->B",
			"A->C",
			"B->A",
			"B->C",
			"A->C",
		}

		moves := towersofhanoi.TowerOfHanoi(n, fromPeg, toPeg, auxPeg)

		for i := 0; i < len(moves); i++ {
			if moves[i] != expectedMoves[i] {
				t.Errorf("Move %d is incorrect. Expected: %s, Got: %s", i, expectedMoves[i], moves[i])
			}
		}
	})
}
