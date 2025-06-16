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
