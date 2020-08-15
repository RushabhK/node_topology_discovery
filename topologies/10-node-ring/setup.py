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
    d = {"name": "M-1", "ip_address": "localhost", "port": str(ports[0])}
    d["neighbors"] = []
    d["neighbors"].append({"ip_address": "localhost", "port": str(ports[1])})
    d["neighbors"].append({"ip_address": "localhost", "port": str(ports[nodes - 1])})
    path_to_config = os.path.join(current_dir, "M-1/config.json")
    with open(path_to_config, "w+") as outfile:
        json.dump(d, outfile)

    for i in range(2, nodes + 1):
        d = {"name": "M-%s" % i, "ip_address": "localhost", "port": str(ports[i - 1])}
        d["neighbors"] = []
        d["neighbors"].append({"ip_address": "localhost", "port": str(ports[i - 2])})
        d["neighbors"].append({"ip_address": "localhost", "port": str(ports[i % nodes])})
        path_to_config = os.path.join(current_dir, "M-%s/config.json" % i)
        with open(path_to_config, "w+") as outfile:
            json.dump(d, outfile)


def main():
    nodes = 10
    copy_binaries(nodes)
    ports = get_free_ports(nodes, 8501)
    print("PORTS:", ports)
    create_configs(nodes, ports)
    # start_servers(nodes)


if __name__ == '__main__':
    main()
