#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sys/resource.h>
#include <cstdlib>

int main(int argc, char *argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: ./program filename" << std::endl;
        return 1;
    }
    
    const char* filename = argv[1];
    std::ifstream file(filename);
    
    if (!file) {
        std::cerr << "Cannot open the file: " << filename << std::endl;
        return 1;
    }
    
    std::string domain;
    while (std::getline(file, domain)) {
        std::string cmd = "./xray_linux_amd64 ws --browser http://" + domain + "/ --html-output " + domain + ".html";
        system(cmd.c_str());
    }
    
    struct rlimit rl;
    rl.rlim_cur = 3 * 512;
    rl.rlim_max = 3 * 512;
    if (setrlimit(RLIMIT_RSS, &rl) == -1) {
        std::cerr << "Error in setrlimit" << std::endl;
        return 1;
    }
    
    return 0;
}
