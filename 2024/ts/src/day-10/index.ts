import { readFileToMatrixNums } from "../../utils/utils"

type Visit = {
	[key: string]: string[]
}
export async function partOne(filePath: string): Promise<number> {
	let matrix: number[][] = []
	try {
		matrix = await readFileToMatrixNums(filePath, '')
	} catch (err) {
		console.error(err)
	}

	let visit: Visit = {}
	for (let i = 0; i < matrix.length; i++) {
		for (let j = 0; j < matrix[i].length; j++) {
			if (matrix[i][j] === 0) {
				let k = `${i},${j}`
				visit[k] = []
				dfs(matrix, i, j, 0, k, visit, false)
			}
		}
	}

	let res = 0
	Object.entries(visit).forEach(([k, v]: [k: string, v: string[]]) => {
		res += v.length
	});
	return res
}

function dfs(
	matrix: number[][],
	x: number,
	y: number,
	c: number,
	k: string,
	visit: Visit,
	two: boolean
) {
	if (x < 0 || y < 0 || x > matrix.length - 1 || y > matrix[0].length - 1) {
		return
	}
	if (matrix[x][y] !== c) {
		return
	}

	if (matrix[x][y] === 9 && c === 9) {
		let newKey = `${x},${y}`

		if (!two && visit[k]?.includes(newKey)) {
			return
		}

		visit[k].push(newKey)
		return
	}

	dfs(matrix, x, y - 1, c + 1, k, visit, two)
	dfs(matrix, x, y + 1, c + 1, k, visit, two)
	dfs(matrix, x - 1, y, c + 1, k, visit, two)
	dfs(matrix, x + 1, y, c + 1, k, visit, two)
}

export async function partTwo(filePath: string): Promise<number> {
	let matrix: number[][] = []
	try {
		matrix = await readFileToMatrixNums(filePath, '')
	} catch (err) {
		console.error(err)
	}

	let visit: Visit = {}
	for (let i = 0; i < matrix.length; i++) {
		for (let j = 0; j < matrix[i].length; j++) {
			if (matrix[i][j] === 0) {
				let k = `${i},${j}`
				visit[k] = []
				dfs(matrix, i, j, 0, k, visit, true)
			}
		}
	}

	let res = 0
	Object.entries(visit).forEach(([k, v]: [k: string, v: string[]]) => {
		res += v.length
	});
	return res
}
