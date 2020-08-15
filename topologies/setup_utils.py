import os
import shutil
from multiprocessing import Process
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
    os.system(cmd)


def start_servers(nodes):
    current_dir = os.getcwd()
    processes = []
    for i in range(1, nodes + 1):
        node_dir = os.path.join(current_dir, "M-%s" % i)
        os.chdir(node_dir)
        p = Process(target=run_server, args=("./bin/server", "M-%s" % i))
        p.start()
        processes.append(p)

    for p in processes:
        p.join()
