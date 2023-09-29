#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    printf("location of code : %p\n", main);
    printf("location of heap : %p\n", malloc(100e6));
    int x = 3;
    printf("location of stack: %p\n", &x);

    printf("---------------------\n");
    int *xx = malloc(10 * sizeof(int));
   

    xx[15]=200;
    for (int i=0; i< 20; i++){
        //xx[i]=i+1;
        printf("index=%d, value= %d \n", i,xx[i]);

    }
    printf("sizeof (xx) =%lu, and its value is  %p \n", sizeof(xx), xx); 


    printf("------------------------\n");
    int days[] = {1,2,3,4,5,6,7};
    int *ptr = days;
    
    printf("Header %d \n",  ptr[-4] );
    printf("size of days[] %lu in bytes\n", sizeof(days));
    printf("size of int %lu in bytes\n", sizeof(int));
    printf("size of days[] %lu \n", sizeof(days)/sizeof(int));
    printf("%lu\n", sizeof(days)/sizeof(days[0]));
    printf("size of ptr in bytes %lu\n", sizeof(ptr)); 
    printf("%d %d\n", *ptr, *days); 
    return 0;
}