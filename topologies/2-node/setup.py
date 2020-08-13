import os
import shutil

node_map = {1: "node-A", 2: "node-B", 3: "node-C", 4: "node-D"}

def copy_binaries(nodes):
    current_dir = os.getcwd()
    os.chdir("../../")
    os.system("make build")
    for n in range(1, nodes + 1):
        path_to_machine_dir = os.path.join(current_dir, node_map[n])
        print("Path to machine dir", path_to_machine_dir)
        path_to_bin = os.path.join(path_to_machine_dir, "bin")
        print("Creating dir: ", path_to_bin)
        os.makedirs(path_to_bin, exist_ok=True)
        shutil.copy("./bin/server", path_to_bin)
    os.chdir(current_dir)

def main():
    nodes = 2
    copy_binaries(nodes)


if __name__ == '__main__':
    main()
