import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-10'


test("Part ONE EX1: ", async () => {
	const res = await partOne('day-10/ex1.txt')
	expect(res).toBe(1)
})

test("Part ONE EX1: ", async () => {
	const res = await partOne('day-10/ex2.txt')
	expect(res).toBe(36)
})

test("Part ONE IN: ", async () => {
	const res = await partOne('day-10/in.txt')
	expect(res).toBe(682)
})

test("Part TWO EX1: ", async () => {
	const res = await partTwo('day-10/ex2.txt')
	expect(res).toBe(81)
})

test("Part Two IN: ", async () => {
	const res = await partTwo('day-10/in.txt')
	expect(res).toBe(1511)
})
