# Code example
```c 
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <assert.h>
#include "common.h"
#include <unistd.h>
int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "usage: cpu <string>\n");
        exit(1);
    }
    char *str = argv[1];
    while (1) {
        Spin(1);
        printf("(%d) %s\n",(int) getpid(), str);
    }
    return 0;
}
```

# Compile example
`gcc -o bin/cpu cpu.c`
# Content
## cpu.c
Takes an argument and prints it over and over again every 1s. For c-source folder run `./bin/cpu a` to see. 


