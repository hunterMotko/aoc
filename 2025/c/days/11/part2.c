#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_HASH 18000
#define M_COL 25
#define C_COL 24

long long graph[MAX_HASH][26];

int hash(char *in);
long long dfs(int st, int end);
void clear_g();

int main(int argc, char *argv[]) {
  memset(graph, 0, sizeof(graph));
  clear_g();
  char *path = (argc > 1) ? argv[1] : "ex2";
  FILE *fp = fopen(path, "r");
  if (!fp)
    return 1;

  char line[256];
  while (fgets(line, sizeof(line), fp)) {
    char *ptr;
    char *p_str = strtok_r(line, ":", &ptr);
    if (!p_str)
      continue;

    int u = hash(p_str);
    char *c_str;
    while ((c_str = strtok_r(NULL, " \n\r", &ptr))) {
      int v = hash(c_str);
      graph[u][graph[u][C_COL]++] = v;
    }
  }
  fclose(fp);

  int svr = hash("svr");
  int fft = hash("fft");
  int dac = hash("dac");
  int out = hash("out");

  long long p1 = dfs(svr, fft);
  printf("p1: %lld\n", p1);
  clear_g();

  long long p2 = dfs(fft, dac);
  printf("p2: %lld\n", p2);
  clear_g();

  long long p3 = dfs(dac, out);
  printf("p3: %lld\n", p2);
  long long total = p1 * p2 * p3;
  // 417190406827152
  printf("Total - %lld\n", total);

  return 0;
}

int hash(char *in) {
  int out = 0;
  while (*in == ' ')
    in++;
  for (int i = 0; i < 3 && in[i] && in[i] != ':'; i++) {
    out = out * 26 + (in[i] - 'a');
  }
  return out;
}

long long dfs(int st, int end) {
  if (st == end)
    return 1;

  if (graph[st][M_COL] != -1)
    return graph[st][M_COL];

  long long sum = 0;
  int children = (int)graph[st][C_COL];
  for (int i = 0; i < children; i++) {
    sum += dfs((int)graph[st][i], end);
  }
  return graph[st][M_COL] = sum;
}

void clear_g() {
  for (int i = 0; i < MAX_HASH; i++)
    graph[i][M_COL] = -1;
}
