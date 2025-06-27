import { readFileToMatrix } from "../../utils/utils"

type Loc = {
	i: number;
	j: number;
}

type Ant = {
	[key: string]: Loc[]
}

export async function partOne(filePath: string): Promise<number> {
	let lines: string[][] = []
	try {
		lines = await readFileToMatrix(filePath, '')
	} catch (err) {
		console.error(err)
	}
	let antennas = getAntennaLocs(lines)
	let antiNodes = getAntiNodes(antennas, lines.length, lines[0]?.length, false)
	return antiNodes.size
}

function getAntennaLocs(lines: string[][]): Ant {
	let ants: Ant = {}
	for (let i = 0; i < lines.length; i++) {
		for (let j = 0; j < lines[i].length; j++) {
			let cur = lines[i][j]
			if (cur && cur !== '.') {
				if (!ants[cur]) {
					ants[cur] = [{ i, j }]
					continue
				}
				ants[cur].push({ i, j })
			}
		}
	}
	return ants
}

function getAntiNodes(ants: Ant, rows: number, cols: number, two: boolean): Set<number> {
	let antiNodes = new Set<number>()
	for (const [k, locs] of Object.entries(ants)) {
		let l = locs.length
		for (let i = 0; i < l - 1; i++) {
			for (let j = i + 1; j < l; j++) {
				let di = locs[j].i - locs[i].i
				let dj = locs[j].j - locs[i].j
				let x = locs[i].i - di
				let y = locs[i].j - dj

				while (x >= 0 && x < rows && y >= 0 && y < cols) {
					antiNodes.add(x * rows + y)
					if (!two) break
					x -= di
					y -= dj
				}

				x = locs[j].i + di
				y = locs[j].j + dj
				while (x >= 0 && x < rows && y >= 0 && y < cols) {
					antiNodes.add(x * rows + y)
					if (!two) break
					x += di
					y += dj
				}
				if (two) {
					antiNodes.add(locs[i].i * rows + locs[i].j)
					antiNodes.add(locs[j].i * rows + locs[j].j)
				}
			}
		}
	}
	return antiNodes
}

export async function partTwo(filePath: string): Promise<number> {
	let lines: string[][] = []
	try {
		lines = await readFileToMatrix(filePath, '')
	} catch (err) {
		console.error(err)
	}
	let antennas = getAntennaLocs(lines)
	let antiNodes = getAntiNodes(antennas, lines.length, lines[0]?.length, true)
	return antiNodes.size
}
