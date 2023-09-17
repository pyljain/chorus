import unittest
from io import StringIO
import sys
from first import add_numbers


class Testing(unittest.TestCase):
    def test_add_numbers(self):
        expected_output = 'The sum of 10 and 20 is 30.'

        # Redirect standard output to StringIO
        output = StringIO()
        sys.stdout = output

        # Call the function
        add_numbers(10, 20)

        # Get the output
        outputStr = output.getvalue().strip()

        # Check the output
        self.assertEqual(outputStr, expected_output)


if __name__ == '__main__':
    unittest.main()