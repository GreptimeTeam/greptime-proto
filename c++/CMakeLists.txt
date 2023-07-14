cmake_minimum_required(VERSION 3.22)

project(greptime-proto-cpp)

file(GLOB_RECURSE OUTPUT_SOURCES "*.h" "*.cc")

add_library(${PROJECT_NAME} ${OUTPUT_SOURCES})

target_include_directories(${PROJECT_NAME} PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

target_link_libraries(${PROJECT_NAME} INTERFACE gRPC::grpc gRPC::grpc++ gRPC::grpc++_reflection)