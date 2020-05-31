#include<stdio.h>

int check_equal(int x, int y){
	return !(x ^ y);
}
void main(){
	printf("3 == 4 ? %d\n", check_equal(3, 4));
	printf("3 == 3 ? %d\n", check_equal(3, 3));
}
