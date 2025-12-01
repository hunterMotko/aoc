import { readFileStr } from "../../utils/utils"

// If the stone is engraved with the number 0, it is replaced by a stone 
// engraved with the number 1.
//
// If the stone is engraved with a number that has an even number of digits, 
// it is replaced by two stones. The left half of the digits are engraved on 
// the new left stone, and the right half of the digits are engraved on the 
// new right stone. (The new numbers don't keep extra leading zeroes: 1000 
// would become stones 10 and 0.)
//
// If none of the other rules apply, the stone is replaced by a new stone; 
// the old stone's number multiplied by 2024 is engraved on the new stone.
export async function partOne(filePath: string, blinks: number): Promise<number> {
	let stones: string[] = []
	try {
		stones = await readFileStr(filePath)
		stones = stones[0].split(' ')
	} catch (err) {
		console.error(err)
	}

	for (let i = 0; i < blinks; i++) {
		stones = magic(stones)
	}

	return stones.length
}

function magic(stones: string[]): string[] {
	let temp: string[] = []
	for (let i of stones) {
		if (parseInt(i) === 0) {
			temp.push('1')
		} else if (i.length % 2 === 0) {
			let mid = i.length / 2
			let left = i.slice(0, mid)
			let right = i.slice(mid, i.length)
			let p = Number(right) + ''
			temp.push(left, p)
		} else {
			let t = parseInt(i) * 2024
			temp.push(t + '')
		}
	}

	return temp
}

function memoization(stones: string[], blink: number): number {
	const memos = Array.from({ length: blink + 1 }).map(() => new Map<string, number>());

	function transform(stone: string, remainingBlink: number): number {
		if (remainingBlink === 0) {
			return 1;
		}

		const memoized = memos[remainingBlink].get(stone);
		if (memoized) return memoized;

		const transformed: string[] = [];
		if (stone === "0") {
			transformed.push("1");
		} else if (stone.length % 2 === 0) {
			transformed.push(
				stone.slice(0, stone.length / 2),
				parseInt(stone.slice(stone.length / 2)).toString()
			);
		} else {
			transformed.push((parseInt(stone) * 2024).toString());
		}

		const res = transformed.reduce(
			(sum, newStone) => sum + transform(newStone, remainingBlink - 1),
			0
		);

		memos[remainingBlink].set(stone, res);

		return res;
	}

	let count = 0;
	stones.forEach((stone) => (count += transform(stone, blink)));

	return count;
}

export async function partTwo(filePath: string, blinks: number): Promise<number> {
	let stones: string[] = []
	try {
		stones = await readFileStr(filePath)
		stones = stones[0].split(' ')
	} catch (err) {
		console.error(err)
	}

	return memoization(stones, blinks)
}

