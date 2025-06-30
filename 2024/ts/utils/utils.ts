import { readFile } from 'fs/promises'
import path from 'path'

export async function readFileToMatrix(filePath: string, sep = " "): Promise<string[][]> {
	const result: string[][] = []
	try {
		const fp = path.join(process.cwd(), 'src', filePath)
		const file = await readFile(fp, { encoding: 'utf8' })
		const lines = file.split('\n')
		for (const line of lines) {
			if (line !== "") {
				result.push(line.split(sep))
			}
		}
	} catch (err) {
		throw err
	}
	return result
}

export async function readFileToMatrixNums(filePath: string, sep = " "): Promise<number[][]> {
	const result: number[][] = []
	try {
		const fp = path.join(process.cwd(), 'src', filePath)
		const file = await readFile(fp, { encoding: 'utf8' })
		const lines = file.split('\n')
		for (const line of lines) {
			if (line !== "") {
				let row = []
				for (const ch of line.split(sep)) {
					row.push(parseInt(ch))
				}
				result.push(row)
			}
		}
	} catch (err) {
		throw err
	}
	return result
}

export async function readFileStr(filePath: string): Promise<string[]> {
	try {
		const fp = path.join(process.cwd(), 'src', filePath)
		const file = await readFile(fp, { encoding: 'utf8' })
		return file.split('\n')
	} catch (err) {
		throw err
	}
}

export async function readFileNums(filePath: string): Promise<number[]> {
	try {
		const fp = path.join(process.cwd(), 'src', filePath)
		const file = await readFile(fp, { encoding: 'utf8' })
		let arr = file.split('')
		return arr.map(Number)
	} catch (err) {
		throw err
	}
}
