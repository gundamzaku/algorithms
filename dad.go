package main
import "fmt"

/*
分治算法的实现
 */

func main()  {
	var low int32;
	var high int32;
	var sum int32;

	var data = [16]int32{13,-3,-25,20,-3,-16,-23,18 ,20,-7,12,-5,-22,15,-4,7};
	low,high,sum = FindMaximumSubarray(data,0,15);
	fmt.Println(low);
	fmt.Println(high);
	fmt.Println(sum);
}

func FindMaxCrossingSubarray(a [16]int32,low int32,mid int32,high int32) (int32,int32,int32){

	var left_sum int32  = 0;
	var sum int32 = 0;
	var max_left int32 = 0;

	for i := mid;i>low ;i--  {
		sum = sum+a[i];
		if(sum>left_sum){
			left_sum = sum;
			max_left = i;
		}
	}
	//fmt.Println("left_sum is : ",left_sum);
	var right_sum int32 = 0;
	sum = 0;
	var max_right int32 = 0;

	for j:=mid+1;j<high;j++{

		sum = sum+a[j];
		if(sum>right_sum){
			right_sum = sum;
			max_right = j;
		}
	}
	//fmt.Println("right_sum is : ",right_sum);
	return max_left,max_right,left_sum+right_sum;

}

func FindMaximumSubarrayLeft(a [16]int32,low int32,high int32)(int32,int32,int32){

	var left_low int32;
	var left_high int32;
	var left_sum int32;

	if(high == low){
		return low,high,a[low];
	}else {
		mid := (low+high)/2;
		left_low,left_high,left_sum = FindMaximumSubarrayLeft(a,low,mid);

		return left_low,left_high,left_sum;

	}
}

func FindMaximumSubarrayRight(a [16]int32,low int32,high int32)(int32,int32,int32){

	var right_low int32;
	var right_high int32;
	var right_sum int32;

	if(high == low){
		return low,high,a[low];
	}else {
		mid := (low+high)/2;
		fmt.Println(mid+1,high);
		right_low,right_high,right_sum = FindMaximumSubarrayRight(a,mid+1,high);

		return right_low,right_high,right_sum;
	}
}

func FindMaximumSubarray(a [16]int32,low int32,high int32)(int32,int32,int32){
	fmt.Println("low is ",low,"high is ",high);
	var left_low int32;
	var left_high int32;
	var left_sum int32;

	var right_low int32;
	var right_high int32;
	var right_sum int32;

	var cross_low int32;
	var cross_high int32;
	var cross_sum int32;
	if(high == low){
		fmt.Println("return ",low,high);
		return low,high,a[low];
	}else {
		fmt.Println("high is ",high);
		mid := (low+high)/2;
		fmt.Println("mid is ",mid);
		left_low,left_high,left_sum = FindMaximumSubarray(a,low,mid);
		fmt.Println(left_low,left_high,left_sum);
		fmt.Println(mid,high);
		right_low,right_high,right_sum = FindMaximumSubarray(a,mid+1,high);

		cross_low,cross_high,cross_sum = FindMaxCrossingSubarray(a,low,mid,high);

		fmt.Println("这是最后一步？");
		if(left_sum>=right_sum && left_sum>=cross_sum){
			return left_low,left_high,left_sum;
		}else if(right_sum>=left_sum && right_sum>=cross_sum){
			return right_low,right_high,right_sum;
		}else{
			return cross_low,cross_high,cross_sum;
		}

	}
}