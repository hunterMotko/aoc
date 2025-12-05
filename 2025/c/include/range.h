#include <stddef.h>

typedef struct {
  long int l;
  long int r;
} Range;

typedef struct {
  Range *data;
  size_t len;
  size_t cap;
} RangeList;

#define INIT_CAP 10

void init_list(RangeList *list);
void add_range(RangeList *list, long int l, long int r);
