import sys, os

cwd = os.getcwd()
parent = os.path.abspath(os.path.join(os.path.join(cwd, os.pardir), os.pardir))
sys.path.append(cwd)
sys.path.append(parent)

from topologies.result_utils import aggregate_results_with_custom_machines


def main():
    nodes = 4
    aggregate_results_with_custom_machines(nodes)


if __name__ == '__main__':
    main()
