import { readFileNums } from "../../utils/utils"

export async function partOne(filePath: string): Promise<number> {
	let nums: number[] = []
	try {
		nums = await readFileNums(filePath)
	} catch (err) {
		console.error
	}

	let diskMap = createDiskMap(nums)
	moveFileBlocks(diskMap)
	return fileSystemChecksum(diskMap)
}

export async function partTwo(filePath: string): Promise<number> {
	let nums: number[] = []
	try {
		nums = await readFileNums(filePath)
	} catch (err) {
		console.error
	}

	let diskMap = createDiskMap(nums)
	console.log(diskMap)
	slidingFileBlocks(diskMap)
	console.log(diskMap)
	return fileSystemChecksum(diskMap)
}

function createDiskMap(nums: number[]): number[] {
	let diskMap: number[] = []
	let j = 0;
	for (let i = 0; i < nums.length; i++) {
		if (i % 2 === 1) {
			for (let k = 0; k < +nums[i]; k++) {
				diskMap.push(-1)
			}
			continue
		}
		for (let k = 0; k < +nums[i]; k++) {
			diskMap.push(j)
		}
		j++
	}
	return diskMap
}

function moveFileBlocks(diskMap: number[]) {
	let i = 0
	let j = diskMap.length - 1
	while (i <= j) {
		let left = diskMap[i]
		let right = diskMap[j]
		if (left === -1 && right !== -1) {
			diskMap[i] = right
			diskMap[j] = left
			i++
			j--
		}
		if (left !== -1) i++
		if (right === -1) j--
	}
}

function slidingFileBlocks(diskMap: number[]) {
	for (let i = diskMap.length - 1; i >= 0;) {
		if (diskMap[i] === -1) {
			i--
			continue
		}
		let len = findLength(diskMap, i, false)
		for (let j = 0; j < i; j++) {
			if (diskMap[j] !== -1) continue
			let l = findLength(diskMap, j, true)
			if (l >= len) {
				for (let k = 0; k < len; k++) {
					let n = diskMap[i - k]
					diskMap[i - k] = diskMap[j + k]
					diskMap[j + k] = n
				}
				break
			}
		}
		i = i - len
	}
}

function findLength(diskMap: number[], i: number, asc: boolean) {
	let count = 1;
	let j = asc ? i + 1 : i - 1;
	const targetValue = diskMap[i]; // The value we are looking for a contiguous block of
	// Ensure j stays within bounds while checking
	while (j >= 0 && j < diskMap.length && diskMap[j] === targetValue) {
		count++;
		asc ? j++ : j--;
	}
	return count;
}

function fileSystemChecksum(diskMap: number[]): number {
	let res = 0
	for (let i = 0; i < diskMap.length; i++) {
		let cur = diskMap[i]
		if (cur && cur !== -1) {
			res += i * cur
		}
	}
	return res
}
