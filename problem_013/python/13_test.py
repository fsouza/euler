import unittest
from problem13 import slice_digits, sum_digits_of

class Problem13TestCase(unittest.TestCase):

    def test_can_sum_digits_of_a_number(self):
        expected = 9
        got = sum_digits_of(27)
        self.assertEquals(expected, got)

    def test_can_slice_digits_of_a_number(self):
        expected = "1234"
        got = slice_digits(123456789, 4)

unittest.main()
