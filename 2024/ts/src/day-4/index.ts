import { readFileToMatrix } from "../../utils/utils"


const dirs = [
	[0, 1],   // Horizontal right
	[0, -1],  // Horizontal left
	[1, 0],   // Vertical down
	[-1, 0],  // Vertical up
	[1, 1],   // Diagonal down-right
	[1, -1],  // Diagonal down-left
	[-1, 1],  // Diagonal up-right
	[-1, -1]  // Diagonal up-left
]

const word = "XMAS"

export async function partOne(path: string): Promise<number> {
	let grid = await readFileToMatrix(path, '')
	let rows = grid.length
	let cols = grid[0].length
	let res = 0

	for (let i = 0; i < rows; i++) {
		for (let j = 0; j < cols; j++) {
			if (grid[i][j] === "X") {
				res += walk(grid, i, j, rows, cols)
			}
		}
	}
	return res
}

function walk(grid: string[][], i: number, j: number, rows: number, cols: number) {
	let res = 0
	for (const [x, y] of dirs) {
		let ci = i
		let cj = j
		let match = true
		for (let k = 1; k < word.length; k++) {
			if (ci + x >= 0 && ci + x < rows && cj + y >= 0 && cj + y < cols) {
				if (grid[ci + x][cj + y] === word[k]) {
					ci += x
					cj += y
				} else {
					match = false
					break
				}
			} else {
				match = false
				break
			}
		}

		if (match) {
			res++
		}
	}
	return res
}


export async function partTwo(path: string): Promise<number> {
	let grid = await readFileToMatrix(path, '')
	let rows = grid.length
	let cols = grid[0].length
	let res = 0

	for (let i = 0; i < rows; i++) {
		for (let j = 0; j < cols; j++) {
			if (grid[i][j] === "A") {
				res += isX(grid, i, j, rows, cols)
			}
		}
	}
	return res
}

function isX(grid: string[][], i: number, j: number, rows: number, cols: number): number {
	if (!isDiagonal(grid, i - 1, j - 1, i + 1, j + 1, rows, cols)) return 0
	if (!isDiagonal(grid, i - 1, j + 1, i + 1, j - 1, rows, cols)) return 0
	return 1
}

function isDiagonal(
	grid: string[][],
	ax: number,
	ay: number,
	bx: number,
	by: number,
	rows: number,
	cols: number
): boolean {
	if (ax < 0 || ax >= rows || ay < 0 || ay >= cols) {
		return false
	}
	if (bx < 0 || bx >= rows || by < 0 || by >= cols) {
		return false
	}

	let a = grid[ax][ay]
	let b = grid[bx][by]

	if (a === "M" && b === "S") return true
	if (a === "S" && b === "M") return true

	return false
}
