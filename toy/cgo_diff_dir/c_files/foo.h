#ifndef FOO_H 
#define FOO_H 
#include <stdio.h> 
#include <stdlib.h>
#ifdef __cplusplus 
extern "C" 
{ 
#endif
    extern int Num;
    extern void foo(); 
    extern int f_add(int a, int b); 
#ifdef __cplusplus 
} 
#endif 
#endif  /*FOO_H*/
