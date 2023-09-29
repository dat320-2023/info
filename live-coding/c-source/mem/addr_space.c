#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    int x = 3;
    int unit=1;
    int *xx = malloc(unit * sizeof(int));
    printf("----Before malloc-------\n");
    printf("location of code : %p\n", main);
    printf("location of heap : %p\n", xx);
    printf("location of stack: %p\n", &x);

    
    printf("----After malloc-------\n");
    int *yy = malloc(unit * sizeof(int));
    int *zz = malloc(unit * sizeof(int));
    printf("location of code : %p\n", main);
    printf("location xx in the heap : %p\n", xx);
    printf("location yy in the heap : %p\n", yy);
    printf("location yy in the heap : %p\n", zz);

    printf("location of stack: %p\n", &x);

    return 0;
}