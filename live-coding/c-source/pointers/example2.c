#include <stdio.h>

int * increment_in_stack(int v){
        int c=v;
        c+=1;
        int *d=&c;
        return d;
    }
    
int main()
{
    int *a;
    printf("the address of a is %p and the value there is %d \n", a, *a);
    a=increment_in_stack(10);
    printf("Now the address of a is %p and the value there is %d \n", a, *a);
    printf("Kind of works, but nor really \n");
    printf("---------------- \n");

    int *b;
    printf("the address of b is %p and the value there is %d \n", b, *b);
    b=increment_in_stack(50);
    printf("Now the address of a is %p and the value there is %d \n", a, *a);
    printf("Now the address of b is %p and the value there is %d \n", b, *b);
    printf("Now the address of a is %p and the value there is %d \n", a, *a);
    printf("WHAT !!!!!!!!!!!!\n");
    int i=0;
    while (i<10)
    {
       i++;
       printf("..\n");
    }
    printf("Now the address of a is %p and the value there is %d \n", a, *a);
     printf("Now the address of b is %p and the value there is %d \n", b, *b);

    printf("NO REALLY WHAT, THIS IS AN UNDEFINED BEHAVIOR !!!!!!!!!!!!\n");
    
    return 0;
}