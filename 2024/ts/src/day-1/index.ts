import { file } from "bun";
import { readFileToMatrix } from "../../utils/utils"

export async function partOne(path: string) {
	const fileContents: string[][] = await readFileToMatrix(path, '  ')
	if (!fileContents.length) {
		return 0
	}

	const left: number[] = fileContents
		.map(item => item[0])
		.filter((item): item is string => item !== "")
		.map(item => parseInt(item))
		.sort((a, b) => a - b);

	const right: number[] = fileContents
		.map(item => item[1])
		.filter((item): item is string => item !== "")
		.map(item => parseInt(item))
		.sort((a, b) => a - b);

	let res = 0
	for (let i = 0; i < left.length; i++) {
		res += Math.abs(right[i] - left[i])
	}

	return res
}

// export async function partTwo(path: string) {
// 	const fileContents: string[][] = await readFileToMatrix(path, '  ')
// 	if (!fileContents.length) {
// 		return 0
// 	}
// 	let occ: number[][] = []
// 	for (let i = 0; i < fileContents.length; i++) {
// 		let cur = fileContents[i][0] as string
// 		let count = 0
// 		for (let j = 0; j < fileContents.length; j++) {
// 			if (+cur === +fileContents[j][1]) {
// 				count++
// 			}
// 		}
// 		occ.push([+cur, count])
// 	}
// 	let res = 0
// 	for (let i of occ) {
// 		res += i[0] * i[1]
// 	}
// 	return res
// }

export async function partTwo(path: string) {
	const fileContents: string[][] = await readFileToMatrix(path, '  ')
	if (!fileContents.length) {
		return 0
	}
	const counts = new Map<number, number>();
	// First pass: Populate the map with counts of values in fileContents[j][1]
	for (let i = 0; i < fileContents.length; i++) {
		const num = +fileContents[i][1]; // Coerce once
		counts.set(num, (counts.get(num) || 0) + 1);
	}

	let res = 0;
	// Second pass: Calculate res using the counts
	for (let i = 0; i < fileContents.length; i++) {
		const cur = +fileContents[i][0]; // Coerce once
		const count = counts.get(cur) || 0; // Get count from the map
		res += cur * count;
	}
	return res;
}
