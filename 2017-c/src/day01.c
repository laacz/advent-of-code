#include "utils.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int part1(char *input) {
  int ret = 0;

  for (int i = 0; i < strlen(input); i++) {
    char next = input[(i + 1) % strlen(input)];

    if (next == input[i]) {
      ret += input[i] - '0';
    }
  }
  return ret;
}

int part2(char *input) {
  int ret = 0;
  int len = strlen(input);

  for (int i = 0; i < len; i++) {
    char next = input[(i + len / 2) % len];

    if (next == input[i]) {
      ret += input[i] - '0';
    }
  }
  return ret;
}

int main() {
  assert(part1("1122") == 3);
  assert(part1("1111") == 4);
  assert(part1("1234") == 0);
  assert(part1("91212129") == 9);

  assert(part2("1212") == 6);
  assert(part2("1221") == 0);
  assert(part2("123425") == 4);
  assert(part2("123123") == 12);
  assert(part2("12131415") == 4);

  char *input = read_file("../data/01.txt");

  printf("part1: %d\n", part1(input));
  printf("part2: %d\n", part2(input));

  free(input);

  return 0;
}
