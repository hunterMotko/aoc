import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-1'

test("PART 1", async () => {
	let res = await partOne('day-1/ex1.txt')
	expect(res).toBe(11)

	let in_res = await partOne('day-1/in.txt')
	expect(in_res).toBe(1579939)
})

test("PART 2", async () => {
	let res = await partTwo('day-1/ex1.txt')
	expect(res).toBe(31)

	let in_res = await partTwo('day-1/in.txt')
	expect(in_res).toBe(20351745)
})
