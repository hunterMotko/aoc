import { readFileStr } from "../../utils/utils";

// export async function partOne(path: string) {
// 	const contents = await readFileStr(path)
// 	let total = 0
// 	for (let line of contents) {
// 		let matches = line.matchAll(/(mul\(\d+,\d+\))/g)
// 		let arr = Array.from(matches)
// 		for (let match of arr) {
// 			let ints = match[0].split(/[mul(,)]/).filter(i => i !== "")
// 			if (ints.length === 2) {
// 				total += +ints[0] * +ints[1]
// 			}
// 		}
// 	}
// 	return total
// }

export async function partOne(path: string): Promise<number> {
	const contents = await readFileStr(path);
	let total = 0;
	const mulRegex = /mul\((\d+),(\d+)\)/g;
	for (const line of contents) {
		for (const match of line.matchAll(mulRegex)) {
			const num1 = parseInt(match[1], 10);
			const num2 = parseInt(match[2], 10);
			total += num1 * num2;
		}
	}
	return total;
}

export async function partTwo(path: string) {
	const contents = await readFileStr(path);
	let total = 0;
	const op = /(?:mul\((\d+),(\d+)\))|do\(\)|don't\(\)/g;
	const mulRegex = /mul\((\d+),(\d+)\)/g;
	let doCheck = true
	for (const line of contents) {
		for (const match of line.matchAll(op)) {
			if (match[0] === 'do()') {
				doCheck = true
			} else if (match[0] === 'don\'t()') {
				doCheck = false
			}
			if (match[0].includes('mul') && doCheck) {
				const num1 = parseInt(match[1], 10);
				const num2 = parseInt(match[2], 10);
				total += num1 * num2;
				continue
			}
		}
	}
	return total;
}
