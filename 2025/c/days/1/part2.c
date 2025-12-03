#include "../../include/parse.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  const char *path; 
	(argc > 1) ? path = argv[1] : "ex1";


  FileContent content = parse_by_line(path);
  if (content.lines == NULL) {
    fprintf(stderr, "Failed to read file\n");
    return 1;
  }

  int start = 50;
  int zeros = 0;
  char prev_dir = 'R';
  for (size_t i = 0; i < content.length; i++) {
    char dir;
    int point;
    if (sscanf(content.lines[i], "%c%d", &dir, &point) == 2) {
      if (dir != prev_dir) {
        start = (100 - start) % 100;
        prev_dir = dir;
      }
      start += point;
      zeros += start / 100;
			start %= 100;
    }
  }

  printf("Zeros: %d\n", zeros);
  free_file_content(content);

  return 0;
}
