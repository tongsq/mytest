#include<stdio.h>
int strlonger(char *s, char *t){
	return strlen(s) > strlen(t);
}
void main(){
	char * str1 = "";
	char * str2 = "hello";
	printf("str1: %s, str1 length:%u\n", str1, strlen(str1));
	printf("str2: %s, str2 length:%u\n", str2, strlen(str2));
	printf("strlonger(str1, str2) :%d\n", strlonger(str1, str2));
}
