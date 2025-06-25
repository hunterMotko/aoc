import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-7'


test("Part ONE EX1: ", async () => {
	const res = await partOne('day-7/ex1.txt')
	expect(res).toBe(3749)
})

test("Part ONE IN: ", async () => {
	const res = await partOne('day-7/in.txt')
	expect(res).toBe(2314935962622)
})

test("Part TWO EX1: ", async () => {
	const res = await partTwo('day-7/ex1.txt')
	expect(res).toBe(11387)
})

test("Part Two IN: ", async () => {
	const res = await partTwo('day-7/in.txt')
	expect(res).toBe(401477450831495)
})
