import { readFile } from 'fs/promises'
import path from 'path'

type Stationary = {
	rules: string[][];
	updates: string[][];
}
async function readFileWithManual(fileName: string): Promise<Stationary> {
	const result: Stationary = {
		rules: [],
		updates: []
	}
	try {
		const fp = path.join(process.cwd(), 'src', fileName)
		const file = await readFile(fp, { encoding: 'utf8' })
		const lines = file.split('\n\n')
		for (let i = 0; i < lines.length; i++) {
			for (let j of lines[i].split('\n')) {
				if (j === "") break
				if (i === 0) {
					result.rules.push(j.split('|'))
				} else {
					result.updates.push(j.split(','))
				}
			}
		}
	} catch (err) {
		console.error("readFile error ", err)
		throw err
	}
	return result
}

export async function partOne(fileName: string): Promise<number> {
	try {
		const { rules, updates } = await readFileWithManual(fileName)
		const pages: number[] = []
		for (let i = 0; i < updates.length; i++) {
			if (updates[i] === undefined) throw new Error('Updates undefinded')

			if (rightOrder(updates[i], rules)) {
				let mid = Math.floor(updates[i].length / 2)
				let val = parseInt(updates[i][mid])
				if (!isNaN(val)) {
					pages.push(val)
				}
			}
		}
		return pages.reduce((a, b) => a + b, 0)
	} catch (err) {
		console.error('Part one error ', err)
		throw err
	}
}

export async function partTwo(fileName: string): Promise<number> {
	try {
		const { rules, updates } = await readFileWithManual(fileName)
		const pages: number[] = []
		const ruleOrders: { [key: string]: string[] } = {}

		for (let i = 0; i < rules.length; i++) {
			const [x, y] = rules[i]
			if (!ruleOrders[x]) {
				ruleOrders[x] = [y]
				continue
			}
			ruleOrders[x].push(y)
		}

		for (let update of updates) {
			if (!rightOrder(update, rules)) {
				for (let i = 0; i < update.length; i++) {
					for (let j = i + 1; j < update.length; j++) {
						let cur = update[j]
						let curRule = ruleOrders[update[i]]
						if (curRule && curRule.includes(cur)) {
							update[j] = update[i]
							update[i] = cur
						}
					}
				}
				let mid = Math.floor(update.length / 2)
				let val = parseInt(update[mid])
				pages.push(val)
			}
		}
		return pages.reduce((a, b) => a + b, 0)
	} catch (err) {
		console.error('Part one error ', err)
		throw err
	}
}

function rightOrder(update: string[], rules: string[][]) {
	for (let i = 0; i < rules.length; i++) {
		const [x, y] = rules[i]
		let a = update.indexOf(x)
		let b = update.indexOf(y)
		if (a === -1 || b === -1) {
			continue
		} else if (a > b) {
			return false
		}
	}
	return true
}

/**
 * Checks if the elements in 'update' array are in the correct order based on 'rules'.
 * This function has been optimized for efficiency.
 *
 * Time Complexity: O(L + R) on average.
 * - Creating `updateIndexMap`: O(L) where L is the length of the `update` array.
 * - Iterating through `rules` and performing Map lookups: O(R) where R is the number of rules.
 * Space Complexity: O(L) for the `updateIndexMap`.
 *
 * @param update - The array of strings to check.
 * @param rules - An array of [predecessor, successor] string pairs defining order.
 * @returns true if the update array satisfies all rules, false otherwise.
 */
function correctOrder(update: string[], rules: string[][]): boolean {
	// Create a map to store the index of each element in the update array.
	// This allows for O(1) average time lookups for an element's position.
	const updateIndexMap = new Map<string, number>();
	for (let i = 0; i < update.length; i++) {
		updateIndexMap.set(update[i], i);
	}
	// Iterate through each rule to check if the update array violates any order.
	for (let i = 0; i < rules.length; i++) {
		const [predecessor, successor] = rules[i];
		// Get the indices of the predecessor and successor in the current update array.
		const indexPredecessor = updateIndexMap.get(predecessor);
		const indexSuccessor = updateIndexMap.get(successor);
		// If either the predecessor or successor is not found in the update array,
		// this specific rule does not apply to this update, so we continue to the next rule.
		if (indexPredecessor === undefined || indexSuccessor === undefined) {
			continue;
		}
		// If the predecessor appears AFTER the successor in the update array,
		// the order is incorrect according to this rule.
		if (indexPredecessor > indexSuccessor) {
			return false;
		}
	}
	// If no rules are violated after checking all of them, the order is correct.
	return true;
}

/**
 * Processes a file containing rules and updates, reorders updates based on rules (if necessary),
 * and calculates a sum from the middle element of the reordered updates.
 * This function incorporates optimizations to reduce time and space complexity.
 * Overall Time Complexity:
 * - `readFileWithManual`: Assumed to be efficient, complexity depends on file I/O and parsing.
 * - Building `ruleOrders`: O(R), where R is the number of rules. `Set.add()` is O(1) average.
 * - Processing `updates`: O(U * (L + R + L^2)), where U is number of updates, L is max update length.
 * - `rightOrder`: O(L + R) on average.
 * - Reordering loop: O(L^2) due to nested loops. The `Set.has()` check inside is O(1) average.
 * - `pages.reduce`: O(U).
 * Overall average: O(file_read_time + R + U * (L + R + L^2)).
 * The dominant factor for large `L` is still the $O(L^2)$ reordering loop if many updates require it.
 * Space Complexity:
 * - `rules`, `updates`: O(R_total_length + U_total_length), where R_total_length is sum of lengths of all rule arrays,
 * and U_total_length is sum of lengths of all update arrays.
 * - `pages`: O(U) for storing middle values.
 * - `ruleOrders`: O(R + R_total_length) for storing the Sets.
 * Overall average: O(R_total_length + U_total_length).
 * @param fileName The path to the input file.
 * @returns A Promise resolving to the sum of "pages" values.
 */
export async function partTwoHalf(fileName: string): Promise<number> {
	try {
		// Read rules and updates from the file.
		const { rules, updates } = await readFileWithManual(fileName);
		const pages: number[] = [];
		// Optimize ruleOrders: use Set<string> for the values to ensure O(1) average time
		// complexity for `has()` checks later when checking rule precedence.
		const ruleOrders: { [key: string]: Set<string> } = {};
		for (let i = 0; i < rules.length; i++) {
			const [predecessor, successor] = rules[i];
			// If the predecessor key doesn't exist yet, initialize it with a new Set.
			if (!ruleOrders[predecessor]) {
				ruleOrders[predecessor] = new Set<string>();
			}
			// Add the successor to the Set of elements that must follow the predecessor.
			ruleOrders[predecessor].add(successor);
		}
		// Process each update array.
		for (let update of updates) {
			// Check if the current update array is already in the correct order.
			// This uses the optimized `rightOrder` function.
			if (!rightOrder(update, rules)) {
				// If the update is NOT in the right order, attempt to reorder it.
				// The following nested loops perform a specific type of pairwise swap.
				// This part remains O(L^2) in time complexity for a single update,
				// but the internal `has()` check is now O(1) average.
				for (let i = 0; i < update.length; i++) {
					for (let j = i + 1; j < update.length; j++) {
						let currentElementAtJ = update[j]; // The element at the later position
						let rulesForElementAtI = ruleOrders[update[i]]; // Rules where update[i] is a predecessor
						// Check if there are rules for `update[i]` and if `update[i]` must precede `currentElementAtJ`.
						// `rulesForElementAtI.has(currentElementAtJ)` is O(1) on average.
						if (rulesForElementAtI && rulesForElementAtI.has(currentElementAtJ)) {
							// If `update[i]` should come before `currentElementAtJ`, and `currentElementAtJ`
							// is currently after `update[i]` (which is always true because `j > i`),
							// then swap them to enforce the rule.
							// This is a standard element swap.
							update[j] = update[i]; // Move element from `i`'s position to `j`'s position
							update[i] = currentElementAtJ; // Move original `j`'s element to `i`'s position
						}
					}
				}
				// After potential reordering, calculate the middle element's value.
				// Note: The original logic only adds to 'pages' if reordering occurred.
				const midIndex = Math.floor(update.length / 2);
				const middleValueString = update[midIndex];
				// Safely parse the integer value, handling cases where it might not be a valid number.
				const val = parseInt(middleValueString);
				if (isNaN(val)) {
					console.warn(`Warning: Could not parse "${middleValueString}" (from update: [${update}]) to a number. Skipping this value.`);
					continue; // Skip adding this value if it's not a valid number.
				}
				pages.push(val);
			}
		}
		// Sum all collected page values.
		return pages.reduce((a, b) => a + b, 0);
	} catch (err) {
		console.error('Part two error: ', err);
		throw err; // Re-throw the error for upstream handling.
	}
}
