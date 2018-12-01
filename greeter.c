#include "greeter.h"
#include <stdio.h>

int greet(struct Greetee *g, char *out) {
    int n;

    // sprintf does write a trailing NULL. Remember
    // those?
    n = sprintf(out,
        "Greetings, %s from %d! All your base are belong to us! Ha haha ha ha.",
        g->name,
        g->year);
    
    // sprintf returns the number of bytes written
    // to the buffer.
    return n;
}