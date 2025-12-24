#include <stdio.h>
#include <stdlib.h>

#define SHAPES 6
#define BUF_SIZE 128

int main(int argc, char *argv[]) {
  char *path = (argc > 1) ? argv[1] : "ex1";
  FILE *fp = fopen(path, "r");
  if (!fp)
    return perror("file error"), 1;

  int shape[SHAPES];
  char buf[BUF_SIZE];
  for (int i = 0; i < SHAPES; i++) {
    if (!fgets(buf, BUF_SIZE, fp))
      break;

    for (int j = 0; j < 3; j++) {
      if (fgets(buf, BUF_SIZE, fp)) {
        for (int k = 0; buf[k] && k < 3; k++) {
          if (buf[k] == '#')
            shape[i]++;
        }
      }
    }
    fgets(buf, BUF_SIZE, fp);
  }

  long res = 0;
  int x, y;
  while (fscanf(fp, "%dx%d:", &x, &y) == 2) {
    int area = 0;
    for (int i = 0; i < SHAPES; i++) {
      int count;
      if (fscanf(fp, "%d", &count) == 1) {
        area += count * shape[i];
      }
    }
    if (area < (x * y)) {
      res++;
    }
  }
	fclose(fp);

  printf("%ld\n", res);
  return 0;
}
