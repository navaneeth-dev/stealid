# CMakeList.txt : CMake project for stealid-implant, include source and define
# project specific logic here.
#

# Add source to this project's executable.
add_executable (stealid-implant "stealid-implant.cpp" "stealid-implant.h")

if (CMAKE_VERSION VERSION_GREATER 3.12)
  set_property(TARGET stealid-implant PROPERTY CXX_STANDARD 20)
endif()

# TODO: Add tests and install targets if needed.
