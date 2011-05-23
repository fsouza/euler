all: problem

run:
	@./run.sh $(problem)

problem:
	@python scripts/new_problem.py $(problem)
