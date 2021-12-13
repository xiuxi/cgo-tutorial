#include "greeter.h"
#include <stdio.h>

int greet(const char *name, int year, char *out) {
    int n;

    // sprintf does write a trailing NULL. Remember
    // those?
    n = sprintf(out,
        "Greetings, %s from %d! All your base are belong to us! Ha haha ha ha.",
        name,
        year);
    
    // sprintf returns the number of bytes written
    // to the buffer.
    return n;
}