# ValidateSubSeq-TestCase

A) Validating Sub Sequence example from Algo Expert. (https://www.algoexpert.io/questions/validate-subsequence)
  2 sets of array where need to determine if the 2nd array is a subset of the 1st array with testcase.


B) Ultilizing TestCase steps:
  1. Create the Test file by running from VS code console.
      >Go Generate Unit Test For Function
  
  2. Fill in the variables for the Test Case where below comment prompt.
    // TODO: Add test cases.
		{
			name: "Hello",
			args: args{
				array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
				sequence: []int{25, 10},
			},
			want: true,
		},
    
   3. The testcase will Pass/Fail base on your want (true/false) when "RUN and Debug".
  
  To understand how to run debug using testcase, do refer to youtube vid for steps.
  https://www.youtube.com/watch?v=6r08zGi38Tk
