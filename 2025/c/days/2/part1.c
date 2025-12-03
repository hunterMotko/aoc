#include <stdio.h>
#include <stdlib.h>
#include <string.h>

long int check_is_double(long int n);
long int check_range_doubles(long int s, long int e);

int main() {
  const char *in_path = "days/2/in1";
  FILE *fp = fopen(in_path, "r");
  if (fp == NULL) {
    printf("Error opening file");
  }

  char line[1024];
  char *token;
  long int total_doubles = 0;

  while (fgets(line, sizeof line, fp) != NULL) {
    token = strtok(line, ",");
    while (token != NULL) {
      long int start;
      long int end;
      sscanf(token, "%ld-%ld", &start, &end);
      total_doubles += check_range_doubles(start, end);
      token = strtok(NULL, ",");
    }
  }

	printf("not first try?: %ld\n", total_doubles);
  fclose(fp);

  return 0;
}

long int check_range_doubles(long int s, long int e) {
	long int total = 0;
  while (s <= e) {
    if (check_is_double(s) == 1) {
			total += s;
		}
    s++;
  }
  return total;
}

long int check_is_double(long int n) {
	// skip single digits
  if (n <= 10) return 0;
	long temp = n;
	int len = 0;
	long divisor = 1;
	while (temp >0) {
		temp /= 10;
		len++;
	}
	// can skip odd lengths
	if (len % 2 != 0) return 0;

	int half = len /2;
	for (int i = 0; i < half; i++) {
		divisor *= 10;
	}

	long int right = n / divisor;
	long int left = n % divisor;
	return right == left;
}
