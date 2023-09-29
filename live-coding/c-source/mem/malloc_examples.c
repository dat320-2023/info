#include <stdlib.h>
#include <stdio.h>
#include <string.h>
int main(int argc, char *argv[]) {



int *x = malloc(sizeof(int));
printf("the address of x = %p \n", &x);
printf("the value of x = %p \n", x);
printf("the value that x point to is %d  \n", *x);

int a=1;
double b=3;
printf("a= %d and b= %d\n",a, (int)b );

int len1, len2;
  
//initializing the strings
char string1[] = "Hello";
char string2[] = {'c','o','m','p','u','t','e','r','\0'};

//calculating the length of the two strings
len1 = strlen(string1);
len2 = strlen(string2);

//displaying the values
printf("Length of string1 is: %d \n", len1);
printf("Length of string2 is: %d \n", len2);



printf("----------- \n");

char *src = "hello";
char *dst = (char *) malloc(strlen(src)); // too small
printf("source length=%d, destination length= %d \n",strlen(src), strlen(dst));
strcpy(dst, src);
printf("Hey: %s\n", dst);
printf("source length=%d, destination length= %d \n",strlen(src), strlen(dst));
   
return 0;
}
