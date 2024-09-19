#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Map {
  size_t size;
  size_t cap;
  void **buckets;
} Map;
const size_t BASE = 0x811c9dc5;
const size_t PRIME = 0x01000193;
size_t hash(Map *m, char *str) {
  size_t initial = BASE;
  while (*str) {
    initial ^= *str++ * PRIME;
  }
  return initial & (m->cap - 1);
}
Map init(size_t cap) {
  Map m = {0, cap};
  m.buckets = malloc(sizeof(void *) * m.cap);
  assert(m.buckets != NULL);
  return m;
}
void put(Map *m, char *str, void *value) {
  m->size++;
  m->buckets[hash(m, str)] = value;
}
void *get(Map *m, char *str) { return m->buckets[hash(m, str)]; }

int main(void) {
  Map m = init(1024);
  double d1 = 25.0;
  double d2 = 50.0;
  put(&m, "key1", (void *)&d1);
  put(&m, "key2", (void *)&d2);
  printf("key1=%f;key2=%f\n", *(double *)get(&m, "key1"),
         *(double *)get(&m, "key2"));
  free(m.buckets);
  return EXIT_SUCCESS;
}
