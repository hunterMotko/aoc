import { expect, test } from 'bun:test'
import { partOne, partTwo } from 'day-11'


// test("Part ONE EX1: ", async () => {
// 	const res = await partOne('day-11/ex1.txt', 6)
// 	expect(res).toBe(22)
// })
//
// test("Part ONE EX1: ", async () => {
// 	const res = await partOne('day-11/ex1.txt', 25)
// 	expect(res).toBe(55312)
// })
//
// test("Part ONE IN: ", async () => {
// 	const res = await partOne('day-11/in.txt', 25)
// 	expect(res).toBe(183435)
// })


test("Part Two IN: ", async () => {
	const res = await partTwo('day-11/in.txt', 75)
	expect(res).toBe(218279375708592)
})
