#ifndef PARSE_H
#define PARSE_H

#include <stdio.h>

typedef struct {
    char** lines;
    size_t length;
} FileContent;

FileContent parse_by_line(const char *filename);
void free_file_content(FileContent content);

#endif // PARSE_H
