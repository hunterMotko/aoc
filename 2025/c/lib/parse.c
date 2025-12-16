#include "../include/parse.h"
#include <stdlib.h>
#include <string.h>

#define INITIAL_CAPACITY 10

FileContent parse_by_line(const char *filename) {
  FileContent result = {NULL, 0};
  FILE *file = NULL;
  char *line = NULL;
  size_t len = 0;
  ssize_t read;
  file = fopen(filename, "r");
  if (file == NULL) {
    perror("Error opening file");
    return result; 
  }

  size_t capacity = INITIAL_CAPACITY;
  result.lines = (char **)malloc(capacity * sizeof(char *));
  if (result.lines == NULL) {
    fclose(file);
    return result;
  }

  while ((read = getline(&line, &len, file)) != -1) {
    if (result.length >= capacity) {
      capacity *= 2;
      char **temp = (char **)realloc(result.lines, capacity * sizeof(char *));
      if (temp == NULL) {
        free(line);
        free_file_content(result);
        fclose(file);
        return (FileContent){NULL, 0};
      }
      result.lines = temp;
    }
    if (read > 0 && line[read - 1] == '\n') {
      line[read - 1] = '\0';
    }
    result.lines[result.length] = strdup(line); 
    if (result.lines[result.length] == NULL) {
      free(line);
      free_file_content(result);
      fclose(file);
      return (FileContent){NULL, 0};
    }
    result.length++;
  }

  free(line);
  fclose(file);
  return result;
}

void free_file_content(FileContent content) {
  if (content.lines) {
    for (size_t i = 0; i < content.length; i++) {
      free(content.lines[i]);
    }
    free(content.lines);
  }
}
