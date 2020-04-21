// Compile with go library
// $ cc hello.c ./static_c_linking.a -lpthread

#include "static_c_linking.h"

int main(void)
{
  Hello();
  return 0;
}
