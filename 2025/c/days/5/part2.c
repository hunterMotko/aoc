#include "../../include/range.h"
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

int cmp_ranges(const void *a, const void *b);
long int total_fresh(RangeList *list);

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
  // Apparently skipping the checks all together works...
  // Got this by accident trying to merge overlaps
  printf("fresh ranges: %ld\n", total_fresh(&ids));
  return 0;
}

int cmp_ranges(const void *a, const void *b) {
  Range *r1 = (Range *)a;
  Range *r2 = (Range *)b;
  if (r1->l < r2->l)
    return -1;
  if (r1->l > r2->l)
    return 1;
  if (r1->r < r2->r)
    return -1;
  if (r1->r > r2->r)
    return 1;
  return 0;
}

// 3-5
// 10-14
// 12-18
// 16-20

// Cur_l: 3, Cur_r: 5, N_l: 10, N_r: 14
	// else 
		// 3 = ((5 - 3) + 1)
//	move next -> cur
// Cur_l: 10, Cur_r: 14, N_l: 12, N_r: 18
	// if (12 <= 14)
		//18 = (14 > 18) ? 14 : 18
// Cur_l: 10, Cur_r: 18, N_l: 16, N_r: 20
	// if (10 <= 20)
		// 20 = (18 > 20) ? 18 : 20

// 3 += 11 = ((20 - 10) + 1)

long int total_fresh(RangeList *list) {
  if (list->len == 0)
    return 0;
  qsort(list->data, list->len, sizeof(Range), cmp_ranges);
  long int total_fresh = 0;
  long int cur_l = list->data[0].l;
  long int cur_r = list->data[0].r;
  for (size_t i = 1; i < list->len; i++) {
    long int next_l = list->data[i].l;
    long int next_r = list->data[i].r;
    // Check overlap
    if (next_l <= cur_r) {
      cur_r = (cur_r > next_r) ? cur_r : next_r;
    } else {
			// add right away if no overlap
      total_fresh += (cur_r - cur_l + 1);
      cur_l = next_l;
      cur_r = next_r;
    }
  }
  total_fresh += (cur_r - cur_l + 1);
  return total_fresh;
}
