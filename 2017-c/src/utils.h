#ifndef UTILS_H
#define UTILS_H

// Don't do MAX(a++, b) or you're doomed for a debugging session nobody asked
// for. Keywords: double evaluation.
#define MAX(a, b) ((a) > (b) ? (a) : (b))
#define MIN(a, b) ((a) < (b) ? (a) : (b))

char *read_file(const char *path);

#endif
