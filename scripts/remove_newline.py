#!/usr/bin/env python3

import sys
import os

if __name__ == "__main__":
    data = []
    sep = " "
    while True:
        line = sys.stdin.readline()
        if not line:
            break
        
        data.append(line.rstrip(os.linesep))

    sys.stdout.write(" ".join(data))