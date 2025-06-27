import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-8'


test("Part ONE EX1: ", async () => {
	const res = await partOne('day-8/ex1.txt')
	expect(res).toBe(14)
})

test("Part ONE IN: ", async () => {
	const res = await partOne('day-8/in.txt')
	expect(res).toBe(265)
})

test("Part TWO EX1: ", async () => {
	const res = await partTwo('day-8/ex1.txt')
	expect(res).toBe(34)
})

test("Part Two IN: ", async () => {
	const res = await partTwo('day-8/in.txt')
	expect(res).toBe(962)
})
