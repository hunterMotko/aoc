#include "../../include/parse.h"
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned long max_joltage(const char *bank);

int main() {
  FileContent con = parse_by_line("days/3/in");

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

unsigned long max_joltage(const char *bank) {
  // 12 digit long joltage-------
  unsigned long max_found = 0;
  int jolt, pos = -1;
  int len = (int)strlen(bank);

  for (int i = 0; i < 12; i++) {
    jolt = 0;
    for (int j = pos + 1; j < len - 12 + i + 1; j++) {
			int cur = bank[j] - '0';
			if (cur > jolt) {
				jolt = cur;
				pos = j;
			}
    }
		max_found = (max_found * 10) + jolt;
  }

  return max_found;
}
