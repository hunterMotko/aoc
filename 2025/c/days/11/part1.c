#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 26^3 is ~17,576. 
// 18000 provides a safe buffer for 3-letter hashes.
#define MAX_HASH 18000
#define MAX_CHILDREN 24
#define M_IDX 25

long long graph[MAX_HASH][26] = {0};

int hash(char *in);
void dfs(int start, int end); 

int main(int argc, char *argv[]) {
  char *filename = (argc > 1) ? argv[1] : "input.txt";
  FILE *fp = fopen(filename, "r");
  if (!fp)
    return printf("Error opening %s\n", filename), 1;

  char line[100];
  while (fgets(line, sizeof(line), fp)) {
    char *ptr;
    char *key = strtok_r(line, ":", &ptr);
    if (!key)
      continue;
    int parent = hash(key);
    char *child;
    while ((child = strtok_r(NULL, " \n\r", &ptr))) {
      int child_id = hash(child);
      int count = (int)graph[parent][MAX_CHILDREN];
      if (count < MAX_CHILDREN) {
        graph[parent][count] = child_id;
        graph[parent][MAX_CHILDREN]++;
      }
    }
  }
  fclose(fp);

  int start = hash("you");
  int end = hash("out");
  graph[end][M_IDX] = 1;
  dfs(start, end);
  long long total = graph[start][M_IDX];
  printf("Total Paths: %lld\n", total);

  return 0;
}

int hash(char *in) {
  int out = 0;
  while (*in == ' ')
    in++; 
  for (int i = 0; i < 3 && in[i] != '\0' && in[i] != ':'; i++) {
    out = out * 26 + (in[i] - 'a');
  }
  return out;
}

void dfs(int start, int end) {
  if (graph[start][M_IDX] != 0 || start == end)
    return;

  long long res = 0;
  int num_children = (int)graph[start][MAX_CHILDREN];
  for (int i = 0; i < num_children; i++) {
    int v = (int)graph[start][i];
    dfs(v, end);
    if (graph[v][M_IDX] > 0) {
      res += graph[v][M_IDX];
    }
  }
  graph[start][M_IDX] = (res == 0) ? -1 : res;
}
