#include <stdio.h>

#define WIDTH 141

int main() {
  long int state[WIDTH];
  long int res = 1;

  int c = 0, i = 0;
  while ((c = getchar()) != '\n')
    state[i++] = (c == 'S');

  i = 0;
  while((c = getchar()) != EOF) {
    if (c == '\n') {
      i = 0;
      continue;
    }
    if ((c == '^') && (state[i] > 0)) {
      res += state[i];
      state[i+1] += state[i];
      state[i-1] += state[i];
      state[i] = 0;
    }
    i++;
  }
  printf("Result: %ld\n", res);
}
