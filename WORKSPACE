http_archive(
    name = "io_bazel_rules_go",
    sha256 = "4d8d6244320dd751590f9100cf39fd7a4b75cd901e1f3ffdfd6f048328883695",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.9.0/rules_go-0.9.0.tar.gz",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "e3dadf036c769d1f40603b86ae1f0f90d11837116022d9b06e4cd88cae786676",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.8/bazel-gazelle-0.8.tar.gz",
)

git_repository(
    name = "org_pubref_rules_protobuf",
    commit = "023cd8ddf51d8a52fadcb46883025d9bd190750a",
    remote = "https://github.com/pubref/rules_protobuf",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_rules_dependencies",
    "go_register_toolchains",
    "go_repository",
)
load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")

go_rules_dependencies()

go_register_toolchains()

go_proto_repositories()

git_repository(
    name = "protoc_gen_persist",
    commit = "2abdcc3ff6028b9e1d2ef63f1c1fe0ed172338f5",
    remote = "https://github.com/tcncloud/protoc-gen-persist",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    commit = "d682213848ed68c0a260ca37d6dd5ace8423f5ba",
    importpath = "github.com/sirupsen/logrus",
)

go_repository(
    name = "com_google_cloud_go",
    commit = "d3a4b58ba5b65453c14062db6a651131a89f0f6e",
    importpath = "cloud.google.com/go",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "d670f9405373e636a5a2765eea47fac0c9bc91a4",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "io_opencensus_go",
    commit = "22f2851c619f086d3cc2845c4ec1873aea3e24e4",
    importpath = "go.opencensus.io",
)

go_repository(
    name = "org_golang_google_api",
    commit = "d2c53ea20b719a26e291430d99eaaf6c9a8eb68c",
    importpath = "google.golang.org/api",
)

go_repository(
    name = "org_golang_google_appengine",
    commit = "150dc57a1b433e64154302bdc40b6bb8aefa313a",
    importpath = "google.golang.org/appengine",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "a8101f21cf983e773d0c1133ebc5424792003214",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "1cd234627e6f392ade0527d593eb3fe53e832d4a",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "b3c9a1d25cfbbbab0ff4780b71c4f54e6e92a0de",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "ab555f366c4508dbe0802550b1b20c46c5c18aa0",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "30785a2c434e431ef7c507b54617d6a951d5f2b4",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "810d7000345868fc619eb81f46307107118f4ae1",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "e19ae1496984b1c655b8044a65c0300a3c878dd3",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "fbec762f837dc349b73d1eaa820552e2ad177942",
    importpath = "golang.org/x/tools",
)

go_repository(
    name = "com_github_burntsushi_toml",
    commit = "a368813c5e648fee92e5f6c30e3944ff9d5e8895",
    importpath = "github.com/BurntSushi/toml",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    tag = "master",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "1cd234627e6f392ade0527d593eb3fe53e832d4a",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "master",
)