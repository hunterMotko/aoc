import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-2'

test("PART 1: EX", async () => {
	let res = await partOne('day-2/ex1.txt')
	expect(res).toBe(2)
})

test("PART 1: IN", async () => {
	let in_res = await partOne('day-2/in.txt')
	expect(in_res).toBe(442)
})

test("PART 2: EX", async () => {
	let res = await partTwo('day-2/ex1.txt')
	expect(res).toBe(4)
})

test("PART 2: IN", async () => {
	let in_res = await partTwo('day-2/in.txt')
	expect(in_res).toBe(493)
})
