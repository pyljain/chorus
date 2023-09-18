import unittest
import io
import sys
from first import add_numbers

class Testing(unittest.TestCase):
    def test_addNumbers(self):
        capturedOutput = io.StringIO()  # Create StringIO object
        sys.stdout = capturedOutput  # Redirect stdout to StringIO object
        add_numbers(5, 10)  # Call the function
        sys.stdout = sys.__stdout__  # Reset redirect
        output = capturedOutput.getvalue().strip()  # Get the captured output
        expected_output = 'The sum of 5 and 10 is 15'
        self.assertEqual(output, expected_output)  # Check if the output matches the expected output

if __name__ == '__main__':
    unittest.main()