package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "resource-consumer",
    library = ":go_default_library",
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "resource_consumer.go",
        "resource_consumer_handler.go",
        "utils.go",
    ],
    tags = ["automanaged"],
    deps = ["//test/images/resource-consumer/common:go_default_library"],
)
