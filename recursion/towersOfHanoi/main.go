/*
	PROBLEM 1:
	The Towers of Hanoi is a mathematical puzzle. It consists of three rods (or pegs or
	towers) and a number of disks of different sizes which can slide onto any rod. The puzzle starts
	with the disks on one rod in ascending order of size, the smallest at the top, thus making a conical
	shape.
	The objective of the puzzle is to move the entire stack to another rod, satisfying the following rules:
	• Only one disk may be moved at a time.
	• Each move consists of taking the upper disk from one of the rods and sliding it onto
	another rod, on top of the other disks that may already be present on that rod.
	• No disk may be placed on top of a smaller disk.


	Algorithm
	• Move the top n – 1 disks from Source to Auxiliary tower,
	• Move the n
	th disk from Source to Destination tower,
	• Move the n – 1 disks from Auxiliary tower to Destination tower.
	• Transferring the top n – 1 disks from Source to Auxiliary tower can again be thought
	of as a fresh problem and can be solved in the same manner. Once we solve Towers
	of Hanoi with three disks, we can solve it with any number of disks with the above
	algorithm.
*/

package main

import "fmt"

func TowerOfHanoi(n int, fromPeg, toPeg, auxPeg string) {
	if n == 1 {
		fmt.Printf("From : %v To : %v \n", fromPeg, toPeg)
		return
	}
	TowerOfHanoi(n-1, fromPeg, auxPeg, toPeg)
	fmt.Printf("From : %v To : %v \n", fromPeg, toPeg)
	TowerOfHanoi(n-1, auxPeg, toPeg, fromPeg)
}

func main() {
	TowerOfHanoi(5, "A", "B", "C")
}
