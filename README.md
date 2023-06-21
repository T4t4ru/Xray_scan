# Project

## Compile and run

To compile and download the C++ code, follow the instructions:

1. Assemble the code into a file, example `main.cpp `.
2. Open the terminal and go to the headmistress with the site `main.cpp `.
3. Compile the code marked `g++': `g++ -o program main.cpp `.
4. Register the program with the name as administrator: `./program 1.txt `.

Make sure you have the `g++` compiler installed. If not, you can install it using the package manager for your Linux distribution, example `apt` for Ubuntu:

## Project Description

This is a C++ program that performs the following actions:

1. Reads the file specified in the command line argument. Each line of the file must contain a domain.
2. For each domain, the command `./xray_linux_amd64 ws --browser http://<domain>/ --html-output <domain> is run.html`.
3. Sets a limit on memory usage.

## Restrictions and Warnings

Please note that the `system` function in C++ may present potential security issues. You may need to replace this function with safer options such as `fork` and `exec'.

The program has been tested and developed for use on Linux. If you are using a different operating system, you will need to adapt the code.
The information provided here is for educational purposes only. It is not intended to be used in illegal activities or actions that violate the rights of other people or organizations.

As the author, I am not responsible for any illegal actions related to the use of this information, and I do not support the illegal use of hacking skills and tools.

Users who use this information for illegal purposes do so at their own risk and should be prepared for possible legal consequences.

It is strongly recommended to use the acquired knowledge only for educational or legal purposes, for example, to strengthen the protection of your own computer system.

## License

This project is distributed under the MIT license. See the `LICENSE` file for more information.
