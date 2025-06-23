import { expect, test } from 'bun:test'
import { partOne, partTwo, partTwoHalf } from 'day-5'

test("PART 1: EX", async () => {
	let res = await partOne('day-5/ex1.txt')
	expect(res).toBe(143)
})

test("PART 1: IN", async () => {
	let in_res = await partOne('day-5/in.txt')
	expect(in_res).toBe(5091)
})

test("PART 2: EX", async () => {
	let res = await partTwo('day-5/ex1.txt')
	expect(res).toBe(123)
})

test("PART 2: IN", async () => {
	let in_res = await partTwo('day-5/in.txt')
	expect(in_res).toBe(4681)
})

test("PART 2-1/2: EX", async () => {
	let res = await partTwoHalf('day-5/ex1.txt')
	expect(res).toBe(123)
})

test("PART 2-1/2: IN", async () => {
	let in_res = await partTwoHalf('day-5/in.txt')
	expect(in_res).toBe(4681)
})
