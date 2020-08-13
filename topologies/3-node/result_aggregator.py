def main():
    nodes = 3
    node_map = {1: "A", 2: "B", 3: "C", 4: "D"}
    fw = open("results.txt", "w+")
    line_break = "______________________________________________________________________________________________________\n\n"
    file_content = ""
    for i in range(1, nodes + 1):
        result_file = "node-%s/result.txt" % node_map[i]
        with open(result_file) as fr:
            file_content += fr.read()
            file_content += line_break
    fw.write(file_content)
    fw.close()


if __name__ == '__main__':
    main()
