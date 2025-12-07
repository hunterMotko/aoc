#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LEN 10000 
#define ROWS 4
#define COLS 1000

int main(int argc, char *argv[]) {
  const char *path;
  (argc > 1) ? path = argv[1] : "ex1";

  FILE *fp = fopen(path, "r");
  if (fp == NULL) {
    perror("open file error");
    return 1;
  }

  char line[MAX_LINE_LEN];
  int row = 0;
  int col = 0;
  int ops_row = 0;
  int grid[ROWS][COLS];
  char opr[COLS];

  while (fgets(line, sizeof(line), fp) != NULL) {
    line[strcspn(line, "\n")] = 0;
    char *token = strtok(line, " ");
    int temp_col = 0;
    while (token != NULL) {
      printf("token: %s, token[0]: %c\n", token, token[0]);
      if (token[0] == '*' || token[0] == '+') {
        opr[temp_col++] = token[0];
        ops_row = 1;
      } else {
        grid[row][temp_col++] = atoi(token);
      }
      token = strtok(NULL, " ");
    }
    if (col == 0)
      col = temp_col;
    if (ops_row == 0)
      row++;
  }

	long int total = 0;
  for (int c = 0; c < col; c++) {
		char op = opr[c];
		long int tt = 0;
		if (op == '*') tt = 1;

    for (int r = 0; r < row; r++) {
			if (op == '*') {
				tt *= grid[r][c];
			} else { 
				tt += grid[r][c];
			}
    }
		printf("TT: %ld\n", tt);
		total += tt;
  }

	printf("TOTAL: %ld\n", total);

  fclose(fp);
  return 0;
}
