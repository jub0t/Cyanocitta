#include <iostream>
#include <fstream>
#include <string>
#include <sys/resource.h>
#include <unistd.h>

// Function to get the current CPU usage of a process
double getCPUUsage(pid_t pid)
{
    struct rusage usage;
    if (getrusage(RUSAGE_SELF, &usage) == 0)
    {
        return usage.ru_utime.tv_sec + usage.ru_utime.tv_usec / 1000000.0;
    }
    return -1.0;
}

// Function to get the current RAM usage of a process
long getRAMUsage(pid_t pid)
{
    std::string statPath = "/proc/" + std::to_string(pid) + "/statm";
    std::ifstream statFile(statPath);
    if (statFile.is_open())
    {
        long size, resident, share, text, lib, data, dt;
        statFile >> size >> resident >> share >> text >> lib >> data >> dt;
        statFile.close();
        // Return the resident set size (RAM usage) in kilobytes
        return resident * sysconf(_SC_PAGESIZE) / 1024;
    }
    return -1;
}
