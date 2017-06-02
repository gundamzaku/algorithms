/*
 ============================================================================
 Name        : algorithm.c
 Author      : 
 Version     :
 Copyright   : Your copyright notice
 Description : Hello World in C, Ansi-style
 ============================================================================
 */

#include <stdio.h>
#include <stdlib.h>

typedef struct{
	int low;
	int high;
	int sum;
}fdata;

int g = 0;

fdata FindMaxCrossingSubarray(int a[],int low,int mid,int high){
	printf("FindMaxCrossingSubarray传入:low_%d,mid_%d,high_%d\n",low,mid,high);
	int i;
	int sum = 0;
	int left_sum = 0;
	fdata cross;

	for(i=mid;i>low;i--){
		sum = sum+a[i];
		if(sum>left_sum){
			left_sum = sum;
			cross.low = i;
		}
	}

	int right_sum = 0;
	sum = 0;
	int j = 0;
	for(j=mid+1;j<high;j++){
		sum = sum+a[j];
		if( sum>right_sum ){
			right_sum = sum;
			cross.high = j;
		}
	}
	cross.sum = left_sum+right_sum;
	return cross;

}

fdata FindMaximumSubarray(int a[],int low,int high){

	printf("FindMaximumSubarray传入:low_%d,high_%d\n",low,high);
	int mid = 0;
	fdata left;
	fdata right;
	fdata cross;

	if(high == low){
		left.high = high;
		left.low = low;
		left.sum = a[low];
		return left ;
	}else{
		mid = (low+high)/2;
		left = FindMaximumSubarray(a,low,mid);
		right = FindMaximumSubarray(a,mid+1,high);
		cross = FindMaxCrossingSubarray(a,low,mid,high);
		printf("left:high_%d ,low_%d ,sum_%d \n",left.high,left.low,left.sum);
		printf("right:high_%d ,low_%d ,sum_%d \n",right.high,right.low,right.sum);
		printf("cross:high_%d ,low_%d ,sum_%d \n",cross.high,cross.low,cross.sum);
		printf("进行%d次比较_left_sum:%d_right_sum:%d_cross_sum:%d\n\n",g++,left.sum,right.sum,cross.sum);
		if(left.sum>=right.sum && left.sum>=cross.sum){
			return left;
		}else if(right.sum>=left.sum && right.sum>=cross.sum){
			return right;
		}else{
			return cross;
		}
	}
}


int main(void) {
	fdata cross;
	cross.low = 0;
	cross.high = 0;
	cross.sum = 0;

	int data[16] = {13,-3,-25,20,-3,-16,-23,18 ,20,-7,12,-5,-2,19,-4,28};
	cross = FindMaximumSubarray(data,0,15);
	printf("%d\n",cross.low);
	printf("%d\n",cross.high);
	printf("%d\n",cross.sum);
	return EXIT_SUCCESS;
}

