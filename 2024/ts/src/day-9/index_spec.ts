//
import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-9'


test("Part ONE EX1: ", async () => {
	const res = await partOne('day-9/ex1.txt')
	expect(res).toBe(1928)
})

test("Part ONE IN: ", async () => {
	const res = await partOne('day-9/in.txt')
	expect(res).toBe(6262891638328)
})

test("Part TWO EX1: ", async () => {
	const res = await partTwo('day-9/ex1.txt')
	expect(res).toBe(2858)
})

test("Part Two IN: ", async () => {
	const res = await partTwo('day-9/in.txt')
	expect(res).toBe(6287317016845)
})
