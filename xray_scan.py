import os
import subprocess
import resource
import argparse

# Создаем парсер аргументов
parser = argparse.ArgumentParser()
parser.add_argument('-d', '--domains', help='Файл со списком доменов', required=True)
args = parser.parse_args()

max_memory = 3 * 512

# Здесь мы открываем файл, указанный в аргументе -d
with open(args.domains) as f:
    for s in map(str.strip, f):
        subprocess.run(["./xray_linux_amd64", "ws", "--browser", f"http://{s}/", "--html-output", f"{s}.html"])

resource.setrlimit(resource.RLIMIT_RSS, (max_memory, max_memory))
