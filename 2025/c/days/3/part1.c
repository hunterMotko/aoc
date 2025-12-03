#include "../../include/parse.h"
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max_joltage(const char *bank);

int main(int argc, char *argv[]) {
  const char *path; 
	(argc > 1) ? path = argv[1] : "ex1";
  FileContent con = parse_by_line(path);

  size_t i = 0;
  long int total = 0;
  while (i < con.length) {
    long int temp = (long int)max_joltage(con.lines[i]);
    total += temp;
    printf("%ld\n", temp);
    i++;
  }
  printf("TOTAL JOLTAGE: %ld\n", total);

  free_file_content(con);
  return 0;
}

int max_joltage(const char *bank) {
  size_t len = strlen(bank);
  int max_found = 0;
  for (size_t i = 0; i < len; i++) {
    size_t j = i + 1;
    while (j < len) {
      int jolts = (bank[i] - '0') * 10 + (bank[j] - '0');
      if (jolts > max_found) {
        max_found = jolts;
      }
      j++;
    }
  }
  return max_found;
}
unsigned long max_joltage_OL(const char *bank) {
  size_t len = strlen(bank);
  if (len < 2)
    return 0;

  unsigned long max_joltage_value = 0;
  char max_digit = bank[0];

  for (size_t i = 1; i < len; i++) {
    int current_joltage = (max_digit - '0') * 10 + (bank[i] - '0');
    if (current_joltage > max_joltage_value) {
      max_joltage_value = current_joltage;
    }

    if (max_joltage_value == 99) {
      return 99;
    }

    if (bank[i] > max_digit) {
      max_digit = bank[i];
    }
  }
  return max_joltage_value;
}
