#include <ctype.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_BUTTONS 64
#define MAX_POS 64
#define FALSE false
#define TRUE true

int64_t matrix[MAX_POS][MAX_BUTTONS];
int64_t joltages[MAX_POS];
int64_t presses[MAX_BUTTONS];
int reduced_on_pos[MAX_BUTTONS];
int n_btns, n_pos, min_presses, depth = 0;

int64_t parse_number(char **s) {
  int64_t res = 0;
  while (**s && !isdigit(**s))
    (*s)++;
  while (isdigit(**s)) {
    res = res * 10 + (**s - '0');
    (*s)++;
  }
  return res;
}

void solve(int joltage, int btn_idx) {
  if (btn_idx < 0) {
    if (joltage < min_presses)
      min_presses = joltage;
    return;
  }

  if (joltage >= min_presses)
    return;

  int pos = reduced_on_pos[btn_idx];
  if (pos >= 0) {
    int64_t val = joltages[pos];
    for (int i = btn_idx + 1; i < n_btns; i++) {
      val -= matrix[pos][i] * presses[i];
    }

    int64_t factor = matrix[pos][btn_idx];
    if (factor < 0) {
      factor = -factor;
      val = -val;
    }

    if (val < 0 || val % factor != 0)
      return;

    int64_t needed = val / factor;
    if (joltage + needed < min_presses) {
      presses[btn_idx] = needed;
      solve(joltage + (int)needed, btn_idx - 1);
    }
  } else {
    presses[btn_idx] = 0;
    int temp = joltage;
    while (temp < min_presses) {
      solve(temp, btn_idx - 1);
      presses[btn_idx]++;
      temp++;
    }
  }
}

int main(int argc, char *argv[]) {
  const char *path = (argc > 1) ? argv[1] : "ex1";
  FILE *fp = fopen(path, "r");
  if (!fp) {
    fprintf(stderr, "File not found: %s\n", path);
    return 1;
  }

  char line[4096];
  int64_t total = 0;
  while (fgets(line, sizeof(line), fp)) {
    if (strlen(line) < 10)
      continue;

    memset(matrix, 0, sizeof(matrix));
    memset(reduced_on_pos, -1, sizeof(reduced_on_pos));
    char *s = line;
    while (*s != '[')
      s++;
    s++;
    char *end_bracket = strchr(s, ']');
    n_pos = (int)(end_bracket - s);
    s = end_bracket + 1;
    n_btns = 0;
    while (*s == ' ' && *(s + 1) == '(')
      s++;
    while (*s == '(') {
      s++;
      while (1) {
        int64_t target_row = parse_number(&s);
        matrix[target_row][n_btns] = 1;
        if (*s == ')') {
          s++;
          break;
        }
        s++; // skip comma
      }
      n_btns++;
      while (*s == ' ')
        s++;
    }

    while (*s != '{')
      s++;
    for (int j = 0; j < n_pos; j++) {
      s++;
      joltages[j] = parse_number(&s);
    }

    bool seen[MAX_POS] = {0};
    for (int b = 0; b < n_btns; b++) {
      int p = -1;
      for (int i = 0; i < n_pos; i++) {
        if (!seen[i] && matrix[i][b] != 0) {
          p = i;
          break;
        }
      }

      if (p != -1) {
        reduced_on_pos[b] = p;
        seen[p] = TRUE;
        int64_t factor = matrix[p][b];
        for (int i = 0; i < n_pos; i++) {
          if (!seen[i] && matrix[i][b] != 0) {
            int64_t cur = matrix[i][b];
            matrix[i][b] = 0;
            for (int b2 = b + 1; b2 < n_btns; b2++) {
              matrix[i][b2] = factor * matrix[i][b2] - cur * matrix[p][b2];
            }
            joltages[i] = factor * joltages[i] - cur * joltages[p];
          }
        }
      }
    }

    min_presses = 10000;
    solve(0, n_btns - 1);
    total += min_presses;
    printf("LINE TOTAL: %d\n", min_presses);
  }
  fclose(fp);
  printf("\nTOTAL: %lld\n", total);
  return 0;
}
