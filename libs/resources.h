#include <iostream>
#include <fstream>
#include <string>
#include <sys/resource.h>
#include <unistd.h>
#include <sys/sysinfo.h>

struct Memory {
public:
  double v; // Virtual memory
  double p; // Pyhsical memory
};

Memory getMemoryByPid(pid_t pid) {
  Memory mem;

  return mem;
}

struct MemoryInfo {
  long long swap;
  long long mem_unit;
  long long total_ram;
};

long long get_total_vritual_memory(){
  struct sysinfo memInfo;
  sysinfo (&memInfo);

  long long totalVirtualMem = memInfo.totalram;
  
  totalVirtualMem += memInfo.totalswap;
  totalVirtualMem *= memInfo.mem_unit;

  return totalVirtualMem;
}

long long get_total_physical_memory() {
  struct sysinfo memInfo;
  long long totalPhysMem = memInfo.totalram;
  
  totalPhysMem *= memInfo.mem_unit;
  return totalPhysMem;
}
