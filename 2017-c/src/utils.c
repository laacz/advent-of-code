#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int is_space(const char c) {
  return c == ' ' || c == '\t' || c == '\n' || c == '\r';
}

char *read_file(const char *path) {
  FILE *f = fopen(path, "r");
  if (!f)
    return NULL;

  fseek(f, 0, SEEK_END);
  long size = ftell(f);
  fseek(f, 0, SEEK_SET);

  char *buf = malloc(size + 1);
  if (!buf) {
    fclose(f);
    return NULL;
  }

  fread(buf, 1, size, f);
  buf[size] = '\0';
  fclose(f);

  // trim
  char *end = buf + strlen(buf) - 1;
  while (end > buf && is_space(*end))
    end--;
  *(end + 1) = '\0';

  char *start = buf;
  while (is_space(*start))
    start++;
  if (start != buf)
    memmove(buf, start, strlen(start) + 1);

  return buf;
}
