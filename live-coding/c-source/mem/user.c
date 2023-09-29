#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    printf("location of code : %p\n", main);
    printf("location of heap : %p\n", malloc(100e6));
    int x = 3;
    printf("location of stack: %p\n", &x);

    printf("---------------------\n");
    int *xx = malloc(10 * sizeof(int));
    
    xx[0]=150;
    free(xx);
    //free(xx);
    printf(" allocated and freed xx \n");
    //double *yy=malloc(100 *sizeof(double));
    for (int i=0; i< 10; i++){
        xx[i]=i+1;
        printf("index=%d, value= %d , address of value %p \n", i,xx[i], &xx[i]);
        //if (i==9){
        //    free(xx);
       // }
        

    }
    printf("xx: %p\n", xx);
    printf("sizeof (xx) =%lu, and its value is  %p \n", sizeof(xx), xx); 
    // printf("------------------------\n");
    // int days[] = {1,2,3,4,5,6,7};
    // days[30]=155;
    // printf("days[30] = %d \n", days[30]);

    return 0;
}