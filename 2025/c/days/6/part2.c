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

  char lines[10][4096];
  size_t rows = 0;

  while (fgets(lines[rows], 4096, fp)) {
    size_t len = strlen(lines[rows]);
    if (lines[rows][len - 1] == '\n') {
      lines[rows][len - 1] = '\0';
    }
    ++rows;
  }

  long int total = 0;
  size_t idx = 0;
  long int nums[10];
  for (int c = strlen(lines[0]) - 1; c >= 0; c--) {
    long int n = 0;
    int n_ok = 0;

    for (size_t r = 0; r < rows; r++) {
      char ch = lines[r][c];
      if (ch >= '0' && ch <= '9') {
        n = n * 10 + ch - '0';
        n_ok = 1;
      } else if (ch == '*') {
        long int x = n;
        for (size_t i = 0; i < idx; i++) {
          x *= nums[i];
        }
        total += x;
      } else if (ch == '+') {
        long int x = n;
        for (size_t i = 0; i < idx; i++) {
          x += nums[i];
        }
        total += x;
      }
    }

    if (n_ok) {
      nums[idx++] = n;
    } else {
      idx = 0;
    }
  }

  printf("TOTAL: %ld\n", total);

  fclose(fp);
  return 0;
}
