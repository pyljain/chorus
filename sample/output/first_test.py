import unittest
import sys
from io import StringIO
import first

class TestAddNumbers(unittest.TestCase):

    def test_add_numbers(self):
        """Test the function add_numbers."""
        # Save the current value of sys.stdout
        old_stdout = sys.stdout
        # Temporarily redirect sys.stdout to a StringIO object
        sys.stdout = StringIO()
        first.add_numbers(3, 5)
        # Get the printed output
        output = sys.stdout.getvalue().strip()
        # Put back sys.stdout to its original setting
        sys.stdout = old_stdout
        # Define the expected output
        expected_output = 'The sum of 3 and 5 is 8.'
        # Check if the produced output matches the expected one
        self.assertEqual(output, expected_output)

if __name__ == '__main__':
    unittest.main()
