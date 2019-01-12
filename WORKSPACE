load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7be7dc01f1e0afdba6c8eb2b43d2fa01c743be1b9273ab1eaf6c233df078d705",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.5/rules_go-0.16.5.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "7949fc6cc17b5b191103e97481cf8889217263acf52e00b560683413af204fcb",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.16.0/bazel-gazelle-0.16.0.tar.gz"],
)

# load go rules
load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

# load gazelle
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# external dependencies

go_repository(
    name = "com_github_asaskevich_govalidator",
    commit = "ccb8e960c48f04d6935e72476ae4a51028f9e22f",
    importpath = "github.com/asaskevich/govalidator",
)

go_repository(
    name = "com_github_docker_go_units",
    commit = "47565b4f722fb6ceae66b95f853feed578a4a51c",
    importpath = "github.com/docker/go-units",
)

go_repository(
    name = "com_github_globalsign_mgo",
    commit = "eeefdecb41b842af6dc652aaea4026e8403e62df",
    importpath = "github.com/globalsign/mgo",
)

go_repository(
    name = "com_github_go_openapi_analysis",
    commit = "e2f3fdbb7ed0e56e070ccbfb6fc75b288a33c014",
    importpath = "github.com/go-openapi/analysis",
)

go_repository(
    name = "com_github_go_openapi_errors",
    commit = "7a7ff1b7b8020f22574411a32f28b4d168d69237",
    importpath = "github.com/go-openapi/errors",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    commit = "ef5f0afec364d3b9396b7b77b43dbe26bf1f8004",
    importpath = "github.com/go-openapi/jsonpointer",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    commit = "8483a886a90412cd6858df4ea3483dce9c8e35a3",
    importpath = "github.com/go-openapi/jsonreference",
)

go_repository(
    name = "com_github_go_openapi_loads",
    commit = "74628589c3b94e3526a842d24f46589980f5ab22",
    importpath = "github.com/go-openapi/loads",
)

go_repository(
    name = "com_github_go_openapi_runtime",
    commit = "41e24cc66d7af6af39eb9b5a6418e901bcdd333c",
    importpath = "github.com/go-openapi/runtime",
)

go_repository(
    name = "com_github_go_openapi_spec",
    commit = "5b6cdde3200976e3ecceb2868706ee39b6aff3e4",
    importpath = "github.com/go-openapi/spec",
)

go_repository(
    name = "com_github_go_openapi_strfmt",
    commit = "e471370ae57ac74eaf0afe816a66e4ddd7f1b027",
    importpath = "github.com/go-openapi/strfmt",
)

go_repository(
    name = "com_github_go_openapi_swag",
    commit = "1d29f06aebd59ccdf11ae04aa0334ded96e2d909",
    importpath = "github.com/go-openapi/swag",
)

go_repository(
    name = "com_github_go_openapi_validate",
    commit = "d2eab7d93009e9215fc85b2faa2c2f2a98c2af48",
    importpath = "github.com/go-openapi/validate",
)

go_repository(
    name = "com_github_gobuffalo_envy",
    commit = "801d7253ade1f895f74596b9a96147ed2d3b087e",
    importpath = "github.com/gobuffalo/envy",
)

go_repository(
    name = "com_github_gobuffalo_packd",
    commit = "eca3b8fd66872a76119b189bbe0f58f776a2a39f",
    importpath = "github.com/gobuffalo/packd",
)

go_repository(
    name = "com_github_gobuffalo_packr",
    commit = "679459352e18b4c74274a979d695fe13b68730c1",
    importpath = "github.com/gobuffalo/packr",
)

go_repository(
    name = "com_github_gobuffalo_syncx",
    commit = "558ac7de985fc4f4057bff27c7bdf99e92fe0750",
    importpath = "github.com/gobuffalo/syncx",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "c6ca198ec95c841fdb89fc0de7496fed11ab854e",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "com_github_joho_godotenv",
    commit = "23d116af351c84513e1946b527c88823e476be13",
    importpath = "github.com/joho/godotenv",
)

go_repository(
    name = "com_github_joonix_log",
    commit = "d2d3f2f4a80658c67a0bc44021dcac30f3017c06",
    importpath = "github.com/joonix/log",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    commit = "5c8c8bd35d3832f5d134ae1e1e375b69a4d25242",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
)

go_repository(
    name = "com_github_mailru_easyjson",
    commit = "60711f1a8329503b04e1c88535f419d0bb440bff",
    importpath = "github.com/mailru/easyjson",
)

go_repository(
    name = "com_github_markbates_oncer",
    commit = "bf2de49a0be218916e69a11d22866e6cd0a560f2",
    importpath = "github.com/markbates/oncer",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "3536a929edddb9a5b34bd6861dc4a9647cb459fe",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "ba968bfe8b2f7e042a574c888954fccecfa385b4",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    commit = "0bcb03f4b4d0a9428594752bd2a3b9aa0a9d4bd4",
    importpath = "github.com/PuerkitoBio/purell",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    commit = "de5bf2ad457846296e2031421a34e2568e304e35",
    importpath = "github.com/PuerkitoBio/urlesc",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    commit = "e1e72e9de974bd926e5c56f83753fba2df402ce5",
    importpath = "github.com/sirupsen/logrus",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "51d6538a90f86fe93ac480b35f37b2be17fef232",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "ff983b9c42bc9fbf91556e191cc8efb585c16908",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "915654e7eabcea33ae277abbecf52f0d8b7a9fdc",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "48ac38b7c8cbedd50b1613c0fccacfc7d88dfcdf",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    commit = "d87f08a7d80821c797ffc8eb8f4e01675f378736",
    importpath = "github.com/rogpeppe/go-internal",
)
