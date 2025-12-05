#include "../../include/parse.h"
#include <stddef.h>
#include <stdio.h>
#include <string.h>

#define MAXLINE 1024

typedef struct {
  int x;
  int y;
} Dirs;

const Dirs directions[] = {
    {-1, 0},  // Up
    {1, 0},   // Down
    {0, -1},  // Left
    {0, 1},   // Right
    {-1, -1}, // Up-Left
    {-1, 1},  // Up-Right
    {1, -1},  // Down-Left
    {1, 1}    // Down-Right
};

int check_adjacent(char **lines, int yl, int xl, int i, int j);
void clean_up_removed(char **lines, int yl, int xl);

int main(int argc, char *argv[]) {
  const char *path;
  (argc > 1) ? path = argv[1] : "ex1";

  FileContent diagram = parse_by_line(path);
  int i = 0;
  int ll = (int)strlen(diagram.lines[0]);
  long int rolls = 0;

	while (1) {
		long int temp_rolls = 0;
		while (i < (int)diagram.length) {
			int j = 0;
			while (j < ll) {
				if (diagram.lines[i][j] == '@' &&
					check_adjacent(diagram.lines, (int)diagram.length, ll, i, j) == 1) {
					++temp_rolls;
				}
				j++;
			}
			i++;
		}

		clean_up_removed(diagram.lines, (int)diagram.length, ll);
		if (temp_rolls == 0) {
			break;
		} else {
			rolls += temp_rolls;
			i = 0;
		}
	}

  printf("Can Access: %ld\n", rolls);
  return 0;
}

int check_adjacent(char **lines, int yl, int xl, int y, int x) {
  int adjacents = 0;
  for (int k = 0; k < 8; k++) {
    int ny = y + directions[k].y;
    int nx = x + directions[k].x;
    // check in bounds
    if (ny >= 0 && ny < yl && nx >= 0 && nx < xl) {
      char c = lines[ny][nx];
      if (c == '@' || c == 'x')
        ++adjacents;
    }
  }
  int res = 0;
  if (adjacents < 4) {
		res = 1;
		lines[y][x] = 'x';
  }
  return res;
}

void clean_up_removed(char **lines, int yl, int xl) {
	for (int i = 0; i < yl; i++) {
		for (int j = 0; j < xl; j++) {
			if (lines[i][j] == 'x') lines[i][j] = '.';
		}
	}
}
