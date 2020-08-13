def main():
    nodes = 10
    fw = open("results.txt", "w+")
    line_break = "______________________________________________________________________________________________________\n\n"
    file_content = ""
    for i in range(1, nodes + 1):
        result_file = "M-%s/result.txt" % i
        with open(result_file) as fr:
            file_content += fr.read()
            file_content += line_break
    fw.write(file_content)
    fw.close()


if __name__ == '__main__':
    main()
