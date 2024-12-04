#include <stdlib.h>

typedef struct {
  int key;
  int count;
  struct ht_entry *next;
} ht_entry;

typedef struct {
  ht_entry **table;
  int size;
  int count;
} hash_table;

hash_table *create_hash(int size) {
  hash_table *table = (hash_table *)malloc(sizeof(hash_table));
  table->size = size;
  table->count = 0;
  table->table = (ht_entry **)calloc(size, sizeof(ht_entry *));
  return table;
}

int hash(int key, int size) { return key % size; }

void ht_insert(hash_table *table, int key) {
  int index = hash(key, table->size);
  ht_entry *entry = table->table[index];
  if (entry == NULL) {
    entry = (ht_entry *)malloc(sizeof(ht_entry));
    entry->key = key;
    entry->count = 1;
    table->table[index] = entry;
  } else {
    entry->count++;
  }
  table->count++;
}

int ht_get_count(hash_table *table, int key) {
  int index = hash(key, table->size);
  ht_entry *entry = table->table[index];
  if (entry != NULL) {
    return entry->count;
  } else {
    return 0; // key not found
  }
}
