#include <stdio.h>
#include <stdlib.h>

#define MAX_BUTTONS 25
#define MAX_LIGHTS 64
typedef unsigned long long mask;
#define HIGH 2147483647

const mask base = 1;

typedef struct {
  mask state;
  mask buttons[MAX_BUTTONS];
  int n_btns;
} Machine;

int count_set_bits(unsigned long long n);
int solve(Machine *m);
int parse_line(Machine *m);

int main() {
  Machine m;
  int total = 0;
  while (parse_line(&m) == 0) {
    total += solve(&m);
  }
  printf("Total: %d\n", total);
  return 0;
}

int count_set_bits(unsigned long long n) {
  int count = 0;
  while (n > 0) {
    if ((n & 1) == 1)
      count++;
    n = n >> 1;
  }
  return count;
}

int solve(Machine *m) {
  int count = HIGH;
  unsigned long long n_btns = (base << m->n_btns);
  for (unsigned long long i = 0; i < n_btns; i++) {
    mask cur_lights = m->state;
    int presses = 0;
    for (int b = 0; b < m->n_btns; b++) {
      int press = (i >> b) & 1;
      if (press) {
        // Toggle lights - exclusive OR (^)
        cur_lights = cur_lights ^ m->buttons[b];
        presses++;
      }
    }
    if (cur_lights == 0) {
      if (presses < count) {
        count = presses;
      }
    }
  }
  return count;
}

int parse_line(Machine *m) {
  int c;
  m->state = 0;
  m->n_btns = 0;
  while ((c = getchar()) != EOF && c != '[')
    ;

  if (c == EOF)
    return -1;

  int i = 0;
  while ((c = getchar()) != ']') {
    if (c == '#') {
      m->state |= (base << i);
    }
    if (c == '.' || c == '#')
      i++;
  }

  mask cur_btn = 0;
  char buf[10];
  int bi = 0;
  while ((c = getchar()) != EOF && c != '\n' && c != '{') {
    if (c == '(') {
      cur_btn = 0;
      bi = 0;
    } else if (c >= '0' && c <= '9') {
      buf[bi++] = c;
    } else if (c == ',' || c == ')') {
      if (bi > 0) {
        buf[bi] = 0;
        cur_btn |= (base << atoi(buf));
        bi = 0;
      }
      if (c == ')') {
        m->buttons[m->n_btns++] = cur_btn;
      }
    }
  }
  if (c == '{')
    while ((c = getchar()) != EOF && c != '\n')
      ;

  return 0;
}
