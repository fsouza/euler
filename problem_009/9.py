# -*- coding: utf-8 -*-

def try_to_find(a_elements, b_elements, c_elements):
    final = None
    for a in a_elements:
        for b in b_elements:
            for c in c_elements:
                if a < b < c and (a + b + c == 1000) and ((a ** 2 + b ** 2) == c ** 2):
                    final = a * b * c
                    break

            if final is not None:
                break

        if final is not None:
            break

    return final

if __name__ == '__main__':
    everything = range(1001)
    final = try_to_find(everything, everything, everything)

    print final
