package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "register.go",
        "types.go",
        "zz_generated.conversion.go",
        "zz_generated.deepcopy.go",
    ],
    tags = ["automanaged"],
    deps = [
        "k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "k8s.io/apimachinery/pkg/conversion:go_default_library",
        "k8s.io/apimachinery/pkg/runtime:go_default_library",
        "k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "k8s.io/sample-apiserver/pkg/apis/wardle:go_default_library",
    ],
)
