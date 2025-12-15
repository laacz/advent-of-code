#include "utils.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int data[16][16];
  int cols[16];
  int rows;
} Grid;

Grid parse(char *input) {
  Grid g = {0};
  char *copy = strdup(input);
  char *line = strtok(copy, "\n");

  while (line) {
    char *p = line;
    int val, n, col = 0;

    while (sscanf(p, "%d%n", &val, &n) == 1) {
      g.data[g.rows][col++] = val;
      p += n;
    }

    g.cols[g.rows] = col;
    g.rows++;

    line = strtok(NULL, "\n");
  }

  free(copy);

  return g;
}

int part1(Grid data) {
  int ret = 0;

  for (int row = 0; row < data.rows; row++) {
    int max = data.data[row][0];
    int min = data.data[row][0];

    for (int col = 0; col < data.cols[row]; col++) {
      max = MAX(data.data[row][col], max);
      min = MIN(data.data[row][col], min);
    }

    ret += max - min;
  }

  return ret;
}

int part2(Grid data) {
  int ret = 0;

  for (int row = 0; row < data.rows; row++) {
    for (int col1 = 0; col1 < data.cols[row]; col1++) {
      for (int col2 = col1 + 1; col2 < data.cols[row]; col2++) {
        int a = MAX(data.data[row][col1], data.data[row][col2]);
        int b = MIN(data.data[row][col1], data.data[row][col2]);

        if (a % b == 0) {
          ret += a / b;
        }
      }
    }
  }

  return ret;
}

int main() {
  assert(part1(parse("5 1 9 5\n7 5 3\n2 4 6 8")) == 18);
  assert(part2(parse("5 9 2 8\n9 4 7 3\n3 8 6 5")) == 9);

  char *input = read_file("../data/02.txt");
  Grid grid = parse(input);

  printf("part1: %d\n", part1(grid));
  printf("part2: %d\n", part2(grid));

  free(input);

  return 0;
}
