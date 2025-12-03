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

  printf("INPUT LEN: %zu\n", content.length);
  int start = 50;
  int zeros = 0;

  for (size_t i = 0; i < content.length; i++) {
    char dir;
    int point;
    sscanf(content.lines[i], "%c%d", &dir, &point);
		int offset = 0;

		if (dir == 'L') {
			offset = -point;
		} else if (dir == 'R') {
			offset = point;
		}

		start = (start + offset) % 100;

		if (start == 0) zeros++;
  }

	printf("Zeros: %d\n", zeros);
  free_file_content(content);

  return 0;
}
