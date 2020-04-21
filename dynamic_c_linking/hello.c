// Compile with go library
// $ cc hello.c ./static_c_linking -lpthread

#include "dynamic_c_linking.h"

int main(void)
{
  Hello();
  return 0;
}
