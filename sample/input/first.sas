%macro addNumbers(num1, num2);
  /* Declare a local variable to hold the sum */
  %let sum = %eval(&num1 + &num2);

  /* Print the result to the log */
  %put The sum of &num1 and &num2 is &sum.;
%mend;