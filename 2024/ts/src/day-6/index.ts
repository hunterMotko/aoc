import { isTypeOnlyImportOrExportDeclaration } from "typescript";
import { readFileToMatrix } from "../../utils/utils";

export async function partOne(filePath: string): Promise<number> {
	try {
		const grid = await readFileToMatrix(filePath, '')
		const start = findStartPos(grid)
		if (!start) {
			throw new Error("start not found");

		}
		return Object.keys(traverseGrid(grid, start)).length
	} catch (err) {
		console.error(err);
		throw err;
	}
}

export async function partTwo(filePath: string): Promise<number> {
	try {
		const grid = await readFileToMatrix(filePath, '')
		const start = findStartPos(grid)
		if (!start) {
			throw new Error("start not found");
		}

		return traverseTwo(grid, start)
	} catch (err) {
		console.error(err);
		throw err;
	}
}

type Pos = {
	x: number;
	y: number;
	dir: string;
}

function findStartPos(grid: string[][]): Pos | null {
	for (let i = 0; i < grid.length; i++) {
		for (let j = 0; j < grid.length; j++) {
			let cur = grid[i][j]
			if (cur === "^" || cur === ">" || cur === "<" || cur === "v") {
				return { x: i, y: j, dir: cur }
			}
		}
	}
	return null
}

type Directions = {
	[key: string]: [number, number]
}

type Visited = {
	[key: string]: boolean
}

const dirs: Directions = {
	"^": [-1, 0],
	">": [0, 1],
	"v": [1, 0],
	"<": [0, -1]
}

const turnDir: { [key: string]: string } = {
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

function traverseGrid(grid: string[][], start: Pos): Visited {
	let { x, y, dir } = start;
	let visited: Visited = {
		[`${x},${y}`]: true
	};
	let rows = grid.length;
	let cols = grid[0].length;
	while (true) {
		const [i, j] = dirs[dir]
		let nx = x + i
		let ny = y + j
		if (nx >= 0 && nx < rows && ny >= 0 && ny < cols) {
			let nextMove = grid[nx][ny]
			if (nextMove === '#') {
				dir = turnDir[dir] as string
				continue
			}
			let key = `${nx},${ny}`
			if (!visited[key]) {
				visited[key] = true
			}
			x = nx
			y = ny
		} else {
			break
		}
	}
	return visited
}

function traverseTwo(grid: string[][], start: Pos): number {
	let pos = start
	let { x, y, dir } = pos;

	let visited: Visited = {
		[`${x},${y}`]: true
	};
	let rows = grid.length;
	let cols = grid[0].length;
	while (true) {
		const [i, j] = dirs[dir]
		let nx = x + i
		let ny = y + j
		if (!(nx >= 0 && nx < rows && ny >= 0 && ny < cols)) {
			break
		}

		let nextMove = grid[nx][ny]
		if (nextMove === '#') {
			dir = turnDir[dir] as string
			continue
		}
		let key = `${nx},${ny}`
		if (!visited[key]) {
			visited[key] = true
		}
		x = nx
		y = ny
	}

	let loops = 0
	for (let key in visited) {
		let curPos = key.split(',')
		grid[+curPos[0]][+curPos[1]] = '#'
		if (checkLoop(grid, start)) {
			loops++
		}
		grid[+curPos[0]][+curPos[1]] = '.'
	}

	return loops
}

function checkLoop(grid, start) {
	let { x, y, dir } = start
	let visited: Visited = {
		[`${x},${y},${dir}`]: true
	}
	let rows = grid.length;
	let cols = grid[0].length;
	let v = false
	while (true) {
		const [i, j] = dirs[dir]
		let nx = x + i
		let ny = y + j
		if (!(nx >= 0 && nx < rows && ny >= 0 && ny < cols)) {
			break
		}
		if (visited[`${nx},${ny},${dir}`]) {
			v = true
			break
		}
		let nextMove = grid[nx][ny]
		if (nextMove === '#') {
			dir = turnDir[dir] as string
			continue
		}
		let key = `${nx},${ny},${dir}`
		if (!visited[key]) {
			visited[key] = true
		}
		x = nx
		y = ny
	}
	return v
}

