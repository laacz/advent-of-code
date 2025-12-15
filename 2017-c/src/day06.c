#include "utils.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int *data;
  size_t size;
  int seen[10000][16];
  size_t pos;
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

int seen_before(IntArray *memory) {
  for (size_t i = 0; i < memory->pos; i++) {
    if (memcmp(memory->seen[i], memory->data, memory->size * sizeof(int)) == 0)
      return 1;
  }
  return 0;
}

void record(IntArray *memory) {
  memcpy(memory->seen[memory->pos++], memory->data, memory->size * sizeof(int));
}

void redistribute(IntArray *memory) {
  int max = 0, max_idx = 0;
  for (size_t i = 0; i < memory->size; i++) {
    if (memory->data[i] > max) {
      max = memory->data[i];
      max_idx = i;
    }
  }

  int blocks = max;
  memory->data[max_idx] = 0;

  for (int i = 0; i < blocks; i++) {
    max_idx = (max_idx + 1) % memory->size;
    memory->data[max_idx]++;
  }
}

int part1(IntArray *memory) {
  int ret = 0;

  while (!seen_before(memory)) {
    record(memory);
    redistribute(memory);
    ret++;
  }
  return ret;
}

int part2(IntArray *memory) {
  int ret = 0;
  int target[16];
  memcpy(target, memory->data, memory->size * sizeof(int));

  do {
    redistribute(memory);
    ret++;
  } while (memcmp(memory->data, target, memory->size * sizeof(int)) != 0);

  return ret;
}

int main() {
  IntArray test_memory = parse("0 2 7 0");
  assert(part1(&test_memory) == 5);
  assert(part2(&test_memory) == 4);
  free(test_memory.data);

  char *input = read_file("../data/06.txt");
  IntArray memory = parse(input);

  printf("part1: %d\n", part1(&memory));
  printf("part2: %d\n", part2(&memory));

  free(input);
  free(memory.data);
  return 0;
}
