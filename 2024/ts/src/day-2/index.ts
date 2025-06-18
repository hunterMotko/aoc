import { readFileToMatrix } from "../../utils/utils";

export async function partOne(path: string) {
	const fileCon: string[][] = await readFileToMatrix(path, ' ')
	let count = 0
	for (let i = 0; i < fileCon.length; i++) {
		let nArr = fileCon[i].map(Number)
		if (isSafe(nArr)) {
			count++
		}
	}

	return count
}

export async function partTwo(path: string) {
	const fileCon: string[][] = await readFileToMatrix(path, ' ')
	let count = 0
	let del = 0
	for (let i = 0; i < fileCon.length; i++) {
		let nArr = fileCon[i].map(Number)
		if (isSafe(nArr)) {
			count++
		} else if (withDeletion(nArr)) {
			del++
		}
	}

	console.log(count, del)
	return count + del
}

function withDeletion(arr: number[]): boolean {
	for (let i = 0; i < arr.length; i++) {
		if (useDelete(arr, i)) {
			return true
		}
	}
	return false
}

function useDelete(arr: number[], i: number): boolean {
	if (i === arr.length - 1) {
		return isSafe(arr.slice(0, i));
	}
	return isSafe([...arr.slice(0, i), ...arr.slice(i + 1)])
}

function isSafe(arr: number[]): boolean {
	let asc = true

	for (let i = 0; i < arr.length - 1; i++) {
		let cur = arr[i] as number, next = arr[i + 1] as number
		if (i === 0 && cur > next) {
			asc = false
		}
		if (asc) {
			let dif = next - cur
			if (cur > next || dif < 1 || dif > 3) {
				return false
			}
		} else {
			let dif = cur - next
			if (cur < next || dif < 1 || dif > 3) {
				return false
			}
		}
	}

	return true
}
