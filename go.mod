module namespacelabs.dev/foundation

go 1.21

toolchain go1.21.5

require (
	buf.build/gen/go/namespace/cloud/grpc/go v1.3.0-20240202094328-7e6c2a95449c.2
	buf.build/gen/go/namespace/cloud/protocolbuffers/go v1.32.0-20240202094328-7e6c2a95449c.1
	cloud.google.com/go/artifactregistry v1.14.1
	cloud.google.com/go/container v1.26.0
	cuelang.org/go v0.4.3
	filippo.io/age v1.0.0
	github.com/agext/levenshtein v1.2.3
	github.com/andybalholm/brotli v1.0.4
	github.com/aws/aws-sdk-go-v2 v1.21.0
	github.com/aws/aws-sdk-go-v2/config v1.18.16
	github.com/aws/aws-sdk-go-v2/credentials v1.13.16
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.15.13
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.6
	github.com/aws/aws-sdk-go-v2/service/eks v1.21.2
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.13
	github.com/aws/aws-sdk-go-v2/service/s3 v1.30.6
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.6
	github.com/aws/smithy-go v1.14.2
	github.com/bcicen/jstream v1.0.1
	github.com/bradleyjkemp/cupaloy v2.3.0+incompatible
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/cespare/xxhash/v2 v2.2.0
	github.com/charmbracelet/bubbles v0.14.0
	github.com/charmbracelet/bubbletea v0.23.1
	github.com/charmbracelet/glamour v0.5.0
	github.com/charmbracelet/lipgloss v0.6.0
	github.com/compose-spec/compose-go v1.14.0
	github.com/containerd/console v1.0.3
	github.com/containerd/containerd v1.7.6
	github.com/containerd/nerdctl v1.2.1
	github.com/containerd/stargz-snapshotter/estargz v0.14.3
	github.com/creack/pty v1.1.18
	github.com/docker/buildx v0.11.2
	github.com/docker/cli v24.0.6+incompatible
	github.com/docker/docker v24.0.6+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.5.0
	github.com/dustin/go-humanize v1.0.1
	github.com/exaring/otelpgx v0.5.3
	github.com/fsnotify/fsevents v0.1.1
	github.com/fsnotify/fsnotify v1.6.0
	github.com/getsentry/sentry-go v0.16.0
	github.com/go-errors/errors v1.4.2
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.5
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/protobuf v1.5.3
	github.com/google/go-cmp v0.6.0
	github.com/google/go-containerregistry v0.19.1
	github.com/google/uuid v1.4.0
	github.com/gorilla/csrf v1.7.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.1
	github.com/gravitational/teleport/api v0.0.0-20231212204825-da589355d4ea
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/vault-client-go v0.4.3
	github.com/hpcloud/tail v1.0.0
	github.com/iancoleman/strcase v0.2.0
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/jackc/pgx/v5 v5.5.3
	github.com/jet/kube-webhook-certgen v1.5.2
	github.com/jhump/protoreflect v1.12.1-0.20220417024638-438db461d753
	github.com/jpillora/chisel v1.9.1
	github.com/keikoproj/aws-auth v0.4.0
	github.com/klauspost/compress v1.16.5
	github.com/kr/text v0.2.0
	github.com/maruel/panicparse/v2 v2.3.0
	github.com/mattn/go-isatty v0.0.17
	github.com/mattn/go-zglob v0.0.3
	github.com/miekg/dns v1.1.50
	github.com/minio/minio-go/v7 v7.0.47
	github.com/moby/buildkit v0.12.3
	github.com/moby/patternmatcher v0.5.0
	github.com/morikuni/aec v1.0.0
	github.com/muesli/cancelreader v0.2.2
	github.com/muesli/reflow v0.3.0
	github.com/natefinch/atomic v1.0.1
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.1.0-rc5
	github.com/philopon/go-toposort v0.0.0-20170620085441-9be86dbd762f
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.16.0
	github.com/protocolbuffers/txtpbfmt v0.0.0-20220608084003-fc78c767cd6a
	github.com/rs/zerolog v1.20.0
	github.com/sirupsen/logrus v1.9.3
	github.com/slack-go/slack v0.12.3
	github.com/soheilhy/cmux v0.1.5
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.14.0
	github.com/tailscale/hujson v0.0.0-20221223112325-20486734a56a
	github.com/tonistiigi/fsutil v0.0.0-20230629203738-36ef4d8c0dbb
	go.lsp.dev/jsonrpc2 v0.10.0
	go.lsp.dev/protocol v0.12.0
	go.lsp.dev/uri v0.3.0
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.45.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.42.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0
	go.opentelemetry.io/otel/exporters/prometheus v0.42.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.19.0
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/sdk/metric v1.19.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/atomic v1.10.0
	go.uber.org/automaxprocs v1.5.1
	go.uber.org/zap v1.25.0
	golang.org/x/crypto v0.17.0
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb
	golang.org/x/mod v0.11.0
	golang.org/x/net v0.19.0
	golang.org/x/oauth2 v0.11.0
	golang.org/x/sync v0.3.0
	golang.org/x/sys v0.15.0
	golang.org/x/term v0.15.0
	google.golang.org/api v0.136.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230920204549-e6e6cdab5c13
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231009173412-8bfb1ae86b6c
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.32.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v3 v3.0.1
	gotest.tools v2.2.0+incompatible
	helm.sh/helm/v3 v3.13.1
	honnef.co/go/tools v0.4.2
	inet.af/tcpproxy v0.0.0-20231102063150-2862066fc2a9
	k8s.io/api v0.28.3
	k8s.io/apiextensions-apiserver v0.28.3
	k8s.io/apimachinery v0.28.3
	k8s.io/client-go v0.28.3
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2
	namespacelabs.dev/go-filenotify v0.0.0-20220511192020-53ea11be7eaa
	namespacelabs.dev/go-ids v0.0.0-20221124082625-9fc72ee06af7
	namespacelabs.dev/integrations v0.0.0-20240204232225-854c7fa4534a
	sigs.k8s.io/controller-runtime v0.16.3
	sigs.k8s.io/yaml v1.3.0
)

require (
	cloud.google.com/go v0.110.8 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.2 // indirect
	cloud.google.com/go/longrunning v0.5.1 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/AdamKorcz/go-118-fuzz-build v0.0.0-20230306123547-8075edf89bb0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/Masterminds/squirrel v1.5.4 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Microsoft/hcsshim v0.11.0 // indirect
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.41 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.35 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.21.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.5 // indirect
	github.com/aymanbagabas/go-osc52 v1.2.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beevik/etree v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/cockroachdb/apd/v2 v2.0.1 // indirect
	github.com/containerd/cgroups v1.1.0 // indirect
	github.com/containerd/continuity v0.4.2 // indirect
	github.com/containerd/fifo v1.1.0 // indirect
	github.com/containerd/go-cni v1.1.9 // indirect
	github.com/containerd/ttrpc v1.2.2 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/containernetworking/cni v1.1.2 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/distribution/v3 v3.0.0-20230214150026-36d8c594d7aa // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/docker/go v1.5.1-1.0.20160303222718-d30aec9fd63c // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fvbommel/sortorder v1.1.0 // indirect
	github.com/go-gorp/gorp/v3 v3.1.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-piv/piv-go v1.11.0 // indirect
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gosuri/uitable v0.0.4 // indirect
	github.com/gravitational/trace v1.3.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.2 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/in-toto/in-toto-golang v0.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jpillora/sizestr v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jxskiss/base62 v1.1.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.1 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mattn/go-shellwords v1.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.21 // indirect
	github.com/miekg/pkcs11 v1.1.1 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/signal v0.7.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
	github.com/muesli/ansi v0.0.0-20221106050444-61f0cd9a192a // indirect
	github.com/muesli/termenv v0.13.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opencontainers/runc v1.1.7 // indirect
	github.com/opencontainers/runtime-spec v1.1.0-rc.2 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	github.com/rootless-containers/rootlesskit v1.1.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/rubenv/sql-migrate v1.5.2 // indirect
	github.com/russellhaering/gosaml2 v0.9.1 // indirect
	github.com/russellhaering/goxmldsig v1.4.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sahilm/fuzzy v0.1.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.4.0 // indirect
	github.com/segmentio/asm v1.1.3 // indirect
	github.com/segmentio/encoding v0.3.4 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/shibumi/go-pathspec v1.3.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/theupdateframework/notary v0.7.0 // indirect
	github.com/tonistiigi/units v0.0.0-20180711220420-6950e57a87ea // indirect
	github.com/tonistiigi/vt100 v0.0.0-20230623042737-f9a4f7ef6531 // indirect
	github.com/vbatts/tar-split v0.11.3 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	github.com/yuin/goldmark-emoji v1.0.1 // indirect
	go.lsp.dev/pkg v0.0.0-20210717090340-384b27a52fb2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.40.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.42.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiserver v0.28.3 // indirect
	k8s.io/cli-runtime v0.28.2 // indirect
	k8s.io/component-base v0.28.3 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9 // indirect
	k8s.io/kubectl v0.28.2 // indirect
	oras.land/oras-go v1.2.4 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.5-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.3-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)
