import { readFileStr } from "../../utils/utils"

export async function partOne(filePath: string): Promise<number> {
	let lines: string[] = []
	try {
		lines = await readFileStr(filePath)
	} catch (err) {
		throw err
	}
	let res = 0
	for (const line of lines) {
		const eq = line.split(': ')
		if (eq.length !== 2) break
		const digits = eq[1]?.split(' ').map(Number)
		if (!digits || digits.length === 0) break
		let ans = Number(eq[0])

		if (checkEquation(digits, ans, 0, 0, "+", false)) {
			res += ans
		}
	}

	return res
}

function checkEquation(digits: number[], ans: number, n: number, i: number, op: string, two: boolean): boolean {
	if (i >= digits.length) {
		return n === ans
	}

	let cur = digits[i]
	if (cur === undefined) {
		return false
	}
	if (op === "+") {
		n += cur
	} else if (op === "*") {
		n *= cur
	} else {
		n = parseInt(`${n}${cur}`)
	}
	let ops: string[]
	if (two) {
		ops = ["+", "*", "|"]
	} else {
		ops = ["+", "*"]
	}
	for (let o of ops) {
		if (checkEquation(digits, ans, n, i + 1, o, two)) {
			return true
		}
	}
	return false
}

export async function partTwo(filePath: string): Promise<number> {
	let lines: string[] = []
	try {
		lines = await readFileStr(filePath)
	} catch (err) {
		throw err
	}
	let res = 0
	for (const line of lines) {
		const eq = line.split(': ')
		if (eq.length !== 2) break
		let ans = Number(eq[0])
		const digits = eq[1]?.split(' ').map(Number)
		if (!digits || digits.length === 0) break

		if (checkEquation(digits, ans, 0, 0, "+", true)) {
			res += ans
		}
	}

	return res
}

