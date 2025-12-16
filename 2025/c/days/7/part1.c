#include <stdio.h>

#define WIDTH 141

int main() {
  long int state[WIDTH];
  int c = 0, i = 0, res = 0;
  while ((c = getchar()) != '\n') {
		state[i++] = (c == 'S');
		printf("C: %c, state: %ld, I: %d\n", c, state[i], i);
	}
	printf("\n---------\n");
  i = 0;
  while((c = getchar()) != EOF) {
    if (c == '\n') {
      i = 0;
      continue;
    }
		printf("C: %c, state: %ld, I: %d\n", c, state[i], i);
    if ((c == '^') && (state[i] > 0)) {
      res++;
      state[i+1] += state[i];
      state[i-1] += state[i];
      state[i] = 0;
    }
    i++;
  }
  printf("Result: %d\n", res);
}
