#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct {
  int x;
  int y;
} Point;

Point coords(int n) {
  if (n == 1)
    return (Point){0, 0};

  int ring = 0;
  while ((2 * ring + 1) * (2 * ring + 1) < n) {
    ring++;
  }

  int ring_start = (2 * ring - 1) * (2 * ring - 1) + 1;
  int offset = n - ring_start;
  int side_len = 2 * ring;
  int side = offset / side_len;
  int pos = offset % side_len - ring + 1;

  switch (side) {
  case 0:
    return (Point){ring, pos};
  case 1:
    return (Point){-pos, ring};
  case 2:
    return (Point){-ring, -pos};
  case 3:
    return (Point){pos, -ring};
  default:
    return (Point){0, 0};
  }
}

int part1(int index) {
  Point c = coords(index);
  return abs(c.x) + abs(c.y);
}

int part2(int target) {
  // woulda go with define, but this is more interesting
  // this is for offset (no negative indices for arrays in C)
  enum {
    SIZE = 101,
    MID = SIZE / 2,
  };

  int grid[SIZE][SIZE] = {0};
  grid[MID][MID] = 1;

  int x = 0, y = 0;
  int dx = 1, dy = 0;
  int steps = 1, step_count = 0, turns = 0;

  while (1) {
    x += dx;
    y += dy;
    step_count++;

    int sum = 0;
    for (int i = -1; i <= 1; i++)
      for (int j = -1; j <= 1; j++)
        sum += grid[MID + x + i][MID + y + j];

    grid[MID + x][MID + y] = sum;
    if (sum > target)
      return sum;

    if (step_count == steps) {
      // turn left
      step_count = 0;
      int tmp = dx;
      dx = -dy;
      dy = tmp;
      if (++turns == 2) {
        turns = 0;
        steps++;
      }
    }
  }
}

int main() {
  assert(part1(1) == 0);
  assert(part1(12) == 3);
  assert(part1(23) == 2);
  assert(part1(1024) == 31);

  assert(part2(1) == 2);
  assert(part2(2) == 4);
  assert(part2(3) == 4);
  assert(part2(4) == 5);
  assert(part2(5) == 10);
  assert(part2(11) == 23);
  assert(part2(23) == 25);
  assert(part2(747) == 806);

  printf("part1: %d\n", part1(361527));
  printf("part2: %d\n", part2(361527));

  return 0;
}
