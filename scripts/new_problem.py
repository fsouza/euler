import os
import sys

ROOT_DIR = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
if __name__ == '__main__':
    problem = int(sys.argv[1])
    print 'Generating structure for problem %d... ' % problem,
    problem_dir = os.path.join(ROOT_DIR, 'problem_%03d' % problem)
    file_path = os.path.join(problem_dir, '%d.go' % problem)
    os.mkdir(problem_dir)
    open(file_path, 'w').close()
    print 'Done.'
