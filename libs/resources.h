#include <iostream>
#include <fstream>
#include <string>
#include <sys/resource.h>
#include <unistd.h>
#include <sys/sysinfo.h>
#include <algorithm>
#include <charconv>
#include <fstream>
#include <regex>
#include <string>

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
struct ProcessStatus {
  long long anon_mem;
  // Add other desired fields here
};

static constexpr auto PROCFS_PATH = "/proc/%1%/status";

/// Extracts a single number from a string using regex
long long extractNumber(const std::string& str) noexcept {
  std::regex reg("([+-]?\\d+)");
  std::smatch match;
  if (!std::regex_search(str, match, reg))
    throw std::runtime_error("Failed to find number in string");
  return std::stoll(match[1]);
}

ProcessStatus get_process_status(pid_t pid) {
  auto path = PROCFS_PATH.arg(pid);
  std::ifstream fs(path);
  if (!fs.is_open())
    throw std::runtime_error("Could not open '" + path + "'");

  ProcessStatus ps{};
  std::string line;
  bool foundAnonMem = false;
  while (std::getline(fs, line).good()) {
    if (line.find("RssAnon:") != std::string::npos) {
      ps.anon_mem = extractNumber(line.substr(line.find(':') + 1));
      foundAnonMem = true;
    }
    // Parse other lines to fill other desired fields
  }

  if (!foundAnonMem)
    throw std::runtime_error("Could not find 'RssAnon' field in '" + path + "'");
  return ps;
}
