import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-6'


test("Part ONE EX1: ", async () => {
	const res = await partOne('day-6/ex1.txt')
	expect(res).toBe(41)
})

test("Part ONE IN: ", async () => {
	const res = await partOne('day-6/in.txt')
	expect(res).toBe(5162)
})

test("Part TWO EX1: ", async () => {
	const res = await partTwo('day-6/ex1.txt')
	expect(res).toBe(6)
})

test("Part Two IN: ", async () => {
	const res = await partTwo('day-6/in.txt')
	expect(res).toBe(1909)
})
