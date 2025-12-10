#include <math.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#define PAIRS 1000

typedef struct {
  long long x;
  long long y;
  long long z;
} point;

typedef struct {
  int dis;
  size_t i;
  size_t j;
} pair;

int distance(point a, point b);
int cmp(const void *a, const void *b);
int cmp_rev(const void *a, const void *b);
int *init_circuits(size_t l);
int find(int *circs, int a);
void union_set(int *circs, int a, int b);

int main(int argc, char *argv[]) {
  char *path;
  (argc > 1) ? path = argv[1] : "ex1";

  FILE *fp = fopen(path, "r");
  if (fp == NULL) {
    perror("Error opening file");
    return 1;
  }

  point pts[PAIRS];
  size_t p_len = 0;
  long long x, y, z;
  while (fscanf(fp, "%lld,%lld,%lld", &x, &y, &z) == 3) {
    point p = {x, y, z};
    pts[p_len] = p;
    ++p_len;
  }
  fclose(fp);

  pair *pairs;
  size_t n_pairs = (p_len * (p_len - 1)) / 2;
  pairs = malloc(n_pairs * sizeof(pair));
  size_t b_len = 0;
  for (size_t i = 0; i < p_len - 1; i++) {
    for (size_t j = i + 1; j < p_len; j++) {
      pairs[b_len].i = i;
      pairs[b_len].j = j;
      pairs[b_len].dis = distance(pts[i], pts[j]);
      b_len++;
    }
  }

  qsort(pairs, b_len, sizeof(pair), cmp);
  int *circs = init_circuits(p_len);
  for (size_t i = 0; i < PAIRS; i++) {
    union_set(circs, pairs[i].i, pairs[i].j);
  }

  int *size_sets = calloc(p_len, sizeof(int));
  for (size_t i = 0; i < p_len; i++) {
    int root = find(circs, i);
    size_sets[root]++;
  }

  qsort(size_sets, p_len, sizeof(size_sets[0]), cmp_rev);
  unsigned long long total_sets = 1;
  for (size_t i = 0; i < 3; i++) {
    total_sets *= size_sets[i];
  }
  printf("Total sets: %llu\n", total_sets);

  free(size_sets);
  free(circs);
  free(pairs);
  return 0;
}

int cmp(const void *a, const void *b) {
  pair *p1 = (pair *)a;
  pair *p2 = (pair *)b;
  return p1->dis - p2->dis;
}

int cmp_rev(const void *a, const void *b) { 
	const int *x = (const int *)a;
	const int *y = (const int *)b;
	if (*x > *y) {
		return -1;
	} else if (*x < *y) {
		return 1;
	} else {
		return 0;
	}
}

int distance(point a, point b) {
  long dx = a.x - b.x;
  long dy = a.y - b.y;
  long dz = a.z - b.z;
  return (int)sqrt((double)(dx * dx + dy * dy + dz * dz));
}

int *init_circuits(size_t p_len) {
  int *circs = malloc(p_len * sizeof(int));
  for (size_t i = 0; i < p_len; i++) {
    circs[i] = -1;
  }
  return circs;
}

int find(int *circs, int a) {
  if (circs[a] < 0)
    return a;
  else
    return circs[a] = find(circs, circs[a]);
}

void union_set(int *circs, int a, int b) {
  int root_a = find(circs, a);
  int root_b = find(circs, b);
  if (root_a != root_b)
    circs[root_b] = root_a;
}
