#include <unistd.h>
#include <signal.h>
#include <stdlib.h>

void signal_handler(int signum) {
    exit(0);
}

int main(void) {
    signal(SIGINT, signal_handler);
    while(1)
    {
        sleep(3600);
    }
    return 0;
}
