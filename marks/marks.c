#include<stdio.h>
int main() {
    int math,chem,phy,comp,eng;
     printf("Enter marks of 5 sunjects:\n");
     scanf("%d%d%d%d%d",&math,&chem,&phy,&comp,&eng); 
   int total , average;
    total=math+comp+phy+comp+eng; 
    average=total/5;
    printf("total marks %d  %d",total,average);
} 