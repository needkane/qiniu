// foo.c 
#include <stdio.h> 
#include "add.h" 
int Num = 8; 

void foo()
{ 
    printf("First line.\n"); 
}

int f_add(int a, int b) 
{ 
    return add(a,b);
}
