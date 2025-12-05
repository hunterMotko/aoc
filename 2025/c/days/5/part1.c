#include "../../include/range.h"
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

int is_fresh(RangeList *ids, long int id);

int main(int argc, char *argv[]) {
  const char *path;
  (argc > 1) ? path = argv[1] : "ex1";

  FILE *fp = fopen(path, "r");

  RangeList ids;
  init_list(&ids);
  char line[100];

  while (fgets(line, sizeof(line), fp) != NULL) {
    if (line[0] == '\n' || line[0] == '\r' || line[0] == '\0') {
      // Break out one ranges are over to
      // change scan
      break;
    }
    long int l, r;
    if (sscanf(line, "%ld-%ld", &l, &r) == 2)
      add_range(&ids, l, r);
  }

  long int total_fresh = 0;
  long int id;
  while (fscanf(fp, "%ld", &id) == 1) {
    if (is_fresh(&ids, id))
      total_fresh++;
  }

  printf("Total fresh: %ld\n", total_fresh);
  return 0;
}

int is_fresh(RangeList *ids, long int id) {
  for (size_t i = 0; i < ids->len; i++) {
    if (id >= ids->data[i].l && id <= ids->data[i].r)
      return 1;
  }
  return 0;
}
