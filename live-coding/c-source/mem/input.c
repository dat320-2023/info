#include <stdio.h>
#include <stdlib.h>

int main() {
   int n = 4, i, *p, s = 0;
   p = (int*) malloc(n * sizeof(int));
   free(p);
   
   if(p == NULL) {
      printf("\nError! memory not allocated.");
      exit(0);
   }

   for(i = 0; i < n; i++) {
    printf("\nEnter elements %d / %d of array to stored at %p : ", i+1,n,p+i);
    scanf("%d", p + i);
    s += *(p + i);
   }
   printf("\nSum : %d \n", s);
   return 0;
}