#include "../include/range.h"
#include <stdio.h>
#include <stdlib.h>

void init_list(RangeList *list) {
  list->data = (Range *)malloc(INIT_CAP * sizeof(Range));
  list->len = 0;
  list->cap = INIT_CAP;
  if (list->data == NULL) {
    fprintf(stderr, "Failed to allocate memory for ranges");
  }
}

void add_range(RangeList *list, long int l, long int r) {
  if (list->len >= list->cap) {
    list->cap *= 2;
    list->data = (Range *)realloc(list->data, list->cap * sizeof(Range));
    if (list->data == NULL) {
      fprintf(stderr, "Failed to allocate memory for ranges");
    }
  }
  list->data[list->len].l = l;
  list->data[list->len].r = r;
  list->len++;
}
