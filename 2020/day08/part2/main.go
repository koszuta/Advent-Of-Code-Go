package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 8: Handheld Halting ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/8#part2
 */

type instruction struct {
	op  string
	arg int
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// Parse the instructions
	instructions := make([]instruction, 0, 0)
	for scanner.Scan() {
		// Instructions are composed of an operation and an argument
		// Operations: acc, jmp, nop
		// Arguments: signed integers
		line := scanner.Text()
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, instruction{op, arg})
	}

	// Exactly one nop/jmp operation is incorrect in the program
	// We need to flip it to get the program to terminate
	// We'll iterate the instructions, flip an op, and see if it fixes the infinite loop
	for i := 0; i < len(instructions); i++ {

		// Skip acc ops since they're all correct
		if instructions[i].op != "acc" {
			// Flip nop to jmp and vice versa
			if instructions[i].op == "nop" {
				instructions[i].op = "jmp"
			} else {
				instructions[i].op = "nop"
			}

			acc, pc := 0, 0
			seen := make(map[int]struct{})
			for pc < len(instructions) {
				_, alreadySeen := seen[pc]
				if alreadySeen {
					break // stop the program; it's in a loop
				}

				instr := instructions[pc]
				seen[pc] = struct{}{}
				switch instr.op {
				case "nop": // increment the program counter
					pc++
				case "acc": // increase or decrease the accumulator by the argument
					acc += instr.arg
					pc++
				case "jmp": // offset the program counter by the argument
					pc += instr.arg
				}
			}

			// Check if the program terminated
			if pc == len(instructions) {
				log.Printf("the value of the accumulator was %d when the program terminated\n", acc)
				break
			}

			// Reset the instructions
			if instructions[i].op == "nop" {
				instructions[i].op = "jmp"
			} else {
				instructions[i].op = "nop"
			}
		}
	}
}
