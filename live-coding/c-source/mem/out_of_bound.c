#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    int *x= malloc(sizeof(int)); // need space for 1 int
    int *arr= malloc(sizeof(int)*100); // need space for 100 ints

    *x= 120; // value that x points to is 120
    arr[90]=50;
    arr[91]=6;
    int in_bound=99;
    arr[in_bound]=75;  
    printf("size=100, arr[90]= %d, arr[91]=%d, arr[%d]=%d \n", arr[90], arr[91],in_bound, arr[in_bound] );

    int out_of_bound=101;
    arr[out_of_bound]=222;
    printf("size=100, arr[90]= %d, arr[91]=%d, arr[%d]=%d \n", arr[90], arr[91],out_of_bound, arr[out_of_bound] );

    printf("---free-------\n");

    //free(arr);
    printf(" arr[90]= %d, arr[91]=%d, arr[%d]=%d \n", arr[90], arr[91],in_bound, arr[in_bound] );
    printf(" arr[90]= %d, arr[91]=%d, arr[%d]=%d \n", arr[90], arr[91],out_of_bound, arr[out_of_bound] );


    return 0;
}