#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  long x;
  long y;
} point;

int main(int argc, char *argv[]) {
  char *path;
  (argc > 1) ? path = argv[1] : "ex1";
  FILE *fp = fopen(path, "r");
  if (fp == NULL) {
    perror("Error opening file");
    return 1;
  }

  int lines = 8;
  if (strcmp(path, "days/9/in") == 0) {
    lines = 496;
  }

  point points[lines];
  char *line = NULL;
  size_t len = 0;
  ssize_t read;
  int i = 0;
  while ((read = getline(&line, &len, fp)) != EOF) {
    long x;
    long y;
    sscanf(line, "%ld,%ld", &x, &y);
    point p = {x, y};
    points[i++] = p;
  }

	long max = 0;
  for (int j = 0; j < i - 1; j++) {
    point cur = points[j];
    for (int k = j + 1; k < i; k++) {
      point next = points[k];
      long dx = (cur.x > next.x) ? cur.x - next.x + 1 : next.x - cur.x + 1;
      long dy = (cur.y > next.y) ? cur.y - next.y + 1 : next.y - cur.y + 1;
			if ((dx*dy) > max) {
				max = dx*dy;
			}
    }
  }

	printf("Max: %ld\n", max);

  if (line)
    free(line);

  return EXIT_SUCCESS;
}
