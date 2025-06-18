import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-4'

test("PART 1", async () => {
	let res = await partOne('day-4/ex1.txt')
	expect(res).toBe(18)
	let in_res = await partOne('day-4/in.txt')
	expect(in_res).toBe(2504)
})

test("PART 2", async () => {
	let res = await partTwo('day-4/ex1.txt')
	expect(res).toBe(9)
	let in_res = await partTwo('day-4/in.txt')
	expect(in_res).toBe(1923)
})
