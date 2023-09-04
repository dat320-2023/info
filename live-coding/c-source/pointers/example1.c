#include <stdio.h>
int main()
{
printf("-----Step 1-------\n");
int *a; 
printf("the address of a = %p \n", &a);
printf("the content of a = %p \n", a);
printf("the value that the content of a is pointing at *a = %d (undefined)\n", *a);
printf("-----Step 2-------\n");
int b=3;
printf("the address of b (&b) = %p \n", &b);
printf("the content of b = %d \n", b);
printf("-----Step 3-------\n");
printf("make the content of 'a' point the value b: \n");
*a=b;
printf("the address of a = %p \n", &a);
printf("the address of b (&b) = %p \n", &b);
printf("the content of a = %p \n", a);
printf("the value that the content of a is pointing at *a = %d = b = %d\n", *a, b);
printf("-----Step 4-------\n");
int *lost_value=a;
b++;
printf("make the address of a be the address of b: \n");
a=&b;
printf("the address of a = %p \n", &a);
printf("the content of a = %p = &b=%p \n", a,&b);
printf("the value that the address of a is pointing to *a = %d = b = %d\n", *a, b);
printf("!!!!! lost pointer !!!!!!!!\n");
printf("address %p, content %p, value %d", &lost_value, lost_value, *lost_value);
return 0;
}