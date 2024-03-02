#include "./resources.h"
#include <sched.h>
#include <iostream>
#include <sys/types.h>
#include <unistd.h>

int main() {
    pid_t pid = getpid();
    long long virt_mem = get_total_vritual_memory();
    long long phys_mem = get_total_physical_memory();
    
    std::cout << "[Process: "<< pid << "] " << "TOTAL VIRTUAl RAM(KB): " << virt_mem << " PHYSICAL RAM(KB): " << phys_mem << "\n";

  return 0;
}
