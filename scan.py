import os
import subprocess
import resource

max_memory = 3 * 512

with open("1.txt") as f:
    for s in map(str.strip, f):
        subprocess.run(["./xray_linux_amd64", "ws", "--browser", f"http://{s}/", "--html-output", f"{s}.html"])

resource.setrlimit(resource.RLIMIT_RSS, (max_memory, max_memory))
