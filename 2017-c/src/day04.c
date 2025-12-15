#include "utils.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int has_dupes(char words[16][64], int num_words) {
  for (int i = 0; i < num_words; i++) {
    for (int j = i + 1; j < num_words; j++) {
      if (strcmp(words[i], words[j]) == 0) {
        return 1;
      }
    }
  }

  return 0;
}

int part1(char *passwords) {
  char *copy = strdup(passwords);
  char *line = strtok(copy, "\n");
  int valid = 0;

  while (line != NULL) {
    char words[16][64];
    int count = 0;

    char word[64];
    int n;
    char *p = line;

    while (sscanf(p, "%s%n", word, &n) == 1) {
      strcpy(words[count++], word);
      p += n;
    }

    valid += !has_dupes(words, count);
    line = strtok(NULL, "\n");
  }

  free(copy);
  return valid;
}

// for qsort
int cmpchar(const void *a, const void *b) {
  return *(const char *)a - *(const char *)b;
}

int part2(char *passwords) {
  char *copy = strdup(passwords);
  char *line = strtok(copy, "\n");
  int valid = 0;

  while (line != NULL) {
    char words[16][64];
    int count = 0;

    char word[64];
    int n;
    char *p = line;

    while (sscanf(p, "%s%n", word, &n) == 1) {
      qsort(word, strlen(word), sizeof(char), cmpchar);
      strcpy(words[count++], word);
      p += n;
    }

    valid += !has_dupes(words, count);
    line = strtok(NULL, "\n");
  }

  free(copy);
  return valid;
}

int main() {
  assert(part1("aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa") == 2);
  assert(part2("abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii "
               "ooii oooi oooo\noiii ioii iioi iiio") == 3);

  char *input = read_file("../data/04.txt");

  printf("part1: %d\n", part1(input));
  printf("part2: %d\n", part2(input));

  free(input);
  return 0;
}
