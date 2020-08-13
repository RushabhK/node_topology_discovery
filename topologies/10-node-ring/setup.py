import json

import os
import shutil
import subprocess
from socket import *


def get_free_ports(n, start_with):
    ip = gethostbyname("localhost")
    print('Starting scan on host: ', ip)
    open_ports = []

    for i in range(start_with, 65535):
        if len(open_ports) == n:
            return open_ports
        s = socket(AF_INET, SOCK_DGRAM)

        conn = s.connect_ex((ip, i))
        if conn == 0:
            open_ports.append(i)
        s.close()
    raise Exception("Not enough free ports found.")


def copy_binaries(nodes):
    current_dir = os.getcwd()
    os.chdir("../../")
    os.system("make build")
    for n in range(1, nodes + 1):
        path_to_machine_dir = os.path.join(current_dir, "M-%s" % n)
        print("Path to machine dir", path_to_machine_dir)
        path_to_bin = os.path.join(path_to_machine_dir, "bin")
        print("Creating dir: ", path_to_bin)
        os.makedirs(path_to_bin, exist_ok=True)
        shutil.copy("./bin/server", path_to_bin)
    os.chdir(current_dir)


def run_server(cmd, machine):
    print("Running server for", machine)
    f = open(machine+".log", "w+")
    return subprocess.Popen([cmd], stdout=f), f


def start_servers(nodes):
    current_dir = os.getcwd()
    processes = []
    for i in range(1, nodes + 1):
        node_dir = os.path.join(current_dir, "M-%s" % i)
        os.chdir(node_dir)
        processes.append(run_server("./bin/server", "M-%s" % i))

    for p, f in processes:
        p.wait()
        f.seek(0)
        print(f.read())
        print("==================")
        f.close()


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
    ports = get_free_ports(nodes, 8301)
    print("PORTS:", ports)
    create_configs(nodes, ports)
    # start_servers(nodes)


if __name__ == '__main__':
    main()
