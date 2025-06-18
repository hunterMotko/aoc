import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-3'

test("PART 1", async () => {
	let res = await partOne('day-3/ex1.txt')
	expect(res).toBe(161)
	let in_res = await partOne('day-3/in.txt')
	expect(in_res).toBe(184511516)
})

test("PART 2", async () => {
	let res = await partTwo('day-3/ex2.txt')
	expect(res).toBe(48)

	let in_res = await partTwo('day-3/in.txt')
	expect(in_res).toBe(90044227)
})
