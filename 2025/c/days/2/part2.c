#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int cheating_invaild(long int n);
int is_invalid(long int n);
long int check_repeats(long int s, long int e);

int main() {
  const char *in_path = "days/2/in1";
  FILE *fp = fopen(in_path, "r");
  if (fp == NULL) {
    printf("Error opening file");
		return 1;
  }

  char line[1024];
  char *token;
  long int total_repeats = 0;

  while (fgets(line, sizeof line, fp) != NULL) {
    token = strtok(line, ",");
    while (token != NULL) {
      long int start;
      long int end;
      sscanf(token, "%ld-%ld", &start, &end);
      total_repeats += check_repeats(start, end);
      token = strtok(NULL, ",");
    }
  }

  printf("not first try?: %ld\n", total_repeats);
  fclose(fp);

  return 0;
}

int cheating_invaild(long int n){
    char str[32];
    sprintf(str, "%ld", n);

    int len = strlen(str);
    if(len % 2) return 0;
    return !strncmp(str, str+len/2, len/2);
}

long int check_repeats(long int s, long int e) {
  long int total = 0;
  while (s <= e) {
    if (is_invalid(s) == 1) {
      total += s;
    }
    s++;
  }
  return total;
}

void compute_prefix_function(char *str, int len, int *pi) {
  pi[0] = 0;
  int j = 0;
  for (int i = 1; i < len; i++) {
    while (j > 0 && str[i] != str[j]) {
      j = pi[j - 1];
    }
    if (str[i] == str[j]) {
      j++;
    }
    pi[i] = j;
  }
}

int is_invalid(long int n) {
  char s[20];
  int len;
  len = sprintf(s, "%ld", n);
  if (len < 2) {
    return 0;
  }

  int pi[20];
  compute_prefix_function(s, len, pi);
  int j = pi[len - 1];
  long int k = len - j;

	if (j > 0 && len % k == 0) {
    return 1;
  }

  return 0;
}
