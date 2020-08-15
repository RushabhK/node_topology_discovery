import json

import os
import sys

cwd = os.getcwd()
parent = os.path.abspath(os.path.join(os.path.join(cwd, os.pardir), os.pardir))
sys.path.append(cwd)
sys.path.append(parent)

from topologies.setup_utils import get_free_ports, copy_binaries


def create_configs(nodes, ports):
    current_dir = os.getcwd()

    for i in range(1, nodes + 1):
        d = {"name": "M-%s" % i, "ip_address": "localhost", "port": str(ports[i - 1])}
        d["neighbors"] = []
        for j in range(1, nodes + 1):
            if j != i:
                d["neighbors"].append({"ip_address": "localhost", "port": str(ports[j - 1])})
        path_to_config = os.path.join(current_dir, "M-%s/config.json" % i)
        with open(path_to_config, "w+") as outfile:
            json.dump(d, outfile)


def main():
    nodes = 10
    copy_binaries(nodes)
    ports = get_free_ports(nodes, 8401)
    print("PORTS:", ports)
    create_configs(nodes, ports)
    # start_servers(nodes)


if __name__ == '__main__':
    main()
