#ifndef PARSE_H
#define PARSE_H

#include <stdio.h>

// A struct to hold both the data (array of lines) and its size.
typedef struct {
    char** lines;
    size_t length;
} FileContent;

/**
 * @brief Opens and reads a file line-by-line into a dynamically allocated array of strings.
 * * @param filename The path to the file to read.
 * @return FileContent A struct containing the array of lines and the total line count.
 * Returns a struct with lines=NULL and count=0 on failure.
 */
FileContent parse_by_line(const char *filename);

/**
 * @brief Frees the memory allocated by parse_by_line. MUST be called after use.
 * * @param content The FileContent struct returned by parse_by_line.
 */
void free_file_content(FileContent content);

#endif // PARSE_H
