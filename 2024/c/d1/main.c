#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../include/ht.h"

int cmp(const void *x, const void *y) {
  return *((int *)x) - *((int *)y);
}

int part1(FILE *fp) {
  char *line = NULL;
  size_t len = 0;
  int left[1000];
  int right[1000];
  char *val = NULL;
  char *save = NULL;
  int i = 0;
  while (getline(&line, &len, fp) > 1) {
    val = strtok_r(line, " \n", &save);
    left[i] = atoi(val);
    val = strtok_r(NULL, " \n", &save);
    right[i] = atoi(val);
    i++;
  }
  qsort(left, i, sizeof(int), cmp);
  qsort(right, i, sizeof(int), cmp);
  int sum = 0;
  for(int j = 0; j < i; j++) {
    sum += abs(left[j] - right[j]);
  }
  fclose(fp);
  if (line) free(line);
  return sum;
}

int part2(FILE *fp) {
  char *line = NULL;
  size_t len = 0;
  ht *hash = NULL;
}

int main() {
  FILE *fp;
  fp = fopen("in", "r");
  if (fp == NULL) exit(EXIT_FAILURE);
  // int i = part1(fp);
  exit(EXIT_SUCCESS);
}
