/*
 *数字存的是补码
 *float存储格式：最高位1表示负数,后面8位表示2的指数，需要做+127移位处理，后23位为尾数
 *linux采用小端存储，低位数字存在低字节
 **/
#include<stdio.h>

typedef unsigned char * byte_pointer;

void show_bytes(byte_pointer start, int len) {
	int i;
	for (i = 0; i < len; i++)
		printf(" %.2x", start[i]);
	printf("\n");
}

void show_int(int x) {
	show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x){
	show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x){
	show_bytes((byte_pointer) &x, sizeof(void *));
}

void main(){
	int i = -2;
	float f = 1.5;
	void * x;
	show_int(i);
	printf("%.2f", f);
	show_float(f);
	show_pointer(x);
	float f2 = 17.0/3;
	show_float(f2);
}
