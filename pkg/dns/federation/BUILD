package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["federation.go"],
    tags = ["automanaged"],
    deps = ["//pkg/util/validation:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["federation_test.go"],
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)
