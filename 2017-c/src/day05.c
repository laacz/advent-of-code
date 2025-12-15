#include "utils.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int *data;
  size_t size;
} IntArray;

IntArray parse(const char *input) {
  IntArray arr = {0};
  size_t cap = 0;
  int val, n;
  const char *p = input;

  while (sscanf(p, "%d%n", &val, &n) == 1) {
    if (arr.size >= cap) {
      cap = cap ? cap * 2 : 8;
      arr.data = realloc(arr.data, cap * sizeof(int));
    }
    arr.data[arr.size++] = val;
    p += n;
  }

  return arr;
}

int part1(IntArray *jumps) {
  int ret = 0;
  int pos = 0;
  int *offsets = malloc(jumps->size * sizeof(int));
  memcpy(offsets, jumps->data, jumps->size * sizeof(int));

  while (pos >= 0 && pos < jumps->size) {
    int jump = offsets[pos];

    offsets[pos] += 1;
    pos += jump;
    ret += 1;
  }

  free(offsets);
  return ret;
}

int part2(IntArray *jumps) {
  int ret = 0;
  int pos = 0;
  int *offsets = malloc(jumps->size * sizeof(int));
  memcpy(offsets, jumps->data, jumps->size * sizeof(int));

  while (pos >= 0 && pos < jumps->size) {
    int jump = offsets[pos];

    offsets[pos] += (jump >= 3) ? -1 : 1;
    pos += jump;
    ret += 1;
  }

  free(offsets);
  return ret;
}

int main() {
  IntArray test_offsets = parse("0\n3\n0\n1\n-3");
  assert(part1(&test_offsets) == 5);
  assert(part2(&test_offsets) == 10);

  char *input = read_file("../data/05.txt");
  IntArray jumps = parse(input);

  printf("part1: %d\n", part1(&jumps));
  printf("part2: %d\n", part2(&jumps));

  free(test_offsets.data);
  free(jumps.data);
  free(input);
  return 0;
}
