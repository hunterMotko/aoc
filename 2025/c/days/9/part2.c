#include <stdio.h>
#include <stdlib.h>

#define MAX_POINTS 500
#define MAX(x, y) (((x) > (y)) ? (x) : (y))
#define MIN(x, y) (((x) < (y)) ? (x) : (y))

typedef struct {
  int x, y;
} point;

typedef struct {
  int x, y1, y2;
} v_edge;

typedef struct {
  int y, x1, x2;
} h_edge;

long long area(point a, point b) {
  return (long long)(abs(a.x - b.x) + 1) * (abs(a.y - b.y) + 1);
}
int cmp_v(const void *a, const void *b) {
  return ((v_edge *)a)->x - ((v_edge *)b)->x;
}
int in_poly(const int x, const int y);
int is_valid_rec(point a, point b);

point pts[MAX_POINTS];
int p_len = 0;
v_edge v_edges[MAX_POINTS];
int v_len = 0;
h_edge h_edges[MAX_POINTS];
int h_len = 0;

int main(int argc, char *argv[]) {
  const char *path = (argc > 1) ? argv[1] : "ex1";
  FILE *fp = fopen(path, "r");
  if (!fp) {
    fprintf(stderr, "File not found: %s\n", path);
    return 1;
  }

  int x, y;
  while (fscanf(fp, "%d,%d", &x, &y) == 2) {
    pts[p_len++] = (point){x, y};
  }
  fclose(fp);

  for (int i = 0; i < p_len; ++i) {
    const int j = (i + 1) % p_len;
    if (pts[i].x == pts[j].x) {
      v_edges[v_len].x = pts[i].x;
      v_edges[v_len].y1 = MIN(pts[i].y, pts[j].y);
      v_edges[v_len].y2 = MAX(pts[i].y, pts[j].y);
      ++v_len;
    } else {
      h_edges[h_len].y = pts[i].y;
      h_edges[h_len].x1 = MIN(pts[i].x, pts[j].x);
      h_edges[h_len].x2 = MAX(pts[i].x, pts[j].x);
      ++h_len;
    }
  }

  qsort(v_edges, v_len, sizeof(v_edge), cmp_v);

  long long res = 0;
  for (int i = 0; i < p_len - 1; ++i) {
    for (int j = i + 1; j < p_len; ++j) {
      if (is_valid_rec(pts[i], pts[j])) {
        long long a = area(pts[i], pts[j]);
        if (a > res)
          res = a;
      }
    }
  }

  printf("Answer: %lld\n", res);
  return 0;
}

int in_poly(const int x, const int y) {
  for (int i = 0; i < h_len; ++i) {
    if (h_edges[i].y == y && x >= h_edges[i].x1 && x <= h_edges[i].x2)
      return 1;
  }
  int crossings = 0;
  for (int i = 0; i < v_len; ++i) {
    if (v_edges[i].x > x)
      break;
    if (v_edges[i].x == x && y >= v_edges[i].y1 && y <= v_edges[i].y2)
      return 1;
    if (v_edges[i].x < x && y >= v_edges[i].y1 && y < v_edges[i].y2)
      ++crossings;
  }
  printf("x: %d, y: %d,Crossing: %d, bitwise: %d\n", x, y, crossings,
         crossings & 1);
  // bitwise AND is sweet
  return crossings & 1;
}

int is_valid_rec(point a, point b) {
  int x1 = MIN(a.x, b.x);
  int x2 = MAX(a.x, b.x);
  int y1 = MIN(a.y, b.y);
  int y2 = MAX(a.y, b.y);

  if (!in_poly(x1, y1) || !in_poly(x2, y1) || !in_poly(x1, y2) ||
      !in_poly(x2, y2)) {
    return 0;
  }
  for (int i = 0; i < p_len; ++i) {
    int py = pts[i].y;
    if (py > y1 && py < y2) {
      if (!in_poly(x1, py) || !in_poly(x2, py))
        return 0;
    }
  }
  for (int i = 0; i < p_len; ++i) {
    int px = pts[i].x;
    if (px > x1 && px < x2) {
      if (!in_poly(px, y1) || !in_poly(px, y2))
        return 0;
    }
  }
  for (int i = 0; i < v_len; ++i) {
    if (v_edges[i].x > x1 && v_edges[i].x < x2) {
      if (v_edges[i].y1 < y1 && v_edges[i].y2 > y2)
        return 0;
    }
  }
  for (int i = 0; i < h_len; ++i) {
    if (h_edges[i].y > y1 && h_edges[i].y < y2) {
      if (h_edges[i].x1 < x1 && h_edges[i].x2 > x2)
        return 0;
    }
  }
  return 1;
}
