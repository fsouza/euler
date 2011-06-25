import os
import sys

MAKEFILE_TEMPLATE = r'''
include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
		%(problem)d.go

include $(GOROOT)/src/Make.pkg
'''

def fill_file(path, problem):
    if path.endswith('.go'):
        content = "package main"
    else:
        content = MAKEFILE_TEMPLATE % { 'problem' : problem }

    fp = open(path, 'w')
    fp.write(content)
    fp.close()

ROOT_DIR = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
if __name__ == '__main__':
    problem = int(sys.argv[1])
    print 'Generating structure for problem %d... ' % problem,

    problem_dir = os.path.join(ROOT_DIR, 'problem_%03d' % problem)
    os.mkdir(problem_dir)

    files = "%(problem)d.go %(problem)d_test.go Makefile" % { 'problem' : problem }
    files = files.split()
    file_paths = [os.path.join(problem_dir, filename) for filename in files]

    for path in file_paths:
        fill_file(path, problem)

    print 'done.'
