#include <stdio.h>

#define FILENAME "file.txt"

void readFile() {
  FILE *fp = NULL;
  char ch;
  int count = 0;
  fp = fopen(FILENAME, "r");
  if (fp == NULL) {
    printf("FIle does not exist");
    return;
  }
  while ((ch = fgetc(fp)) != EOF) {
    if (ch == '\n') {
      count++;
    }
  }
  fclose(fp);
  fp = NULL;
  printf("Total number of lines are %d\n", count);
}
